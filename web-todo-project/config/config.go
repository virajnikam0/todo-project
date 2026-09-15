package config

import (
	"log"
	"os"
	"encoding/json"
)

type Config struct {
	Env  string `json:"env"`
	Port string `json:"port"`
}

func GetConfig()(*Config) {

	var cfg Config
	// taking field from json file with encoding
	file, err := os.Open("C:\\Users\\shubh\\Documents\\VN_study\\web-todo-project\\config\\local.json")
	if err != nil {
		log.Fatalf("configuration are not loaded: %v\n",err)
	}

	if err := json.NewDecoder(file).Decode(&cfg);err!=nil{
		log.Fatalf("file having error: %v \n",err)
	}

	return &cfg

}