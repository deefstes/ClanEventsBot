package main

import (
	"fmt"
	"net/http"
)

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.RequestURI)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
