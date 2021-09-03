package mandelbrot

// MandelbrotFunction Iterate the Mandelbrot function f(z) =  z^2 + c over the given number of iterations by returning f(f(...)).
// Using the result you can start iterating for further iterations.
func MandelbrotFunction(z, c complex128, iterations int) IterationResult {
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
