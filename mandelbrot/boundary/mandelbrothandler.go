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

func createMandelbrotImage(dto dto.Fractal) ([]byte, error) {

	fractal := entity.Fractal{
		C0: complex(dto.C0, dto.C0i), C1: complex(dto.C1, dto.C1i),
		Width: dto.Width, Height: dto.Height,
		MaxIterations: dto.MaxIterations,
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
