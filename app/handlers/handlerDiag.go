package handlers

import (
	constants "assignment_1/app"
	"assignment_1/app/structs"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var start time.Time = time.Now()

/*
Function to handle what type of request for the application to handle from uni
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It have been modified to fit this application */
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGetRequestDiag(w, r)
	default:
		http.Error(w, "Method not supported. Currentlyonly GET are supported", http.StatusNotImplemented)
	}
}

/*
Handler for REST get for Country
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It have been modified to fit this application*/
func handlerGetRequestDiag(w http.ResponseWriter, r *http.Request) {
	respUni, errUni := http.Get(constants.READ_ALL_UNI)
	respCount, errCount := http.Get(constants.READ_ALL_COUNTRIES + "no")
	if errUni != nil || errCount != nil {
		log.Fatal(errUni)
		log.Fatal(errCount)
	}

	diag := structs.Diag{
		Universitiapi: respUni.StatusCode,
		Counteriesapi: respCount.StatusCode,
		Version:       constants.VERSION,
		Uptime:        time.Duration.Seconds(time.Since(start)),
	}

	w.Header().Add("contet-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(diag)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}
