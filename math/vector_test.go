package math

import (
	"testing"
)

//
//var (
//	a Vector
//	b Vector
//	c Vector
//)

var (
	a = Vector{1, 2, 3}
	b = Vector{4, 5, 6}
	c = Vector{5, 6}
)

func TestVector_Dot(t *testing.T) {
	ab := a.InnerProduct(b)
	if ab[0] != 4 || ab[1] != 10 || ab[2] != 18 {
		t.Error("Error a . b")
	}
	ac := a.InnerProduct(c)
	if ac != nil {
		t.Error("Error processing mismatch")
	}
}
