package main

import (
	"goges/conf"
	"goges/kordis"
)

func main() {
	err := conf.LoadEnv()
	if err != nil {
		panic(err)
	}

	credentials := kordis.GetMygesCredentials()
	err = credentials.Connect()
	if err != nil {
		panic(err)
	}
}
