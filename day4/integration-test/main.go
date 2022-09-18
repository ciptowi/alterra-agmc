package main

import (
	"fmt"
	"integration-test/config"
	"integration-test/models"
	"integration-test/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//connect to postgresql
	e := routers.New()
	db, err := config.Setup()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("DB Connected")
	db.AutoMigrate(models.User{})
	fmt.Println("DB Migrated")

	p := godotenv.Load()
	if p != nil {
		panic(err.Error())
	}
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
