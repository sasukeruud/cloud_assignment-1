package handlers

import (
	read "assignment_1/app/readData"
	"assignment_1/app/structs"
	"encoding/json"
	"net/http"
	"path"
)

/*
Function to handle what type of request for the application to handle from uni
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo */
func UniHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGetRequestUni(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET and 	POST are supported", http.StatusNotImplemented)
	}
}

/*
Handler for REST get for Uni
some of the code is taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It have been modified to fit this application*/
func handlerGetRequestUni(w http.ResponseWriter, r *http.Request) {
	search := path.Base(r.URL.Path)

	var uni []structs.Uni
	json.Unmarshal(read.ReadUniAPI(search), &uni)

	w.Header().Add("contet-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(uni)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
