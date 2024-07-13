// first project while learning the Go language
package main

import (
	"fmt"
)

func main() {
	var conferenceName = "My Conference"
	const conferenceTickets = 30
	var remainingTickets = 30

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total available tickets is %v  and remaining tickets are %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference..")

	var userName string
	var userTickets int
	fmt.Println()
	fmt.Println("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Println("Enter the number of tickets you want to book ")
	fmt.Scan(&userTickets)

	if userTickets > conferenceTickets || userTickets > remainingTickets {
		fmt.Println("Invalid input")
		return

	} else {
		remainingTickets--
	}

	fmt.Printf("Congrats %v you have successfully booked %v tickets\n", userName, userTickets)

}
