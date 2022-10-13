package conf

import (
	"github.com/joho/godotenv"
	"os"
)

const USERNAME_ENV = "username"
const PASSWORD_ENV = "password"

type mygesCredentials struct {
	Username string
	Password string
}

func GetMygesCredentials() mygesCredentials {
	return mygesCredentials{
		Username: os.Getenv(USERNAME_ENV),
		Password: os.Getenv(PASSWORD_ENV),
	}
}

func LoadEnv() {
	godotenv.Load(".env")

}
