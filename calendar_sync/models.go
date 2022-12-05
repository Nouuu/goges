package calendar_sync

import (
	"github.com/golang-module/carbon/v2"
	"goges/kordis"
	"google.golang.org/api/calendar/v3"
	"log"
	"os"
	"time"
)

type Event struct {
	// StartDate: The start date of the event.
	StartDate time.Time `json:"startDate"`
	// EndDate: The end date of the event.
	EndDate time.Time `json:"endDate"`
	// Title: The title of the event.
	Title string `json:"title"`
	// Description: The description of the event.
	Description string `json:"description"`
	// Location: The location of the event.
	Location string `json:"location"`
}

func FromGoogleCalendarEvent(event *calendar.Event) *Event {
	startDate, err := time.Parse(time.RFC3339, event.Start.DateTime)
	if err != nil {
		log.Fatalf("Error parsing start date: %v", err)
	}
	endDate, err := time.Parse(time.RFC3339, event.End.DateTime)
	if err != nil {
		log.Fatalf("Error parsing end date: %v", err)
	}

	return &Event{
		StartDate:   startDate,
		EndDate:     endDate,
		Title:       event.Summary,
		Description: event.Description,
		Location:    event.Location,
	}
}

func FromGoogleCalendarEvents(events []*calendar.Event) []*Event {
	result := make([]*Event, len(events))
	for i, event := range events {
		result[i] = FromGoogleCalendarEvent(event)
	}
	return result
}

func ToGoogleCalendarEvent(event *Event) *calendar.Event {
	return &calendar.Event{ // TODO add color in function of location
		Start: &calendar.EventDateTime{
			DateTime: event.StartDate.Format(time.RFC3339),
			TimeZone: event.StartDate.Location().String(),
		},
		End: &calendar.EventDateTime{
			DateTime: event.EndDate.Format(time.RFC3339),
			TimeZone: event.EndDate.Location().String(),
		},
		Summary:     event.Title,
		Description: event.Description,
		Location:    event.Location,
	}
}

func ToGoogleCalendarEvents(events []*Event) []*calendar.Event {
	result := make([]*calendar.Event, len(events))
	for i, event := range events {
		result[i] = ToGoogleCalendarEvent(event)
	}
	return result
}

func FromKordisEvent(event *kordis.AgendaEvent, c *carbon.Carbon) *Event {
	result := &Event{
		StartDate: c.
			CreateFromTimestamp(event.StartDate / 1000).
			SetTimezone(os.Getenv("TZ")).
			Carbon2Time(),
		EndDate: c.
			CreateFromTimestamp(event.EndDate / 1000).
			SetTimezone(os.Getenv("TZ")).
			Carbon2Time(),
		Description: "Intervenant : " + event.Teacher + "\n",
	}
	if event.Rooms != nil && len(event.Rooms) > 0 {
		result.Location = event.Rooms[0].Campus // TODO: use futur location map
		result.Description += "Salle : \n"
		for _, room := range event.Rooms {
			result.Description += "- " + room.Name + "\n"
		}
	}

	return result
}

func FromKordisEvents(events []*kordis.AgendaEvent, c *carbon.Carbon) []*Event {
	result := make([]*Event, len(events))
	for i, event := range events {
		result[i] = FromKordisEvent(event, c)
	}
	return result
}
