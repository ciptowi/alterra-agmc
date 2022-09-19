package config

import (
	"fmt"
	"integration-test/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Karena .env tidak bisa dalam mode test, sementara database hardcode disini
Apabila sudah bisa lebih baik pakai env variable, lebih aman
*/
// const (
// 	Host = "localhost"
// 	User = "postgres"
// 	Pass = "postgres"
// 	Port = "5432"
// 	DB   = "testing"
// )

func Setup() (*gorm.DB, error) {
	// fmt.Println("RUN db connection") // log test
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
	// fmt.Println("RUN db migration for testing") // log test
	db.AutoMigrate(&models.User{}, &models.Book{})
	return
}
