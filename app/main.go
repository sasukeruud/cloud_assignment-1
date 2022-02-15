package main

import (
	"fmt"
	"net/http"
	"os"
	//"./structs"
	//"strings"
	//"log"
)

/*
* Function to get a spesific port to run the application on.
* It's needed for the application to run on heroku where to application is run on.*/
func getport() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helllo, %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(getport(), nil)
}
