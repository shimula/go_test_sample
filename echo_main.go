package main

import (
	"net/http"
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)


func index(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}


func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/users/:id", handler.GetUsers)
	e.Run(standard.New(":8888"))
}
