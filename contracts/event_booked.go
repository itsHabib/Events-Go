package contracts

// EventBookedEvent is emitted whenever an event is booked
type EventBookedEvent struct {
	EventID string `json:"eventId"`
	UserID string `json:"userId"`
}

// EventName returns the topic name
func (c *EventBookedEvent) EventName() string {
	return "eventBooked"
}