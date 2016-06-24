package main

import (
	"github.com/shimula/go_test_sample/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.GET("/", H)
	e.Run(standard.New(":8888"))
}
