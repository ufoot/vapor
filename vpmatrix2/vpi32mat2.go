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
	"ufoot.org/vapor/vpnumber"
	"ufoot.org/vapor/vpsys"
)

// I32Mat2 is a matrix containing 2x2 int32 values.
// Can hold the values of a point in a plane.
type I32Mat2 [4]int32

// I32Mat2New creates a new matrix containing 4x4 int32 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func I32Mat2New(i1, i2, i3, i4 int32) *I32Mat2 {
	return &I32Mat2{i1, i2, i3, i4}
}

// I32Mat2Identity creates a new identity matrix.
func I32Mat2Identity() *I32Mat2 {
	return &I32Mat2{vpnumber.I32Const1, vpnumber.I32Const0, vpnumber.I32Const0, vpnumber.I32Const1}
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *I32Mat2) ToI64() *I64Mat2 {
	var ret I64Mat2

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *I32Mat2) ToX32() *X32Mat2 {
	var ret X32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.I32ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *I32Mat2) ToX64() *X64Mat2 {
	var ret X64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.I32ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *I32Mat2) ToF32() *F32Mat2 {
	var ret F32Mat2

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *I32Mat2) ToF64() *F64Mat2 {
	var ret F64Mat2

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *I32Mat2) Set(col, row int, val int32) {
	mat[col*2+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *I32Mat2) Get(col, row int) int32 {
	return mat[col*2+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *I32Mat2) MarshalJSON() ([]byte, error) {
	var tmpArray [2][2]int32

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*2+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal I32Mat2")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *I32Mat2) UnmarshalJSON(data []byte) error {
	var tmpArray [2][2]int32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal I32Mat2")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*2+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *I32Mat2) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *I32Mat2) Add(op *I32Mat2) *I32Mat2 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *I32Mat2) Sub(op *I32Mat2) *I32Mat2 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// I32Mat2Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func I32Mat2Add(mata, matb *I32Mat2) *I32Mat2 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// I32Mat2Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func I32Mat2Sub(mata, matb *I32Mat2) *I32Mat2 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}
