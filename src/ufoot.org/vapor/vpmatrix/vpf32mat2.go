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

// F32Mat2 is a matrix containing 2x2 float32 values.
// Can hold the values of a point in a plane.
type F32Mat2 [4]float32

// F32Mat2New creates a new matrix containing 2x2 float32 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F32Mat2New(f1, f2, f3, f4 float32) *F32Mat2 {
	return &F32Mat2{f1, f2, f3, f4}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F32Mat2) ToI32() *I32Mat2 {
	var ret I32Mat2

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F32Mat2) ToI64() *I64Mat2 {
	var ret I64Mat2

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F32Mat2) ToX32() *X32Mat2 {
	var ret X32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F32Mat2) ToX64() *X64Mat2 {
	var ret X64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *F32Mat2) ToF64() *F64Mat2 {
	var ret F64Mat2

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F32Mat2) Set(col, row int, val float32) {
	mat[col*2+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F32Mat2) Get(col, row int) float32 {
	return mat[col*2+row]
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat2) Add(op *F32Mat2) *F32Mat2 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat2) Sub(op *F32Mat2) *F32Mat2 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat2) MulScale(factor float32) *F32Mat2 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat2) DivScale(factor float32) *F32Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.F32Div(v, factor)
	}

	return mat
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat2) MulComp(op *F32Mat2) *F32Mat2 {
	*mat = *F32Mat2MulComp(mat, op)

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F32Mat2) IsSimilar(op *F32Mat2) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// F32Mat2Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat2Add(mata, matb *F32Mat2) *F32Mat2 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F32Mat2Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat2Sub(mata, matb *F32Mat2) *F32Mat2 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F32Mat2MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat2MulScale(mat *F32Mat2, factor float32) *F32Mat2 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F32Mat2DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat2DivScale(mat *F32Mat2, factor float32) *F32Mat2 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F32Mat2IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func F32Mat2IsSimilar(mata, matb *F32Mat2) bool {
	return mata.IsSimilar(matb)
}

// MulComp multiplies two matrices (composition).
// It modifies the matrix, and returns a pointer on it.
func F32Mat2MulComp(a, b *F32Mat2) *F32Mat2 {
	var ret F32Mat2

	for c := 0; c < 2; c++ {
		for r := 0; r < 2; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1))
		}
	}

	return &ret
}
