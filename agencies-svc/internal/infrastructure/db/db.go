package db

import (
	"fmt"
	"log"

	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/agencies-svc/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	connectToDB()
	migrateTables()
}

func connectToDB() {
	dsn := config.EnvValues.DbUrl
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to DB. Error:", err)
	}
}
func migrateTables() {

	if err := DB.AutoMigrate(&models.KsebSectionCode{}); err != nil {
		log.Fatal("Couldn't migrate models.KsebSectionCode. Error:", err)
	}

	if err := DB.AutoMigrate(&models.UserKsebConsumerNumber{}); err != nil {
		log.Fatal("Couldn't migrate models.KsebConsumerNumber. Error:", err)
	}

	if err := DB.AutoMigrate(&models.KsebComplaint{}); err != nil {
		log.Fatal("Couldn't migrate models.KsebComplaint. Error:", err)
	}

	fmt.Println("Migrated tables successfully")
}
