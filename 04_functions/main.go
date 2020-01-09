package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(greeting("novak"))
	fmt.Println(getSum(3, 5))
}

func getSum(num int, num2 int) int {
	return num + num2

}
