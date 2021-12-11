package main

import (
	"fmt"
	"net/http"

	"github.com/mayurvpatil/go-server/pkg/handler"
)

const portNumber = ":8080"

// main function
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println("starting server on port number " + portNumber)
	http.ListenAndServe(":8080", nil)
}
