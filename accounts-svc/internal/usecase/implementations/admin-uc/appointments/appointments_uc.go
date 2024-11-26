package appointmentsuc

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	hashpassword "github.com/AbdulRahimOM/go-utils/hashPassword"
	"github.com/AbdulRahimOM/go-utils/helper"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	repo "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/repository/interface"
	usecase "github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/usecase/interface"
	requests "github.com/AbdulRahimOM/gov-services-app/internal/common-dto/request"
	"github.com/AbdulRahimOM/gov-services-app/internal/project/data"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
	"github.com/atotto/clipboard"
)

const (
	ksebBottomOfficeRank = 8
)

type AppointmentUseCase struct {
	adminRepo repo.IAdminRepo
}

func NewAppointmentUseCase(adminRepo repo.IAdminRepo) usecase.IAppointmentUC {
	return &AppointmentUseCase{
		adminRepo: adminRepo,
	}
}

// CreateChildOffice
func (uc *AppointmentUseCase) CreateChildOffice(adminID int32, proposedChildOffice *requests.ProposedOffice) (int32, string, error) {
	officeDetails, err := uc.adminRepo.GetOfficeDetailsByAdminID(adminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get office id of admin")
	}
	_, respCode, err := uc.checkIfOfficeCanCreateSubOffice(officeDetails.ID, officeDetails.DeptID)
	if err != nil {
		return 0, respCode, fmt.Errorf("while checking if office can create sub office, error: %v", err)
	}

	//check if the proposed office name already exists
	doExist, err := uc.adminRepo.CheckIfOfficeNameExists(&proposedChildOffice.Name)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to check if office name exists: %v", err)
	}
	if doExist {
		return 0, respcode.NameExists, fmt.Errorf("office name already exists")
	}

	//create office
	childOfficeID, err := uc.adminRepo.AddChildOffice(&models.Office{
		Name:             proposedChildOffice.Name,
		Address:          proposedChildOffice.Address,
		DeptID:           officeDetails.DeptID,
		Rank:             officeDetails.Rank + 1,
		SuperiorOfficeID: officeDetails.ID,
	})
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("error in adding child office: %v", err)
	}

	return childOfficeID, "", nil
}

// AppointAttender
func (uc *AppointmentUseCase) AppointAttender(adminID int32, appointee *requests.Appointee) (int32, string, error) {
	officeID, err := uc.adminRepo.GetOfficeIDByAdminID(adminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get office id of admin")
	}

	//check if admin is above rank 2
	rankOfAdmin, err := uc.adminRepo.GetRankOfOffice(officeID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get rank of admin")
	}
	if rankOfAdmin < 2 {
		return 0, respcode.Unauthorized, fmt.Errorf("admin rank should not be 1 or 2")
	}

	//generate unique username
	var doExist bool
	var randomUserName string
	tryCount := 0
	for doExist || tryCount < 3 {
		randomUserName = generateUsername(appointee.FirstName, appointee.LastName, 3+tryCount)
		//check if username already exists
		doExist, err = uc.adminRepo.CheckIfAdminUsernameExists(&randomUserName)
		if err != nil {
			return 0, respcode.DBError, fmt.Errorf("@db: failed to check if admin username exists: %v", err)
		}
		tryCount++
	}
	if doExist {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to generate unique username")
	}

	//generate random password
	randomPassword := helper.GenerateRandomAlphanumeric(10)
	//hash password
	hashedPassword, err := hashpassword.Hashpassword(randomPassword)
	if err != nil {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to hash password: %v", err)
	}

	//create admin
	adminId, err := uc.adminRepo.AddSubAdmin(&models.Admin{
		FName:       appointee.FirstName,
		LName:       appointee.LastName,
		Email:       appointee.Email,
		PhoneNumber: appointee.PhoneNumber,
		OfficeID:    officeID,
		Username:    randomUserName,
		HashedPW:    hashedPassword,
		Designation: data.Designation_Attender,
	})
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("error in adding sub admin: %v", err)
	}

	//send credentials to sub admin
	err = sendCredentialsToSubAdmin(appointee.Email, randomUserName, randomPassword)
	if err != nil {
		return 0, respcode.OtherInternalError, fmt.Errorf("failed to send credentials to sub admin: %v", err)
	}

	return adminId, "", nil
}

func generateUsername(firstName, lastName string, length int) string {
	//generate username
	//remove spaces and convert to lowercase
	firstName = strings.ReplaceAll(strings.ToLower(firstName), " ", "")
	lastName = strings.ReplaceAll(strings.ToLower(lastName), " ", "")

	//generate a 5 alphanumeric character suffix
	suffix := helper.GenerateRandomNumeric(length)
	username := fmt.Sprintf("%s%s.%s", firstName, lastName, suffix)
	return username
}

func sendCredentialsToSubAdmin(email, username, password string) error {
	fmt.Println("sending credentials to sub admin...")

	if config.EnvValues.Environment == "development" {
		copyText := `{
	"username": "` + username + `",
	"password": "` + password + `"
}`
		err := clipboard.WriteAll(copyText)
		if err != nil {
			fmt.Println("failed to copy to clipboard: ", err)
		}
	}
	sendEmail(email, "Credentials", "Username: "+username+"\nPassword: "+password)

	return nil
}
func sendEmail(to, subject, body string) error {
	from := config.Emailing.FromEmail
	auth := smtp.PlainAuth("", from, config.Emailing.AppPassword, config.Emailing.SmtpServerAddress)
	addr := fmt.Sprintf("%s:%s", config.Emailing.SmtpServerAddress, config.Emailing.SmtpsPort)

	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body))

	err := smtp.SendMail(addr, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

// AppointChildOfficeHead
func (uc *AppointmentUseCase) AppointChildOfficeHead(adminID int32, childOfficeID int32, appointee *requests.Appointee) (int32, string, error) {
	//check if the post is occupied
	isOccupied, err := uc.adminRepo.CheckOccupancyByDesignation(childOfficeID, data.Designation_OfficeHead)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't check if post is occupied: %v", err)
	}
	if isOccupied {
		return 0, respcode.PostOccupied, fmt.Errorf("post is already occupied")
	}

	officeID, err := uc.adminRepo.GetOfficeIDByAdminID(adminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get office id of admin")
	}

	//check if admin is allowed to appoint child office head
	superiorOfficeOfChildOffice, err := uc.adminRepo.GetSuperiorOfficeIdByOfficeId(childOfficeID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get superior office id of child office")
	}

	if superiorOfficeOfChildOffice != officeID {
		return 0, respcode.Unauthorized, fmt.Errorf("admin is not the superior of the child office")
	}

	//generate unique username
	var doExist bool
	var randomUserName string
	tryCount := 0
	for doExist || tryCount < 3 {
		randomUserName = generateUsername(appointee.FirstName, appointee.LastName, 3+tryCount)
		//check if username already exists
		doExist, err = uc.adminRepo.CheckIfAdminUsernameExists(&randomUserName)
		if err != nil {
			return 0, respcode.DBError, fmt.Errorf("@db: failed to check if admin username exists: %v", err)
		}
		tryCount++
	}
	if doExist {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to generate unique username")
	}

	//generate random password
	randomPassword := helper.GenerateRandomAlphanumeric(10)
	//hash password
	hashedPassword, err := hashpassword.Hashpassword(randomPassword)
	if err != nil {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to hash password: %v", err)
	}

	//create admin
	adminId, err := uc.adminRepo.AddSubAdmin(&models.Admin{
		FName:       appointee.FirstName,
		LName:       appointee.LastName,
		Email:       appointee.Email,
		PhoneNumber: appointee.PhoneNumber,
		OfficeID:    childOfficeID,
		Username:    randomUserName,
		HashedPW:    hashedPassword,
		Designation: data.Designation_OfficeHead,
		CreatedBy:   adminID,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("error in adding sub admin: %v", err)
	}

	//send credentials to sub admin
	err = sendCredentialsToSubAdmin(appointee.Email, randomUserName, randomPassword)
	if err != nil {
		return 0, respcode.OtherInternalError, fmt.Errorf("failed to send credentials to sub admin: %v", err)
	}

	return adminId, "", nil
}

// AppointChildOfficeDeputyHead
func (uc *AppointmentUseCase) AppointChildOfficeDeputyHead(adminID int32, childOfficeID int32, appointee *requests.Appointee) (int32, string, error) {
	officeID, err := uc.adminRepo.GetOfficeIDByAdminID(adminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get office id of admin")
	}

	//check if admin is allowed to appoint child office deputy head
	superiorOfficeOfChildOffice, err := uc.adminRepo.GetSuperiorOfficeIdByOfficeId(childOfficeID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: couldn't get superior office id of child office")
	}

	if superiorOfficeOfChildOffice != officeID {
		return 0, respcode.Unauthorized, fmt.Errorf("admin is not the superior of the child office")
	}

	//generate unique username
	var doExist bool
	var randomUserName string
	tryCount := 0
	for doExist || tryCount < 3 {
		randomUserName = generateUsername(appointee.FirstName, appointee.LastName, 3+tryCount)
		//check if username already exists
		doExist, err = uc.adminRepo.CheckIfAdminUsernameExists(&randomUserName)
		if err != nil {
			return 0, respcode.DBError, fmt.Errorf("@db: failed to check if admin username exists: %v", err)
		}
		tryCount++
	}
	if doExist {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to generate unique username")
	}

	//generate random password
	randomPassword := helper.GenerateRandomAlphanumeric(10)
	//hash password
	hashedPassword, err := hashpassword.Hashpassword(randomPassword)
	if err != nil {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to hash password: %v", err)
	}

	//create admin
	adminId, err := uc.adminRepo.AddSubAdmin(&models.Admin{
		FName:       appointee.FirstName,
		LName:       appointee.LastName,
		Email:       appointee.Email,
		PhoneNumber: appointee.PhoneNumber,
		OfficeID:    officeID,
		Username:    randomUserName,
		HashedPW:    hashedPassword,
		Designation: data.Designation_DeputyOfficeHead,
	})
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("error in adding sub admin: %v", err)
	}

	//send credentials to sub admin
	err = sendCredentialsToSubAdmin(appointee.Email, randomUserName, randomPassword)
	if err != nil {
		return 0, respcode.OtherInternalError, fmt.Errorf("failed to send credentials to sub admin: %v", err)
	}

	return adminId, "", nil
}

func (uc *AppointmentUseCase) checkIfOfficeCanCreateSubOffice(officeID int32, deptID int32) (bool, string, error) {
	// check if admin is above rank 2
	rankOfOffice, err := uc.adminRepo.GetRankOfOffice(officeID)
	if err != nil {
		return false, respcode.DBError, fmt.Errorf("@db: couldn't get rank of admin")
	}
	if rankOfOffice <= 2 {
		return false, respcode.Unauthorized, fmt.Errorf("admin rank should not be 1 or 2")
	} else if rankOfOffice >= ksebBottomOfficeRank {
		return false, respcode.Unauthorized, fmt.Errorf("admin's office's heirarchial rank should not be %d or above", ksebBottomOfficeRank)
	}

	switch deptID {
	case 5: //KSEB
		if rankOfOffice == data.KSEB_Bottom_Office_Rank {
			return false, respcode.Unauthorized, fmt.Errorf("admin belongs to bottom office. No further sub office can be created")
		}
	}
	return true, "", nil
}
