package main

import "fmt"

func main() {
	// Go 3.1 Write a simple Go program that prints the formatted string containing the type of the variable and the value of the variable for different types of variables
	var (
		text       string  = "GoLang"
		num        int     = 3
		rating     float64 = 7.256
		isEligible bool    = true
	)

	fmt.Printf("%T: %s\n", text, text)
	fmt.Printf("%T: %d\n", num, num)
	fmt.Printf("%T: %.2f\n", rating, rating)
	fmt.Printf("%T: %t\n", isEligible, isEligible)

	/* Go 3.2. Declare a variable of type integer with an initial value of 0 and print its value.
	Then assign a value to it and print the current value. Lastly, reassign a different value to it and print the final value. */
	var score int
	fmt.Println(score)

	score = 3
	fmt.Println(score)

	score = 7
	fmt.Println(score)

	/* Go 4 Write a Go program that calculates the speed of a car given the car's distance traveled and the time taken.
	The function should accept 2 parameters - distance and time, and return the speed as a float64.
	The resulting speed should be formatted to 2 decimal places.
	The equation for calculating the speed - speed = distance / time

	Use the following inputs to use in your program as inputs to your function
	1. Distance - 165 km, Time - 3.5 hours
	2. Distance - 60 km, Time - 2 hours */

	fmt.Printf("%.2f\n", getCarSpeed(165, 3.5))
	fmt.Printf("%.2f\n", getCarSpeed(60, 2))
}

func getCarSpeed(distance, time float64) float64 {
	var speed float64
	speed = distance / time
	return speed
}
