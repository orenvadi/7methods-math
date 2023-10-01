package methods

import (
	"fmt"
	"math"
)

func SimpleIterationMethod(f func(float64) float64, x0, tol float64, maxIterations int) float64 {
	iterations := 0

	fmt.Printf("%-12s %-12s %-12s\n", "Iteration", "x", "Tolerance")

	for iterations < maxIterations {
		iterations++
		x1 := f(x0)
		tolerance := math.Abs(x1 - x0)

		fmt.Printf("%-12d %-12.6f %-12.6f\n", iterations, x1, tolerance)

		if tolerance < tol {
			return x1
		}

		x0 = x1
	}

	fmt.Println("Maximum number of iterations reached.")
	return x0
}
