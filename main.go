package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting app on port: %s", portNumber))

	http.ListenAndServe(":8080", nil)
}
