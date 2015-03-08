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

// F64Vec2 is a vector containing 2 float64 values.
// Can hold the values of a point in a plane.
type F64Vec2 [2]float64

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec2) Add(op *F64Vec2) *F64Vec2 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec2) Sub(op *F64Vec2) *F64Vec2 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec2) MulScale(factor float64) *F64Vec2 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec2) DivScale(factor float64) *F64Vec2 {
	for i, v := range vec {
		vec[i] = vpnumber.F64Div(vec[i], v)
	}

	return vec
}

// SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F64Vec2) SumSq() float64 {
	var sq float64

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

func (vec *F64Vec2) Length() float64 {
	return math.Sqrt(vec.SumSq())
}

func (vec *F64Vec2) Normalize() *F64Vec2 {
	vec.DivScale(vec.Length())

	return vec
}

func F64Vec2Add(veca, vecb *F64Vec2) *F64Vec2 {
	var ret F64Vec2 = *veca

	_ = ret.Add(vecb)

	return &ret
}

func F64Vec2Sub(veca, vecb *F64Vec2) *F64Vec2 {
	var ret F64Vec2 = *veca

	_ = ret.Sub(vecb)

	return &ret
}

func F64Vec2MulScale(vec *F64Vec2, factor float64) *F64Vec2 {
	var ret F64Vec2 = *vec

	_ = ret.MulScale(factor)

	return &ret
}

func F64Vec2DivScale(vec *F64Vec2, factor float64) *F64Vec2 {
	var ret F64Vec2 = *vec

	_ = ret.DivScale(factor)

	return &ret
}

func F64Vec2SumSq(vec *F64Vec2) float64 {
	return vec.SumSq()
}

func F64Vec2Length(vec *F64Vec2) float64 {
	return vec.Length()
}

func F64Vec2Normalize(vec *F64Vec2) *F64Vec2 {
	var ret F64Vec2 = *vec

	_ = ret.Normalize()

	return &ret
}
