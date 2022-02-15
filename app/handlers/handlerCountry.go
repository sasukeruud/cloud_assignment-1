package handlers

import (
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
 * Function to handle what type of request for the application to handle from uni*/
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlerPostRequestCountry(w, r)
	case http.MethodGet:
		handlerGetRequestCountry(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET and 	POST are supported", http.StatusNotImplemented)
	}
}

/*
* Handler for REST post for Country*/
func handlerPostRequestCountry(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	country := structs.Country{}

	err := decoder.Decode(&country)
	if err != nil {
		http.Error(w, "Error during decoding", http.StatusInternalServerError)
		return
	}

	output, err := json.MarshalIndent(country, "", "")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(output))

	http.Error(w, "OK", http.StatusOK)
}

/*
* Handler for REST get for Country*/
func handlerGetRequestCountry(w http.ResponseWriter, r *http.Request) {
	country := structs.Country{
		Name:           "Norge",
		Aplha_two_code: "NO",
		Region:         "Europe",
		Capital:        "Oslo"}

	w.Header().Add("contet-type", "application/country")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(country)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
