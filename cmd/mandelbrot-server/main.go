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
		fractalDto := FractalDto{
			C0: -1.5, C0i: -1, C1: 0.5, C1i: 1,
			Width: 700, Height: 700, MaxIterations: 100,
		}
		var err error
		err = c.ShouldBindQuery(&fractalDto)
		if err != nil {
			log.Print(fractalDto)
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
			return
		}
		log.Print(fractalDto)
		c0 := complex(fractalDto.C0, fractalDto.C0i)
		c1 := complex(fractalDto.C1, fractalDto.C1i)
		fractal := mandelbrot.Fractal{
			C0: c0, C1: c1,
			Width: fractalDto.Width, Height: fractalDto.Height,
			MaxIterations: fractalDto.MaxIterations,
		}
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
