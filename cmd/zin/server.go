package main

import (
	"fmt"
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
	name := c.Query("name")
	if name == "" {
		name = "world"
	}
	c.JSON(http.StatusOK, zin.H{"message": fmt.Sprintf("Hello, %s!! via GET", name)})
}
func postHello(c *zin.Context) {
	c.JSON(http.StatusOK, zin.H{"message": "Hello, world!! via POST"})
}
func postBye(c *zin.Context) {
	c.JSON(http.StatusOK, zin.H{"message": "Bye, world!! via POST"})
}
