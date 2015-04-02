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
)

// I64Vec4 is a vector containing 4 int64 values.
// Can hold the values of a point in a plane.
type I64Vec4 [4]int64

// I64Vec4New creates a new vector containing 4 int64 values.
func I64Vec4New(i1, i2, i3, i4 int64) *I64Vec4 {
	return &I64Vec4{i1, i2, i3, i4}
}

// ToI32 converts the vector to an int32 vector.
func (vec *I64Vec4) ToI32() *I32Vec4 {
	var ret I32Vec4

	for i, v := range vec {
		ret[i] = int32(v)
	}

	return &ret
}

// ToX32 converts the vector to a fixed point number vector on 32 bits.
func (vec *I64Vec4) ToX32() *X32Vec4 {
	var ret X32Vec4

	for i, v := range vec {
		ret[i] = vpnumber.I64ToX32(v)
	}

	return &ret
}

// ToX64 converts the vector to a fixed point number vector on 64 bits.
func (vec *I64Vec4) ToX64() *X64Vec4 {
	var ret X64Vec4

	for i, v := range vec {
		ret[i] = vpnumber.I64ToX64(v)
	}

	return &ret
}

// ToF32 converts the vector to a float32 vector.
func (vec *I64Vec4) ToF32() *F32Vec4 {
	var ret F32Vec4

	for i, v := range vec {
		ret[i] = float32(v)
	}

	return &ret
}

// ToF64 converts the vector to a float64 vector.
func (vec *I64Vec4) ToF64() *F64Vec4 {
	var ret F64Vec4

	for i, v := range vec {
		ret[i] = float64(v)
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (vec *I64Vec4) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal([4]int64(*vec))
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal I64Vec4")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (vec *I64Vec4) UnmarshalJSON(data []byte) error {
	var tmpArray [4]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal I64Vec4")
	}

	*vec = I64Vec4(tmpArray)

	return nil
}

// String returns a readable form of the vector.
func (vec *I64Vec4) String() string {
	buf, err := vec.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I64Vec4) Add(op *I64Vec4) *I64Vec4 {
	for i, v := range op {
		vec[i] += v
	}

	return vec
}

// Sub substracts operand from the vector.
// It modifies the vector, and returns a pointer on it.
func (vec *I64Vec4) Sub(op *I64Vec4) *I64Vec4 {
	for i, v := range op {
		vec[i] -= v
	}

	return vec
}

// Neg changes the sign of all vector members.
// It modifies the vector, and returns a pointer on it.
func (vec *I64Vec4) Neg() *I64Vec4 {
	for i, v := range vec {
		vec[i] = -v
	}

	return vec
}

// I64Vec4Add adds two vectors.
// Args are left untouched, a pointer on a new object is returned.
func I64Vec4Add(veca, vecb *I64Vec4) *I64Vec4 {
	var ret = *veca

	_ = ret.Add(vecb)

	return &ret
}

// I64Vec4Sub substracts vector b from vector a.
// Args are left untouched, a pointer on a new object is returned.
func I64Vec4Sub(veca, vecb *I64Vec4) *I64Vec4 {
	var ret = *veca

	_ = ret.Sub(vecb)

	return &ret
}

// I64Vec4Neg changes the sign of all vector members.
// Arg is left untouched, a pointer on a new object is returned.
func I64Vec4Neg(vec *I64Vec4) *I64Vec4 {
	var ret = *vec

	_ = ret.Neg()

	return &ret
}
