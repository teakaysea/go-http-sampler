package main

import (
	"fmt"
	"net/http"
)

type handler struct {
	method string
	f      handlerFunc
}
type engine struct {
}
type context struct {
	w http.ResponseWriter
	r *http.Request
}

type handlerFunc func(*context)

func newEngine() *engine {
	return new(engine)
}

func newContext(w http.ResponseWriter, r *http.Request) *context {
	c := new(context)
	c.w = w
	c.r = r
	return c
}
func (c *context) status(code int) {
	c.w.WriteHeader(code)
}
func (c *context) writeBody(str string) {
	fmt.Fprintln(c.w, str)
}
func hello(c *context) {
	c.status(http.StatusOK)
	c.writeBody("Hello, world!!")
}
func main() {
	e := newEngine()
	e.GET("/hello", hello)

	http.ListenAndServe(":8080", nil)
}

func (e *engine) handleFunc(path, method string, f handlerFunc) {
	h := handler{method: method, f: f}
	http.HandleFunc(path, h.handlerFunc)
}
func (e *engine) GET(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodGet, f)
}
func (h handler) handlerFunc(w http.ResponseWriter, r *http.Request) {
	if h.method != r.Method {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := newContext(w, r)
	h.f(c)
}
