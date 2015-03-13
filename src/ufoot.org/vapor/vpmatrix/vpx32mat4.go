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

// X32Mat4 is a matrix containing 4x4 fixed point 32 bit values.
// Can be used in 3D matrix transformations.
type X32Mat4 [16]vpnumber.X32

// X32Mat4New creates a new matrix containing 4x4 fixed point 32 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X32Mat4New(x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16 vpnumber.X32) *X32Mat4 {
	return &X32Mat4{x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *X32Mat4) ToI32() *I32Mat4 {
	var ret I32Mat4

	for i, v := range mat {
		ret[i] = vpnumber.X32ToI32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int32 matrix.
func (mat *X32Mat4) ToI64() *I64Mat4 {
	var ret I64Mat4

	for i, v := range mat {
		ret[i] = vpnumber.X32ToI64(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X32Mat4) ToX64() *X64Mat4 {
	var ret X64Mat4

	for i, v := range mat {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X32Mat4) ToF32() *F32Mat4 {
	var ret F32Mat4

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X32Mat4) ToF64() *F64Mat4 {
	var ret F64Mat4

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X32Mat4) Set(col, row int, val vpnumber.X32) {
	mat[col*4+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X32Mat4) Get(col, row int) vpnumber.X32 {
	return mat[col*4+row]
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat4) Add(op *X32Mat4) *X32Mat4 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat4) Sub(op *X32Mat4) *X32Mat4 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat4) MulScale(factor vpnumber.X32) *X32Mat4 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat4) DivScale(factor vpnumber.X32) *X32Mat4 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X32Mat4) IsSimilar(op *X32Mat4) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X32Mat4) MulComp(op *X32Mat4) *X32Mat4 {
	*mat = *X32Mat4MulComp(mat, op)

	return mat
}

// MulVec performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *X32Mat4) MulVec(vec *X32Vec4) *X32Vec4 {
	var ret X32Vec4

	for i:= range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + vpnumber.X32Mul(mat.Get(2, i), vec[2]) + vpnumber.X32Mul(mat.Get(3, i), vec[3])
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *X32Mat4) MulVecPos(vec *X32Vec3) *X32Vec3 {
	var ret X32Vec3

	for i:= range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + vpnumber.X32Mul(mat.Get(2, i), vec[2]) + mat.Get(3, i)
	}

	return &ret
}

// MulVecDir performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X32Mat4) MulVecDir(vec *X32Vec3) *X32Vec3 {
	var ret X32Vec3

	for i:= range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + vpnumber.X32Mul(mat.Get(2, i), vec[2])
	}

	return &ret
}

// X32Mat4Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat4Add(mata, matb *X32Mat4) *X32Mat4 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X32Mat4Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat4Sub(mata, matb *X32Mat4) *X32Mat4 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X32Mat4MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat4MulScale(mat *X32Mat4, factor vpnumber.X32) *X32Mat4 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X32Mat4DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32Mat4DivScale(mat *X32Mat4, factor vpnumber.X32) *X32Mat4 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X32Mat4IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func X32Mat4IsSimilar(mata, matb *X32Mat4) bool {
	return mata.IsSimilar(matb)
}

// X32Mat4MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X32Mat4MulComp(a, b *X32Mat4) *X32Mat4 {
	var ret X32Mat4

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			ret.Set(c, r, vpnumber.X32Mul(a.Get(0, r), b.Get(c, 0))+vpnumber.X32Mul(a.Get(1, r), b.Get(c, 1))+vpnumber.X32Mul(a.Get(2, r), b.Get(c, 2))+vpnumber.X32Mul(a.Get(3, r), b.Get(c, 3)))
		}
	}

	return &ret
}
