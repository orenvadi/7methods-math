package main

import (
	"fmt"
	"github.com/orenvadi/7methods/methods"
	"math"
)

// деления пополам, Ньютона, простой итерации, метод хорд, метод секущих,
func f(x float64) float64 {
	return math.Pow(x, 5) - x + 0.2
}

func derivativeF(x float64) float64 {
	return math.Pow(x, 4)*5 - 1
}

func main() {
	a := -2.0 // Lower bound of the interval
	b := 2.0  // Upper bound of the interval
	x0 := (a + b) / 2
	//x0 := a
	tol := 0.0001 // Tolerance
	maxIterations := 20

	fmt.Println(derivativeF(100.0))

	// Half method
	fmt.Println("\n---------------------------Half----------------------------")
	root := methods.HalfMethod(f, a, b, tol)
	fmt.Printf("\nApproximate root: %.6f\n", root)
	fmt.Println("\n-----------------------------------------------------------")

	// Newton method
	fmt.Println("\n---------------------------Newton----------------------------")
	rootNewton := methods.NewtonMethod(f, derivativeF, x0, tol)
	fmt.Printf("\nApproximate root: %.6f\n", rootNewton)
	fmt.Println("\n-------------------------------------------------------------")

	// Simple Iteration method
	fmt.Println("\n---------------------------SimpleIteration----------------------------")
	rootSimpleIter := methods.SimpleIterationMethod(f, x0, tol, maxIterations)
	fmt.Printf("\nApproximate root: %.6f\n", rootSimpleIter)
	fmt.Println("\n----------------------------------------------------------------------")

	// Chord method
	fmt.Println("\n---------------------------Chord----------------------------")
	rootChord := methods.ChordMethod(f, x0, a, tol, maxIterations)
	fmt.Printf("\nApproximate rootChord: %.6f\n", rootChord)
	fmt.Println("\n------------------------------------------------------------")

	// Secant method
	fmt.Println("\n---------------------------Secant----------------------------")
	rootSecant, err := methods.SecantMethod(f, x0, a, tol, maxIterations)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Approximate root: %.8f\n", rootSecant)
	}
	fmt.Println("\n-------------------------------------------------------------")
}
