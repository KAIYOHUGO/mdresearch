package matrix

type (
	Matrix [][]float64
	Vector []float64
)

func MultMatrix(a, b Matrix) Matrix {
	return multmatrix(a, b)
}

func CalcFn(a Matrix, b func(float64) float64) Matrix {
	return calcfn(a, b)
}

func VecToMat(a []float64) Matrix {
	return vectomat(a)
}

func T(a Matrix) Matrix {
	return t(a)
}
