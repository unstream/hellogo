package entity

type IterationColor struct {
	Iteration int
	Color     string
}

type Fractal struct {
	C0               complex128
	C1               complex128
	Width            int
	Height           int
	MaxIterations    int
	ImageCompression int
	IterationColors  []IterationColor
}
