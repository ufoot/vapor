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

package vpvec4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
	"math"
)

// F32 is a vector containing 4 float32 values.
// Can be used in 3D matrix transformations.
type F32 [Size]float32

// F32New creates a new vector containing 4 float32 values.
func F32New(f1, f2, f3, f4 float32) *F32 {
	return &F32{f1, f2, f3, f4}
}

// F32AxisX returns a new vector representing the X axis.
func F32AxisX() *F32 {
	return &F32{vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0}
}

// F32AxisY returns a new vector representing the Y axis.
func F32AxisY() *F32 {
	return &F32{vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0}
}

// F32AxisZ returns a new vector representing the Z axis.
func F32AxisZ() *F32 {
	return &F32{vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0}
}

// F32FromVec3 creates a new vector from a smaller one,
// by appending a value at its end.
func F32FromVec3(vec *vpvec3.F32, f float32) *F32 {
	return &F32{vec[0], vec[1], vec[2], f}
}

// ToVec3 creates a smaller vector by removing the last value.
func (vec *F32) ToVec3() *vpvec3.F32 {
	return &vpvec3.F32{vec[0], vec[1], vec[2]}
}

// ToI32 converts the vector to an int32 vector.
func (vec *F32) ToI32() *I32 {
	var ret I32

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int64 vector.
func (vec *F32) ToI64() *I64 {
	var ret I64

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *F32) ToX32() *X32 {
	var ret X32

	for i, v := range vec {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *F32) ToX64() *X64 {
	var ret X64

	for i, v := range vec {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *F32) ToF64() *F64 {
	var ret F64

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// String returns a readable form of the vector.
func (vec *F32) String() string {
	buf, err := json.Marshal(vec)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F32) Add(op *F32) *F32 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F32) Sub(op *F32) *F32 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F32) Neg() *F32 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F32) MulScale(factor float32) *F32 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F32) DivScale(factor float32) *F32 {
	for i, v := range vec {
		vec[i] = vpnumber.F32Div(v, factor)
	}

	return vec
}

// Lerp performs a linear interpolation with another vector.
func (vec *F32) Lerp(op *F32, beta float32) *F32 {
	switch {
	case beta <= vpnumber.F32Const0:
		return vec
	case beta >= vpnumber.F32Const1:
		*vec = *op
		return vec
	}

	vec.MulScale(vpnumber.F32Const1 - beta)
	vec.Add(F32MulScale(op, beta))

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F32) SqMag() float32 {
	var sq float32

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F32) Length() float32 {
	return float32(math.Sqrt(float64(vec.SqMag())))
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *F32) Normalize() *F32 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *F32) IsSimilar(op *F32) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
func (vec *F32) Dot(op *F32) float32 {
	var dot float32

	for i, v := range op {
		dot += vec[i] * v
	}

	return dot
}

// F32Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F32Add(veca, vecb *F32) *F32 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F32Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F32Sub(veca, vecb *F32) *F32 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F32Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func F32Neg(vec *F32) *F32 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// F32MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32MulScale(vec *F32, factor float32) *F32 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F32DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32DivScale(vec *F32, factor float32) *F32 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F32Lerp performs a linear interpolation between 2 vectors.
func F32Lerp(veca, vecb *F32, beta float32) *F32 {
	var ret = *veca

	ret.Lerp(vecb, beta)

	return &ret
}

// F32Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F32Normalize(vec *F32) *F32 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}
