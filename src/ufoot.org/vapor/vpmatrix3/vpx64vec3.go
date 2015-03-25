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

package vpmatrix3

import (
	"encoding/json"
	"ufoot.org/vapor/vpmath"
	"ufoot.org/vapor/vpnumber"
	"ufoot.org/vapor/vpsys"
)

// X64Vec3 is a vector containing 3 fixed point 64 bit values.
// Can hold the values of a point in space.
type X64Vec3 [3]vpnumber.X64

// X64Vec3New creates a new vector containing 3 fixed point 64 bit values.
func X64Vec3New(x1, x2, x3 vpnumber.X64) *X64Vec3 {
	return &X64Vec3{x1, x2, x3}
}

// ToI32 converts the vector to an int32 vector.
func (vec *X64Vec3) ToI32() *I32Vec3 {
	var ret I32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X64ToI32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int32 vector.
func (vec *X64Vec3) ToI64() *I64Vec3 {
	var ret I64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X64ToI64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 64 bits.
func (vec *X64Vec3) ToX32() *X32Vec3 {
	var ret X32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *X64Vec3) ToF32() *F32Vec3 {
	var ret F32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *X64Vec3) ToF64() *F64Vec3 {
	var ret F64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (vec *X64Vec3) MarshalJSON() ([]byte, error) {
	var tmpArray [3]int64

	for i := range tmpArray {
		tmpArray[i] = int64(vec[i])
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X64Vec3")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (vec *X64Vec3) UnmarshalJSON(data []byte) error {
	var tmpArray [3]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X64Vec3")
	}

	for i := range tmpArray {
		vec[i] = vpnumber.X64(tmpArray[i])
	}

	return nil
}

// String returns a readable form of the vector.
func (vec *X64Vec3) String() string {
	buf, err := vec.ToF64().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) Add(op *X64Vec3) *X64Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) Sub(op *X64Vec3) *X64Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) Neg() *X64Vec3 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) MulScale(factor vpnumber.X64) *X64Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Mul(v, factor)
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) DivScale(factor vpnumber.X64) *X64Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.X64Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *X64Vec3) SqMag() vpnumber.X64 {
	var sq vpnumber.X64

	for _, v := range vec {
		sq += vpnumber.X64Mul(v, v)
	}

	return sq
}

// Length returns the length of the vector.
func (vec *X64Vec3) Length() vpnumber.X64 {
	return vpmath.X64Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) Normalize() *X64Vec3 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *X64Vec3) IsSimilar(op *X64Vec3) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
// It modifies the vector, and returns a pointer on it.
func (vec *X64Vec3) Dot(op *X64Vec3) *X64Vec3 {
	for i, v := range op {
		vec[i] = vpnumber.X64Mul(vec[i], v)
	}

	return vec
}

// X64Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3Add(veca, vecb *X64Vec3) *X64Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// X64Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3Sub(veca, vecb *X64Vec3) *X64Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// X64Vec3Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func X64Vec3Neg(vec *X64Vec3) *X64Vec3 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// X64Vec3MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3MulScale(vec *X64Vec3, factor vpnumber.X64) *X64Vec3 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// X64Vec3DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3DivScale(vec *X64Vec3, factor vpnumber.X64) *X64Vec3 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// X64Vec3SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func X64Vec3SqMag(vec *X64Vec3) vpnumber.X64 {
	return vec.SqMag()
}

// X64Vec3Length returns the length of a vector.
func X64Vec3Length(vec *X64Vec3) vpnumber.X64 {
	return vec.Length()
}

// X64Vec3Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func X64Vec3Normalize(vec *X64Vec3) *X64Vec3 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// X64Vec3IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func X64Vec3IsSimilar(veca, vecb *X64Vec3) bool {
	return veca.IsSimilar(vecb)
}

// X64Vec3Dot returns the dot products of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func X64Vec3Dot(veca, vecb *X64Vec3) *X64Vec3 {
	var ret = *veca

	_ = ret.Dot(vecb)

	return &ret
}
