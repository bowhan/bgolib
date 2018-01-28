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