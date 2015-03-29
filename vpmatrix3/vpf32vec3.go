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

package vpmatrix3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmatrix2"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
	"math"
)

// F32Vec3 is a vector containing 3 float32 values.
// Can hold the values of a point in space.
type F32Vec3 [3]float32

// F32Vec3New creates a new vector containing 3 float32 values.
func F32Vec3New(f1, f2, f3 float32) *F32Vec3 {
	return &F32Vec3{f1, f2, f3}
}

// F32VecFromVec2 creates a new vector from a smaller one,
// by appending a value at its end.
func F32Vec3FromVec2(vec *vpmatrix2.F32Vec2, f float32) *F32Vec3 {
	return &F32Vec3{vec[0], vec[1], f}
}

// F32VecToVec2 creates a smaller vector by removing the last value.
func (vec *F32Vec3) ToVec2() *vpmatrix2.F32Vec2 {
	return &vpmatrix2.F32Vec2{vec[0], vec[1]}
}

// ToI32 converts the vector to an int32 vector.
func (vec *F32Vec3) ToI32() *I32Vec3 {
	var ret I32Vec3

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the vector to an int64 vector.
func (vec *F32Vec3) ToI64() *I64Vec3 {
	var ret I64Vec3

	for i, v := range vec {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *F32Vec3) ToX32() *X32Vec3 {
	var ret X32Vec3

	for i, v := range vec {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *F32Vec3) ToX64() *X64Vec3 {
	var ret X64Vec3

	for i, v := range vec {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *F32Vec3) ToF64() *F64Vec3 {
	var ret F64Vec3

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (vec *F32Vec3) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal([3]float32(*vec))
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F32Vec3")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (vec *F32Vec3) UnmarshalJSON(data []byte) error {
	var tmpArray [3]float32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F32Vec3")
	}

	*vec = F32Vec3(tmpArray)

	return nil
}

// String returns a readable form of the vector.
func (vec *F32Vec3) String() string {
	buf, err := vec.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) Add(op *F32Vec3) *F32Vec3 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) Sub(op *F32Vec3) *F32Vec3 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) Neg() *F32Vec3 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// MulScale multiplies all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) MulScale(factor float32) *F32Vec3 {
	for i, v := range vec {
		vec[i] = v * factor
	}

	return vec
}

// DivScale divides all values of the vector by factor.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) DivScale(factor float32) *F32Vec3 {
	for i, v := range vec {
		vec[i] = vpnumber.F32Div(v, factor)
	}

	return vec
}

// SqMag returns the sum of the squares of all values.
// It is used to calculate length, it is faster than the complete
// length calculation, as it does not perform a square root.
func (vec *F32Vec3) SqMag() float32 {
	var sq float32

	for _, v := range vec {
		sq += v * v
	}

	return sq
}

// Length returns the length of the vector.
func (vec *F32Vec3) Length() float32 {
	return float32(math.Sqrt(float64(vec.SqMag())))
}

// Normalize scales the vector so that its length is 1.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) Normalize() *F32Vec3 {
	vec.DivScale(vec.Length())

	return vec
}

// IsSimilar returns true if vectors are approximatively the same.
// This is a workarround to ignore rounding errors.
func (vec *F32Vec3) IsSimilar(op *F32Vec3) bool {
	ret := true
	for i, v := range vec {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// Dot returns the the dot product of two vectors.
func (vec *F32Vec3) Dot(op *F32Vec3) float32 {
	var dot float32

	for i, v := range op {
		dot += vec[i] * v
	}

	return dot
}

// Cross returns the the cross product of two vectors.
// It modifies the vector, and returns a pointer on it.
func (vec *F32Vec3) Cross(op *F32Vec3) *F32Vec3 {
	*vec = *F32Vec3Cross(vec, op)

	return vec
}

// F32Vec3Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec3Add(veca, vecb *F32Vec3) *F32Vec3 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// F32Vec3Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec3Sub(veca, vecb *F32Vec3) *F32Vec3 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// F32Vec3Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func F32Vec3Neg(vec *F32Vec3) *F32Vec3 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}

// F32Vec3MulScale multiplies all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec3MulScale(vec *F32Vec3, factor float32) *F32Vec3 {
	var ret = *vec

	_ = ret.MulScale(factor)

	return &ret
}

// F32Vec3DivScale divides all values of a vector by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec3DivScale(vec *F32Vec3, factor float32) *F32Vec3 {
	var ret = *vec

	_ = ret.DivScale(factor)

	return &ret
}

// F32Vec3Normalize scales a vector so that its length is 1.
// Arg is left untouched, a pointer on a new object is returned.
func F32Vec3Normalize(vec *F32Vec3) *F32Vec3 {
	var ret = *vec

	_ = ret.Normalize()

	return &ret
}

// F32Vec3Cross returns the the cross product of two vectors.
// Args are left untouched, a pointer on a new object is returned.
func F32Vec3Cross(veca, vecb *F32Vec3) *F32Vec3 {
	var ret = F32Vec3{veca[1]*vecb[2] - veca[2]*vecb[1], veca[2]*vecb[0] - veca[0]*vecb[2], veca[0]*vecb[1] - veca[1]*vecb[0]}

	return &ret
}
