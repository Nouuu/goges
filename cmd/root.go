package cmd

import (
	"github.com/nouuu/goges/calendar_sync"
	"github.com/nouuu/goges/conf"
	"github.com/nouuu/goges/google_api"
	"github.com/nouuu/goges/kordis"
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

	log.Println("Getting Google Calendar service...")
	googleCalendar, err := google_api.CalendarClientService(config)
	if err != nil {
		log.Fatalf("error getting Google Calendar service: %v", err)
	}
	log.Println("Google Calendar service loaded")

	log.Println("Getting kordis api service...")
	kordisApi, err := kordis.GetMygesApi(config)
	if err != nil {
		log.Fatalf("error getting kordis api service: %v", err)
	}
	log.Println("Kordis api service loaded")

	switch config.Mode {
	case "scheduler":
		break
	case "sync":
		err := calendar_sync.Sync(config.PlanningDaysSync, googleCalendar, kordisApi)
		if err != nil {
			log.Fatalf("error syncing: %v", err)
		}

	default:
		log.Fatalf("unknown mode: %s", config.Mode)
	}
}
