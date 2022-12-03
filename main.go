package main

import (
	"github.com/nouuu/goges/conf"
	"github.com/nouuu/goges/kordis"
)

func main() {
	conf.LoadEnv()

	credentials := kordis.GetMygesCredentials()

	kordis.Connect(&credentials)
}
