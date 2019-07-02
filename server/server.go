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

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println("REQ:", string(reqBody))
		fmt.Println("RES:", string(resBody))
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(port))

}