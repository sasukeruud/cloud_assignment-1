package handlers

import (
	read "assignment_1/app/readData"
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

/*
Function to handle what type of request for the application to handle from uni
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo */
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGetRequestCountry(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET are supported", http.StatusNotImplemented)
	}
}

/*
Handler for REST get for Country
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It have been modified to fit this application*/
func handlerGetRequestCountry(w http.ResponseWriter, r *http.Request) {
	search := path.Base(r.URL.Path)

	var country []structs.Country

	fmt.Println(search)
	//json.Unmarshal(read.ReadCountriesAPI(search), &country)
	read.ReadCountriesAPI(search)

	w.Header().Add("contet-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(country)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
