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

// F64Mat3 is a matrix containing 3x3 float64 values.
// Can hold the values of a point in space.
type F64Mat3 [9]float64

// F64Mat3New creates a new matrix containing 3x3 float64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F64Mat3New(f1, f2, f3, f4, f5, f6, f7, f8, f9 float64) *F64Mat3 {
	return &F64Mat3{f1, f2, f3, f4, f5, f6, f7, f8, f9}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F64Mat3) ToI32() *I32Mat3 {
	var ret I32Mat3

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F64Mat3) ToI64() *I64Mat3 {
	var ret I64Mat3

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F64Mat3) ToX32() *X32Mat3 {
	var ret X32Mat3

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F64Mat3) ToX64() *X64Mat3 {
	var ret X64Mat3

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *F64Mat3) ToF32() *F32Mat3 {
	var ret F32Mat3

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// Add adds operand to the matrix.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat3) Add(op *F64Mat3) *F64Mat3 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat3) Sub(op *F64Mat3) *F64Mat3 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat3) MulScale(factor float64) *F64Mat3 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat3) DivScale(factor float64) *F64Mat3 {
	for i, v := range mat {
		mat[i] = vpnumber.F64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrixs are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F64Mat3) IsSimilar(op *F64Mat3) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// F64Mat3Add adds two matrixs.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat3Add(mata, matb *F64Mat3) *F64Mat3 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F64Mat3Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat3Sub(mata, matb *F64Mat3) *F64Mat3 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F64Mat3MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat3MulScale(mat *F64Mat3, factor float64) *F64Mat3 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F64Mat3DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat3DivScale(mat *F64Mat3, factor float64) *F64Mat3 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F64Mat3IsSimilar returns true if matrixs are approximatively the same.
// This is a workarround to ignore rounding errors.
func F64Mat3IsSimilar(mata, matb *F64Mat3) bool {
	return mata.IsSimilar(matb)
}
