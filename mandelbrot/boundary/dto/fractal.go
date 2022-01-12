package dto

type IterationColor struct {
	Iteration int    `query:"iteration" json:"iteration"`
	Color     string `query:"color" json:"color"`
}

type Fractal struct {
	C0               float64 `query:"c0" json:"c0"`
	C0i              float64 `query:"c0i" json:"c0i"`
	C1               float64 `query:"c1" json:"c1"`
	C1i              float64 `query:"c1i" json:"c1i"`
	Width            int     `query:"w" json:"w"`
	Height           int     `query:"h" json:"h"`
	MaxIterations    int     `query:"max_iterations" json:"max_iterations"`
	ImageCompression int     `query:"image_compression" json:"image_compression"`
	I0Color          string  `query:"i0_color" json:"i0_color"`
	I0Iteration      int     `query:"i0_iteration" json:"i0_iteration"`
	I1Color          string  `query:"i1_color" json:"i1_color"`
	I1Iteration      int     `query:"i1_iteration" json:"i1_iteration"`

	IterationColors []IterationColor `query:"iteration_colors" json:"iteration_colors"`
}
