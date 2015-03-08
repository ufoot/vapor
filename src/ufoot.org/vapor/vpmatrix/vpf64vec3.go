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

// F64Vec3 is a vector containing 3 float64 values.
// Can hold the values of a point in space.
type F64Vec3 [3]float64

// F64Vec3New creates a new vector containing 3 float64 values.
func F64Vec3New(f1, f2, f3 float64) *F64Vec3 {
	return &F64Vec3{f1, f2, f3}
}

// Add adds operand to the vector.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec3) Add(op *F64Vec3) *F64Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec3) Sub(op *F64Vec3) *F64Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec3) MulScale(factor float64) *F64Vec3 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec3) DivScale(factor float64) *F64Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.F64Div(vec[i], v)
	}

	return vec
}

// SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F64Vec3) SumSq() float64 {
	var sq float64

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F64Vec3) Length() float64 {
	return math.Sqrt(vec.SumSq())
}

// Normalize scales the vector so that its length is 1.
// It modifies it, and returns a pointer on it.
func (vec *F64Vec3) Normalize() *F64Vec3 {
	vec.DivScale(vec.Length())

	return vec
}

// F64Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec3Add(veca, vecb *F64Vec3) *F64Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F64Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec3Sub(veca, vecb *F64Vec3) *F64Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F64Vec3MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec3MulScale(vec *F64Vec3, factor float64) *F64Vec3 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F64Vec3DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec3DivScale(vec *F64Vec3, factor float64) *F64Vec3 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F64Vec3SumSq returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func F64Vec3SumSq(vec *F64Vec3) float64 {
	return vec.SumSq()
}

// F64Vec3Length returns the length of a vector.
func F64Vec3Length(vec *F64Vec3) float64 {
	return vec.Length()
}

// F64Vec3Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F64Vec3Normalize(vec *F64Vec3) *F64Vec3 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}
