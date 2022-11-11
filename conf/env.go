package conf

import (
	"github.com/joho/godotenv"
)

const USERNAME_ENV = "username"
const PASSWORD_ENV = "password"

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
