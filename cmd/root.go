package cmd

import (
	"github.com/nouuu/goges/calendar_sync"
	"github.com/nouuu/goges/conf"
	"github.com/nouuu/goges/google_api"
	"github.com/nouuu/goges/kordis"
	"github.com/nouuu/goges/scheduler"
	"log"
)

func Root() {
	log.Println("Starting goges...")
	log.Println("Loading environment variables...")
	config, err := conf.LoadEnv()
	if err != nil {
		log.Fatalf("error loading env: %v", err)
	}
	log.Println("Environment variables loaded")

	launchCmd(config)
}

func launchCmd(config *conf.Config) {
	switch config.Mode {
	case "scheduler":
		err := scheduler.Main(config)
		if err != nil {
			log.Fatalf("error in scheduler: %v", err)
		}
	case "sync":
		err := calendar_sync.Sync(config.PlanningDaysSync, google_api.GetCalendarService(config), kordis.GetKordisApi(config))
		if err != nil {
			log.Fatalf("error syncing: %v", err)
		}

	default:
		log.Fatalf("unknown mode: %s", config.Mode)
	}

}
