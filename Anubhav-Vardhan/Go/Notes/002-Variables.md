``` Go
package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	/*
		var				=> We're about to declare a new variable
		card			=> The name of the variable
		string			=> Only a "string" will ever be assigned to this variable
							(Go is a static language)
		"Ace of Spades"	=> Assign the value "Ace of Spades" to card variable
	*/

	card := "Ace of Spades" //Alternative way to declare variable

	fmt.Println(card)

	/*
		Go variable types
		bool	=> true false
		string	=> "hi" "How's it going"
		int		=> 0 -10000 99999
		float64	=> 10.000001 0.00009 -100.003

		There are many more types but these are the fundamental
	*/
}
```