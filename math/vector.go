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

// An alias for an integer slice
type Vector []int

// Inner product of two vectors
func (a Vector) InnerProduct(b Vector) Vector {
	if len(a) != len(b) {
		return nil
	}
	c := make(Vector, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] * b[i]
	}
	return c
}

