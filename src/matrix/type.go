package matrix

func vectomat(a Vector) Matrix {
	matrix := make(Matrix, 1)
	matrix[0] = a
	return matrix
}

func mattovec(a Matrix) Vector {
	vector := make(Vector, 0)
	for _, row := range a {
		vector = append(vector, row...)
	}
	return vector[:len(vector):len(vector)]
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
