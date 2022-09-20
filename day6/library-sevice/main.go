package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
