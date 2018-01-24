package helper

import (
	"log"
	"os"
	"path/filepath"
	"encoding/json"
)

import "../model"



func ReadConfig(conf *model.Config, configFileName string) {
	// configFileName := "config.json"
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