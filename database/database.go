package database

import (
	"fmt"
	"os"

	"github.com/Arun-Kumar21/distributed-object-storage/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load()

	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + 
			" password=" + os.Getenv("POSTGRES_PASSWORD") + 
			" dbname=" + os.Getenv("POSTGRES_DB") + 
			" port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	
	db, err := 	gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect with database")
	}

	db.AutoMigrate(&models.File{}, &models.User{})
	DB = db
	fmt.Println("Database Connected Successfully")
}