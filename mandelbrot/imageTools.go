package mandelbrot

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"time"
)

func CreateJpgImage(img image.Image) (jpg []byte, err error) {
	start := time.Now()
	buffer := new(bytes.Buffer)
	options := new(jpeg.Options)
	options.Quality = 100
	jpg = nil
	if err = jpeg.Encode(buffer, img, options); err != nil {
		log.Println("unable to encode image.")
	} else {
		log.Print("Encoding image took ", time.Since(start))
		jpg = buffer.Bytes()
	}
	return jpg, err
}

// Convert a pixel array into an actual image

func Image(fractalArray [][]int, maxIterations int) image.Image {

	start := time.Now()
	width := len(fractalArray[0])
	height := len(fractalArray)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			col := uint8(0)
			if fractalArray[y][x] == maxIterations {
				col = 255
			} else {
				col = uint8(fractalArray[y][x] * 255 / maxIterations)
			}

			img.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: col,
				A: 255,
			})
		}
	}
	log.Print("Creating image took ", time.Since(start))
	return img
}
