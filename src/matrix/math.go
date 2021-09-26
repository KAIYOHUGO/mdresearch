package matrix

func multmatrix(a, b Matrix) Matrix {
	matrix := New(len(a), len(b[0]))
	for arow := range a {
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

func calcfn(a Matrix, b func(float64) float64) Matrix {
	matrix := New(len(a), len(a[0]))
	for r, row := range a {
		for c, value := range row {
			matrix[r][c] = b(value)
		}
	}
	return matrix
}
