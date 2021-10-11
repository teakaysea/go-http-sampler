package main

import (
	"fmt"
	"log"
	"net/http"
)

var handlerMap map[string]handler

type handler struct {
	path   string
	method string
	f      http.HandlerFunc
}

func (h handler) key() string {
	return fmt.Sprintf("%s %s", h.method, h.path)
}
func hello(w http.ResponseWriter, r *http.Request) {
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
func init() {
	handlerMap = map[string]handler{}
}
func main() {
	addRoute(http.MethodGet, "/", hello)

	http.ListenAndServe(":8080", nil)
}
func newHandler(path, method string, f http.HandlerFunc) handler {
	return handler{path: path, method: method, f: f}
}
func addRoute(path, method string, f http.HandlerFunc) {
	h := newHandler(path, method, f)
	handlerMap[h.key()] = h
	http.HandleFunc(path, getHandler(h))
}
func getHandler(h handler) http.HandlerFunc {
	return nil
}
