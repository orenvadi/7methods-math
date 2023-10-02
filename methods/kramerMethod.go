package methods

import "fmt"

// determinant calculates the determinant of a square matrix.
func determinant(matrix [][]float64) float64 {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0.0
	for i := 0; i < n; i++ {
		subMatrix := make([][]float64, n-1)
		for j := range subMatrix {
			subMatrix[j] = make([]float64, n-1)
		}
		for j := 1; j < n; j++ {
			for k := 0; k < n; k++ {
				if k < i {
					subMatrix[j-1][k] = matrix[j][k]
				} else if k > i {
					subMatrix[j-1][k-1] = matrix[j][k]
				}
			}
		}
		sign := 1.0
		if i%2 == 1 {
			sign = -1.0
		}
		det += sign * matrix[0][i] * determinant(subMatrix)
	}
	return det
}

// cramerMethod solves a system of linear equations using Cramer's Rule.
func KramerMethod(A [][]float64, b []float64) ([]float64, error) {
	n := len(A)
	if n == 0 || len(A[0]) != n || len(b) != n {
		return nil, fmt.Errorf("input is not a square matrix or incompatible dimensions")
	}

	// Calculate the determinant of the coefficient matrix A
	detA := determinant(A)

	if detA == 0 {
		return nil, fmt.Errorf("the determinant of the coefficient matrix is zero, Cramer's Rule cannot be applied")
	}

	// Initialize a slice to store the solutions
	x := make([]float64, n)

	// Calculate the determinants of matrices obtained by replacing each column with the constants vector b
	for i := 0; i < n; i++ {
		// Create a copy of matrix A to avoid modifying the original
		A_copy := make([][]float64, n)
		for j := range A_copy {
			A_copy[j] = make([]float64, n)
			copy(A_copy[j], A[j])
		}

		// Replace the i-th column with the constants vector b
		for j := 0; j < n; j++ {
			A_copy[j][i] = b[j]
		}

		// Calculate the determinant of the modified matrix and store it as the solution
		x[i] = determinant(A_copy) / detA
	}

	return x, nil
}
