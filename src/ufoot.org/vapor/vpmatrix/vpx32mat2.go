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

// X32Mat2 is a matrix containing 2x2 fixed point 32 bit values.
// Can hold the values of a point in a plane.
type X32Mat2 [4]vpnumber.X32

// X32Mat2New creates a new matrix containing 2x2 fixed point 32 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X32Mat2New(x1, x2, x3, x4 vpnumber.X32) *X32Mat2 {
	return &X32Mat2{x1, x2, x3, x4}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *X32Mat2) ToI32() *I32Mat2 {
	var ret I32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X32ToI32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int32 matrix.
func (mat *X32Mat2) ToI64() *I64Mat2 {
	var ret I64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X32ToI64(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X32Mat2) ToX64() *X64Mat2 {
	var ret X64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X32Mat2) ToF32() *F32Mat2 {
	var ret F32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X32Mat2) ToF64() *F64Mat2 {
	var ret F64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// Add adds operand to the matrix.
// It modifies it, and returns a pointer on it.
func (mat *X32Mat2) Add(op *X32Mat2) *X32Mat2 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies it, and returns a pointer on it.
func (mat *X32Mat2) Sub(op *X32Mat2) *X32Mat2 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies it, and returns a pointer on it.
func (mat *X32Mat2) MulScale(factor vpnumber.X32) *X32Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies it, and returns a pointer on it.
func (mat *X32Mat2) DivScale(factor vpnumber.X32) *X32Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrixs are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X32Mat2) IsSimilar(op *X32Mat2) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// X32Mat2Add adds two matrixs.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat2Add(mata, matb *X32Mat2) *X32Mat2 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X32Mat2Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat2Sub(mata, matb *X32Mat2) *X32Mat2 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X32Mat2MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat2MulScale(mat *X32Mat2, factor vpnumber.X32) *X32Mat2 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X32Mat2DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat2DivScale(mat *X32Mat2, factor vpnumber.X32) *X32Mat2 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X32Mat2IsSimilar returns true if matrixs are approximatively the same.
// This is a workarround to ignore rounding errors.
func X32Mat2IsSimilar(mata, matb *X32Mat2) bool {
	return mata.IsSimilar(matb)
}
