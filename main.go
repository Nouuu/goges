package main

import (
	"fmt"
	"goges/conf"
)

func main() {
	conf.LoadEnv()

	credentials := conf.GetMygesCredentials()

	fmt.Printf("username = %s, password = %s", credentials.Username, credentials.Password)
}
