package control

import (
	"bytes"
	"github.com/unstream/hellogo/mandelbrot/entity"
	"gopkg.in/go-playground/colors.v1"
	"image"
	"image/color"
	"image/png"
	"log"
	"time"
)

func createColorMap(fractal entity.Fractal) []color.NRGBA {
	colorMap := make([]color.NRGBA, fractal.MaxIterations+1)
	log.Println(fractal)
	c0, err0 := colors.ParseHEX(fractal.IterationColors[0].Color)
	if err0 != nil {
		c0, err0 = colors.ParseHEX("#000000")
	}
	c1, err1 := colors.ParseHEX(fractal.IterationColors[1].Color)
	if err1 != nil {
		c1, err1 = colors.ParseHEX("#00ff00")
	}

	for i := 0; i <= fractal.MaxIterations; i++ {
		colorMap[i] = color.NRGBA{
			R: uint8(((fractal.MaxIterations-i)*int(c0.ToRGB().R) + i*int(c1.ToRGB().R)) / fractal.MaxIterations),
			G: uint8(((fractal.MaxIterations-i)*int(c0.ToRGB().G) + i*int(c1.ToRGB().G)) / fractal.MaxIterations),
			B: uint8(((fractal.MaxIterations-i)*int(c0.ToRGB().B) + i*int(c1.ToRGB().B)) / fractal.MaxIterations),
			A: 255,
		}
	}
	return colorMap
}

func CreatePngImage(img image.Image, compressionLevel png.CompressionLevel) (pngimage []byte, err error) {
	start := time.Now()
	buffer := new(bytes.Buffer)
	pngimage = nil
	encoder := png.Encoder{
		CompressionLevel: compressionLevel,
	}

	if err = encoder.Encode(buffer, img); err != nil {
		log.Println("unable to encode image.")
	} else {
		log.Print("Encoding image took ", time.Since(start))
		pngimage = buffer.Bytes()
	}
	return pngimage, err
}

// Convert a pixel array into an actual image
func Image16(fractalArray [][]int, fractal entity.Fractal) image.Image {
	start := time.Now()
	width := len(fractalArray[0])
	height := len(fractalArray)
	img := image.NewNRGBA64(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			col := uint16(0)
			i := fractalArray[y][x]
			if i == fractal.MaxIterations {
				col = 65535
			} else {
				col = uint16(i * 65536 / fractal.MaxIterations)
			}

			img.Set(x, y, color.RGBA64{
				R: 0,
				G: 0,
				B: col,
				A: 65535,
			})
		}
	}
	log.Print("Creating image took ", time.Since(start))
	return img
}

// Convert a pixel array into an actual image
func Image(fractalArray [][]int, fractal entity.Fractal) image.Image {
	start := time.Now()
	width := len(fractalArray[0])
	height := len(fractalArray)

	colorMap := createColorMap(fractal)

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := fractalArray[y][x]
			img.Set(x, y, colorMap[i])
		}
	}
	log.Print("Creating image took ", time.Since(start))
	return img
}
