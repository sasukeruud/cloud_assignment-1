package handlers

import (
	read "assignment_1/app/readData"
	"assignment_1/app/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
 * Function to handle what type of request for the application to handle from uni
 * code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo */
func UniHandler(w http.ResponseWriter, r *http.Request) {
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
* Handler for REST post for Uni
* code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
* It have been modified to fit this application*/
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
* Handler for REST get for Uni
* code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
* It have been modified to fit this application*/
func handlerGetRequestUni(w http.ResponseWriter, r *http.Request) {
	var uni []structs.Uni
	json.Unmarshal(read.ReadUniAPI("norwegian"), &uni)
	for i := 0; i < len(uni); i++ {
		fmt.Print(uni[i].Name + "\n")
	}

	w.Header().Add("contet-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(uni)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
