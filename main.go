package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unstream/hellogo/mandelbrot/boundary"
	"net/http"
	"os"
	"strconv"
)

const version = "V1"

// content holds our static web server content.
//go:embed web/build
var content embed.FS

// Initialization of go-echo server
var contentHandler = echo.WrapHandler(http.FileServer(http.FS(content)))

// The embedded files will all be in the '/static' folder so need to rewrite the request (could also do this with fs.Sub)
var contentRewrite = middleware.Rewrite(map[string]string{"/*": "/web/build/$1"})

var configPath string

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "MandelbrotServer",
	Long:  `Mandelbrot Server`,
}

func init() {
	serverCmd := &cobra.Command{
		Use:   "run",
		Short: "server",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	viper.SetDefault("port", "8081")
	rootCmd.AddCommand(serverCmd)
}

func main() {

	viper.AutomaticEnv() // Automatically search for environment variables
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	name := "Mandelbrot-Server"
	e := echo.New()
	e.HideBanner = true
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
	var port = 8081
	if viper.IsSet("port") {
		port = viper.GetInt("port")
	}
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}
