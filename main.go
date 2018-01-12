package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO separate handlers  into its own packages

type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}

type eventServiceHandler struct{}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter,
	r *http.Request) {
}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter,
	r *http.Request) {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter,
	r *http.Request) {

}

func main() {
	log.Fatal(ServeAPI(":8081"))
}

func ServeAPI(endpoint string) error {
	handler := &eventServiceHandler{}
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()

	// Searching for events via id or name
	// events/id/123 or /events/name/rock
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.findEventHandler)

	// Retrieving all events at once
	// /events
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	// Creating a new event
	// /events
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)

}
