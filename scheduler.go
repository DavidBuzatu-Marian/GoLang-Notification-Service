package main

import (
	"fmt"
	"time"
)

func Schedule(repeatInterval time.Duration) {
	fmt.Println("Schedule")
	for {
		SendNotification("Test", "Test message from GoLang")
		<-time.After(repeatInterval * time.Second)
	}
}
