package handlers

import (
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
 * Function to handle what type of request for the application to handle from uni*/
func uniHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlerPostRequestUni(w, r)
	case http.MethodGet:
		handlerGetRequestUni(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET and 	POST are supported", http.StatusNotImplemented)
	}
}

/*
* Handler for REST post for Uni*/
func handlerPostRequestUni(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	uni := structs.Uni{}

	err := decoder.Decode(&uni)
	if err != nil {
		http.Error(w, "Error during decoding", http.StatusInternalServerError)
		return
	}

	output, err := json.MarshalIndent(uni, "", "")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(output))

	http.Error(w, "OK", http.StatusOK)
}

/*
* Handler for REST get for Uni*/
func handlerGetRequestUni(w http.ResponseWriter, r *http.Request) {
	uni := structs.Uni{}

	w.Header().Add("contet-type", "application/diag")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(uni)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}

func postUni() {

}

func getCountry() {

}

func postCountry() {

}
