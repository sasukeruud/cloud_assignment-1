package handlers

import (
	"fmt"
	"net/http"
)

func DeafultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		defaultPage(w, r)
	}
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
