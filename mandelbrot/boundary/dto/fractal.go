package dto

type Fractal struct {
	C0            float64 `query:"c0" json:"c0"`
	C0i           float64 `query:"c0i" json:"c0i"`
	C1            float64 `query:"c1" json:"c1"`
	C1i           float64 `query:"c1i" json:"c1i"`
	Width         int     `query:"w" json:"w"`
	Height        int     `query:"h" json:"h"`
	MaxIterations int     `query:"max_iterations" json:"max_iterations"`
}
