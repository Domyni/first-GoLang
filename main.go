package main

import "fmt"

func main() {

	/* Go 5, Write a Go program that contains a function `isFizzBuzz`.
	The function should accept 1 integer parameter.
	If the number is divisible by 3, then return "fizz".
	If the number is divisible by 5, then return "buzz".
	If the number is divisible by both 3 and 5, then return "fizzbuzz".
	If the number doesn't match any of the conditions, then return "no go".
	Use the following inputs provided to test your program with - 3, 5, 15, 26.*/

	fmt.Println(isFizzBuzz(-3))
	fmt.Println(isFizzBuzz(5))
	fmt.Println(isFizzBuzz(15))
	fmt.Println(isFizzBuzz(26))

}

func isFizzBuzz(n int) string {

	if n%3 == 0 && n%5 == 0 {
		return "fizzbuzz"
	} else if n%3 == 0 {
		return "fizz"
	} else if n%5 == 0 {
		return "buzz"
	} else {
		return "no go"
	}

}
