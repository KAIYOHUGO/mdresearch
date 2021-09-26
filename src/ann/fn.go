package ann

import (
	"math"
	"math/rand"
)

func inithe(in, on int) func(int, int) float64 {
	std := math.Sqrt(2 / float64(in))
	return func(_, _ int) float64 {
		return math.Mod(rand.NormFloat64()*std, 1)
	}
}
