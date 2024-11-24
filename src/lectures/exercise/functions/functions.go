//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello", name)
}

func message() string {
	return "Hello"
}

func sum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func number() int {
	return 1
}

func numbers() (int, int) {
	return 1, 2
}

func main() {

	greeting("Fatomi")
	fmt.Println(message())
	fmt.Println(sum(1, 2, 3))
	fmt.Println(number())
	a, b := numbers()
	fmt.Println(a, b)

}
