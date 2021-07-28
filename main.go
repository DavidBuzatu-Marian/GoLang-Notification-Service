package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	MongoURI string
}

var DAY time.Duration = 86400
var config Config

func ReadConfig() {
	file, _ := os.Open("./config/default.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config = Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error reading config file!")
		os.Exit(1)
	}
}

func main() {
	ReadConfig()
	Schedule(5)
}
