package main

import (
	"github.com/nouuu/goges/conf"
	"github.com/nouuu/goges/scheduler"
)

func main() {
	err := conf.LoadEnv()
	if err != nil {
		panic(err)
	}

	err = scheduler.Main()
	if err != nil {
		panic(err)
	}
}
