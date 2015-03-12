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

// F32Mat4 is a matrix containing 4x4 float32 values.
// Can be used in 3D matrix transformations.
type F32Mat4 [16]float32

// F32Mat4New creates a new matrix containing 4x4 float32 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F32Mat4New(f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16 float32) *F32Mat4 {
	return &F32Mat4{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F32Mat4) ToI32() *I32Mat4 {
	var ret I32Mat4

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F32Mat4) ToI64() *I64Mat4 {
	var ret I64Mat4

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F32Mat4) ToX32() *X32Mat4 {
	var ret X32Mat4

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F32Mat4) ToX64() *X64Mat4 {
	var ret X64Mat4

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *F32Mat4) ToF64() *F64Mat4 {
	var ret F64Mat4

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F32Mat4) Set(col, row int, val float32) {
	mat[col*4+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F32Mat4) Get(col, row int) float32 {
	return mat[col*4+row]
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat4) Add(op *F32Mat4) *F32Mat4 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat4) Sub(op *F32Mat4) *F32Mat4 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat4) MulScale(factor float32) *F32Mat4 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat4) DivScale(factor float32) *F32Mat4 {
	for i, v := range mat {
		mat[i] = vpnumber.F32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F32Mat4) IsSimilar(op *F32Mat4) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat4) MulComp(op *F32Mat4) *F32Mat4 {
	*mat = *F32Mat4MulComp(mat, op)

	return mat
}

// MulCol performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// It modifies the vector, and returns a pointer on it.
func (mat *F32Mat4) MulVec(vec *F32Vec4) *F32Vec4 {
	var ret F32Vec4
	var i int

	for i, _ = range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2] + mat.Get(3, i)*vec[3]
	}

	return &ret
}

// MulCol1 performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
// It modifies the vector, and returns a pointer on it.
func (mat *F32Mat4) MulVec1(vec *F32Vec3) *F32Vec3 {
	var ret F32Vec3
	var i int

	for i, _ = range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2] + mat.Get(3, i)
	}

	return &ret
}

// F32Mat4Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat4Add(mata, matb *F32Mat4) *F32Mat4 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F32Mat4Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat4Sub(mata, matb *F32Mat4) *F32Mat4 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F32Mat4MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat4MulScale(mat *F32Mat4, factor float32) *F32Mat4 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F32Mat4DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat4DivScale(mat *F32Mat4, factor float32) *F32Mat4 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F32Mat4IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func F32Mat4IsSimilar(mata, matb *F32Mat4) bool {
	return mata.IsSimilar(matb)
}

// MulComp multiplies two matrices (composition).
// It modifies the matrix, and returns a pointer on it.
func F32Mat4MulComp(a, b *F32Mat4) *F32Mat4 {
	var ret F32Mat4

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1)+a.Get(2, r)*b.Get(c, 2)+a.Get(3, r)*b.Get(c, 3))
		}
	}

	return &ret
}
