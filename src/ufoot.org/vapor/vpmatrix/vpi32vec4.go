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

import (
	"ufoot.org/vapor/vpnumber"
)

// I32Vec4 is a vector containing 4 int32 values.
// Can hold the values of a point in a plane.
type I32Vec4 [4]int32

// I32Vec4New creates a new vector containing 4 int32 values.
func I32Vec4New(i1, i2, i3, i4 int32) *I32Vec4 {
	return &I32Vec4{i1, i2, i3, i4}
}

// ToI64 converts the vector to an int64 vector.
func (vec *I32Vec4) ToI64() *I64Vec4 {
	var ret I64Vec4

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *I32Vec4) ToX32() *X32Vec4 {
	var ret X32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.I32ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *I32Vec4) ToX64() *X64Vec4 {
	var ret X64Vec4

	for i, v := range vec {
		ret[i] = vpnumber.I32ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *I32Vec4) ToF32() *F32Vec4 {
	var ret F32Vec4

	for i, v := range vec {
		ret[i] = float32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *I32Vec4) ToF64() *F64Vec4 {
	var ret F64Vec4

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I32Vec4) Add(op *I32Vec4) *I32Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I32Vec4) Sub(op *I32Vec4) *I32Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// I32Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func I32Vec4Add(veca, vecb *I32Vec4) *I32Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// I32Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func I32Vec4Sub(veca, vecb *I32Vec4) *I32Vec4 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}
