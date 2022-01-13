package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/unstream/hellogo/mandelbrot/boundary"
	"net/http"
	"os"
)

// content holds our static web server content.
//go:embed web/build
var content embed.FS

// Initialization of go-echo server
var contentHandler = echo.WrapHandler(http.FileServer(http.FS(content)))

// The embedded files will all be in the '/static' folder so need to rewrite the request (could also do this with fs.Sub)
var contentRewrite = middleware.Rewrite(map[string]string{"/*": "/web/build/$1"})

func main() {

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
	e.GET("/*", contentHandler, contentRewrite)

	APIRoute := e.Group("/api")
	// grouping routes for version 1.0 API
	v1route := APIRoute.Group("/v1")
	v1route.GET("/mandelbrot", boundary.MandelbrotGetHandler)
	v1route.POST("/mandelbrot", boundary.MandelbrotPostHandler)

	// firing up the server
	e.Logger.Fatal(e.Start(":8081"))
}
