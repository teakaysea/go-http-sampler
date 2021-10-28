package zin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	path string
	m    map[string]handlerFunc
}

type handlers map[string]handler

type Engine struct {
	hs handlers
}
type Context struct {
	w http.ResponseWriter
	r *http.Request
}

type handlerFunc func(*Context)

func NewEngine() *Engine {
	e := new(Engine)
	e.hs = map[string]handler{}
	return e
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	c := new(Context)
	c.w = w
	c.r = r
	return c
}
func (c *Context) Status(code int) {
	c.w.WriteHeader(code)
}
func (c *Context) WriteBody(str string) {
	fmt.Fprintln(c.w, str)
}

type H map[string]interface{}

func (c *Context) JSON(status int, h H) {
	c.Status(status)
	if b, err := json.Marshal(h); err == nil {
		c.WriteBody(string(b))
	}
}
func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get(key)
}
func Run() {

	port := "8080"
	log.Printf("starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to start server on port %s", port)
	}
}

func (e *Engine) GET(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodGet, f)
}
func (e *Engine) POST(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodPost, f)
}
func (e *Engine) PUT(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodPut, f)
}
func (e *Engine) DELETE(path string, f handlerFunc) {
	e.handleFunc(path, http.MethodDelete, f)
}
func (e *Engine) handleFunc(path, method string, f handlerFunc) {
	h, ok := e.handler(path, method, f)
	if !ok {
		http.HandleFunc(path, h.handlerFunc)
	}
}
func (e *Engine) handler(path, method string, f handlerFunc) (handler, bool) {
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
		c.WriteBody("404 page not found")
		return
	}
	f(c)
}
