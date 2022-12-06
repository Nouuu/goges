package calendar_sync

func GetEventsToRemove(kordisEvents []*Event, googleEvents []*Event) []*Event {
	var result []*Event
	for _, googleEvent := range googleEvents {
		if !Contains(kordisEvents, googleEvent) {
			result = append(result, googleEvent)
		}
	}
	return result
}

func GetEventsToAdd(kordisEvents []*Event, googleEvents []*Event) []*Event {
	var result []*Event
	for _, kordisEvent := range kordisEvents {
		if !Contains(googleEvents, kordisEvent) {
			result = append(result, kordisEvent)
		}
	}
	return result
}

func Contains(events []*Event, event *Event) bool {
	for _, e := range events {
		if Equals(e, event) {
			return true
		}
	}
	return false
}
