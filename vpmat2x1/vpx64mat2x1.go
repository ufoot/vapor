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

package vpmat2x1

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
)

// X64Mat2x1 is a matrix containing 2x1 fixed point 64 bit values.
// Can hold the values of a point in a plane.
type X64Mat2x1 [2]vpnumber.X64

// X64Mat2x1New creates a new matrix containing 2x1 fixed point 64 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X64Mat2x1New(x1, x2 vpnumber.X64) *X64Mat2x1 {
	return &X64Mat2x1{x1, x2}
}

// X64Mat2x1Identity creates a new identity matrix.
func X64Mat2x1Identity() *X64Mat2x1 {
	return &X64Mat2x1{vpnumber.X64Const1, vpnumber.X64Const0}
}

// X64Mat2x1Trans creates a new translation matrix.
func X64Mat2x1Trans(x vpnumber.X64) *X64Mat2x1 {
	return &X64Mat2x1{vpnumber.X64Const1, x}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *X64Mat2x1) ToI32() *I32Mat2x1 {
	var ret I32Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.X64ToI32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int32 matrix.
func (mat *X64Mat2x1) ToI64() *I64Mat2x1 {
	var ret I64Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.X64ToI64(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X64Mat2x1) ToX32() *X32Mat2x1 {
	var ret X32Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X64Mat2x1) ToF32() *F32Mat2x1 {
	var ret F32Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X64Mat2x1) ToF64() *F64Mat2x1 {
	var ret F64Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X64Mat2x1) Set(col, row int, val vpnumber.X64) {
	mat[col+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X64Mat2x1) Get(col, row int) vpnumber.X64 {
	return mat[col+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *X64Mat2x1) MarshalJSON() ([]byte, error) {
	var tmpArray [2][1]int64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = int64(mat[col+row])
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X64Mat2x1")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *X64Mat2x1) UnmarshalJSON(data []byte) error {
	var tmpArray [2][1]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X64Mat2x1")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col+row] = vpnumber.X64(tmpArray[col][row])
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *X64Mat2x1) String() string {
	buf, err := mat.ToF64().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2x1) Add(op *X64Mat2x1) *X64Mat2x1 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2x1) Sub(op *X64Mat2x1) *X64Mat2x1 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2x1) MulScale(factor vpnumber.X64) *X64Mat2x1 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2x1) DivScale(factor vpnumber.X64) *X64Mat2x1 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X64Mat2x1) IsSimilar(op *X64Mat2x1) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2x1) MulComp(op *X64Mat2x1) *X64Mat2x1 {
	*mat = *X64Mat2x1MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *X64Mat2x1) Det() vpnumber.X64 {
	return mat.Get(0, 0)
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2x1) Inv() *X64Mat2x1 {
	*mat = *X64Mat2x1Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 2x1 matrix,
// MulVecPos performs a multiplication of a vector by a 2x1 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *X64Mat2x1) MulVecPos(vec vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Mul(mat.Get(0, 0), vec) + mat.Get(1, 0)
}

// MulVecDir performs a multiplication of a vector by a 2x1 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X64Mat2x1) MulVecDir(vec vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Mul(mat.Get(0, 0), vec)
}

// X64Mat2x1Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2x1Add(mata, matb *X64Mat2x1) *X64Mat2x1 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X64Mat2x1Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2x1Sub(mata, matb *X64Mat2x1) *X64Mat2x1 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X64Mat2x1MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2x1MulScale(mat *X64Mat2x1, factor vpnumber.X64) *X64Mat2x1 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X64Mat2x1DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2x1DivScale(mat *X64Mat2x1, factor vpnumber.X64) *X64Mat2x1 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X64Mat2x1MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2x1MulComp(a, b *X64Mat2x1) *X64Mat2x1 {
	var ret X64Mat2x1

	ret.Set(0, 0, vpnumber.X64Mul(a.Get(0, 0), b.Get(0, 0)))
	ret.Set(1, 0, vpnumber.X64Mul(a.Get(0, 0), b.Get(1, 0))+a.Get(1, 0))

	return &ret
}

// X64Mat2x1Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func X64Mat2x1Inv(mat *X64Mat2x1) *X64Mat2x1 {
	ret := X64Mat2x1{
		vpnumber.X64Const1,
		-mat.Get(1, 0),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
