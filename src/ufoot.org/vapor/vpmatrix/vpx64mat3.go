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
	"ufoot.org/vapor/vpnumber"
)

// X64Mat3 is a matrix containing 3x3 fixed point 64 bit values.
// Can hold the values of a point in space.
type X64Mat3 [9]vpnumber.X64

// X64Mat3New creates a new matrix containing 3x3 fixed point 64 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X64Mat3New(x1, x2, x3, x4, x5, x6, x7, x8, x9 vpnumber.X64) *X64Mat3 {
	return &X64Mat3{x1, x2, x3, x4, x5, x6, x7, x8, x9}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *X64Mat3) ToI32() *I32Mat3 {
	var ret I32Mat3

	for i, v := range mat {
		ret[i] = vpnumber.X64ToI32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int32 matrix.
func (mat *X64Mat3) ToI64() *I64Mat3 {
	var ret I64Mat3

	for i, v := range mat {
		ret[i] = vpnumber.X64ToI64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X64Mat3) ToX32() *X32Mat3 {
	var ret X32Mat3

	for i, v := range mat {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X64Mat3) ToF32() *F32Mat3 {
	var ret F32Mat3

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X64Mat3) ToF64() *F64Mat3 {
	var ret F64Mat3

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X64Mat3) Set(col, row int, val vpnumber.X64){
	mat[col*3+row]=val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X64Mat3) Get(col, row int) vpnumber.X64 {
	return mat[col*3+row]
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat3) Add(op *X64Mat3) *X64Mat3 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat3) Sub(op *X64Mat3) *X64Mat3 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat3) MulScale(factor vpnumber.X64) *X64Mat3 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat3) DivScale(factor vpnumber.X64) *X64Mat3 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X64Mat3) IsSimilar(op *X64Mat3) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// X64Mat3Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat3Add(mata, matb *X64Mat3) *X64Mat3 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X64Mat3Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat3Sub(mata, matb *X64Mat3) *X64Mat3 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X64Mat3MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat3MulScale(mat *X64Mat3, factor vpnumber.X64) *X64Mat3 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X64Mat3DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat3DivScale(mat *X64Mat3, factor vpnumber.X64) *X64Mat3 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X64Mat3IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func X64Mat3IsSimilar(mata, matb *X64Mat3) bool {
	return mata.IsSimilar(matb)
}
