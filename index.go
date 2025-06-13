package main

import (
	"gorm.io/gorm"
	"homestay.com/nguyenduy/config"
	"homestay.com/nguyenduy/routes"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := routes.InitRoute(db)
	router.Run()
}
