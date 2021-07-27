package main

import (
	"fmt"
	"time"

	"github.com/DavidBuzatu-Marian/GoLang-Mongo-Service/collector.go"
)

func Schedule(repeatInterval time.Duration) {
	for {
		// SendNotification("Test", "Test message from GoLang")
		people := collector.CollectBirthdays()
		for _, val := range people {
			fmt.Println(val)
		}
		<-time.After(repeatInterval * time.Second)
	}
}
