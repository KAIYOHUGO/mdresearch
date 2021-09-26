package ann

import (
	"research/src/matrix"
)

func (s *ANN) backproperror(e matrix.Vector) {
	s.layer[len(s.layer)-1].Error = matrix.VecToMat(e)
	for i := len(s.layer) - 1; i > 0; i-- {
		s.layer[i-1].Error = matrix.T(matrix.MultMatrix(s.layer[i].Weights, s.layer[i].Error))
	}
}

// ? how calc partical
func (s *ANN) backprop() {
	// TODO add backprop the weight
}
