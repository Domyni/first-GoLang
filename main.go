package main

import "fmt"

func main() {

	//	Given the following map
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuff": {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
	}

	/*	Write a Go program that declares 4 blank slices for each of the keys in the map.
		Go over the map and append/insert the values of each key as a separate element into their respective slices.
		Once done, print the length of each of the slices. */
	sliceLen(houses)

	/* Write a Go program that loops from 0-15.
	For each iteration, multiply the current iteration value by 2 and append the result to a slice.
	Print the last 4 elements of the slice once done.
	Use slice range (:) to print the elements. */
	lastFour()
}

func sliceLen(input map[string][]string) {

	for _, v := range input {
		var arr []interface{}
		for _, elem := range v {
			arr = append(arr, elem)
		}
		fmt.Println(len(arr))
	}
}

func lastFour() {

	var arr []interface{}
	for i := 0; i <= 15; i++ {
		arr = append(arr, i*2)
	}
	fmt.Println(arr[len(arr)-4:])
}
