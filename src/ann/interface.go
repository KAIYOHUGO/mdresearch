package ann

import "research/src/matrix"

// interface to init
type IANN interface {
	// init weight
	// layer,node,toNode=>init weight
	Init(int, int) float64
	// active node, like relu or sigmoid
	ActFn(float64) float64
	// calc loss value, like mse
	// predict,desire=>loss value
	LosFn([]float64, []float64) float64
	// for optimizatio, like sgd or adam
	OptFn()
}

// config to init
type Config struct {
	// input layer node number
	Input uint
	// hidden layer node number
	Hidden []uint
	// output layer node number
	Output uint
}

// ANN model.
type ANN struct {
	// interfaces store interface from init.
	interfaces IANN
	// config store setting from init.
	config *Config
	// layer store nn model
	layer []*Layer
}

// for model trai & test data
type Data struct {
	Input  []float64
	Output []float64
}

// --------\
// -weight----nodes
// --------/

// NN Layer
type Layer struct {
	Weights matrix.Matrix
	// ! nodes should always nodes[0][...]
	Nodes matrix.Matrix
	// ! error should always error[0][...]
	Error matrix.Matrix
}

// create new ANN network
func New(config *Config, interfaces IANN) *ANN {
	return newann(config, interfaces)
}

// TODO make this work
// train model
func (s *ANN) Train(n, groupsize uint, data []Data) {

}

// just run the model & get the predict
func (s *ANN) Predict(value []float64) []float64 {
	return s.predict(value)
}

// calc loss by data & get loss value list
func (s *ANN) Loss(data []Data) []float64 {
	return s.loss(data)
}
