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

// X64Vec2 is a vector containing 2 fixed point 64 bit values.
// Can hold the values of a point in a plane.
type X64Vec2 [2]vpnumber.X64

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
		vec[i] = vpnumber.X64Mul(vec[i], v)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) DivScale(factor vpnumber.X64) *X64Vec2 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Div(vec[i], v)
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
	return vpnumber.F64ToX64(math.Sqrt(vpnumber.X64ToF64(vec.SumSq())))
}

// Normalize scales the vector so that its length is 1.
// It modifies it, and returns a pointer on it.
func (vec *X64Vec2) Normalize() *X64Vec2 {
	vec.DivScale(vec.Length())

	return vec
}

func X64Vec2Add(veca, vecb *X64Vec2) *X64Vec2 {
	var ret X64Vec2 = *veca

	_ = ret.Add(vecb)

	return &ret
}

func X64Vec2Sub(veca, vecb *X64Vec2) *X64Vec2 {
	var ret X64Vec2 = *veca

	_ = ret.Sub(vecb)

	return &ret
}

func X64Vec2MulScale(vec *X64Vec2, factor vpnumber.X64) *X64Vec2 {
	var ret X64Vec2 = *vec

	_ = ret.MulScale(factor)

	return &ret
}

func X64Vec2DivScale(vec *X64Vec2, factor vpnumber.X64) *X64Vec2 {
	var ret X64Vec2 = *vec

	_ = ret.DivScale(factor)

	return &ret
}

func X64Vec2SumSq(vec *X64Vec2) vpnumber.X64 {
	return vec.SumSq()
}

func X64Vec2Length(vec *X64Vec2) vpnumber.X64 {
	return vec.Length()
}

func X64Vec2Normalize(vec *X64Vec2) *X64Vec2 {
	var ret X64Vec2 = *vec

	_ = ret.Normalize()

	return &ret
}
