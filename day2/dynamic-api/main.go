package main

import (
	"dynamic/config"
	"dynamic/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":5000"))
}
