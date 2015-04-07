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
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
)

// X64 is a vector containing 4 fixed point 64 bit values.
// Can be used in 3D matrix transformations.
type X64 [Size]vpnumber.X64

// X64New creates a new vector containing 4 fixed point 64 bit values.
func X64New(x1, x2, x3, x4 vpnumber.X64) *X64 {
	return &X64{x1, x2, x3, x4}
}

// X64AxisX returns a new vector representing the X axis.
func X64AxisX() *X64 {
	return &X64{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0}
}

// X64AxisY returns a new vector representing the Y axis.
func X64AxisY() *X64 {
	return &X64{vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0}
}

// X64AxisZ returns a new vector representing the Z axis.
func X64AxisZ() *X64 {
	return &X64{vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0}
}

// X64FromVec3 creates a new vector from a smaller one,
// by appending a value at its end.
func X64FromVec3(vec *vpvec3.X64, x vpnumber.X64) *X64 {
	return &X64{vec[0], vec[1], vec[2], x}
}

// ToVec3 creates a smaller vector by removing the last value.
func (vec *X64) ToVec3() *vpvec3.X64 {
	return &vpvec3.X64{vec[0], vec[1], vec[2]}
}

// ToI32 converts the vector to an int32 vector.
func (vec *X64) ToI32() *I32 {
	var ret I32

	for i, v := range vec {
		ret[i] = vpnumber.X64ToI32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int32 vector.
func (vec *X64) ToI64() *I64 {
	var ret I64

	for i, v := range vec {
		ret[i] = vpnumber.X64ToI64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 64 bits.
func (vec *X64) ToX32() *X32 {
	var ret X32

	for i, v := range vec {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *X64) ToF32() *F32 {
	var ret F32

	for i, v := range vec {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *X64) ToF64() *F64 {
	var ret F64

	for i, v := range vec {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// String returns a readable form of the vector.
func (vec *X64) String() string {
	buf, err := json.Marshal(vec.ToF64())

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X64) Add(op *X64) *X64 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X64) Sub(op *X64) *X64 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X64) Neg() *X64 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X64) MulScale(factor vpnumber.X64) *X64 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X64) DivScale(factor vpnumber.X64) *X64 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Div(v, factor)
	}

	return vec
}

// Lerp performs a linear interpolation with another vector.
func (vec *X64) Lerp(op *X64, beta vpnumber.X64) *X64 {
	switch {
	case beta <= vpnumber.X64Const0:
		return vec
	case beta >= vpnumber.X64Const1:
		*vec = *op
		return vec
	}

	vec.MulScale(vpnumber.X64Const1 - beta)
	vec.Add(X64MulScale(op, beta))

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X64) SqMag() vpnumber.X64 {
	var sq vpnumber.X64

	for _, v := range vec {
		sq += vpnumber.X64Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X64) Length() vpnumber.X64 {
	return vpmath.X64Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *X64) Normalize() *X64 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X64) IsSimilar(op *X64) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
func (vec *X64) Dot(op *X64) vpnumber.X64 {
	var dot vpnumber.X64

	for i, v := range op {
		dot += vpnumber.X64Mul(vec[i], v)
	}

	return dot
}

// X64Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X64Add(veca, vecb *X64) *X64 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X64Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X64Sub(veca, vecb *X64) *X64 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X64Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func X64Neg(vec *X64) *X64 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// X64MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64MulScale(vec *X64, factor vpnumber.X64) *X64 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X64DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64DivScale(vec *X64, factor vpnumber.X64) *X64 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X64Lerp performs a linear interpolation between 2 vectors.
func X64Lerp(veca, vecb *X64, beta vpnumber.X64) *X64 {
	var ret = *veca

	ret.Lerp(vecb, beta)

	return &ret
}

// X64Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X64Normalize(vec *X64) *X64 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}
