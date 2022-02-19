package handlers

import (
	read "assignment_1/app/readData"
	"encoding/json"
	"net/http"
	"strings"
)

/*
Function to handle what type of request for the application to handle from uni
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo */
func UniNeighborHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGetRequestUniNeighbor(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET are supported", http.StatusNotImplemented)
	}
}

/*
Handler for REST get for Uni
some of the code is taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It has been modified to fit this application*/
func handlerGetRequestUniNeighbor(w http.ResponseWriter, r *http.Request) {
	search := strings.SplitAfter(r.URL.Path, "/")

	w.Header().Set("contet-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(read.ReadCountryUniInfo(search[4], search[5]))
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
