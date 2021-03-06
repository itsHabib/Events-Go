package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/itsHabib/cloud-native-go/contracts"

	"github.com/itsHabib/cloud-native-go/lib/msgqueue"

	"github.com/gorilla/mux"
	"github.com/itsHabib/cloud-native-go/lib/persistence"
)

type eventServiceHandler struct {
	dbhandler    persistence.DatabaseHandler
	eventEmitter msgqueue.EventEmitter
}

func newEventHandler(dbhandler persistence.DatabaseHandler,
	eventEmitter msgqueue.EventEmitter) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler:    dbhandler,
		eventEmitter: eventEmitter,
	}
}

// Finds an event by id from db
func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter,
	r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either `+
			`search by id via id/4 or search by name via /name/drakeconcert`)
		return
	}
	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either `+
			`search by id via id/4 or search by name via /name/drakeconcert`)
	}
	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprint(w, "{error: %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset-utf8")
	json.NewEncoder(w).Encode(&event)
}

// Returns all events from db
func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter,
	r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to find all "+
			"available events %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset-utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode events "+
			"to JSON %s}", err)
		return
	}
}

// Creates a new event in db
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter,
	r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while decoding event data %s}", err)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while persisting event %d %s",
			id, err)
		return
	}
	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		LocationID: string(event.Location.ID),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}
	eh.eventEmitter.Emit(&msg)
	fmt.Fprint(w, `{"id":%d}`, id)
}
