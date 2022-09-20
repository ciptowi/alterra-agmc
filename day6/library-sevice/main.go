package main

import (
	"github.com/joho/godotenv"
	"library-sevice/internal/factory"
	"library-sevice/internal/http"
	"os"

	"github.com/labstack/echo/v4"
	"library-sevice/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
func main() {
	database.CreateConnection()

	e := echo.New()

	f := factory.NewFactory()
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
