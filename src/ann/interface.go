package ann

type IANN interface {
	ActFn(float64) float64
	LosFn()
	OptFn()
}

type Config struct {
	Input  uint
	Hidden []uint
	Output uint
}

type ANN struct {
	api    IANN
	config *Config
	layer  []*Layer
}

type Data struct {
	Input  []float64
	Output []float64
}

type Layer struct {
	Weights [][]float64
	// nodes should always nodes[0][...]
	Nodes [][]float64
	Act   bool
}

func New(config *Config, api IANN) *ANN {
	return &ANN{
		api:    api,
		config: config,
		layer:  newlayers(config),
	}
}

func (s *ANN) Train(n uint, data []Data) {

}

func (s *ANN) Predict(value []float64) {

}

func (s *ANN) Loss(data []Data) {

}
