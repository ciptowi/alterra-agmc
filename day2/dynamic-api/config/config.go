package config

import (
	"dynamic/models"
	"fmt"
	"github.com/SbstnErhrdt/env"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() (client *gorm.DB, err error) {
	env.LoadEnvFiles(".env.local")

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	db := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, db, pass)
	ssl := os.Getenv("SQL_SSL")
	if len(ssl) > 0 {
		// extend dsn
		dsn = fmt.Sprintf("%s sslmode=%s", dsn, ssl)
	}
	// Connect to db
	client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		//log.Error("Postgres Client: error:", err)
		panic(err)
		return nil, err
	}
	log.Info("Postgres Client: connected")

	InitMigrate(client)
	return
}

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
	return
}
