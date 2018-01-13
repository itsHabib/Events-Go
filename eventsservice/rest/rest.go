package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itsHabib/cloud-native-go/lib/persistence"
)

// Sets up routes and serves api at given endpoint and given db handler
func ServeAPI(endpoint, tlsEndpoint string,
	databasehandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := NewEventHandler(databasehandler)
	r := mux.NewRouter()

	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

	httpErrChan := make(chan error)
	httpTLSErrChan := make(chan error)
	go func() {
		httpTLSErrChan <- http.ListenAndServeTLS(tlsEndpoint,
			"cert.pem", "key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()

	return httpErrChan, httpTLSErrChan
}
