package ann

import "research/src/matrix"

func newlayers(config *Config) []*Layer {
	layers := make([]*Layer, len(config.Hidden))
	for i := 1; i < len(config.Hidden); i++ {
		layers[i] = newlayer(config.Hidden[i-1], config.Hidden[i])
	}
	layers[0] = newlayer(config.Input, config.Hidden[0])
	layers[len(config.Hidden)-1] = newlayer(config.Hidden[len(config.Hidden)-1], config.Output)
	return layers
}

// nodeN: number of nodes
// weightN: number of weight that every node have
func newlayer(nodeN, weightN uint) *Layer {
	layer := &Layer{
		Weights: make(matrix.Matrix, nodeN),
		Nodes:   make(matrix.Matrix, nodeN),
		Error:   make(matrix.Matrix, nodeN),
	}
	for i := range layer.Weights {
		layer.Weights[i] = make([]float64, weightN)
	}
	return layer
}
