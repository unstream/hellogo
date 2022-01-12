package boundary

import (
	"github.com/labstack/echo/v4"
	"github.com/unstream/hellogo/mandelbrot/boundary/dto"
	"github.com/unstream/hellogo/mandelbrot/control"
	"github.com/unstream/hellogo/mandelbrot/entity"
	"image"
	"image/png"
	"log"
	"net/http"
)

func MandelbrotGetHandler(c echo.Context) error {

	fractalDto := dto.Fractal{
		C0: -1.5, C0i: -1, C1: 0.5, C1i: 1,
		Width: 700, Height: 700, MaxIterations: 100, ImageCompression: 0,
		I0Color: "#000000", I0Iteration: 0,
		I1Color: "#0000dd", I1Iteration: 100,
	}
	if err := c.Bind(&fractalDto); err != nil {
		return err
	}
	log.Println(fractalDto)

	iterationColors := make([]dto.IterationColor, 2)
	iterationColors[0] = dto.IterationColor{
		Iteration: fractalDto.I0Iteration,
		Color:     fractalDto.I0Color,
	}
	iterationColors[1] = dto.IterationColor{
		Iteration: fractalDto.I1Iteration,
		Color:     fractalDto.I1Color,
	}
	fractalDto.IterationColors = iterationColors

	log.Println("IterationColors: ", fractalDto.IterationColors)

	bytes, err := createMandelbrotImage(fractalDto)
	if err != nil {
		return err
	}

	return c.Blob(http.StatusOK, "image/png", bytes)
}

func MandelbrotPostHandler(c echo.Context) error {
	fractalDto := new(dto.Fractal)
	if err := c.Bind(fractalDto); err != nil {
		return err
	}

	bytes, err := createMandelbrotImage(*fractalDto)
	if err != nil {
		return err
	}
	return c.Blob(http.StatusOK, "image/png", bytes)
}

func createMandelbrotImage(dto dto.Fractal) ([]byte, error) {
	var bytes []byte
	var err error
	var img image.Image

	iterationColors := make([]entity.IterationColor, len(dto.IterationColors))
	for i := 0; i < len(dto.IterationColors); i++ {
		iterationColors[i] = entity.IterationColor{
			Iteration: dto.IterationColors[i].Iteration,
			Color:     dto.IterationColors[i].Color,
		}
	}

	fractal := entity.Fractal{
		C0: complex(dto.C0, dto.C0i), C1: complex(dto.C1, dto.C1i),
		Width: dto.Width, Height: dto.Height,
		MaxIterations:    dto.MaxIterations,
		ImageCompression: dto.ImageCompression,
		IterationColors:  iterationColors,
	}
	fractalArray := control.CreateFractal(control.MandelbrotFunction, fractal)
	log.Println("IterationColors: ", dto.IterationColors)
	if fractal.ImageCompression == -1 {
		img = control.Image16(fractalArray, fractal)
		bytes, err = control.CreatePngImage(img, png.NoCompression)
	} else {
		img = control.Image(fractalArray, fractal)
		bytes, err = control.CreatePngImage(img, png.DefaultCompression)
	}
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return bytes, nil
}
