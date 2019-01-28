package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type AppResponse struct {
	Project string `json:"projectName"`
	Email   string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
