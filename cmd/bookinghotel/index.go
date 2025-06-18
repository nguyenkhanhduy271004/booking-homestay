package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/config"
	"homestay.com/nguyenduy/internal/routes"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routes.InitRoute(db)
	router.Run()
}
