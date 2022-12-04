package google_api

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
	"time"
)

func PrintEvents(events []*calendar.Event) {
	if len(events) == 0 {
		fmt.Println("No events found.")
	} else {
		for _, item := range events {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}

func (calendar *GoogleCalendar) GetEvents(start time.Time, end time.Time) (events []*calendar.Event, err error) {
	eventsResult, err := calendar.srv.Events.List(calendar.calendarId).
		TimeMin(start.Format(time.RFC3339)).
		TimeMax(end.Format(time.RFC3339)).
		ShowDeleted(false).
		SingleEvents(true).
		OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}
	return eventsResult.Items, nil
}
