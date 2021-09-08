package main

type FractalDto struct {
	C0            float64 `form:"c0" json:"c0"`
	C0i           float64 `form:"c0i" json:"c0i"`
	C1            float64 `form:"c1" json:"c1"`
	C1i           float64 `form:"c1i" json:"c1i"`
	Width         int     `form:"w" json:"w"`
	Height        int     `form:"h" json:"h"`
	MaxIterations int     `form:"max_iterations" json:"max_iterations" binding:"required"`
}
