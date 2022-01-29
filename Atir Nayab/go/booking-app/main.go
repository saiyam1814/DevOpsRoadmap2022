package main

import (
	"fmt"
	"sync"
	"time"

	// "strconv"
	"booking-app/helper"
)

var wg = sync.WaitGroup{}

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50	// don't allow negetive value to store
	// bookings := []string{}
	// var userData =  make([]UserData,0) // empty map

		type UserData struct {
		firstName string
		// lastName string
		// email string
		// numberOfTickets uint
	}




	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here")

	fmt.Printf("We have total of %v and %v are left\n", conferenceTickets, remainingTickets)


	var userName string
	var userTickets uint

	helper.HelperFunction()
	helperFunction1()


	// userData["firstName"] = userName
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)


	// bookings = append(bookings, userName)

	// for{
		fmt.Println("Enter your name")
		fmt.Scan(&userName) // scan function can assing user's value to userName variable, because it has a pointer to its memory address
	
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		var bookings = UserData {
			firstName: userName,
		}	
		fmt.Println(bookings)
		wg.Add(1)
		go sendTicket()
		fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
		wg.Wait()
	// }

}

func sendTicket (){
	time.Sleep(10 * time.Second)
	fmt.Println("simulation ticket sending to email")
	wg.Done()
}