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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpvec2

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
)

// I32 is a vector containing 2 int32 values.
// Can hold the values of a point in a plane.
type I32 [Size]int32

// I32New creates a new vector containing 2 int32 values.
func I32New(i1, i2 int32) *I32 {
	return &I32{i1, i2}
}

// ToI64 converts the vector to an int64 vector.
func (vec *I32) ToI64() *I64 {
	var ret I64

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *I32) ToX32() *X32 {
	var ret X32

	for i, v := range vec {
		ret[i] = vpnumber.I32ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *I32) ToX64() *X64 {
	var ret X64

	for i, v := range vec {
		ret[i] = vpnumber.I32ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *I32) ToF32() *F32 {
	var ret F32

	for i, v := range vec {
		ret[i] = float32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *I32) ToF64() *F64 {
	var ret F64

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// String returns a readable form of the vector.
func (vec *I32) String() string {
	buf, err := json.Marshal(vec)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I32) Add(op *I32) *I32 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I32) Sub(op *I32) *I32 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *I32) Neg() *I32 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// I32Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func I32Add(veca, vecb *I32) *I32 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// I32Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func I32Sub(veca, vecb *I32) *I32 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// I32Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func I32Neg(vec *I32) *I32 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}
