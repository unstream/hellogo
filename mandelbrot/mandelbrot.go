package mandelbrot

import (
	"image"
	"image/color"
)

type IterationResult struct {
	Z         complex128
	Completed int
}

func Image() image.Image {
	const iterations = 100
	const width, height = 500, 500
	const c0 = complex(-1.5, -1)
	const c1 = complex(0.5, 1)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	dx := (real(c1) - real(c0)) / width
	dy := (imag(c1) - imag(c0)) / height

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := complex(real(c0)+float64(x)*dx, imag(c0)+float64(y)*dy)
			result := Compute(complex(0, 0), c, iterations)
			var col uint8
			if result.Completed == iterations {
				col = 255
			} else {
				col = uint8(result.Completed * 255 / iterations)
			}

			img.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: col,
				A: 255,
			})
		}
	}
	return img
}

// Iterate the Mandelbrot function f(z) =  z^2 + c over the given number of iterations by returning f(f(...)).
// Using the result you can start iterating for further iterations.

func Compute(z, c complex128, iterations int) IterationResult {
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
