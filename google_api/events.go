package google_api

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"google.golang.org/api/calendar/v3"
	"math"
	"time"
)

func PrintEvents(events []*calendar.Event) {
	if len(events) == 0 {
		fmt.Println("No events found.")
	} else {
		for i := range events {
			date := events[i].Start.DateTime
			if date == "" {
				date = events[i].Start.Date
			}
			fmt.Printf("%v (%v)\n", events[i].Summary, date)
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

func (calendar *GoogleCalendar) GetEventsFromNow(days int) (events []*calendar.Event, err error) {
	return calendar.GetEvents(
		carbon.Now().StartOfDay().Carbon2Time(),
		carbon.Now().AddDays(int(math.Max(float64(days), 0))).EndOfDay().Carbon2Time(),
	)
}

func (calendar *GoogleCalendar) RemoveEvent(eventId string) (err error) {
	return calendar.srv.Events.Delete(calendar.calendarId, eventId).Do()
}

func (calendar *GoogleCalendar) RemoveEvents(events []*calendar.Event) (err error) {
	for i := range events {
		err = calendar.RemoveEvent(events[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (calendar *GoogleCalendar) AddEvent(event *calendar.Event) (err error) {
	_, err = calendar.srv.Events.Insert(calendar.calendarId, event).Do()
	return err
}

func (calendar *GoogleCalendar) AddEvents(events []*calendar.Event) (err error) {
	for _, event := range events {
		err = calendar.AddEvent(event)
		if err != nil {
			return err
		}
	}
	return nil
}
