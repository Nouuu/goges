package calendar_sync

import (
	"github.com/golang-module/carbon/v2"
	"github.com/nouuu/goges/kordis"
	"google.golang.org/api/calendar/v3"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func FromGoogleCalendarEvent(event *calendar.Event) *Event {
	startDate, err := time.Parse(time.RFC3339, event.Start.DateTime)
	if err != nil {
		log.Fatalf("Error parsing start date: %v", err)
	}
	endDate, err := time.Parse(time.RFC3339, event.End.DateTime)
	if err != nil {
		log.Fatalf("Error parsing end date: %v", err)
	}

	result := &Event{
		Id:        event.Id,
		StartDate: startDate,
		EndDate:   endDate,
		Title:     event.Summary,
		Location:  event.Location,
		Color:     event.ColorId,
	}

	teacher, rooms := extractTeacherAndRoom(event.Description)

	result.Teacher = teacher
	result.Rooms = make([]*Room, len(rooms))
	for i, room := range rooms {
		split := strings.Split(room, " - ")
		if len(split) == 2 {
			result.Rooms[i] = &Room{
				Name:   split[1],
				Campus: split[0],
			}
		}
	}

	return result
}

func FromGoogleCalendarEvents(events []*calendar.Event) []*Event {
	result := make([]*Event, len(events))
	for i := range events {
		result[i] = FromGoogleCalendarEvent(events[i])
	}
	return result
}

func ToGoogleCalendarEvent(event *Event) *calendar.Event {
	result := &calendar.Event{
		Id: event.Id,
		Start: &calendar.EventDateTime{
			DateTime: event.StartDate.Format(time.RFC3339),
			TimeZone: event.StartDate.Location().String(),
		},
		End: &calendar.EventDateTime{
			DateTime: event.EndDate.Format(time.RFC3339),
			TimeZone: event.EndDate.Location().String(),
		},
		Summary:  event.Title,
		Location: event.Location,
		ColorId:  event.Color,
	}

	if len(event.Teacher) > 0 {
		result.Description = "<span>Intervenant : " + event.Teacher + "</span>"
	}
	if len(event.Rooms) > 0 {
		result.Description = strings.Join([]string{result.Description, "<span>Salle(s) :<ul>"}, "<br>")
		for i := range event.Rooms {
			result.Description += "<li>" + event.Rooms[i].Campus + " - " + event.Rooms[i].Name + "</li>"
		}
		result.Description += "</ul></span>"
	}
	return result
}

func ToGoogleCalendarEvents(events []*Event) []*calendar.Event {
	result := make([]*calendar.Event, len(events))
	for i := range events {
		result[i] = ToGoogleCalendarEvent(events[i])
	}
	return result
}

func FromKordisEvent(event *kordis.AgendaEvent, c *carbon.Carbon) *Event {
	result := &Event{
		Title: event.Name,
		StartDate: c.
			CreateFromTimestamp(event.StartDate / 1000).
			SetTimezone(os.Getenv("TZ")).
			Carbon2Time(),
		EndDate: c.
			CreateFromTimestamp(event.EndDate / 1000).
			SetTimezone(os.Getenv("TZ")).
			Carbon2Time(),
		Teacher: strings.TrimSpace(event.Teacher),
	}
	if event.Rooms != nil && len(event.Rooms) > 0 {
		// Getting campus location from first room
		campus, err := GetCampus(event.Rooms[0].Campus)
		if err == nil && len(campus) > 0 {
			// If campus is found, add it to result
			result.Location = campus[0]
			result.Color = campus[1]
		} else {
			result.Color = "11"
		}

		result.Rooms = make([]*Room, len(event.Rooms))
		for i := range event.Rooms {
			result.Rooms[i] = &Room{
				Name:   event.Rooms[i].Name,
				Campus: event.Rooms[i].Campus,
			}
		}
	} else {
		result.Color = "11"
	}

	return result
}

func FromKordisEvents(events []*kordis.AgendaEvent, c *carbon.Carbon) []*Event {
	result := make([]*Event, len(events))
	for i := range events {
		result[i] = FromKordisEvent(events[i], c)
	}
	return result
}

func extractTeacherAndRoom(html string) (string, []string) {
	intervenantRegexp := regexp.MustCompile(`<span>Intervenant : (.*?)</span>`)
	salleRegexp := regexp.MustCompile(`<li>(.*?)</li>`)

	intervenantMatch := intervenantRegexp.FindStringSubmatch(html)
	var intervenant string
	if len(intervenantMatch) > 1 {
		intervenant = strings.TrimSpace(intervenantMatch[1])
	}

	sallesMatch := salleRegexp.FindAllStringSubmatch(html, -1)
	salles := make([]string, len(sallesMatch))
	for i, salleMatch := range sallesMatch {
		if len(salleMatch) > 1 {
			salle := salleMatch[1]
			salles[i] = salle
		}
	}

	return intervenant, salles
}
