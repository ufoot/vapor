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

// I32Vec3 is a vector containing 3 int32 values.
// Can hold the values of a point in a plane.
type I32Vec3 [3]int32

// I32Vec3New creates a new vector containing 3 int32 values.
func I32Vec3New(i1, i2, i3 int32) *I32Vec3 {
	return &I32Vec3{i1, i2, i3}
}

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *I32Vec3) Add(op *I32Vec3) *I32Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *I32Vec3) Sub(op *I32Vec3) *I32Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// I32Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func I32Vec3Add(veca, vecb *I32Vec3) *I32Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// I32Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func I32Vec3Sub(veca, vecb *I32Vec3) *I32Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}
