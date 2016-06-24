package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type H struct {

}

func (h *H) Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}