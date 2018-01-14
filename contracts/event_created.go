package contracts

import "time"

/*
 Shared Library that contains struct definitions for all possible events.
 Contains (un)serialization instructions for JSON
*/

type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"location_id"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}

// EventName returns the topic name
func (e *EventCreatedEvent) EventName() string {
	return "event.created"
}