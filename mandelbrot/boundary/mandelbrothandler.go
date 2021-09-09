package boundary

import (
	"github.com/labstack/echo"
	"github.com/unstream/hellogo/mandelbrot/boundary/dto"
	"github.com/unstream/hellogo/mandelbrot/control"
	"github.com/unstream/hellogo/mandelbrot/entity"
	"log"
	"net/http"
)

func MandelbrotHandler(c echo.Context) error {

	fractalDto := dto.Fractal{
		C0: -1.5, C0i: -1, C1: 0.5, C1i: 1,
		Width: 700, Height: 700, MaxIterations: 100,
	}

	if err := c.Bind(&fractalDto); err != nil {
		log.Fatal(err)
	}
	bytes, err := createMandelbrotImage(fractalDto)
	if err != nil {
		return err
	}
	return c.Blob(http.StatusOK, "image/jpg", bytes)
}

func createMandelbrotImage(fractalDto dto.Fractal) ([]byte, error) {

	c0 := complex(fractalDto.C0, fractalDto.C0i)
	c1 := complex(fractalDto.C1, fractalDto.C1i)
	fractal := entity.Fractal{
		C0: c0, C1: c1,
		Width: fractalDto.Width, Height: fractalDto.Height,
		MaxIterations: fractalDto.MaxIterations,
	}
	fractalArray := control.CreateFractal(control.MandelbrotFunction, fractal)
	img := control.Image(fractalArray, fractal.MaxIterations)
	bytes, err := control.CreateJpgImage(img)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return bytes, nil
}
