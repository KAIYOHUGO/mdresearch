package ann

import "research/src/matrix"

func newann(config *Config, interfaces IANN) *ANN {
	r := &ANN{
		interfaces: interfaces,
		config:     config,
		layer:      newlayers(config),
	}
	r.init()
	return r
}

func (s *ANN) init() {
	for _, layer := range s.layer {
		matrix.FillFn(layer.Weights, s.interfaces.Init)
	}
}

func (s *ANN) predict(value []float64) []float64 {
	s.layer[0].Nodes = matrix.CalcFn(matrix.T(matrix.MultMatrix(s.layer[0].Weights, matrix.VecToMat(value))), s.interfaces.ActFn)
	for i := 1; i < len(s.layer)-1; i++ {
		s.layer[i].Nodes = matrix.CalcFn(matrix.T(matrix.MultMatrix(s.layer[i].Weights, s.layer[i-1].Nodes)), s.interfaces.ActFn)
	}
	s.layer[len(s.layer)-1].Nodes = matrix.T(matrix.MultMatrix(s.layer[0].Weights, matrix.VecToMat(value)))
	return s.layer[len(s.layer)-1].Nodes[0]
}

func (s *ANN) loss(datas []Data) []float64 {
	losses := make([]float64, len(datas))
	for i, data := range datas {
		losses[i] = s.calcloss(s.predict(data.Input), data.Output)
	}
	return losses
}

func (s *ANN) calcloss(predict, desire []float64) float64 {
	return s.interfaces.LosFn(predict, desire)
}
