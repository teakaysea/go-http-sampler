package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	path string
	m    map[string]handlerFunc
}

func key(path, method string) string {
	return fmt.Sprintf("%s %s", method, path)
}

// func newHandler(path, method string, f handlerFunc) handler {
// 	return handler{key: key(path, method), path: path, method: method, f: f}
// }

type handlers map[string]handler

type engine struct {
	hs handlers
}
type context struct {
	w http.ResponseWriter
	r *http.Request
}

type handlerFunc func(*context)

func newEngine() *engine {
	e := new(engine)
	e.hs = map[string]handler{}
	return e
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
func getHello(c *context) {
	c.status(http.StatusOK)
	c.writeBody("Hello, world!! via GET")
}
func postHello(c *context) {
	c.status(http.StatusOK)
	c.writeBody("Hello, world!! via POST")
}
func main() {
	e := newEngine()
	e.GET("/hello", getHello)
	e.POST("/hello", postHello)

	port := "8080"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to start server on port %s", port)
	}
	log.Printf("server on port %s started", port)
}

func (e *engine) GET(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodGet, f)
}
func (e *engine) POST(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodPost, f)
}
func (e *engine) handleFunc(path, method string, f handlerFunc) {
	h, ok := e.handler(path, method, f)
	if !ok {
		http.HandleFunc(path, h.handlerFunc)
	}
}
func (e *engine) handler(path, method string, f handlerFunc) (handler, bool) {
	h, ok := e.hs[path]
	if !ok {
		h = handler{path: path, m: map[string]handlerFunc{}}
		e.hs[path] = h
	}
	h.m[method] = f
	return h, ok
}
func (h handler) handlerFuncWithContext(method string) (handlerFunc, bool) {
	f, ok := h.m[method]
	return f, ok
}
func (h handler) handlerFunc(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	f, ok := h.handlerFuncWithContext(r.Method)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	f(c)
}
