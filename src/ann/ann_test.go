package ann

import (
	"testing"
)

type TestModel struct {
}

func TestPredict(t *testing.T) {
	// New(&Config{})
}

func TestHe(t *testing.T) {
	l := make([]int, 19)
	for i := 0; i < 1000; i++ {
		v := inithe(1, 10)(0, 0)
		if (v > 1) || (v < -1) {
			t.Fatal(v)
		}
		l[9+int(v*10)]++
	}
	for i, v := range l {
		t.Log(i-9, v)
	}
}
