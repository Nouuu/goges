package cmd

import (
	"goges/internal/conf"
	"goges/internal/kordis"
)

func Go() {
	conf.LoadEnv()

	credentials := kordis.GetMygesCredentials()

	kordis.Connect(&credentials)
}
