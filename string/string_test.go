package string

import (
	"testing"
	"bytes"
)

func TestToBytes(t *testing.T) {
	str := String("hello")
	bys := []byte("hello")
	if bytes.Compare(str.to_bytes(), bys) != 0 {
		t.Error("Test failed")
	}
}
