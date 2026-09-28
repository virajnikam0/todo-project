package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// config structure
type Config struct {
	Env string 
	Port string
}


func MustLoad() *Config{
	var cfg = Config{}

	file,err := os.Open("./local.json")
	if err != nil {
		fmt.Println("File Cannot be opened ")
	}
	
	if err := json.NewDecoder(file).Decode(&cfg); err!= nil{
		fmt.Println("Decoding have problem ")
	}

	return &cfg
}