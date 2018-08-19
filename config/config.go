package config

import (
	"encoding/json"
	"log"
	"os"
)

type cfg struct {
	Service         string `json:"service_name"`
	Host            string `json:"host_ip"`
	MongoIP         string `json:"mongo_ip"`
	MongoDBName     string `json:"mongo_db_name"`
	MongoCollection string `json:"mongo_db_collection"`
}

// C main config
var C = newC()

func newC() cfg {
	var c = cfg{}
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		log.Fatal("error while loading config.json")
	}
	err = json.NewDecoder(configFile).Decode(&c)
	if err != nil {
		log.Fatal("error while pasring config.json")
	}
	return c
}
