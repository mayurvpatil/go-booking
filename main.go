package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("starting server on port number " + portNumber)
	http.ListenAndServe(":8080", nil)
}
