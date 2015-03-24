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

package vpmatrix2

import (
	"encoding/json"
	"ufoot.org/vapor/vpmath"
	"ufoot.org/vapor/vpnumber"
	"ufoot.org/vapor/vpsys"
)

// X32Vec2 is a vector containing 2 fixed point 32 bit values.
// Can hold the values of a point in a plane.
type X32Vec2 [2]vpnumber.X32

// X32Vec2New creates a new vector containing 2 fixed point 32 bit values.
func X32Vec2New(x1, x2 vpnumber.X32) *X32Vec2 {
	return &X32Vec2{x1, x2}
}

// ToI32 converts the vector to an int32 vector.
func (vec *X32Vec2) ToI32() *I32Vec2 {
	var ret I32Vec2

	for i, v := range vec {
		ret[i] = vpnumber.X32ToI32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int32 vector.
func (vec *X32Vec2) ToI64() *I64Vec2 {
	var ret I64Vec2

	for i, v := range vec {
		ret[i] = vpnumber.X32ToI64(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *X32Vec2) ToX64() *X64Vec2 {
	var ret X64Vec2

	for i, v := range vec {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *X32Vec2) ToF32() *F32Vec2 {
	var ret F32Vec2

	for i, v := range vec {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *X32Vec2) ToF64() *F64Vec2 {
	var ret F64Vec2

	for i, v := range vec {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (vec *X32Vec2) MarshalJSON() ([]byte, error) {
	var tmpArray [2]int32

	for i := range tmpArray {
		tmpArray[i] = int32(vec[i])
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X32Vec2")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (vec *X32Vec2) UnmarshalJSON(data []byte) error {
	var tmpArray [2]int32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X32Vec2")
	}

	for i := range tmpArray {
		vec[i] = vpnumber.X32(tmpArray[i])
	}

	return nil
}

// String returns a readable form of the vector.
func (vec *X32Vec2) String() string {
	buf, err := vec.ToF32().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec2) Add(op *X32Vec2) *X32Vec2 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec2) Sub(op *X32Vec2) *X32Vec2 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec2) Neg() *X32Vec2 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec2) MulScale(factor vpnumber.X32) *X32Vec2 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec2) DivScale(factor vpnumber.X32) *X32Vec2 {
	for i, v := range vec {
		vec[i] = vpnumber.X32Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X32Vec2) SqMag() vpnumber.X32 {
	var sq vpnumber.X32

	for _, v := range vec {
		sq += vpnumber.X32Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X32Vec2) Length() vpnumber.X32 {
	return vpmath.X32Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *X32Vec2) Normalize() *X32Vec2 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X32Vec2) IsSimilar(op *X32Vec2) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// X32Vec2Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec2Add(veca, vecb *X32Vec2) *X32Vec2 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X32Vec2Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec2Sub(veca, vecb *X32Vec2) *X32Vec2 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X32Vec2Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func X32Vec2Neg(vec *X32Vec2) *X32Vec2 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// X32Vec2MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec2MulScale(vec *X32Vec2, factor vpnumber.X32) *X32Vec2 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X32Vec2DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Vec2DivScale(vec *X32Vec2, factor vpnumber.X32) *X32Vec2 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X32Vec2SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X32Vec2SqMag(vec *X32Vec2) vpnumber.X32 {
	return vec.SqMag()
}

// X32Vec2Length returns the length of a vector.
func X32Vec2Length(vec *X32Vec2) vpnumber.X32 {
	return vec.Length()
}

// X32Vec2Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X32Vec2Normalize(vec *X32Vec2) *X32Vec2 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X32Vec2IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func X32Vec2IsSimilar(veca, vecb *X32Vec2) bool {
	return veca.IsSimilar(vecb)
}
