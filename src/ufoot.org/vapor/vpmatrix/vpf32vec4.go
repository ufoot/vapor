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

package vpmath

import (
	"math"
	"ufoot.org/vapor/vpnumber"
)

// F32Vec4 is a vector containing 4 float32 values.
// Can be used in 3D matrix transformations.
type F32Vec4 [4]float32

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *F32Vec4) Add(op *F32Vec4) *F32Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *F32Vec4) Sub(op *F32Vec4) *F32Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *F32Vec4) MulScale(factor float32) *F32Vec4 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *F32Vec4) DivScale(factor float32) *F32Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.F32Div(vec[i], v)
	}

	return vec
}

// SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F32Vec4) SumSq() float32 {
	var sq float32

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F32Vec4) Length() float32 {
	return float32(math.Sqrt(float64(vec.SumSq())))
}

// Normalize scales the vector so that its length is 1.
// It modifies it, and returns a pointer on it.
func (vec *F32Vec4) Normalize() *F32Vec4 {
	vec.DivScale(vec.Length())

	return vec
}

// F32Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4Add(veca, vecb *F32Vec4) *F32Vec4 {
	var ret F32Vec4 = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F32Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4Sub(veca, vecb *F32Vec4) *F32Vec4 {
	var ret F32Vec4 = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F32Vec4MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4MulScale(vec *F32Vec4, factor float32) *F32Vec4 {
	var ret F32Vec4 = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F32Vec4DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec4DivScale(vec *F32Vec4, factor float32) *F32Vec4 {
	var ret F32Vec4 = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F32Vec4SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func F32Vec4SumSq(vec *F32Vec4) float32 {
	return vec.SumSq()
}

// F32Vec4Length returns the length of a vector.
func F32Vec4Length(vec *F32Vec4) float32 {
	return vec.Length()
}

// F32Vec4Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F32Vec4Normalize(vec *F32Vec4) *F32Vec4 {
	var ret F32Vec4 = *vec

	_ = ret.Normalize()

	return &ret
}
