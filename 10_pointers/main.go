package main

import "fmt"

func main() {
	//pointer allows you to point to a memory adress of a value
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read var from adress
	fmt.Println(b)

	//change val with pointer
	*b = 20
	fmt.Println(a)
}
