package matrix

import (
	"reflect"
	"testing"
)

func TestMult(t *testing.T) {
	r := MultMatrix(
		Matrix{
			{1},
			{2},
		},
		Matrix{
			{1, 2, 3},
			{6, 2, 3},
		})
	t.Log(r)
	if !reflect.DeepEqual(r, Matrix{{7, 4, 6}, {14, 8, 12}}) {
		t.Fatal("should eq")
	}
}

func TestVecToMat(t *testing.T) {
	o := Vector{
		2, 0, 0, 4, 1, 2, 1, 4,
	}
	r := VecToMat(o)
	t.Log(r)
	if !reflect.DeepEqual(r[0], o) {
		t.Fatal("should eq")
	}
}

func TestT(t *testing.T) {
	o := Matrix{
		{
			6, 2, 3,
		},
	}
	r := T(T(o))
	t.Log(r)
	t.Log(o)
	if !reflect.DeepEqual(r, o) {
		t.Fatal("should eq")
	}
}
