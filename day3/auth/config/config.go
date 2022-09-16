package config

import (
	"auth/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	User := os.Getenv("SQL_USER")
	Pass := os.Getenv("SQL_PASSWORD")
	Host := os.Getenv("SQL_HOST")
	Port := os.Getenv("SQL_PORT")
	DB := os.Getenv("SQL_DATABASE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host,
		Port,
		User,
		DB,
		Pass,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	InitMigrate(db)
	return db, nil
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Book{})
	return
}
