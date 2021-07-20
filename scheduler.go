package main

import (
	"fmt"
	"time"
)

func Schedule(repeatInterval time.Duration) {
	fmt.Println("Schedule")
	for {
		// SendNotification("Test", "Test message from GoLang")
		ConnectToMongo()
		CollectBirthdays()
		<-time.After(repeatInterval * time.Second)
	}
}
