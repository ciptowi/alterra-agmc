package main

import (
	"auth/config"
	"auth/models"
	"auth/routers"
	"fmt"
	"log"
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

	e.Logger.Fatal(e.Start(":5000"))
}
