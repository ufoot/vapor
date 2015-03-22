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
	"math"
	"ufoot.org/vapor/vpnumber"
)

// F32Vec4 is a vector containing 4 float32 values.
// Can be used in 3D matrix transformations.
type F32Vec4 [4]float32

// F32Vec4New creates a new vector containing 4 float32 values.
func F32Vec4New(f1, f2, f3, f4 float32) *F32Vec4 {
	return &F32Vec4{f1, f2, f3, f4}
}

// ToI32 converts the vector to an int32 vector.
func (vec *F32Vec4) ToI32() *I32Vec4 {
	var ret I32Vec4

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int64 vector.
func (vec *F32Vec4) ToI64() *I64Vec4 {
	var ret I64Vec4

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *F32Vec4) ToX32() *X32Vec4 {
	var ret X32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *F32Vec4) ToX64() *X64Vec4 {
	var ret X64Vec4

	for i, v := range vec {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *F32Vec4) ToF64() *F64Vec4 {
	var ret F64Vec4

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec4) Add(op *F32Vec4) *F32Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec4) Sub(op *F32Vec4) *F32Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec4) Neg() *F32Vec4 {
	for i,v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec4) MulScale(factor float32) *F32Vec4 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec4) DivScale(factor float32) *F32Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.F32Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F32Vec4) SqMag() float32 {
	var sq float32

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F32Vec4) Length() float32 {
	return float32(math.Sqrt(float64(vec.SqMag())))
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec4) Normalize() *F32Vec4 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *F32Vec4) IsSimilar(op *F32Vec4) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// F32Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4Add(veca, vecb *F32Vec4) *F32Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F32Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4Sub(veca, vecb *F32Vec4) *F32Vec4 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F32Vec4Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func F32Vec4Neg(vec *F32Vec4) *F32Vec4 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// F32Vec4MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4MulScale(vec *F32Vec4, factor float32) *F32Vec4 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F32Vec4DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4DivScale(vec *F32Vec4, factor float32) *F32Vec4 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F32Vec4SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func F32Vec4SqMag(vec *F32Vec4) float32 {
	return vec.SqMag()
}

// F32Vec4Length returns the length of a vector.
func F32Vec4Length(vec *F32Vec4) float32 {
	return vec.Length()
}

// F32Vec4Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F32Vec4Normalize(vec *F32Vec4) *F32Vec4 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// F32Vec4IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func F32Vec4IsSimilar(veca, vecb *F32Vec4) bool {
	return veca.IsSimilar(vecb)
}
