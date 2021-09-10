package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/unstream/hellogo/mandelbrot/boundary"
	"os"
)

func main() {
	// Initialization of go-echo server
	name := "Mandelbrot-Server"
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${bytes_in} bytes_in | ${bytes_out} bytes_out | ${method} | ${uri} ",
			name,
		),
		CustomTimeFormat: "2006/01/02 15:04:05", // custom readable time format
		Output:           os.Stdout,             // output method
	}))

	// list of endpoint routes
	APIRoute := e.Group("/api")
	// grouping routes for version 1.0 API
	v1route := APIRoute.Group("/v1")
	v1route.GET("/mandelbrot", boundary.MandelbrotHandler)

	// firing up the server
	e.Logger.Fatal(e.Start(":8080"))
}
