package string

import (
	"strconv"
	"math"
	"encoding/binary"
)

type StringLike interface {
	to_bytes() []byte
	to_string() string
	to_uint64() uint64
	to_float64() float64
}

type String string
type Bytes []byte

func (s String) to_bytes() []byte {
	return []byte(s)
}

func (b Bytes) to_bytes() []byte {
	return b
}

func (s String) to_string() string {
	return string(s)
}

func (b Bytes) to_string() string {
	return string(b)
}

func (s String) to_uint64() uint64 {
	r, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return r
}

func (b Bytes) to_uint64() uint64 {
	return binary.LittleEndian.Uint64(b)
}

func (s String) to_float64() float64 {
	r, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		panic(err)
	}
	return r
}

func (b Bytes) to_float64() float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}

//func NewAlignment(args []StringLike) *Alignment {
//	aln := Alignment{
//		QseqID:   args[0].to_string(),
//		SseqID:   args[1].to_string(),
//		Pident:   args[2].to_float64(),
//		Length:   args[3].to_uint64(),
//		Mismatch: args[4].to_uint64(),
//		Gapopen:  args[5].to_uint64(),
//		Qstart:   args[6].to_uint64(),
//		Qend:     args[7].to_uint64(),
//		Sstart:   args[8].to_uint64(),
//		Send:     args[9].to_uint64(),
//		Evalue:   args[10].to_float64(),
//		Bitscore: args[11].to_float64(),
//	}
//	if aln.Sstart < aln.Send {
//		aln.Strand = '+'
//	} else {
//		aln.Strand = '-'
//		t := aln.Sstart
//		aln.Sstart = aln.Send
//		aln.Send = t
//	}
//	return &aln
//}
