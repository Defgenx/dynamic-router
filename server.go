package main

import (
	"github.com/defgenx/dynamic-router/web"
	"log"
	"net/http"
)

const (
	ADDRESS = ":1992"
)

func main() {
	// Execute route handler
	web.MyApp.HandleRoute()
	// Start serving
	log.Println("Starting server on: ", ADDRESS)
	// Init main route
	http.Handle("/", web.MyApp)

	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		log.Fatal("An error occured when trying to serve app on : \n", err)
	}
}
