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

// F64Mat2x1 is a matrix containing 2x1 float64 values.
// Can hold the values of a point in a plane.
type F64Mat2x1 [2]float64

// F64Mat2x1New creates a new matrix containing 2x1 float64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F64Mat2x1New(f1, f2 float64) *F64Mat2x1 {
	return &F64Mat2x1{f1, f2}
}

// F64Mat2x1Identity creates a new identity matrix.
func F64Mat2x1Identity() *F64Mat2x1 {
	return &F64Mat2x1{vpnumber.F64Const1, vpnumber.F64Const0}
}

// F64Mat2x1Trans creates a new translation matrix.
func F64Mat2x1Trans(f float64) *F64Mat2x1 {
	return &F64Mat2x1{vpnumber.F64Const1, f}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F64Mat2x1) ToI32() *I32Mat2x1 {
	var ret I32Mat2x1

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F64Mat2x1) ToI64() *I64Mat2x1 {
	var ret I64Mat2x1

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F64Mat2x1) ToX32() *X32Mat2x1 {
	var ret X32Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F64Mat2x1) ToX64() *X64Mat2x1 {
	var ret X64Mat2x1

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *F64Mat2x1) ToF32() *F32Mat2x1 {
	var ret F32Mat2x1

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F64Mat2x1) Set(col, row int, val float64) {
	mat[col+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F64Mat2x1) Get(col, row int) float64 {
	return mat[col+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F64Mat2x1) MarshalJSON() ([]byte, error) {
	var tmpArray [2][1]float64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F64Mat2x1")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F64Mat2x1) UnmarshalJSON(data []byte) error {
	var tmpArray [2][1]float64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F64Mat2x1")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F64Mat2x1) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2x1) Add(op *F64Mat2x1) *F64Mat2x1 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2x1) Sub(op *F64Mat2x1) *F64Mat2x1 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2x1) MulScale(factor float64) *F64Mat2x1 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2x1) DivScale(factor float64) *F64Mat2x1 {
	for i, v := range mat {
		mat[i] = vpnumber.F64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F64Mat2x1) IsSimilar(op *F64Mat2x1) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2x1) MulComp(op *F64Mat2x1) *F64Mat2x1 {
	*mat = *F64Mat2x1MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F64Mat2x1) Det() float64 {
	return mat.Get(0, 0)
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2x1) Inv() *F64Mat2x1 {
	*mat = *F64Mat2x1Inv(mat)

	return mat
}

// MulVecPos performs a multiplication of a vector by a 2x1 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *F64Mat2x1) MulVecPos(vec float64) float64 {
	return mat.Get(0, 0)*vec + mat.Get(1, 0)
}

// MulVecDir performs a multiplication of a vector by a 2x1 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F64Mat2x1) MulVecDir(vec float64) float64 {
	return mat.Get(0, 0) * vec
}

// F64Mat2x1Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2x1Add(mata, matb *F64Mat2x1) *F64Mat2x1 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F64Mat2x1Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2x1Sub(mata, matb *F64Mat2x1) *F64Mat2x1 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F64Mat2x1MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2x1MulScale(mat *F64Mat2x1, factor float64) *F64Mat2x1 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F64Mat2x1DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2x1DivScale(mat *F64Mat2x1, factor float64) *F64Mat2x1 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F64Mat2x1MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2x1MulComp(a, b *F64Mat2x1) *F64Mat2x1 {
	var ret F64Mat2x1

	ret.Set(0, 0, a.Get(0, 0)*b.Get(0, 0))
	ret.Set(1, 0, a.Get(0, 0)*b.Get(1, 0)+a.Get(1, 0))

	return &ret
}

// F64Mat2x1Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F64Mat2x1Inv(mat *F64Mat2x1) *F64Mat2x1 {
	ret := F64Mat2x1{
		vpnumber.F64Const1,
		-mat.Get(1, 0),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}