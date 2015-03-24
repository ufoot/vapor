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

package vpmatrix4

import (
	"encoding/json"
	"ufoot.org/vapor/vpmath"
	"ufoot.org/vapor/vpnumber"
	"ufoot.org/vapor/vpsys"
)

// X64Vec4 is a vector containing 4 fixed point 64 bit values.
// Can be used in 3D matrix transformations.
type X64Vec4 [4]vpnumber.X64

// X64Vec4New creates a new vector containing 4 fixed point 64 bit values.
func X64Vec4New(x1, x2, x3, x4 vpnumber.X64) *X64Vec4 {
	return &X64Vec4{x1, x2, x3, x4}
}

// ToI32 converts the vector to an int32 vector.
func (vec *X64Vec4) ToI32() *I32Vec4 {
	var ret I32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.X64ToI32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int32 vector.
func (vec *X64Vec4) ToI64() *I64Vec4 {
	var ret I64Vec4

	for i, v := range vec {
		ret[i] = vpnumber.X64ToI64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 64 bits.
func (vec *X64Vec4) ToX32() *X32Vec4 {
	var ret X32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *X64Vec4) ToF32() *F32Vec4 {
	var ret F32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *X64Vec4) ToF64() *F64Vec4 {
	var ret F64Vec4

	for i, v := range vec {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (vec *X64Vec4) MarshalJSON() ([]byte, error) {
	var tmpArray [4]int64

	for i := range tmpArray {
		tmpArray[i] = int64(vec[i])
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X64Vec4")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (vec *X64Vec4) UnmarshalJSON(data []byte) error {
	var tmpArray [4]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X64Vec4")
	}

	for i := range tmpArray {
		vec[i] = vpnumber.X64(tmpArray[i])
	}

	return nil
}

// String returns a readable form of the vecrix.
func (vec *X64Vec4) String() string {
	buf, err := vec.ToF64().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec4) Add(op *X64Vec4) *X64Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec4) Sub(op *X64Vec4) *X64Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec4) Neg() *X64Vec4 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec4) MulScale(factor vpnumber.X64) *X64Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec4) DivScale(factor vpnumber.X64) *X64Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X64Vec4) SqMag() vpnumber.X64 {
	var sq vpnumber.X64

	for _, v := range vec {
		sq += vpnumber.X64Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X64Vec4) Length() vpnumber.X64 {
	return vpmath.X64Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec4) Normalize() *X64Vec4 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X64Vec4) IsSimilar(op *X64Vec4) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// X64Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec4Add(veca, vecb *X64Vec4) *X64Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X64Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec4Sub(veca, vecb *X64Vec4) *X64Vec4 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X64Vec4Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func X64Vec4Neg(vec *X64Vec4) *X64Vec4 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// X64Vec4MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec4MulScale(vec *X64Vec4, factor vpnumber.X64) *X64Vec4 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X64Vec4DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec4DivScale(vec *X64Vec4, factor vpnumber.X64) *X64Vec4 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X64Vec4SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X64Vec4SqMag(vec *X64Vec4) vpnumber.X64 {
	return vec.SqMag()
}

// X64Vec4Length returns the length of a vector.
func X64Vec4Length(vec *X64Vec4) vpnumber.X64 {
	return vec.Length()
}

// X64Vec4Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X64Vec4Normalize(vec *X64Vec4) *X64Vec4 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X64Vec4IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func X64Vec4IsSimilar(veca, vecb *X64Vec4) bool {
	return veca.IsSimilar(vecb)
}
