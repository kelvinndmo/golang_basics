package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("server starting...")
	http.ListenAndServe(":5000", nil)
}
