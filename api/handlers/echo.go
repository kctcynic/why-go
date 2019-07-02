package handlers


import (
	"fmt"
	"net/http"

	"app/bindings"

	"github.com/labstack/echo/v4"
)

// Echo handles API requests on /echo
func Echo(c echo.Context) error {
	var msg *bindings.ErrorResult

	var params = new(bindings.EchoParameters)
	if err := c.Bind(params); err != nil {
		fmt.Println(err)
		msg = new(bindings.ErrorResult)
		msg.ErrorMessage = "Invalid Parameters"
		return c.JSON(http.StatusOK, msg)
	}

	var echoResult = new(bindings.EchoResult)

	echoResult.Output = params.Input

	return c.JSON(http.StatusOK, echoResult)
}
