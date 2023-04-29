package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configType struct {
	PyFilePath        string `json:"pyFilePath"`
	PyInterpreterName string `json:"pyInterpreterName"`
}

func loadConfig() (configType, error) {
	config := configType{}

	f, err := os.Open("config.json")
	if err != nil {
		return config, fmt.Errorf("error opening config.json: %v", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return config, fmt.Errorf("error decoding config.json: %v", err)
	}

	return config, nil
}
