package main

import (
	"flag"
	"fmt"
	"log"

	msgqueue_amqp "github.com/itsHabib/cloud-nativ-go/lib/msgqueue/amqp"
	"github.com/itsHabib/cloud-native-go/eventsservice/rest"
	"github.com/itsHabib/cloud-native-go/lib/configuration"
	"github.com/itsHabib/cloud-native-go/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`,
		"flag to the path to teh configuration json file")
	flag.Parse()
	// extract configuration
	config, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	log.Println("Database connection successful... ")
	// RESTful API start
	httpErrChan, httpTLSChan := rest.ServeAPI(config.RestfulEndpoint,
		config.RestfulTLSEndpoint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httpTLSChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
