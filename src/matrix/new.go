package matrix

func matnew(row, col int) Matrix {
	matrix := make(Matrix, row)
	for i := range matrix {
		matrix[i] = make(Vector, col)
	}
	return matrix
}

func newfill(row, col int, value float64) Matrix {
	matrix := New(row, col)
	for ri := range matrix {
		for ci := range matrix[ri] {
			matrix[ri][ci] = value
		}
	}
	return matrix
}

func newfn(row, col int, fn func(int, int) float64) Matrix {
	matrix := New(row, col)
	return FillFn(matrix, fn)
}
