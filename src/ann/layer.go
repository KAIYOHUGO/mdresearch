package ann

func newlayers(config *Config) []*Layer {
	layers := make([]*Layer, len(config.Hidden))
	for i := 1; i < len(config.Hidden); i++ {
		layers[i] = newlayer(config.Hidden[i-1], config.Hidden[i], true)
	}
	layers[0] = newlayer(config.Input, config.Hidden[0], true)
	layers[len(config.Hidden)-1] = newlayer(config.Hidden[len(config.Hidden)-1], config.Output, false)
	return layers
}

func newlayer(nodeN, weightN uint, act bool) *Layer {
	layer := &Layer{
		Weights: make([][]float64, weightN),
		Nodes:   make([][]float64, nodeN),
		Act:     act,
	}
	for i := range layer.Weights {
		layer.Weights[i] = make([]float64, weightN)
	}
	return layer
}
