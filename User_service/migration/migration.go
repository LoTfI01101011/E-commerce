package main

import (
	"github.com/LoTfI01101011/E-commerce/User_service/internal"
	"github.com/LoTfI01101011/E-commerce/User_service/models"
)

func init() {
	internal.DbConnection()
}
func main() {
	internal.DB.AutoMigrate(&models.User{})
}
