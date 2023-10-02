package methods

func GaussMethod(matrix [][]float64, constants []float64) []float64 {
	n := len(matrix)
	solution := make([]float64, n)

	// Step 1: Forward Elimination
	for i := 0; i < n; i++ {
		// Find the row with the maximum element in the current column
		maxRow := i
		for j := i + 1; j < n; j++ {
			if matrix[j][i] > matrix[maxRow][i] {
				maxRow = j
			}
		}

		// Swap the current row with the row containing the maximum element
		matrix[i], matrix[maxRow] = matrix[maxRow], matrix[i]
		constants[i], constants[maxRow] = constants[maxRow], constants[i]

		// Make the diagonal element of the current row equal to 1
		pivot := matrix[i][i]
		for j := i; j < n; j++ {
			matrix[i][j] /= pivot
		}
		constants[i] /= pivot

		// Eliminate all other elements in the current column
		for j := i + 1; j < n; j++ {
			factor := matrix[j][i]
			for k := i; k < n; k++ {
				matrix[j][k] -= factor * matrix[i][k]
			}
			constants[j] -= factor * constants[i]
		}
	}

	// Step 2: Back Substitution
	for i := n - 1; i >= 0; i-- {
		// Solve for the variable at the current row
		solution[i] = constants[i]
		for j := i + 1; j < n; j++ {
			solution[i] -= matrix[i][j] * solution[j]
		}
	}

	return solution
}
