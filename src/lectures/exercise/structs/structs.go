//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func (myRectangle Rectangle) area() int {
	return myRectangle.length * myRectangle.width
}

func (myRectangle Rectangle) perimeter() int {
	return (myRectangle.length + myRectangle.width) * 2
}

func (myRectangle *Rectangle) doubleSize() {
	myRectangle.length *= 2
	myRectangle.width *= 2
}

func main() {
	myRectangle := Rectangle{10, 5}

	fmt.Println("Area:", myRectangle.area())
	fmt.Println("Perimeter:", myRectangle.perimeter())

	myRectangle.doubleSize()

	fmt.Println("Area:", myRectangle.area())
	fmt.Println("Perimeter:", myRectangle.perimeter())
}
