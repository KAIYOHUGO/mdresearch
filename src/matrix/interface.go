package matrix

type (
	Vector []float64
	Matrix []Vector
)

func MultMatrix(a, b Matrix) Matrix {
	return multmatrix(a, b)
}

func CalcFn(a Matrix, b func(float64) float64) Matrix {
	return calcfn(a, b)
}

func VecToMat(a Vector) Matrix {
	return vectomat(a)
}

func MatToVec(a Matrix) Vector {
	return mattovec(a)
}

func T(a Matrix) Matrix {
	return t(a)
}

func NewFn(row, col int, fn func(int, int) float64) Matrix {
	return newfn(row, col, fn)
}

func NewFill(row, col int, value float64) Matrix {
	return newfill(row, col, value)
}

func New(row, col int) Matrix {
	return matnew(row, col)
}

func FillFn(matrix Matrix, fn func(int, int) float64) Matrix {
	return fillfn(matrix, fn)
}
