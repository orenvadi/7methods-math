package methods

import (
	"fmt"
	"math"
)

func NewtonMethod(f func(float64) float64, df func(float64) float64, x0, tol float64) float64 {
	iterations := 0

	fmt.Printf("%-12s %-12s %-12s\n", "Iteration", "x", "Tolerance")

	for {
		iterations++
		x1 := x0 - f(x0)/df(x0)
		tolerance := math.Abs(x1 - x0)

		fmt.Printf("%-12d %-12.6f %-12.6f\n", iterations, x1, tolerance)

		if tolerance < tol {
			return x1
		}

		x0 = x1
	}
}
