package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/unstream/hellogo/mandelbrot"
	"image/jpeg"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/mandelbrot", func(c *gin.Context) {
		iterations := c.DefaultQuery("iterations", "100")
		log.Print(iterations)
		writeImage(c)
	})
	router.Run()
}

// writeImage encodes an image 'img' in jpeg format and writes it into the context response.
func writeImage(c gin.Context) {
	img := mandelbrot.Image()

	start := time.Now()
	buffer := new(bytes.Buffer)
	options := new(jpeg.Options)
	options.Quality = 100
	if err := jpeg.Encode(buffer, img, options); err != nil {
		log.Println("unable to encode image.")
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Data(http.StatusOK, "image/jpeg", buffer.Bytes())
		log.Print("Encoding image took ", time.Since(start))
	}
}
