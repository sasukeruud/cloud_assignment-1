package handlers

import (
	read "assignment_1/app/readData"
	"encoding/json"
	"fmt"
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
		http.Error(w, "Method not supported. Currentlyonly GET are supported", http.StatusNotImplemented)
	}
}

/*
Handler for REST get for Uni
some of the code is taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It has been modified to fit this application*/
func handlerGetRequestUni(w http.ResponseWriter, r *http.Request) {
	search := path.Base(r.URL.Path)

	if search != "uniinfo" {
		w.Header().Set("contet-type", "application/json")

		encoder := json.NewEncoder(w)

		err := encoder.Encode(read.ReadUniInfo(search))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
		}
	} else {
		fmt.Fprintf(w, "To search enter a valid search parameter \n")
		fmt.Fprintf(w, "eks: "+r.URL.Path[1:]+"norway")
	}
}
