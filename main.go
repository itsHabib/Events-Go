package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/itsHabib/cloud-native-go/eventsservice/rest"
	"github.com/itsHabib/cloud-native-go/lib/configuration"
	"github.com/itsHabib/cloud-native-go/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`,
		"flag to the path to teh configuration json file")
	flag.Parse()
	// extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	// RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
