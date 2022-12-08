package calendar_sync

import (
	"github.com/golang-module/carbon/v2"
	"github.com/nouuu/goges/google_api"
	"github.com/nouuu/goges/kordis"
	"google.golang.org/api/calendar/v3"
	"log"
)

func Sync(days int, googleCalendarClient *google_api.GoogleCalendar, mygesClient *kordis.MygesApi) error {
	if days <= 0 {
		return nil
	}
	c := carbon.NewCarbon()
	log.Printf("Syncing %d days\n", days)

	log.Printf("Retrieving events from kordis...\n")
	kordisEvents, err := mygesClient.GetAgendaFromNow(days)
	if err != nil {
		return err
	}
	log.Printf("Retrieved %d events from kordis\n", len(kordisEvents.Result))

	log.Printf("Retrieving events from google calendar...\n")
	googleEvents, err := googleCalendarClient.GetEventsFromNow(days)
	if err != nil {
		return err
	}
	log.Printf("Retrieved %d events from google calendar\n", len(googleEvents))

	kordisEventsPointer := make([]*kordis.AgendaEvent, len(kordisEvents.Result))
	for i, event := range kordisEvents.Result {
		kordisEventsPointer[i] = &event
	}

	convertedKordisEvents := FromKordisEvents(kordisEventsPointer, &c)

	googleEventsPointer := make([]*calendar.Event, len(googleEvents))
	for i, event := range googleEvents {
		googleEventsPointer[i] = event
	}

	convertedGoogleEvents := FromGoogleCalendarEvents(googleEventsPointer)

	eventsToRemove := GetEventsToRemove(convertedKordisEvents, convertedGoogleEvents)
	log.Printf("Found %d events to remove\n", len(eventsToRemove))
	eventsToAdd := GetEventsToAdd(convertedKordisEvents, convertedGoogleEvents)
	log.Printf("Found %d events to add\n", len(eventsToAdd))

	log.Printf("Removing events...\n")
	err = googleCalendarClient.RemoveEvents(ToGoogleCalendarEvents(eventsToRemove))
	if err != nil {
		return err
	}

	log.Printf("Adding events...\n")
	err = googleCalendarClient.AddEvents(ToGoogleCalendarEvents(eventsToAdd))
	if err != nil {
		return err
	}

	return nil
}
