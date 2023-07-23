package main

import (
	"fmt"
	"net/http"

	"github.com/thanhquy1105/Go-BasicBackend/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf(fmt.Sprintf("Starting app on port: %s", portNumber))
	http.ListenAndServe(":8080", nil)
}
