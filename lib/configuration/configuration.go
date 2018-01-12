package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/itsHabib/cloud-native-go/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8081"
)

type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Config file not found, continuing with default")
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
