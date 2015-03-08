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

// X32Vec4 is a vector containing 4 fixed point 32 bit values.
// Can be used in 3D matrix transformations.
type X32Vec4 [4]vpnumber.X32

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *X32Vec4) Add(op *X32Vec4) *X32Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *X32Vec4) Sub(op *X32Vec4) *X32Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X32Vec4) MulScale(factor vpnumber.X32) *X32Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Mul(vec[i], v)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X32Vec4) DivScale(factor vpnumber.X32) *X32Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Div(vec[i], v)
	}

	return vec
}

// SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X32Vec4) SumSq() vpnumber.X32 {
	var sq vpnumber.X32

	for _, v := range vec {
		sq += vpnumber.X32Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X32Vec4) Length() vpnumber.X32 {
	return vpnumber.F64ToX32(math.Sqrt(vpnumber.X32ToF64(vec.SumSq())))
}

// Normalize scales the vector so that its length is 1.
// It modifies it, and returns a pointer on it.
func (vec *X32Vec4) Normalize() *X32Vec4 {
	vec.DivScale(vec.Length())

	return vec
}

// X32Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec4Add(veca, vecb *X32Vec4) *X32Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X32Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec4Sub(veca, vecb *X32Vec4) *X32Vec4 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X32Vec4MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec4MulScale(vec *X32Vec4, factor vpnumber.X32) *X32Vec4 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X32Vec4DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec4DivScale(vec *X32Vec4, factor vpnumber.X32) *X32Vec4 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X32Vec4SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X32Vec4SumSq(vec *X32Vec4) vpnumber.X32 {
	return vec.SumSq()
}

// X32Vec4Length returns the length of a vector.
func X32Vec4Length(vec *X32Vec4) vpnumber.X32 {
	return vec.Length()
}

// X32Vec4Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X32Vec4Normalize(vec *X32Vec4) *X32Vec4 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}
