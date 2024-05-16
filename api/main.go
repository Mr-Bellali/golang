package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Any("/greaterthan/:id", func(c echo.Context) error {
		n, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid number")
		}

		if n > 5 {
			return c.String(http.StatusOK, fmt.Sprintf("%d is bigger than 5", n))
		} else if n < 5 {
			return c.String(http.StatusOK, fmt.Sprintf("%d is smaller than 5", n))
		}
		return c.String(http.StatusOK, fmt.Sprintf("%d = 5", n))
	})

	e.Logger.Fatal(e.Start(":3000"))
}
