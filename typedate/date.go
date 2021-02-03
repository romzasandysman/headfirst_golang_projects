package main

import (
	"fmt"
	"log"
	"github.com/romzasandysman/calendar"
)


func main() {
	event := calendar.Event{}
	err := event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(1)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}