package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	hostName, _ := os.Hostname()

	msg := fmt.Sprintf("Application Demo B - host: %v\n", hostName)

	fmt.Println(msg)

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, msg)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
