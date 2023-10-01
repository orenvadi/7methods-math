package methods

import (
	"fmt"
	"math"
)

func HalfMethod(f func(float64) float64, a, b, tol float64) float64 {
	if f(a)*f(b) >= 0 {
		fmt.Println("Cannot use bisection method on this interval.")
		return math.NaN()
	}

	iterations := 0
	var c float64

	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Iteration", "a", "b", "Tolerance")

	for {
		c = (a + b) / 2
		iterations++
		fmt.Printf("%-12d %-12.6f %-12.6f %-12.6f\n", iterations, a, b, (b-a)/2)

		if f(c) == 0 || (b-a)/2 < tol {
			break
		}

		if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}

	return c
}
