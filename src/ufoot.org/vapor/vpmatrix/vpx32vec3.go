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

// X32Vec3 is a vector containing 3 fixed point 32 bit values.
// Can hold the values of a point in space.
type X32Vec3 [3]vpnumber.X32

// X32Vec3New creates a new vector containing 3 fixed point 32 bit values.
func X32Vec3New(x1, x2, x3 vpnumber.X32) *X32Vec3 {
	return &X32Vec3{x1, x2, x3}
}

// ToI32 converts the vector to an int32 vector.
func (vec *X32Vec3) ToI32() *I32Vec3 {
	var ret I32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X32ToI32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int32 vector.
func (vec *X32Vec3) ToI64() *I64Vec3 {
	var ret I64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X32ToI64(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *X32Vec3) ToX64() *X64Vec3 {
	var ret X64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *X32Vec3) ToF32() *F32Vec3 {
	var ret F32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *X32Vec3) ToF64() *F64Vec3 {
	var ret F64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec3) Add(op *X32Vec3) *X32Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec3) Sub(op *X32Vec3) *X32Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec3) MulScale(factor vpnumber.X32) *X32Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec3) DivScale(factor vpnumber.X32) *X32Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X32Vec3) SqMag() vpnumber.X32 {
	var sq vpnumber.X32

	for _, v := range vec {
		sq += vpnumber.X32Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X32Vec3) Length() vpnumber.X32 {
	return vpmath.X32Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec3) Normalize() *X32Vec3 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X32Vec3) IsSimilar(op *X32Vec3) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// X32Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec3Add(veca, vecb *X32Vec3) *X32Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X32Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec3Sub(veca, vecb *X32Vec3) *X32Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X32Vec3MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec3MulScale(vec *X32Vec3, factor vpnumber.X32) *X32Vec3 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X32Vec3DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec3DivScale(vec *X32Vec3, factor vpnumber.X32) *X32Vec3 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X32Vec3SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X32Vec3SqMag(vec *X32Vec3) vpnumber.X32 {
	return vec.SqMag()
}

// X32Vec3Length returns the length of a vector.
func X32Vec3Length(vec *X32Vec3) vpnumber.X32 {
	return vec.Length()
}

// X32Vec3Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X32Vec3Normalize(vec *X32Vec3) *X32Vec3 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X32Vec3IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func X32Vec3IsSimilar(veca, vecb *X32Vec3) bool {
	return veca.IsSimilar(vecb)
}
