package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"app/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var portPtr = flag.Int("port", 8888, "port for game to listen on")
	flag.Parse()
  
	fmt.Println("Listening on port", *portPtr)

	//var port string
	port := ":" + strconv.Itoa(*portPtr)


	var ms, err = handlers.NewMemoryStore()
	if err != nil {
		log.Fatalf("Failed to create memory store", err)
	}


	// Echo instance
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println("REQ:", string(reqBody))
		ct := c.Response().Header().Get(echo.HeaderContentType)
		fmt.Println("Result content type is", ct)
		if ct[:6] != "image/" {
			fmt.Println("RES:", string(resBody))
		}
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//
	// Inject our own custom context with memory database
	//
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.CustomContext{Context: c}
			cc.SetDatabase(ms)
			return h(cc)
		}
	})

	// Static Assets
	e.File("/favicon.ico", "images/favicon.ico")

	// Routes
	e.GET("/", handlers.Hello)
	e.POST("/echo", handlers.Echo)

	e.POST("/createTask", handlers.CreateTask)
	e.POST("/startTask", handlers.StartTask)
	e.POST("/finishTask", handlers.FinishTask)

	// Start server
	e.Logger.Fatal(e.Start(port))

}