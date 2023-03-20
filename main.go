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

	var firstName string
	var secondName string
	var email string
	var userTickets int
	var age int

	// ask user about his name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your second name: ")
	fmt.Scan(&secondName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)
	fmt.Printf("How many ticket do you want to book?\n")
	fmt.Scan(&userTickets)


	//var availableTickets int = remainingTickets - userTickets
	//fmt.Printf("user %s %s has %d tickets\n", firstName, secondName, userTickets)
	//fmt.Printf("The numbers of available tickets now is %d\n", availableTickets)
	fmt.Printf("Thank you %s %s for booking %d tickets. You'll receive confirmation email at %s \n",firstName,secondName,userTickets,email)
}
