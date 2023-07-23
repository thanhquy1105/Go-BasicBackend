package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thanhquy1105/Go-BasicBackend/pkg/config"
	"github.com/thanhquy1105/Go-BasicBackend/pkg/handlers"
	"github.com/thanhquy1105/Go-BasicBackend/pkg/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf(fmt.Sprintf("Starting app on port: %s", portNumber))
	http.ListenAndServe(":8080", nil)
}
