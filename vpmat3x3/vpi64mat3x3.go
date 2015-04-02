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

package vpmat3x3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
)

// I64Mat3x3x3 is a matrix containing 3x3 int64 values.
// Can hold the values of a point in a plane.
type I64Mat3x3x3 [9]int64

// I64Mat3x3x3New creates a new matrix containing 3x3 int64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func I64Mat3x3x3New(i1, i2, i3, i4, i5, i6, i7, i8, i9 int64) *I64Mat3x3x3 {
	return &I64Mat3x3x3{i1, i2, i3, i4, i5, i6, i7, i8, i9}
}

// I64Mat3x3x3Identity creates a new identity matrix.
func I64Mat3x3x3Identity() *I64Mat3x3x3 {
	return &I64Mat3x3x3{vpnumber.I64Const1, vpnumber.I64Const0, vpnumber.I64Const0, vpnumber.I64Const0, vpnumber.I64Const1, vpnumber.I64Const0, vpnumber.I64Const0, vpnumber.I64Const0, vpnumber.I64Const1}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *I64Mat3x3x3) ToI32() *I32Mat3x3x3 {
	var ret I32Mat3x3x3

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *I64Mat3x3x3) ToX32() *X32Mat3x3x3 {
	var ret X32Mat3x3x3

	for i, v := range mat {
		ret[i] = vpnumber.I64ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *I64Mat3x3x3) ToX64() *X64Mat3x3x3 {
	var ret X64Mat3x3x3

	for i, v := range mat {
		ret[i] = vpnumber.I64ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *I64Mat3x3x3) ToF32() *F32Mat3x3x3 {
	var ret F32Mat3x3x3

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *I64Mat3x3x3) ToF64() *F64Mat3x3x3 {
	var ret F64Mat3x3x3

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *I64Mat3x3x3) Set(col, row int, val int64) {
	mat[col*3+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *I64Mat3x3x3) Get(col, row int) int64 {
	return mat[col*3+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *I64Mat3x3x3) MarshalJSON() ([]byte, error) {
	var tmpArray [3][3]int64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*3+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal I64Mat3x3x3")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *I64Mat3x3x3) UnmarshalJSON(data []byte) error {
	var tmpArray [3][3]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal I64Mat3x3x3")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*3+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *I64Mat3x3x3) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *I64Mat3x3x3) Add(op *I64Mat3x3x3) *I64Mat3x3x3 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *I64Mat3x3x3) Sub(op *I64Mat3x3x3) *I64Mat3x3x3 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// I64Mat3x3x3Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func I64Mat3x3x3Add(mata, matb *I64Mat3x3x3) *I64Mat3x3x3 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// I64Mat3x3x3Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func I64Mat3x3x3Sub(mata, matb *I64Mat3x3x3) *I64Mat3x3x3 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}
