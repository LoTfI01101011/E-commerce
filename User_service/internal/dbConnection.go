package internal

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	LoadEnv()
}

func DbConnection() {
	dsn := "host=localhost " + "user=" + os.Getenv("User") + " password=" + os.Getenv("Password") + " dbname=" + os.Getenv("DbName") + " port=" + os.Getenv("Port") + " sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connecting the db")
	}
}
