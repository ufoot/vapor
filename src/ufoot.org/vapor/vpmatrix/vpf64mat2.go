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

// F64Mat2 is a matrix containing 2x2 float64 values.
// Can hold the values of a point in a plane.
type F64Mat2 [4]float64

// F64Mat2New creates a new matrix containing 2x2 float64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F64Mat2New(f1,f2,f3,f4 float64) *F64Mat4 {
	return &F64Mat4{f1,f2,f3,f4}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F64Mat2) ToI32() *I32Mat2 {
	var ret I32Mat2

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F64Mat2) ToI64() *I64Mat2 {
	var ret I64Mat2

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F64Mat2) ToX32() *X32Mat2 {
	var ret X32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F64Mat2) ToX64() *X64Mat2 {
	var ret X64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float32 matrix.
func (mat *F64Mat2) ToF32() *F32Mat2 {
	var ret F32Mat2

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// Add adds operand to the matrix.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat2) Add(op *F64Mat2) *F64Mat2 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat2) Sub(op *F64Mat2) *F64Mat2 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat2) MulScale(factor float64) *F64Mat2 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies it, and returns a pointer on it.
func (mat *F64Mat2) DivScale(factor float64) *F64Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.F64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrixs are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F64Mat2) IsSimilar(op *F64Mat2) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// F64Mat2Add adds two matrixs.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2Add(mata, matb *F64Mat2) *F64Mat2 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F64Mat2Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2Sub(mata, matb *F64Mat2) *F64Mat2 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F64Mat2MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2MulScale(mat *F64Mat2, factor float64) *F64Mat2 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F64Mat2DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2DivScale(mat *F64Mat2, factor float64) *F64Mat2 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F64Mat2IsSimilar returns true if matrixs are approximatively the same.
// This is a workarround to ignore rounding errors.
func F64Mat2IsSimilar(mata, matb *F64Mat2) bool {
	return mata.IsSimilar(matb)
}
