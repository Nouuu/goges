package calendar_sync

import (
	"time"
)

type Event struct {
	// Id: the id of the event in the calendar
	Id string `json:"id"`
	// StartDate: The start date of the event.
	StartDate time.Time `json:"startDate"`
	// EndDate: The end date of the event.
	EndDate time.Time `json:"endDate"`
	// Title: The title of the event.
	Title string `json:"title"`
	// Teacher: The teacher of the event.
	Teacher string `json:"teacher"`
	// Location: The location of the event.
	Location string `json:"location"`
	// Rooms: The rooms of the event.
	Rooms []*Room `json:"rooms"`
	// Event Color
	Color string `json:"color"`
}

type Room struct {
	Name   string `json:"name"`
	Campus string `json:"campus"`
}

func Equals(e1, e2 *Event) bool {
	return e1.StartDate.Equal(e2.StartDate) &&
		e1.EndDate.Equal(e2.EndDate) &&
		e1.Title == e2.Title &&
		e1.Teacher == e2.Teacher &&
		e1.Location == e2.Location &&
		equalsRooms(e1.Rooms, e2.Rooms) &&
		e1.Color == e2.Color
}

func equalsRooms(rooms []*Room, rooms2 []*Room) bool {
	if len(rooms) != len(rooms2) {
		return false
	}
	for i, room := range rooms {
		if room.Name != rooms2[i].Name || room.Campus != rooms2[i].Campus {
			return false
		}
	}
	return true
}
