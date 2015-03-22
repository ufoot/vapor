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

// I64Vec3 is a vector containing 3 int64 values.
// Can hold the values of a point in a plane.
type I64Vec3 [3]int64

// I64Vec3New creates a new vector containing 3 int64 values.
func I64Vec3New(i1, i2, i3 int64) *I64Vec3 {
	return &I64Vec3{i1, i2, i3}
}

// ToI32 converts the vector to an int32 vector.
func (vec *I64Vec3) ToI32() *I32Vec3 {
	var ret I32Vec3

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *I64Vec3) ToX32() *X32Vec3 {
	var ret X32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.I64ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *I64Vec3) ToX64() *X64Vec3 {
	var ret X64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.I64ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *I64Vec3) ToF32() *F32Vec3 {
	var ret F32Vec3

	for i, v := range vec {
		ret[i] = float32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *I64Vec3) ToF64() *F64Vec3 {
	var ret F64Vec3

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I64Vec3) Add(op *I64Vec3) *I64Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I64Vec3) Sub(op *I64Vec3) *I64Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *I64Vec3) Neg() *I64Vec3 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// I64Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func I64Vec3Add(veca, vecb *I64Vec3) *I64Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// I64Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func I64Vec3Sub(veca, vecb *I64Vec3) *I64Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// I64Vec3Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func I64Vec3Neg(vec *I64Vec3) *I64Vec3 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}
