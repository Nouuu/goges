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

	credentials, err := kordis.GetMygesCredentials()
	if err != nil {
		panic(err)
	}

	agenda, err := credentials.GetAgendaFromNow(10)
	if err != nil {
		panic(err)
	}
	kordis.PrintAgenda(agenda)
}
