package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/DavidBuzatu-Marian/go_mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func Schedule(repeatInterval time.Duration) {
	for {
		client := go_mongo.ConnectToMongo(config.MongoURI)
		people := go_mongo.CollectBirthdays(client)
		for _, val := range people {
			fmt.Println(val)
		}
		message := createNotificationMessageUsingPeopleInfo(people)
		SendNotification("Birthdays today", message)
		<-time.After(repeatInterval * time.Second)
	}
}

func createNotificationMessageUsingPeopleInfo(people []bson.D) string {
	var message strings.Builder
	for _, person := range people {
		personMap := person.Map()
		fmt.Fprintf(&message, "Name: %s; Email: %s; PhoneNumber: %s", personMap["name"], personMap["email"], personMap["phoneNumber"])
	}
	return message.String()
}
