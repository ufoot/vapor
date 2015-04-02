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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpvec4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
	"github.com/ufoot/vapor/vpvec3"
	"math"
)

// F64Vec4 is a vector containing 4 float64 values.
// Can be used in 3D matrix transformations.
type F64Vec4 [4]float64

// F64Vec4New creates a new vector containing 4 float64 values.
func F64Vec4New(f1, f2, f3, f4 float64) *F64Vec4 {
	return &F64Vec4{f1, f2, f3, f4}
}

// F64Vec4FromVec3 creates a new vector from a smaller one,
// by appending a value at its end.
func F64Vec4FromVec3(vec *vpvec3.F64Vec3, f float64) *F64Vec4 {
	return &F64Vec4{vec[0], vec[1], vec[2], f}
}

// ToVec3 creates a smaller vector by removing the last value.
func (vec *F64Vec4) ToVec3() *vpvec3.F64Vec3 {
	return &vpvec3.F64Vec3{vec[0], vec[1], vec[2]}
}

// ToI32 converts the vector to an int32 vector.
func (vec *F64Vec4) ToI32() *I32Vec4 {
	var ret I32Vec4

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int64 vector.
func (vec *F64Vec4) ToI64() *I64Vec4 {
	var ret I64Vec4

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *F64Vec4) ToX32() *X32Vec4 {
	var ret X32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *F64Vec4) ToX64() *X64Vec4 {
	var ret X64Vec4

	for i, v := range vec {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *F64Vec4) ToF32() *F32Vec4 {
	var ret F32Vec4

	for i, v := range vec {
		ret[i] = float32(v)
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (vec *F64Vec4) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal([4]float64(*vec))
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F64Vec4")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (vec *F64Vec4) UnmarshalJSON(data []byte) error {
	var tmpArray [4]float64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F64Vec4")
	}

	*vec = F64Vec4(tmpArray)

	return nil
}

// String returns a readable form of the vector.
func (vec *F64Vec4) String() string {
	buf, err := vec.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F64Vec4) Add(op *F64Vec4) *F64Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F64Vec4) Sub(op *F64Vec4) *F64Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F64Vec4) Neg() *F64Vec4 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F64Vec4) MulScale(factor float64) *F64Vec4 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F64Vec4) DivScale(factor float64) *F64Vec4 {
	for i, v := range vec {
		vec[i] = vpnumber.F64Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F64Vec4) SqMag() float64 {
	var sq float64

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F64Vec4) Length() float64 {
	return math.Sqrt(vec.SqMag())
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *F64Vec4) Normalize() *F64Vec4 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *F64Vec4) IsSimilar(op *F64Vec4) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
func (vec *F64Vec4) Dot(op *F64Vec4) float64 {
	var dot float64

	for i, v := range op {
		dot += vec[i] * v
	}

	return dot
}

// F64Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec4Add(veca, vecb *F64Vec4) *F64Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F64Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec4Sub(veca, vecb *F64Vec4) *F64Vec4 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F64Vec4Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func F64Vec4Neg(vec *F64Vec4) *F64Vec4 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// F64Vec4MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec4MulScale(vec *F64Vec4, factor float64) *F64Vec4 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F64Vec4DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Vec4DivScale(vec *F64Vec4, factor float64) *F64Vec4 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F64Vec4Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F64Vec4Normalize(vec *F64Vec4) *F64Vec4 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}