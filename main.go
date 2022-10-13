package main

import (
	"goges/conf"
	"goges/kordis"
)

func main() {
	conf.LoadEnv()

	credentials := kordis.GetMygesCredentials()

	kordis.Connect(&credentials)
}
