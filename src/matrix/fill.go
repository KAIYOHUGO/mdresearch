package matrix

func fillfn(matrix Matrix, fn func(int, int) float64) Matrix {
	for ri := range matrix {
		for ci := range matrix[ri] {
			matrix[ri][ci] = fn(ri, ci)
		}
	}
	return matrix
}
