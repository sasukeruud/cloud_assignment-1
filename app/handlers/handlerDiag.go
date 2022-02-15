package handlers

import (
	"assignment_1/app/structs"
	"encoding/json"
	"net/http"
)

/*
 * Function to handle what type of request for the application to handle from uni
 * code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
 * It have been modified to fit this application */
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGetRequestDiag(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET and 	POST are supported", http.StatusNotImplemented)
	}
}

/*
* Handler for REST get for Country
* code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
* It have been modified to fit this application*/
func handlerGetRequestDiag(w http.ResponseWriter, r *http.Request) {
	country := structs.Diag{}

	w.Header().Add("contet-type", "application/country")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(country)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
