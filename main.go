package main

import (
	"fmt"
	"math"

	"github.com/orenvadi/7methods/methods"
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
	// x0 := a
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

	//
	//
	// For algebraic equations

	fmt.Printf("\n\n")
	fmt.Println("+------------------------------------------------------------------------+")
	fmt.Println("|                          Algebraic Equations                           |")
	fmt.Println("+------------------------------------------------------------------------+")
	fmt.Printf("\n\n")
	//
	//

	fmt.Println("\n---------------------------Kramer----------------------------")

	AKram := [][]float64{
		{3, -2, 4},
		{3, 4, -2},
		{2, -1, -1},
	}

	bKram := []float64{21, 9, 10}

	rootsKr, err := methods.KramerMethod(AKram, bKram)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Roots:")
	for i, root := range rootsKr {
		fmt.Printf("x%d = %.2f\n", i+1, root)
	}
	fmt.Println("\n-------------------------------------------------------------")

	fmt.Println("\n---------------------------Gauss----------------------------")
	AGs := [][]float64{
		{3, -2, 4},
		{3, 4, -2},
		{2, -1, -1},
	}

	bGs := []float64{21, 9, 10}

	rootsGs := methods.GaussMethod(AGs, bGs)

	fmt.Println("Roots:")
	for i, root := range rootsGs {
		fmt.Printf("x%d = %.2f\n", i+1, root)
	}
	fmt.Println("\n------------------------------------------------------------")
}
