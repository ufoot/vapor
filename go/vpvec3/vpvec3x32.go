// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
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

package vpvec3

import (
	"encoding/json"
	"github.com/ufoot/vapor/go/vpmath"
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
)

// X32 is a vector containing 3 fixed point 32 bit values.
// Can hold the values of a point in space.
type X32 [Size]vpnumber.X32

// X32UnaryOperator designs funcs such as Neg,
// which operates on one vector and returns another vector.
type X32UnaryOperator func(a *X32) *X32

// X32BinaryOperator designs funcs such as Add, Sub, Min, Max,
// which operates on two vectors and return one.
type X32BinaryOperator func(a, b *X32) *X32

// X32New creates a new vector containing 3 fixed point 32 bit values.
func X32New(x1, x2, x3 vpnumber.X32) *X32 {
	return &X32{x1, x2, x3}
}

// X32AxisX returns a new vector representing the X axis.
func X32AxisX() *X32 {
	return &X32{vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0}
}

// X32AxisY returns a new vector representing the Y axis.
func X32AxisY() *X32 {
	return &X32{vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0}
}

// X32AxisZ returns a new vector representing the Z axis.
func X32AxisZ() *X32 {
	return &X32{vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32FromVec2 creates a new vector from a smaller one,
// by appending a value at its end.
func X32FromVec2(vec *vpvec2.X32, x vpnumber.X32) *X32 {
	return &X32{vec[0], vec[1], x}
}

// ToVec2 creates a smaller vector by removing the last value.
func (vec *X32) ToVec2() *vpvec2.X32 {
	return &vpvec2.X32{vec[0], vec[1]}
}

// ToI32 converts the vector to an int32 vector.
func (vec *X32) ToI32() *I32 {
	var ret I32

	for i, v := range vec {
		ret[i] = vpnumber.X32ToI32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int32 vector.
func (vec *X32) ToI64() *I64 {
	var ret I64

	for i, v := range vec {
		ret[i] = vpnumber.X32ToI64(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *X32) ToX64() *X64 {
	var ret X64

	for i, v := range vec {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *X32) ToF32() *F32 {
	var ret F32

	for i, v := range vec {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *X32) ToF64() *F64 {
	var ret F64

	for i, v := range vec {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// String returns a readable form of the vector.
func (vec *X32) String() string {
	buf, err := json.Marshal(vec.ToF32())

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Add(op *X32) *X32 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Sub(op *X32) *X32 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Neg() *X32 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// Min returns the minimum of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Min(op *X32) *X32 {
	for i, v := range op {
		if vec[i] > v {
			vec[i] = v
		}
	}

	return vec
}

// Max returns the maximum of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Max(op *X32) *X32 {
	for i, v := range op {
		if vec[i] < v {
			vec[i] = v
		}
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) MulScale(factor vpnumber.X32) *X32 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) DivScale(factor vpnumber.X32) *X32 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Div(v, factor)
	}

	return vec
}

// Lerp performs a linear interpolation with another vector.
func (vec *X32) Lerp(op *X32, beta vpnumber.X32) *X32 {
	switch {
	case beta <= vpnumber.X32Const0:
		return vec
	case beta >= vpnumber.X32Const1:
		*vec = *op
		return vec
	}

	vec.MulScale(vpnumber.X32Const1 - beta)
	vec.Add(X32MulScale(op, beta))

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X32) SqMag() vpnumber.X32 {
	var sq vpnumber.X32

	for _, v := range vec {
		sq += vpnumber.X32Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X32) Length() vpnumber.X32 {
	return vpmath.X32Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Normalize() *X32 {
	vec.DivScale(vec.Length())

	return vec
}

// Homogeneous scales the vector so that its latest member is 1.
// This what we want to do when projecting, to have homegemous coords.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Homogeneous() *X32 {
	vec.DivScale(vec[Size-1])

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X32) IsSimilar(op *X32) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
func (vec *X32) Dot(op *X32) vpnumber.X32 {
	var dot vpnumber.X32

	for i, v := range op {
		dot += vpnumber.X32Mul(vec[i], v)
	}

	return dot
}

// Cross returns the the cross product of two vectors.
// It modifies the vector, and returns a pointer on it.
func (vec *X32) Cross(op *X32) *X32 {
	*vec = *X32Cross(vec, op)

	return vec
}

// X32Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Add(veca, vecb *X32) *X32 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X32Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X32Sub(veca, vecb *X32) *X32 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X32Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func X32Neg(vec *X32) *X32 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// X32Min returns the mininum (member-wise) of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Min(veca *X32, vecb *X32) *X32 {
	var ret = *veca

	_ = ret.Min(vecb)

	return &ret
}

// X32Max returns the mininum (member-wise) of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Max(veca *X32, vecb *X32) *X32 {
	var ret = *veca

	_ = ret.Max(vecb)

	return &ret
}

// X32MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32MulScale(vec *X32, factor vpnumber.X32) *X32 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X32DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32DivScale(vec *X32, factor vpnumber.X32) *X32 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X32Lerp performs a linear interpolation between 2 vectors.
func X32Lerp(veca, vecb *X32, beta vpnumber.X32) *X32 {
	var ret = *veca

	ret.Lerp(vecb, beta)

	return &ret
}

// X32Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X32Normalize(vec *X32) *X32 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X32Homogeneous scales the vector so that its latest member is 1.
// This what we want to do when projecting, to have homegemous coords.
// Arg is left untouched, a pointer on a new object is returned.
func X32Homogeneous(vec *X32) *X32 {
	var ret = *vec

	_ = ret.Homogeneous()

	return &ret
}

// X32Cross returns the the cross product of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Cross(veca, vecb *X32) *X32 {
	var ret = X32{vpnumber.X32Mul(veca[1], vecb[2]) - vpnumber.X32Mul(veca[2], vecb[1]), vpnumber.X32Mul(veca[2], vecb[0]) - vpnumber.X32Mul(veca[0], vecb[2]), vpnumber.X32Mul(veca[0], vecb[1]) - vpnumber.X32Mul(veca[1], vecb[0])}

	return &ret
}
