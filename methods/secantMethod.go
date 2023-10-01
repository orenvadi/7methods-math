package methods

import (
	"fmt"
	"math"
)

func SecantMethod(f func(float64) float64, x0, x1, tol float64, maxIter int) (float64, error) {
	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Iteration", "x", "f(x)", "Tolerance")
	for i := 0; i < maxIter; i++ {
		fx0 := f(x0)
		fx1 := f(x1)

		if math.Abs(fx1) < tol {
			fmt.Printf("\nConverged with tolerance: %e\n", tol)
			return x1, nil
		}

		x2 := x1 - (fx1 * (x1 - x0) / (fx1 - fx0))

		fmt.Printf("%-12d %-12.6f %-12.6f %-12.6f\n", i, x2, fx1, math.Abs(x2-x1))

		if math.Abs(x2-x1) < tol {
			fmt.Printf("\nConverged with tolerance: %e\n", tol)
			return x2, nil
		}

		x0 = x1
		x1 = x2
	}

	return 0, fmt.Errorf("Maximum number of iterations reached")
}
