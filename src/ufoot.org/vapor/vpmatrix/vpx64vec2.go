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
	"ufoot.org/vapor/vpmath"
	"ufoot.org/vapor/vpnumber"
)

// X64Vec2 is a vector containing 2 fixed point 64 bit values.
// Can hold the values of a point in a plane.
type X64Vec2 [2]vpnumber.X64

// X64Vec2New creates a new vector containing 2 fixed point 64 bit values.
func X64Vec2New(x1, x2 vpnumber.X64) *X64Vec2 {
	return &X64Vec2{x1, x2}
}

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) Add(op *X64Vec2) *X64Vec2 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) Sub(op *X64Vec2) *X64Vec2 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) MulScale(factor vpnumber.X64) *X64Vec2 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) DivScale(factor vpnumber.X64) *X64Vec2 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Div(v, factor)
	}

	return vec
}

// SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X64Vec2) SumSq() vpnumber.X64 {
	var sq vpnumber.X64

	for _, v := range vec {
		sq += vpnumber.X64Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X64Vec2) Length() vpnumber.X64 {
	return vpmath.X64Sqrt(vec.SumSq())
}

// Normalize scales the vector so that its length is 1.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) Normalize() *X64Vec2 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X64Vec2) IsSimilar(op *X64Vec2) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// X64Vec2Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec2Add(veca, vecb *X64Vec2) *X64Vec2 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X64Vec2Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec2Sub(veca, vecb *X64Vec2) *X64Vec2 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X64Vec2MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec2MulScale(vec *X64Vec2, factor vpnumber.X64) *X64Vec2 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X64Vec2DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec2DivScale(vec *X64Vec2, factor vpnumber.X64) *X64Vec2 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X64Vec2SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X64Vec2SumSq(vec *X64Vec2) vpnumber.X64 {
	return vec.SumSq()
}

// X64Vec2Length returns the length of a vector.
func X64Vec2Length(vec *X64Vec2) vpnumber.X64 {
	return vec.Length()
}

// X64Vec2Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X64Vec2Normalize(vec *X64Vec2) *X64Vec2 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X64Vec2IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func X64Vec2IsSimilar(veca, vecb *X64Vec2) bool {
	return veca.IsSimilar(vecb)
}
