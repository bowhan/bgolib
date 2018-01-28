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

package string

import (
	"strconv"
	"math"
	"encoding/binary"
)

// StringLike An interface
type StringLike interface {
	ToBytes() []byte
	ToString() string
	ToUint64() uint64
	ToFloat64() float64
}

// String alias to string, to implement StringLike interface
type String string

// Bytes alias for byte slice, to implement StringLike interface
type Bytes []byte

// ToBytes Convert String to Btes
func (s String) ToBytes() []byte {
	return []byte(s)
}

// ToBytes Convert Bytes to []byte
func (b Bytes) ToBytes() []byte {
	return b
}

// ToString convert string to String
func (s String) ToString() string {
	return string(s)
}

// ToString convert Bytes to string
func (b Bytes) ToString() string {
	return string(b)
}

// ToUint64 convert String to uint64
func (s String) ToUint64() uint64 {
	r, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return r
}

// ToUint64 convert Bytes to uint64
func (b Bytes) ToUint64() uint64 {
	return binary.LittleEndian.Uint64(b)
}

// ToFloat64 convert String to float64
func (s String) ToFloat64() float64 {
	r, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		panic(err)
	}
	return r
}

// ToFloat64 convert Bytes to float64
func (b Bytes) ToFloat64() float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}
