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

func TestMatrix_MultipleMatrix(t *testing.T) {
	u := NewMatrix(2, 2)
	u.Set(0, 0, 1)
	u.Set(0, 1, 2)
	u.Set(1, 0, 3)
	u.Set(1, 1, 4)
	v := NewMatrix(2, 2)
	v.Set(0, 0, 5)
	v.Set(0, 1, 6)
	v.Set(1, 0, 7)
	v.Set(1, 1, 8)
	uv, _ := u.MultipleMatrix(v)
	vu, _ := v.MultipleMatrix(u)
	if uv.QuickGet(0, 0) != 19 ||
		uv.QuickGet(0, 1) != 22 ||
		uv.QuickGet(1, 0) != 43 ||
		uv.QuickGet(1, 1) != 50 {
		t.Error("matrix multiplication wrong")
	}
	if vu.QuickGet(0, 0) != 23 ||
		vu.QuickGet(0, 1) != 34 ||
		vu.QuickGet(1, 0) != 31 ||
		vu.QuickGet(1, 1) != 46 {
		t.Error("matrix multiplication wrong")
	}

	r := NewMatrix(3, 1)
	_, e := u.MultipleMatrix(r)
	if e == nil {
		t.Error("error expected from unmatched matrix multiplication")
	}
}

func TestMatrix_MultipleMatrix2(t *testing.T) {
	a := NewMatrix(2, 3)
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(0, 2, 3)
	a.Set(1, 0, 4)
	a.Set(1, 1, 5)
	a.Set(1, 2, 6)
	b := NewMatrix(3, 2)
	b.Set(0, 0, 1)
	b.Set(0, 1, 2)
	b.Set(1, 0, 3)
	b.Set(1, 1, 4)
	b.Set(2, 0, 5)
	b.Set(2, 1, 6)
	ab, _ := a.MultipleMatrix(b)
	ba, _ := b.MultipleMatrix(a)
	if ab.NumRow != 2 || ab.NumCol != 2 || ba.NumRow != 3 || ba.NumCol != 3 {
		t.Error("matrix multiplication produces wrong dimention")
	}
	if ab.QuickGet(0, 0) != 22 ||
		ab.QuickGet(0, 1) != 28 ||
		ab.QuickGet(1, 0) != 49 ||
		ab.QuickGet(1, 1) != 64 {
		t.Error("matrix multiplication wrong")
	}
	if ba.QuickGet(0, 0) != 9 ||
		ba.QuickGet(0, 1) != 12 ||
		ba.QuickGet(0, 2) != 15 ||
		ba.QuickGet(1, 0) != 19 ||
		ba.QuickGet(1, 1) != 26 ||
		ba.QuickGet(1, 2) != 33 ||
		ba.QuickGet(2, 0) != 29 ||
		ba.QuickGet(2, 1) != 40 ||
		ba.QuickGet(2, 2) != 51 {
		t.Error("matrix multiplication wrong")
	}
}

func TestMatrix_Transpose(t *testing.T) {
	a := NewMatrix(2, 3)
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(0, 2, 3)
	a.Set(1, 0, 4)
	a.Set(1, 1, 5)
	a.Set(1, 2, 6)
	// 1 2 3
	// 4 5 6
	b := a.Transpose()
	// 1 4
	// 2 5
	// 3 6
	if b.NumRow != 3 || b.NumCol != 2 {
		t.Error("tranpose dimention wrong")
	}
	if b.QuickGet(0, 0) != 1 ||
		b.QuickGet(0, 1) != 4 ||
		b.QuickGet(1, 0) != 2 ||
		b.QuickGet(1, 1) != 5 ||
		b.QuickGet(2, 0) != 3 ||
		b.QuickGet(2, 1) != 6 {
		t.Error("transpose wrong")
	}
}
