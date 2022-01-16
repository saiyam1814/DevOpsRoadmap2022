package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	/*
		var				=> We're about to declare a new variable
		card			=> The name of the variable
		string			=> Only a "string" will ever be assigned to this variable
							(Go is a static language)
		"Ace of Spades"	=> Assign the value "Ace of Spades" to card variable
	*/

	fmt.Println(card)
}
