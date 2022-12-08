package conf

import (
	"github.com/joho/godotenv"
)

const USERNAME_ENV = "username"
const PASSWORD_ENV = "password"

func LoadEnv() {
	godotenv.Load(".env")
}
