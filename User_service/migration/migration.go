package main

import (
	"log"

	"github.com/LoTfI01101011/E-commerce/User_service/internal"
	"github.com/LoTfI01101011/E-commerce/User_service/models"
)

func init() {
	internal.DbConnection()
}
func main() {
	internal.DB.Migrator().DropTable(&models.User{})
	err := internal.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Print(err)
	}
}
