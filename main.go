package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s", r.Method)
	switch r.Method {
	case "GET":
		log.Printf("implementation for GET")
	case "POST":
		log.Printf("implementation for POST")
	default:
		log.Printf("%s not allowed", r.Method)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, world!!")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
