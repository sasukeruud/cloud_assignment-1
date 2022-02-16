package main

import (
	constants "assignment_1/app"
	"assignment_1/app/handlers"
	"fmt"
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

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)

	http.HandleFunc(constants.UNI_PATH, handlers.UniHandler)
	http.HandleFunc(constants.COUNTRY_PATH, handlers.CountryHandler)
	http.HandleFunc(constants.DIAG_PATH, handlers.DiagHandler)

	http.ListenAndServe(getport(), nil)
}
