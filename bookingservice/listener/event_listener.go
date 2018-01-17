package listener

import (
	"log"

	"github.com/itsHabib/cloud-native-go/contracts"
	"gopkg.in/mgo.v2/bson"

	"github.com/itsHabib/cloud-native-go/lib/msgqueue"
	"github.com/itsHabib/cloud-native-go/lib/persistence"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

// ProocessEvents listens in for event.created events
func (p *EventProcessor) ProcessEvents() error {
	log.Println("Listening to events...")

	received, errors, err := p.EventListener.Listen("event.created")
	if err != nil {
		return err
	}

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("recieved error while processing message: %s",
				err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s", e.ID, e)
		p.Database.AddEvent(persistence.Event{ID: bson.ObjectId(e.ID)})
	case *contracts.LocationCreatedEvent:
		log.Printf("location %s created: %s", e.ID, e)
		p.Database.AddLocation(persistence.Location{ID: bson.ObjectId(e.ID)})
	default:
		log.Printf("unknown event: %t", e)
	}
}
