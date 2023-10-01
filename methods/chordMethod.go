package methods

import (
	"fmt"
	"math"
)

func ChordMethod(f func(float64) float64, x0, x1, tol float64, maxIterations int) float64 {
	iterations := 0

	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Iteration", "x0", "x1", "Tolerance")

	for iterations < maxIterations {
		iterations++
		fx0 := f(x0)
		fx1 := f(x1)

		x2 := x1 - (fx1 * (x1 - x0) / (fx1 - fx0))
		tolerance := math.Abs(x2 - x1)

		fmt.Printf("%-12d %-12.6f %-12.6f %-12.6f\n", iterations, x0, x1, tolerance)

		if tolerance < tol {
			return x2
		}

		x0 = x1
		x1 = x2
	}

	fmt.Println("Maximum number of iterations reached.")
	return x1
}
