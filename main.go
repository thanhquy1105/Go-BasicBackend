package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return

	}

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting app on port: %s", portNumber))

	http.ListenAndServe(":8080", nil)
}
