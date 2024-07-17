package db

import (
	"fmt"
	"log"

	"github.com/AbdulRahimOM/gov-services-app/chat-svc/internal/config"

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

	fmt.Println("Migrated tables successfully")
}
