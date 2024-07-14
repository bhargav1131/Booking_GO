// first project while learning the Go language
package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "My Conference"
	const conferenceTickets = 30
	var remainingTickets = 30
	// var bookings [50]string //array
	var bookings []string //slices are dynamic

	fmt.Printf("**Welcome to %v booking application\n", conferenceName)
	fmt.Printf("**Total available tickets is %v  and remaining tickets are %v \n", conferenceTickets, remainingTickets)
	fmt.Println("**Get your tickets here to attend our conference..")

	var firstName string
	var lastName string
	var userTickets int
	var customerID int = 1
	var customerIndex int = 0

	for remainingTickets != 0 {

		//user details
		fmt.Println()
		fmt.Printf("#ID: %v\n", customerID)
		fmt.Println("->Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("->Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("->Enter the number of tickets you want to book ")
		fmt.Scan(&userTickets)
		customerID++

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		//ticket booking logic
		if userTickets > conferenceTickets || userTickets > remainingTickets {
			fmt.Println("Invalid input")
			return

		} else {
			remainingTickets = remainingTickets - userTickets
		}

		//logic to extract only the first name
		onlyFirstName := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			onlyFirstName = append(onlyFirstName, names[0])
		}

		// welcome statements
		fmt.Printf("*Congrats %v! You have successfully booked %v tickets for the upcoming conference\n", onlyFirstName[customerIndex], userTickets)
		customerIndex++
		fmt.Printf("*Current available tickets: %v\n", remainingTickets)
		fmt.Println()
		fmt.Println()
	}
}
