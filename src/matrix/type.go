package matrix

func vectomat(a []float64) Matrix {
	matrix := make(Matrix, 1)
	matrix[0] = a
	return matrix
}

func t(a Matrix) Matrix {
	matrix := make(Matrix, len(a[0]))
	for i := range matrix {
		matrix[i] = make(Vector, len(a))
	}
	for row := range a {
		for col := range a[0] {
			matrix[col][row] = a[row][col]
		}
	}
	return matrix
}
