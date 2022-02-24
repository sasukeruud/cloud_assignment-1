package handlers

import (
	"fmt"
	"net/http"
)

/*
Function to handle access to the deafult API page.
code taken from: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022/-/tree/main/02-JSON-demo
It has been modified to fit this application  */
func DeafultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		defaultPage(w, r)
	}
}

/*
Default page for when the API*/
func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To use the api use the url with either of these:\n")
	fmt.Fprintf(w, "/unisearcher/v1/uniinfo/ \n")
	fmt.Fprintf(w, "/unisearcher/v1/neighbourunis/\n")
	fmt.Fprintf(w, "/unisearcher/v1/diag/\n")
}
