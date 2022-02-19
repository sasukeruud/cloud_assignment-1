package main

import (
	constants "assignment_1/app"
	"assignment_1/app/handlers"
	"net/http"
	"os"
)

/*
Function to get a spesific port to run the application on.
It's needed for the application to run on heroku where to application is run on.*/
func getport() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {
	http.HandleFunc(constants.UNI_PATH, handlers.UniHandler)
	http.HandleFunc(constants.COUNTRY_UNI_PATH, handlers.UniNeighborHandler)
	http.HandleFunc(constants.DIAG_PATH, handlers.DiagHandler)

	http.ListenAndServe(getport(), nil)
}
