package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// stucture for configuration
type Config struct {
	ENV  string `json:"env"`
	PORT string `json:"port"`
}

// get configuration function
func GetConfig() *Config {
	var config Config

	// open file and get data
	file, err := os.Open("urlshortner/config/local.json")

	if err != nil {
		fmt.Println("local.json file having error")
	}

	// decode the file data to structure
	decodeErr := json.NewDecoder(file).Decode(&config)
	if decodeErr != nil {
		fmt.Println("Error during Decoding the structural data ")
	}

	return &config
}
