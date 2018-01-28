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

package math

import (
	"errors"
)

type Matrix struct {
	Values []int
	NumRow int
	NumCol int
}

func NewMatrix(m, n int) Matrix {
	return Matrix{
		Values: make([]int, m*n),
		NumRow: m,
		NumCol: n,
	}
}

func (m Matrix) Get(i, j int) (r int, e error) {
	if i >= m.NumRow || j >= m.NumCol {
		e = errors.New("matrix out of range")
	}
	r = m.Values[i*m.NumCol+j]
	return
}

func (m *Matrix) Set(i, j, n int) (e error) {
	if i >= m.NumRow || j >= m.NumCol {
		e = errors.New("matrix out of range")
	}
	m.Values[i*m.NumCol+j] = n
	return nil
}

func (m *Matrix) Increment(s int) {
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		m.Values[i] += s
	}
}

func (m Matrix) AddScalar(s int) Matrix {
	r := NewMatrix(m.NumRow, m.NumCol)
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		r.Values[i] = m.Values[i] + s
	}
	return r
}

func (m Matrix) Map(f func(int) int) Matrix {
	r := NewMatrix(m.NumRow, m.NumCol)
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		r.Values[i] = f(m.Values[i])
	}
	return r
}

func (m Matrix) AddMatrix(n Matrix) (mn Matrix, e error) {
	if m.NumRow != n.NumRow || m.NumCol != n.NumCol {
		e = errors.New("matrix dimension does not match")
		return
	}
	mn = NewMatrix(m.NumRow, n.NumCol)
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		mn.Values[i] = m.Values[i] + n.Values[i]
	}
	return
}

func (m Matrix) MultipleMatrix(n Matrix) (mn Matrix, e error) {

	return
}
