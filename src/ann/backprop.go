package ann

import "research/src/matrix"

func (s *ANN) predict() {
	for i, layer := range s.layer {
		value := matrix.MultMatrix(s.layer[i-1].Nodes, layer.Weights)
		if layer.Act {
			matrix.CalcFn(value, s.api.ActFn)
		}
		layer.Nodes = value
	}
}
