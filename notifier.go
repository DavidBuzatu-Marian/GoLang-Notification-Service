package main

import (
	"github.com/gen2brain/beeep"
)

func SendNotification(title string, message string) {
	err := beeep.Notify(title, message, "/assets/info-circle-solid.svg")
	if err != nil {
		panic(err)
	}
}
