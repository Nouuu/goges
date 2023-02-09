package calendar_sync

func GetEventsToRemove(kordisEvents []*Event, googleEvents []*Event) []*Event {
	var result []*Event
	for i := range googleEvents {
		if !Contains(kordisEvents, googleEvents[i]) {
			result = append(result, googleEvents[i])
		}
	}
	return result
}

func GetEventsToAdd(kordisEvents []*Event, googleEvents []*Event) []*Event {
	var result []*Event
	for i := range kordisEvents {
		if !Contains(googleEvents, kordisEvents[i]) {
			result = append(result, kordisEvents[i])
		}
	}
	return result
}

func Contains(events []*Event, event *Event) bool {
	for i := range events {
		if Equals(events[i], event) {
			return true
		}
	}
	return false
}
