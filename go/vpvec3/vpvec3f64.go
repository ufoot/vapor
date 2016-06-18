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
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
	"math"
)

// F64 is a vector containing 3 float64 values.
// Can hold the values of a point in space.
type F64 [Size]float64

// F64UnaryOperator designs funcs such as Neg,
// which operates on one vector and returns another vector.
type F64UnaryOperator func(a *F64) *F64

// F64BinaryOperator designs funcs such as Add, Sub, Min, Max,
// which operates on two vectors and return one.
type F64BinaryOperator func(a, b *F64) *F64

// F64New creates a new vector containing 3 float64 values.
func F64New(f1, f2, f3 float64) *F64 {
	return &F64{f1, f2, f3}
}

// F64AxisX returns a new vector representing the X axis.
func F64AxisX() *F64 {
	return &F64{vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0}
}

// F64AxisY returns a new vector representing the Y axis.
func F64AxisY() *F64 {
	return &F64{vpnumber.F64Const0, vpnumber.F64Const1, vpnumber.F64Const0}
}

// F64AxisZ returns a new vector representing the Z axis.
func F64AxisZ() *F64 {
	return &F64{vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1}
}

// F64FromVec2 creates a new vector from a smaller one,
// by appending a value at its end.
func F64FromVec2(vec *vpvec2.F64, f float64) *F64 {
	return &F64{vec[0], vec[1], f}
}

// ToVec2 creates a smaller vector by removing the last value.
func (vec *F64) ToVec2() *vpvec2.F64 {
	return &vpvec2.F64{vec[0], vec[1]}
}

// ToI32 converts the vector to an int32 vector.
func (vec *F64) ToI32() *I32 {
	var ret I32

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int64 vector.
func (vec *F64) ToI64() *I64 {
	var ret I64

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *F64) ToX32() *X32 {
	var ret X32

	for i, v := range vec {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *F64) ToX64() *X64 {
	var ret X64

	for i, v := range vec {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *F64) ToF32() *F32 {
	var ret F32

	for i, v := range vec {
		ret[i] = float32(v)
	}

	return &ret
}

// String returns a readable form of the vector.
func (vec *F64) String() string {
	buf, err := json.Marshal(vec)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Add(op *F64) *F64 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Sub(op *F64) *F64 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Neg() *F64 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// Min returns the minimum of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Min(op *F64) *F64 {
	for i, v := range op {
		if vec[i] > v {
			vec[i] = v
		}
	}

	return vec
}

// Max returns the maximum of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Max(op *F64) *F64 {
	for i, v := range op {
		if vec[i] < v {
			vec[i] = v
		}
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) MulScale(factor float64) *F64 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) DivScale(factor float64) *F64 {
	for i, v := range vec {
		vec[i] = vpnumber.F64Div(v, factor)
	}

	return vec
}

// Lerp performs a linear interpolation with another vector.
func (vec *F64) Lerp(op *F64, beta float64) *F64 {
	switch {
	case beta <= vpnumber.F64Const0:
		return vec
	case beta >= vpnumber.F64Const1:
		*vec = *op
		return vec
	}

	vec.MulScale(vpnumber.F64Const1 - beta)
	vec.Add(F64MulScale(op, beta))

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F64) SqMag() float64 {
	var sq float64

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F64) Length() float64 {
	return math.Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Normalize() *F64 {
	vec.DivScale(vec.Length())

	return vec
}

// Homogeneous scales the vector so that its latest member is 1.
// This what we want to do when projecting, to have homegemous coords.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Homogeneous() *F64 {
	vec.DivScale(vec[Size-1])

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *F64) IsSimilar(op *F64) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
func (vec *F64) Dot(op *F64) float64 {
	var dot float64

	for i, v := range op {
		dot += vec[i] * v
	}

	return dot
}

// Cross returns the the cross product of two vectors.
// It modifies the vector, and returns a pointer on it.
func (vec *F64) Cross(op *F64) *F64 {
	*vec = *F64Cross(vec, op)

	return vec
}

// F64Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F64Add(veca, vecb *F64) *F64 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F64Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F64Sub(veca, vecb *F64) *F64 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F64Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func F64Neg(vec *F64) *F64 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// F64Min returns the mininum (member-wise) of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F64Min(veca *F64, vecb *F64) *F64 {
	var ret = *veca

	_ = ret.Min(vecb)

	return &ret
}

// F64Max returns the mininum (member-wise) of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F64Max(veca *F64, vecb *F64) *F64 {
	var ret = *veca

	_ = ret.Max(vecb)

	return &ret
}

// F64MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64MulScale(vec *F64, factor float64) *F64 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F64DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64DivScale(vec *F64, factor float64) *F64 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F64Lerp performs a linear interpolation between 2 vectors.
func F64Lerp(veca, vecb *F64, beta float64) *F64 {
	var ret = *veca

	ret.Lerp(vecb, beta)

	return &ret
}

// F64Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F64Normalize(vec *F64) *F64 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// F64Homogeneous scales the vector so that its latest member is 1.
// This what we want to do when projecting, to have homegemous coords.
// Arg is left untouched, a pointer on a new object is returned.
func F64Homogeneous(vec *F64) *F64 {
	var ret = *vec

	_ = ret.Homogeneous()

	return &ret
}

// F64Cross returns the the cross product of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F64Cross(veca, vecb *F64) *F64 {
	var ret = F64{veca[1]*vecb[2] - veca[2]*vecb[1], veca[2]*vecb[0] - veca[0]*vecb[2], veca[0]*vecb[1] - veca[1]*vecb[0]}

	return &ret
}
