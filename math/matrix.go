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

// Matrix Matrix a class to represent two dimensional integer
type Matrix struct {
	Values []int
	NumRow int
	NumCol int
}

// NewMatrix Return a new int m x n Matrix
func NewMatrix(m, n int) Matrix {
	return Matrix{
		Values: make([]int, m*n),
		NumRow: m,
		NumCol: n,
	}
}

// Get Return the [i, j] of Matrix m
func (m Matrix) Get(i, j int) (r int, e error) {
	if i >= m.NumRow || j >= m.NumCol {
		e = errors.New("matrix out of range")
	}
	r = m.Values[i*m.NumCol+j]
	return
}

// Get Return the [i, j] of Matrix m without doing boundary check
func (m Matrix) QuickGet(i, j int) int {
	return m.Values[i*m.NumCol+j]
}

// Set Set the value of [i, j] of Matrix m
func (m *Matrix) Set(i, j, n int) (e error) {
	if i >= m.NumRow || j >= m.NumCol {
		e = errors.New("matrix out of range")
	}
	m.Values[i*m.NumCol+j] = n
	return nil
}

// Increment Increment value element-wise of a matrix
func (m *Matrix) Increment(s int) {
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		m.Values[i] += s
	}
}

// AddScalar Add a scalar to a matrix, return a new matrix
func (m Matrix) AddScalar(s int) Matrix {
	r := NewMatrix(m.NumRow, m.NumCol)
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		r.Values[i] = m.Values[i] + s
	}
	return r
}

// Map Apply a function to a matrix, return a new matrix
func (m Matrix) Map(f func(int) int) Matrix {
	r := NewMatrix(m.NumRow, m.NumCol)
	for i := 0; i < m.NumRow*m.NumCol; i++ {
		r.Values[i] = f(m.Values[i])
	}
	return r
}

// MapSelf Apply a function to each member of a matrix
func (m *Matrix) MapSelf(f func(int) int) {
	for i := 0; i < m.NumRow*m.NumCol; i ++ {
		m.Values[i] = f(m.Values[i])
	}
}

// AddMatrix Add one matrix to another, with the same dimension
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

// MultipleMatrix Multiply matrices
func (m Matrix) MultipleMatrix(n Matrix) (mn Matrix, e error) {
	if m.NumCol != n.NumRow {
		e = errors.New("matrix multiplication requires col matches row")
		return
	}
	mn = NewMatrix(m.NumRow, n.NumCol)
	for i := 0; i < m.NumRow; i ++ {
		for j := 0; j < n.NumCol; j++ {
			sum := 0
			for k := 0; k < m.NumCol; k++ {
				sum += m.QuickGet(i, k) * n.QuickGet(k, j)
			}
			mn.Set(i, j, sum)
		}
	}
	return
}

// Transpose Return the tranposed matrix
func (m Matrix) Transpose() Matrix {
	t := NewMatrix(m.NumCol, m.NumRow)
	for i := 0; i < m.NumRow; i ++ {
		for j := 0; j < m.NumCol; j++ {
			t.Set(j, i, m.QuickGet(i, j))
		}
	}
	return t
}
