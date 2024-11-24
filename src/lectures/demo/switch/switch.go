package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch fare := price(); {
		case fare == Economy:
			fmt.Println("I boarded an Economy class flight")
		case fare == Business:
			fmt.Println("I boarded a Business class flight")
			fallthrough
		case fare == FirstClass:
			fmt.Println("I boarded a First Class flight")
		default:
			fmt.Println("I boarded a flight")
	}
}
