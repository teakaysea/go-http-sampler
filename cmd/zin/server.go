package main

import (
	"net/http"

	"github.com/teakaysea/go-http-sampler/zin"
)

func main() {
	e := zin.NewEngine()
	e.GET("/hello", getHello)
	e.POST("/hello", postHello)
	e.POST("/bye", postBye)

	zin.Run()
}
func getHello(c *zin.Context) {
	c.Status(http.StatusOK)
	c.WriteBody("Hello, world!! via GET")
}
func postHello(c *zin.Context) {
	c.Status(http.StatusOK)
	c.WriteBody("Hello, world!! via POST")
}
func postBye(c *zin.Context) {
	c.Status(http.StatusOK)
	c.WriteBody("Bye, world!! via POST")
}
