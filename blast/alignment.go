package alignment

import (
    "strconv"
)

type Alignment struct {
	QseqID   []byte
	SseqID   []byte
	Pident   float64
	Length   int
	Mismatch int
	Gapopen  int
	Qstart   int
	Qend     int
	Sstart   int
	Send     int
	Evalue   float64
	Bitscore float64
	Nreads   int
	Strand   byte
}

func NewAlignment(nreads string, args ... string) *Alignment {
	aln := Alignment{
		QseqID: []byte(args[0]),
		SseqID: []byte(args[1]),
	}
	aln.Pident, _ = strconv.ParseFloat(args[2], 64)
	aln.Length, _ = strconv.Atoi(args[3])
	aln.Mismatch, _ = strconv.Atoi(args[4])
	aln.Gapopen, _ = strconv.Atoi(args[5])
	aln.Qstart, _ = strconv.Atoi(args[6])
	aln.Qend, _ = strconv.Atoi(args[7])
	aln.Sstart, _ = strconv.Atoi(args[8])
	aln.Send, _ = strconv.Atoi(args[9])
	aln.Evalue, _ = strconv.ParseFloat(args[10], 64)
	aln.Bitscore, _ = strconv.ParseFloat(args[11], 64)
	aln.Nreads, _ = strconv.Atoi(nreads)
	if aln.Sstart < aln.Send {
		aln.Strand = '+'
	} else {
		aln.Strand = '-'
		t := aln.Sstart
		aln.Sstart = aln.Send
		aln.Send = t
	}
	return &aln
}
