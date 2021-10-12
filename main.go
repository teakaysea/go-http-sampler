package main

import (
	"fmt"
	"net/http"
)

type handler struct {
	method string
	f      http.HandlerFunc
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, world!!")
}
func main() {
	handleFunc("/hello", http.MethodGet, hello)

	http.ListenAndServe(":8080", nil)
}
func handleFunc(path, method string, f http.HandlerFunc) {
	h := handler{method: method, f: f}
	http.HandleFunc(path, h.handlerFunc)
}
func (h handler) handlerFunc(w http.ResponseWriter, r *http.Request) {
	if h.method != r.Method {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h.f(w, r)
}
