package matrix

func multmatrix(a, b [][]float64) [][]float64 {
	matrix := make(Matrix, len(a))
	for arow := range a {
		matrix[arow] = make(Vector, len(b[0]))
		for acol := range a[0] {
			for bcol := range b[0] {
				for brow := range b {
					matrix[arow][bcol] += a[arow][acol] * b[brow][bcol]
				}
			}
		}
	}
	return matrix
}

func calcfn(a [][]float64, b func(float64) float64) [][]float64 {
	matrix := make(Matrix, len(a))
	for r, row := range a {
		matrix[r] = make(Vector, len(a[0]))
		for c, value := range row {
			matrix[r][c] = b(value)
		}
	}
	return matrix
}
