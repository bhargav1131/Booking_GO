// first project while learning the Go language
package main

import (
	"fmt"
)

func main() {
	conferenceName := "My Conference"
	const conferenceTickets = 30
	var remainingTickets = 30
	// var bookings [50]string //array
	var bookings []string //slices are dynamic

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total available tickets is %v  and remaining tickets are %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference..")

	var firstName string
	var lastName string
	var userTickets int

	fmt.Println()
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter the number of tickets you want to book ")
	fmt.Scan(&userTickets)

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	if userTickets > conferenceTickets || userTickets > remainingTickets {
		fmt.Println("Invalid input")
		return

	} else {
		remainingTickets = remainingTickets - userTickets
	}

	fmt.Printf("Congrats %v %v! you have successfully booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("Available tickets: %v\n", remainingTickets)

	fmt.Println(bookings[0])

}
