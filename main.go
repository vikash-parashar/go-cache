package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {}
