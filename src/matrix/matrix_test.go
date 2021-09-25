package matrix

import (
	"reflect"
	"testing"
)

func TestMult(t *testing.T) {
	r := MultMatrix([][]float64{
		{1},
		{2},
	},
		[][]float64{
			{1, 2, 3},
			{6, 2, 3},
		})
	t.Log(r)
	if !reflect.DeepEqual(r, [][]float64{{7, 4, 6}, {14, 8, 12}}) {
		t.Fatal("should eq")
	}
}

func TestVecToMat(t *testing.T) {
	o := []float64{
		2, 0, 0, 4, 1, 2, 1, 4,
	}
	r := VecToMat(o)
	t.Log(r)
	if !reflect.DeepEqual(r[0], o) {
		t.Fatal("should eq")
	}
}
