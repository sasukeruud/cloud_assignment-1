package handlers

import (
	read "assignment_1/app/readData"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	var query int

	if r.URL.Query().Get("limit") != "" {
		queryTest, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		query = queryTest
	} else {
		query = 0
	}

	if len(search) == 6 && search[5] != "" {
		w.Header().Set("contet-type", "application/json")

		encoder := json.NewEncoder(w)

		err := encoder.Encode(read.ReadCountryUniInfo(search[4], search[5], query))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
		}
	} else {
		fmt.Fprintf(w, "To search enter a valid search parameters \n")
		fmt.Fprintf(w, "eks: "+r.URL.Path[1:]+"norway/science\n")
		fmt.Fprintf(w, "or eks: "+r.URL.Path[1:]+"norway/science?limit=3")
	}

}
