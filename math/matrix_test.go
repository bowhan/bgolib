package math

import "testing"

var (
	mx1 = Matrix{Values: make([]int, 9), NumRow: 3, NumCol: 3}
)

func init() {
	mx1.Set(0, 0, 0)
	mx1.Set(0, 1, 1)
	mx1.Set(0, 2, 2)
	mx1.Set(1, 0, 3)
	mx1.Set(1, 1, 4)
	mx1.Set(1, 2, 5)
	mx1.Set(2, 0, 6)
	mx1.Set(2, 1, 7)
	mx1.Set(2, 2, 8)
}

func TestMatrix_Get(t *testing.T) {
	k := 0
	for i := 0; i < mx1.NumRow; i++ {
		for j := 0; j < mx1.NumCol; j++ {
			v, e := mx1.Get(i, j)
			if e != nil {
				t.Error("Error should be nil")
			}
			if v != k {
				t.Error("Get value wrong")
			}
			k++
		}
	}
}

func TestNewMatrix(t *testing.T) {
	a := NewMatrix(3, 2)
	if a.NumRow != 3 || a.NumCol != 2 || len(a.Values) != 6 {
		t.Error("matrix dimention wrong")
	}
}

func TestMatrix_AddScalar(t *testing.T) {
	mx11 := mx1.AddScalar(3)
	v, _ := mx11.Get(0, 0)
	if v != 3 {
		t.Error("Matrix and scalar addition wrong")
	}
}



func TestMatrix_AddMatrix(t *testing.T) {
	mx2, _ := mx1.AddMatrix(mx1)
	k := 0
	for i := 0; i < mx1.NumRow; i ++ {
		for j := 0; j < mx1.NumCol; j++ {
			if v, e := mx2.Get(i, j); v != k || e != nil {
				t.Error("Error adding two matrices")
			}
			k += 2
		}
	}
}
