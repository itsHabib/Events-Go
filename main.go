package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

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
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	handler := eventServiceHandler{}

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

	http.ListenAndServe(":8181", r)

}
