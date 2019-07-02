// Sample API server
package main

import (
	"flag"
	"fmt"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {

	var portPtr = flag.Int("port", 8888, "port for game to listen on")
	flag.Parse()
  
	fmt.Println("Listening on port", *portPtr)

	//var port string
	port := ":" + strconv.Itoa(*portPtr)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static Assets
	e.File("/favicon.ico", "images/favicon.ico")

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(port))

}