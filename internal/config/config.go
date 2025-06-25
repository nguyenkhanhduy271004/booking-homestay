package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	model "homestay.com/nguyenduy/internal/app/models"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	required := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_NAME"}
	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Missing required environment variable: %s", key)
		}
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to database")

	migrateModels(db)

	return db
}

func migrateModels(db *gorm.DB) {
	models := []interface{}{
		&model.Permission{},
		&model.Role{},
		&model.User{},
		&model.Hotel{},
		&model.RoomType{},
		&model.Room{},
		&model.Staff{},
		&model.Guest{},
		&model.Booking{},
		&model.Payment{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Fatalf("Failed to migrate model %T: %v", m, err)
		}
	}

	log.Println("Database migration completed")
}

func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance for closing:", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close database connection:", err)
	}

	log.Println("Database connection closed")
}
