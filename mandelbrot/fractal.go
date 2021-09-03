package mandelbrot

import (
	"log"
	"time"
)

type Fractal struct {
	C0            complex128 `form:"c0" json:"c0"`
	C1            complex128 `form:"c1" json:"c1"`
	Width         int        `form:"w" json:"w"`
	Height        int        `form:"h" json:"h"`
	MaxIterations int        `form:"max_iterations" json:"max_iterations" binding:"required"`
}

type IterationFunction func(z, c complex128, iterations int) IterationResult

type IterationResult struct {
	z         complex128
	completed int
}

type row struct {
	x   int
	row []int
}

// Use the iteration function f to render a fractal in the area [C0, C1] into a pixel area

func CreateFractal(f IterationFunction, fractal Fractal) [][]int {
	start := time.Now()

	img := make([][]int, fractal.Height)

	dy := (imag(fractal.C1) - imag(fractal.C0)) / float64(fractal.Height)
	queue := make(chan row, fractal.Height)
	for y := range img {
		cy := imag(fractal.C0) + float64(y)*dy
		go computeRow(f, y, cy, fractal, queue)
	}
	i := 0
	for i < fractal.Height {
		row := <-queue
		img[row.x] = row.row
		i++
	}
	log.Print("Computing fractal took ", time.Since(start))
	return img
}

func computeRow(f IterationFunction, y int, cy float64, fractalDef Fractal, c chan row) {
	r := make([]int, fractalDef.Width)
	dx := (real(fractalDef.C1) - real(fractalDef.C0)) / float64(fractalDef.Width)
	for x := range r {
		cx := real(fractalDef.C0) + float64(x)*dx
		r[x] = f(complex(0, 0), complex(cx, cy), fractalDef.MaxIterations).completed
	}
	c <- row{y, r}
}
