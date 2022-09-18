package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func getSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	secret := os.Getenv("JWT_SECRET_KEY")
	return secret
}

var JWT_SECRET = []byte(getSecretKey())
