package db

import (
	"fmt"
	"log"

	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/accounts-svc/internal/domain/models"

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
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Couldn't migrate models.User. Error:", err)
	}

	if err := DB.AutoMigrate(&models.Admin{}); err != nil {
		log.Fatal("Couldn't migrate models.Admin. Error:", err)
	}

	if err := DB.AutoMigrate(&models.Department{}); err != nil {
		log.Fatal("Couldn't migrate models.Department. Error:", err)
	}

	if err := DB.AutoMigrate(&models.Office{}); err != nil {
		log.Fatal("Couldn't migrate models.Office. Error:", err)
	}

	fmt.Println("Migrated tables successfully")
}
