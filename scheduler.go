package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type Person []struct {
	ID          string    `json:"_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Country     string    `json:"country"`
	Birthday    time.Time `json:"birthday"`
}

func Schedule(repeatInterval time.Duration) {
	for {
		people := new(Person)
		err := getDataFromURL("http://localhost:8080/api/person/info/birthdays", people)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(people)
		// message := createNotificationMessageUsingPeopleInfo(people)
		// SendNotification("Birthdays today", message)
		<-time.After(repeatInterval * time.Second)
	}
}

func getDataFromURL(url string, target interface{}) error {
	response, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}

func createNotificationMessageUsingPeopleInfo(people []bson.D) string {
	var message strings.Builder
	for _, person := range people {
		personMap := person.Map()
		fmt.Fprintf(&message, "Name: %s; Email: %s; PhoneNumber: %s", personMap["name"], personMap["email"], personMap["phoneNumber"])
	}
	return message.String()
}
