package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
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
		err := getPersonsWithTodayBirthday("http://localhost:8080/api/person/info/birthdays", people)
		if err != nil {
			log.Fatal(err)
		}
		message := createNotificationMessageUsingPeopleInfo(people)
		SendNotification("Birthdays today", message)
		<-time.After(repeatInterval * time.Second)
	}
}

func getPersonsWithTodayBirthday(url string, target interface{}) error {
	response, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&target)
}

func createNotificationMessageUsingPeopleInfo(people *Person) string {
	var message strings.Builder
	for _, person := range *people {
		fmt.Fprintf(&message, "Name: %s; Email: %s; PhoneNumber: %s", person.Name, person.Email, person.PhoneNumber)
	}
	return message.String()
}
