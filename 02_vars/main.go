package main

import "fmt"

var name = "novak"

func main() {

	name := "Kelvin"
	var age int = 37
	const isCool = true
	size := 1.34543334433

	name, email := "kelvin onkundi", "ndemo@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
}
