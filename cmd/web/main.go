package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zaouldyeck/simple-web-service/pkg/config"
	"github.com/zaouldyeck/simple-web-service/pkg/handlers"
	"github.com/zaouldyeck/simple-web-service/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %d\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
