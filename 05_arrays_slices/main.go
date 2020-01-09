package main

import "fmt"

func main() {
	//Arrays
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "orange"

	// declare and assign arrays
	namesarray := [2]string{"kelvin", "onkundi"}

	//Array slices
	places := []string{"apple", "Orange", "grape"}

	fmt.Println(fruitArray, namesarray, places)
}
