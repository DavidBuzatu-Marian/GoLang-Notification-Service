package main

import (
	"fmt"
	"time"

	"github.com/DavidBuzatu-Marian/go_mongo"
)

func Schedule(repeatInterval time.Duration) {
	for {
		// SendNotification("Test", "Test message from GoLang")
		go_mongo.ConnectToMongo()
		people := go_mongo.CollectBirthdays()
		for _, val := range people {
			fmt.Println(val)
		}
		<-time.After(repeatInterval * time.Second)
	}
}
