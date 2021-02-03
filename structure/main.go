package main

import (
	"fmt"
	"github.com/romzasandysman/magazine"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.Rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
	applyDiscount(subscriber2)
	printInfo(subscriber2)

	empoyee := magazine.Employee{Name: "Joy Carr", Salary: 60000}
	fmt.Println(empoyee.Name)
	fmt.Println(empoyee.Salary)

	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	fmt.Println(address)

	subscriber1.Address = address
	fmt.Println(subscriber1)

	empoyee.Street = "456 Elm St"
	empoyee.City = "Portland"
	empoyee.State = "OR"
	empoyee.PostalCode = "97222"

	fmt.Println(empoyee.State)
}
