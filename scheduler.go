package main

import (
	"fmt"
	"time"
)

func Schedule(repeatInterval time.Duration) {
	for {
		// SendNotification("Test", "Test message from GoLang")
		ConnectToMongo()
		people := CollectBirthdays()
		for _, val := range people {
			fmt.Println(val)
		}
		<-time.After(repeatInterval * time.Second)
	}
}
