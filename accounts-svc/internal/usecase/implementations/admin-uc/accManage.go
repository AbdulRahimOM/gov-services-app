package adminuc

import (
	"fmt"
	"strings"

	hashpassword "github.com/AbdulRahimOM/go-utils/hashPassword"
	helper "github.com/AbdulRahimOM/go-utils/helper"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/dto/request"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	commondto "github.com/AbdulRahimOM/gov-services-app/internal/common-dto"
	respcode "github.com/AbdulRahimOM/gov-services-app/internal/std-response/response-code"
)

// AdminGetAdmins
func (a *AdminUseCase) AdminGetAdmins(adminID int32, searchCriteria *request.AdminSearchCriteria) (*[]commondto.Admin, string, error) {
	admins, err := a.adminRepo.AdminGetAdmins(searchCriteria)
	if err != nil {
		return nil, respcode.DBError, fmt.Errorf("@db: failed to get admins: %v", err)
	}
	return admins, "", nil
}

// AdminAddSubAdmin
func (a *AdminUseCase) AdminAddSubAdmin(req *request.AdminAddSubAdmin) (int32, string, error) {
	//check if current admin rank is superior to proposed sub admin rank
	//get rank by AdminID
	currentAdminRankID, err := a.adminRepo.GetRankByAdminID(req.AdminID)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to get current admin rank: %v", err)
	}
	if currentAdminRankID >= req.NewSubAdmin.RankID {
		return 0, respcode.Unauthorized, fmt.Errorf("current admin rank is not superior to proposed sub admin rank")
	}

	var doExist bool
	var randomUserName string
	tryCount := 0
	for doExist || tryCount < 5{
		randomUserName = generateUsername(req.NewSubAdmin.FirstName, req.NewSubAdmin.LastName, req.NewSubAdmin.DeptID)
		//check if username already exists
		doExist, err = a.adminRepo.CheckIfAdminUsernameExists(&randomUserName)
		if err != nil {
			return 0, respcode.DBError, fmt.Errorf("@db: failed to check if admin username exists: %v", err)
		}
		tryCount++
	}
	if doExist {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to generate unique username")
	}

	randomPassword := helper.GenerateRandomAlphanumeric(10)
	//hash password
	hashedPassword, err := hashpassword.Hashpassword(randomPassword)
	if err != nil {
		return 0, respcode.FailureToGenerate, fmt.Errorf("failed to hash password: %v", err)
	}

	newSubAdmin:=models.Admin{
		FName: req.NewSubAdmin.FirstName,
		LName: req.NewSubAdmin.LastName,
		Email: req.NewSubAdmin.Email,
		PhoneNumber: req.NewSubAdmin.PhoneNumber,
		DeptID: req.NewSubAdmin.DeptID,
		Designation: req.NewSubAdmin.Designation,
		RankID: req.NewSubAdmin.RankID,
		OfficeID: req.NewSubAdmin.OfficeID,
		Username: randomUserName,
		HashedPW: hashedPassword,
	}

	//add sub admin
	newSubAdminID, err := a.adminRepo.AddSubAdmin(&newSubAdmin)
	if err != nil {
		return 0, respcode.DBError, fmt.Errorf("@db: failed to add sub admin: %v", err)
	}

	//send credentials to sub admin
	err = sendCredentialsToSubAdmin(req.NewSubAdmin.Email, randomUserName, randomPassword)
	if err != nil {
		return 0, respcode.OtherInternalError, fmt.Errorf("failed to send credentials to sub admin: %v", err)
	}
	return newSubAdminID, "", nil
}

func generateUsername(firstName, lastName string, deptID int32) string {
	//generate username
	//remove spaces and convert to lowercase
	firstName = strings.ReplaceAll(strings.ToLower(firstName), " ", "")
	lastName = strings.ReplaceAll(strings.ToLower(lastName), " ", "")

	//generate a 5 alphanumeric character suffix
	suffix := helper.GenerateRandomAlphanumeric(5)
	username := fmt.Sprintf("%s.%s.%d.%s", firstName, lastName, deptID, suffix)
	return username
}

func sendCredentialsToSubAdmin(email, username, password string) error {
	fmt.Println("sending credentials to sub admin...")
	fmt.Println("To be implemented...")
	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
	fmt.Println("Email: ", email)
	//send email
	//send sms
	return nil
}
