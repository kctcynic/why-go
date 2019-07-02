package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


// Hello handles GET requests on /
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
  