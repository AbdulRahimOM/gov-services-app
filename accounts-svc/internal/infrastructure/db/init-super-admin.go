package db

import (
	"log"
	"time"

	hashpassword "github.com/AbdulRahimOM/go-utils/hashPassword"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"
	"github.com/AbdulRahimOM/gov-services-app/internal/project/data"
)

// func init() {

// }

const (
	superAdminID          = 1
	superAdminFirstName   = "superAdmin"
	superAdminLastName    = ""
	superAdminRankID      = 1
	superAdminDeptID      = 1
	superAdminDesignation = "superAdmin"
)

func doSuperAdminExists() (bool, error) {
	var count int32
	result := DB.Raw("SELECT COUNT(*) FROM admins WHERE id=1").Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func initSuperAdminIfNotInitialized() {
	//check if super admin is already initialized
	adminExists, err := doSuperAdminExists()
	if err != nil {
		log.Fatal("failed to check if super admin exists: ", err)
	}
	if adminExists {
		log.Println("Super admin exists")
		return
	} else {
		log.Println("Super admin does not exist")
	}

	log.Println("Initializing super admin from environment variables")
	superAdminUsername, superAdminPassword := config.GetSuperAdminCredentials()
	if superAdminUsername == "" || superAdminPassword == "" {
		log.Fatal("SUPER_ADMIN_USERNAME or SUPER_ADMIN_PASSWORD is not set in the environment")
	}

	//seed super admin
	hashedPW, err := hashpassword.Hashpassword(superAdminPassword)
	if err != nil {
		log.Fatal("failed to hash initial super-admin-password: ", err)
	}

	var superAdmin = models.Admin{
		ID:          superAdminID,
		FName:       superAdminFirstName,
		LName:       superAdminLastName,
		Username:    superAdminUsername,
		HashedPW:    hashedPW,
		OfficeID:    1,
		Designation: data.Designation_SuperAdmin,
		CreatedBy:   0,
		CreatedAt:   time.Now(),
		IsActive:    true,
	}

	result := DB.Create(&superAdmin)
	if result.Error != nil {
		log.Fatal("failed to create super admin: ", result.Error)
	}

	log.Println("Super admin initialized successfully")

}
