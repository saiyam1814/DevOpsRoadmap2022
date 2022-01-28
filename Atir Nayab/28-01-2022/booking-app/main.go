package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50	// don't allow negetive value to store
	bookings := []string{}


	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here")

	fmt.Printf("We have total of %v and %v are left\n", conferenceTickets, remainingTickets)


	var userName string
	var userTickets uint

	fmt.Println("Enter your name")
	fmt.Scan(&userName) // scan function can assing user's value to userName variable, because it has a pointer to its memory address

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	bookings = append(bookings, userName)

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}