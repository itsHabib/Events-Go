package contracts

import "github.com/itsHabib/cloud-native-go/lib/persistence"

// LocationCreatedEvent is emitted whenever a location is created
type LocationCreatedEvent struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Country string             `json:"country"`
	Halls   []persistence.Hall `json:"halls"`
}

// EventName retuns the topic name
func (c *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}
