package main

import (
	"fmt"
	"time"
)

var DAY time.Duration = 86400

func main() {
	ReadConfig()
	Schedule(5)
	fmt.Println("Hello from GoLang!")
}
