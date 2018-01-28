/*
 * Copyright (c) Bo Han 2018.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package alignment

import (
	"strconv"
)


/// Alignment from Blast fmt6
type Alignment struct {
	QseqID   string
	SseqID   string
	Pident   float64
	Length   uint64
	Mismatch uint64
	Gapopen  uint64
	Qstart   uint64
	Qend     uint64
	Sstart   uint64
	Send     uint64
	Evalue   float64
	Bitscore float64
	Strand   byte
}

/// return a pointer to an Alignment constructed from strings
func NewAlignment(args ...string) *Alignment {
	aln := Alignment{
		QseqID: args[0],
		SseqID: args[1],
	}
	aln.Pident, _ = strconv.ParseFloat(args[2], 64)
	aln.Length, _ = strconv.ParseUint(args[3], 10, 64)
	aln.Mismatch, _ = strconv.ParseUint(args[4], 10, 64)
	aln.Gapopen, _ = strconv.ParseUint(args[5], 10, 64)
	aln.Qstart, _ = strconv.ParseUint(args[6], 10, 64)
	aln.Qend, _ = strconv.ParseUint(args[7], 10, 64)
	aln.Sstart, _ = strconv.ParseUint(args[8], 10, 64)
	aln.Send, _ = strconv.ParseUint(args[9], 10, 64)
	aln.Evalue, _ = strconv.ParseFloat(args[10], 64)
	aln.Bitscore, _ = strconv.ParseFloat(args[11], 64)
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

/// return a pointer to an Alignment constructed from bytes
func NewAlignmentFromBytes(args ...[]byte) *Alignment {
	aln := Alignment{
		QseqID:   string(args[0]),
		SseqID:   string(args[1]),
		Length:   bytesToUint64(args[3]),
		Mismatch: bytesToUint64(args[4]),
		Gapopen:  bytesToUint64(args[5]),
		Qstart:   bytesToUint64(args[6]),
		Qend:     bytesToUint64(args[7]),
		Sstart:   bytesToUint64(args[8]),
		Send:     bytesToUint64(args[9]),
	}
	aln.Pident, _ = strconv.ParseFloat(string(args[2]), 64)
	aln.Evalue, _ = strconv.ParseFloat(string(args[10]), 64)
	aln.Bitscore, _ = strconv.ParseFloat(string(args[11]), 64)
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

func bytesToUint64(b []byte) uint64 {
	var r uint64
	for i := 0; i < len(b) ; i++ {
		r = 10*r + uint64(b[i]-'0')
	}
	return r
}
