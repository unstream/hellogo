package mandelbrot

import (
	"image"
	"image/color"
	"log"
	"time"
)

type IterationResult struct {
	z         complex128
	completed int
}

type Row struct {
	x   int
	row []int
}

type iterationFunction func(z, c complex128, iterations int) IterationResult

// Convert a pixel array into an actual image

func Image() image.Image {

	const iterations = 500
	const width, height = 700, 700
	const c0 = complex(-1.5, -1)
	const c1 = complex(0.5, 1)
	fractal := CreateFractal(IterationFunction, c0, c1, width, height, iterations)

	start := time.Now()
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			col := uint8(0)
			if fractal[x][y] == iterations {
				col = 255
			} else {
				col = uint8(fractal[x][y] * 255 / iterations)
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

// Use the iteration function f to render a fractal in the area [c0, c1] into a pixel area

func CreateFractal(f iterationFunction, c0, c1 complex128, width, height, iterations int) [][]int {
	start := time.Now()

	img := make([][]int, width)
	dx := (real(c1) - real(c0)) / float64(width)
	dy := (imag(c1) - imag(c0)) / float64(height)
	queue := make(chan Row, height)
	for x := range img {
		cx := real(c0) + float64(x)*dy
		go computeRow(f, c0, x, height, dx, cx, iterations, queue)
	}
	i := 0
	for i < height {
		row := <-queue
		img[row.x] = row.row
		i++
	}
	log.Print("Computing fractal took ", time.Since(start))
	return img
}

func computeRow(f iterationFunction, c0 complex128, x int, height int, dx float64, cx float64, iterations int, c chan Row) {
	row := make([]int, height)
	for y := range row {
		cy := imag(c0) + float64(y)*dx
		row[y] = f(complex(0, 0), complex(cx, cy), iterations).completed
	}
	c <- Row{x, row}
}

// Iterate the Mandelbrot function f(z) =  z^2 + c over the given number of iterations by returning f(f(...)).
// Using the result you can start iterating for further iterations.

func IterationFunction(z, c complex128, iterations int) IterationResult {
	var x = real(z)
	var y = imag(z)
	var completed = iterations
	for i := 0; i < iterations; i++ {
		x2 := x * x
		y2 := y * y
		y = 2*x*y + imag(c)
		x = x2 - y2 + real(c)
		if (x2 + y2) >= 4 {
			completed = i + 1
			break
		}
	}
	return IterationResult{z, completed}
}
