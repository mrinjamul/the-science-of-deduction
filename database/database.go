package database

import (
	"log"

	"github.com/mrinjamul/the-science-of-deduction/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// IsConnected returns the connection status
	IsConnected bool
	IsSQLite    bool
)

func GetDB() *gorm.DB {
	var db *gorm.DB
	var err error
	// Get ENV variables
	// dbHost := os.Getenv("POSTGRES_HOST")
	// dbName := os.Getenv("POSTGRES_DB")
	// dbUser := os.Getenv("POSTGRES_USER")
	// dbPassword := os.Getenv("POSTGRES_PASSWORD")
	// dbPort := os.Getenv("POSTGRES_PORT")
	// if db == nil {
	// 	if dbHost == "" {
	// 		fmt.Println("Environment variable DB_HOST is null.")
	// 		return nil
	// 	}
	// 	if dbName == "" {
	// 		fmt.Println("Environment variable DB_NAME is null.")
	// 		return nil
	// 	}
	// 	if dbUser == "" {
	// 		fmt.Println("Environment variable DB_USERNAME is null.")
	// 		return nil
	// 	}
	// 	if dbPassword == "" {
	// 		fmt.Println("Environment variable DB_PASSWORD is null.")
	// 		return nil
	// 	}

	// 	if dbPort == "" {
	// 		dbPort = "5432"
	// 	}
	// }

	// // Connect to db
	// dest := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
	// 	dbHost, dbUser, dbPassword, dbName, dbPort)
	// db, err := gorm.Open(postgres.Open(dest), &gorm.Config{})

	// if err == nil {
	// 	IsConnected = true
	// } else {
	// 	log.Println("failed to connect database")
	// }

	// if unable to connect to database, create sqlite database
	if !IsConnected {
		// Create sqlite connection
		db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
		if err == nil {
			IsConnected = true
			IsSQLite = true
		} else {
			log.Println("failed to connect database")
		}
	}

	// Migrate the schema
	db.AutoMigrate(&models.CaseFiles{})

	return db
}
