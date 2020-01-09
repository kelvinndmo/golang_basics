package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	// add ides together
	sum := 0
	for _, id := range ids {
		sum += id

	}
	fmt.Println("sum", sum)

	// Range with map
	emails := map[string]string{"name": "ndemokelvinonkundi@gmail.com", "veronica": "vero@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s", k, v)
	}
}
