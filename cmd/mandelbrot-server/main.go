package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unstream/hellogo/mandelbrot"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Service is up")
	})
	router.GET("/mandelbrot", mandelbrotHandler())
	router.Run()
}

func mandelbrotHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		fractal := mandelbrot.Fractal{C0: complex(-1.5, -1), C1: complex(0.5, 1),
			Width: 700, Height: 700, MaxIterations: 100}
		var err error
		err = c.ShouldBindQuery(&fractal)
		if err != nil {
			log.Print(fractal)
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
			return
		}
		log.Print(fractal.MaxIterations)

		fractalArray := mandelbrot.CreateFractal(mandelbrot.MandelbrotFunction, fractal)
		img := mandelbrot.Image(fractalArray, fractal.MaxIterations)
		bytes, err := mandelbrot.CreateJpgImage(img)
		if err != nil {
			log.Print(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Data(http.StatusOK, "image/jpeg", bytes)
	}
}
