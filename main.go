package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket int = 50
	fmt.Printf("%s\n", conferenceName)
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n")

	var userName string
	var userTickets int
	var age int

	// ask user about his name
	fmt.Printf("Your name: ")
	fmt.Scan(&userName)
	fmt.Printf("Your age: ")
	fmt.Scan(&age)
	fmt.Printf("How many ticket do you want to book?\n")
	fmt.Scan(&userTickets)

	var availableTickets int = remainingTickets - userTickets
	fmt.Printf("user %s has %d tickets\n", userName, userTickets)
	fmt.Printf("The numbers of available tickets now is %d\n", availableTickets)
}
