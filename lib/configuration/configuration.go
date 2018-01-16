package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/itsHabib/cloud-native-go/lib/persistence/dblayer"
)

var (
	DBTypeDefault            = dblayer.DBTYPE("mongodb")
	DBConnectionDefault      = "mongodb://127.0.0.1"
	RestfulEPDefault         = "localhost:8081"
	RestfulTLSEPDefault      = "localhost:9091"
	AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
)

type ServiceConfig struct {
	Databasetype       dblayer.DBTYPE `json:"databasetype"`
	DBConnection       string         `json:"dbconnection"`
	RestfulEndpoint    string         `json:"restfulapi_endpoint"`
	RestfulTLSEndpoint string         `json:"restful_tlsendpoint`
	AMQPMessageBroker  string         `json:"amqp_message_broker"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		RestfulTLSEPDefault,
		AMQPMessageBrokerDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Config file not found, continuing with default")
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	if broker := os.Getenv("AMQP_URL"); broker != "" {
		conf.AMQPMessageBroker = broker
	}
	return conf, err
}
