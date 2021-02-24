package main

import "fmt"

func main() {

	/* Go 6 #1 Write a Go program that contains a function `isFizzBuzz`.
	If the number is divisible by 3, then return "fizz".
	If the number is divisible by 5, then return "buzz".
	If the number is divisible by both 3 and 5, then return "fizzbuzz".
	If the number doesn't match any of the conditions, then return "no go".
	This time the function should loop from 1-100 and print the results accordingly. */
	isFizzBuzz()

	// Go 6 #2
	// Write a program that loops from 0-50 and prints the results of the current number in the iteration multiplied by 2.
	// If the result of the multiplication is an odd number, defer printing the resulting. Otherwise, print it immediately.
	// HINT: this is a trick question
	deferOdd()
}

func isFizzBuzz() {

	for n := 1; n <= 100; n++ {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println(n, "fizzbuzz")
		} else if n%3 == 0 {
			fmt.Println(n, "fizz")
		} else if n%5 == 0 {
			fmt.Println(n, "buzz")
		} else {
			fmt.Println(n, "no go")
		}
	}
}

func deferOdd() {

	for n := 0; n <= 50; n++ {
		var result int = n * 2
		if result%2 != 0 {
			defer fmt.Println(result, "Odd")
		} else {
			fmt.Println(result, "Even")
		}
	}
}
