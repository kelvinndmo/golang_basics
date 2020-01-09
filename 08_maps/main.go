package main

import "fmt"

func main() {
	emails := make(map[string]string)

	//asign key values
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["novak"] = "emdl@gmail.com"
	fmt.Println(emails)
	delete(emails, "bob")
	fmt.Println(emails)

	//assign on declaration
	schools := map[string]string{"bob": "bo@gmail.com", "kelvin": "kelvin@gmail.com"}
	fmt.Println(schools)
}
