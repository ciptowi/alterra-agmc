package config

import (
	"fmt"
	"integration-test/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
	fmt.Println("RUN db connection") // log test
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

func TestSetup() (*gorm.DB, error) {
	fmt.Println("RUN db connection for testing") // log test
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	User := os.Getenv("TEST_SQL_USER")
	Pass := os.Getenv("TEST_SQL_PASSWORD")
	Host := os.Getenv("TEST_SQL_HOST")
	Port := os.Getenv("TEST_SQL_PORT")
	DB := os.Getenv("TEST_SQL_DATABASE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host,
		Port,
		User,
		DB,
		Pass,
	)

	db_test, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	TestInitMigrate(db_test)
	return db_test, nil
}

func InitMigrate(db *gorm.DB) {
	fmt.Println("RUN db migration for testing") // log test
	db.AutoMigrate(&models.TestUser{}, &models.TestBook{})
	return
}

func TestInitMigrate(db_test *gorm.DB) {
	fmt.Println("RUN db migration") // log test
	db_test.AutoMigrate(&models.User{}, &models.Book{})
	return
}
