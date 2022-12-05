package main

import (
	"goges/conf"
	"goges/google_api"
	"os"
	"strconv"
	"time"
)

func main() {

	daysToSync, err := strconv.Atoi(os.Getenv(conf.PlanningDaysSyncEnv))
	if err != nil {
		panic(err)
	}

	if daysToSync <= 0 {
		panic("PLANNING_DAYS_SYNC must be a positive integer")
	}

	googleCalendarClient, err := google_api.CalendarClientService()
	if err != nil {
		panic(err)
	}
	events, err := googleCalendarClient.GetEvents(time.Now(), time.Now().AddDate(0, 0, daysToSync))
	if err != nil {
		panic(err)
	}
	google_api.PrintEvents(events)
}
