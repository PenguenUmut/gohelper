package main

import (
	"log"
	"os"
	"path/filepath"
	"encoding/json"
)

import "github.com/PenguenUmut/gohelper"



var conf gohelper.Config

func main() {
	readConfig(&conf)

	log.Printf("config.json Name: ", conf.Name)
	log.Printf("config.json Version: ", conf.Version)
	log.Printf("config.json MyObject.Enabled: ", conf.MyObject.Enabled)
	log.Printf("config.json MyObject.ID: ", conf.MyObject.ID)
}


func readConfig(conf *gohelper.Config) {
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&conf); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}