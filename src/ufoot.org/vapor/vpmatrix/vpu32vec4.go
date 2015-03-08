// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpmatrix

// U32Vec4 is a vector containing 4 uint32 values.
// Can hold the values of a pouint in a plane.
type U32Vec4 [4]uint32

// U32Vec4New creates a new vector containing 4 uint32 values.
func U32Vec4New(i1, i2, i3, i4 uint32) *U32Vec4 {
	return &U32Vec4{i1, i2, i3, i4}
}

// Add adds operand to the vector.
// It modifies it, and returns a pouinter on it.
func (vec *U32Vec4) Add(op *U32Vec4) *U32Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// U32Vec4Add adds two vectors.
// Args are left untouched, a pouinter on a new object is returned.
func U32Vec4Add(veca, vecb *U32Vec4) *U32Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}
