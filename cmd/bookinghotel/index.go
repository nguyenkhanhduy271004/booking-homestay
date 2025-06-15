package main

import (
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/config"
	"homestay.com/nguyenduy/internal/routes"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := routes.InitRoute(db)
	router.Run()
}
