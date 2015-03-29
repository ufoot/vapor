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

package vpmatrix2

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
)

// X64Mat2 is a matrix containing 2x2 fixed point 64 bit values.
// Can hold the values of a point in a plane.
type X64Mat2 [4]vpnumber.X64

// X64Mat2New creates a new matrix containing 2x2 fixed point 64 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X64Mat2New(x1, x2, x3, x4 vpnumber.X64) *X64Mat2 {
	return &X64Mat2{x1, x2, x3, x4}
}

// X64Mat2Identity creates a new identity matrix.
func X64Mat2Identity() *X64Mat2 {
	return &X64Mat2{vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1}
}

// X64Mat2Trans creates a new translation matrix.
func X64Mat2Trans(x vpnumber.X64) *X64Mat2 {
	return &X64Mat2{vpnumber.X64Const1, vpnumber.X64Const0, x, vpnumber.X64Const1}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *X64Mat2) ToI32() *I32Mat2 {
	var ret I32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X64ToI32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int32 matrix.
func (mat *X64Mat2) ToI64() *I64Mat2 {
	var ret I64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X64ToI64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X64Mat2) ToX32() *X32Mat2 {
	var ret X32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X64ToX32(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X64Mat2) ToF32() *F32Mat2 {
	var ret F32Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X64Mat2) ToF64() *F64Mat2 {
	var ret F64Mat2

	for i, v := range mat {
		ret[i] = vpnumber.X64ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X64Mat2) Set(col, row int, val vpnumber.X64) {
	mat[col*2+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X64Mat2) Get(col, row int) vpnumber.X64 {
	return mat[col*2+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *X64Mat2) MarshalJSON() ([]byte, error) {
	var tmpArray [2][2]int64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = int64(mat[col*2+row])
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X64Mat2")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *X64Mat2) UnmarshalJSON(data []byte) error {
	var tmpArray [2][2]int64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X64Mat2")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*2+row] = vpnumber.X64(tmpArray[col][row])
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *X64Mat2) String() string {
	buf, err := mat.ToF64().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2) Add(op *X64Mat2) *X64Mat2 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2) Sub(op *X64Mat2) *X64Mat2 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2) MulScale(factor vpnumber.X64) *X64Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2) DivScale(factor vpnumber.X64) *X64Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.X64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X64Mat2) IsSimilar(op *X64Mat2) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X64IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2) MulComp(op *X64Mat2) *X64Mat2 {
	*mat = *X64Mat2MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *X64Mat2) Det() vpnumber.X64 {
	return vpnumber.X64Mul(mat.Get(0, 0), mat.Get(1, 1)) - vpnumber.X64Mul(mat.Get(0, 1), mat.Get(1, 0))
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *X64Mat2) Inv() *X64Mat2 {
	*mat = *X64Mat2Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 2x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *X64Mat2) MulVec(vec *X64Vec2) *X64Vec2 {
	var ret X64Vec2

	for i := range vec {
		ret[i] = vpnumber.X64Mul(mat.Get(0, i), vec[0]) + vpnumber.X64Mul(mat.Get(1, i), vec[1])
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 2x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *X64Mat2) MulVecPos(vec vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Mul(mat.Get(0, 0), vec) + mat.Get(1, 0)
}

// MulVecDir performs a multiplication of a vector by a 2x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X64Mat2) MulVecDir(vec vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Mul(mat.Get(0, 0), vec)
}

// X64Mat2Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2Add(mata, matb *X64Mat2) *X64Mat2 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X64Mat2Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2Sub(mata, matb *X64Mat2) *X64Mat2 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X64Mat2MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2MulScale(mat *X64Mat2, factor vpnumber.X64) *X64Mat2 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X64Mat2DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2DivScale(mat *X64Mat2, factor vpnumber.X64) *X64Mat2 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X64Mat2MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X64Mat2MulComp(a, b *X64Mat2) *X64Mat2 {
	var ret X64Mat2

	for c := 0; c < 2; c++ {
		for r := 0; r < 2; r++ {
			ret.Set(c, r, vpnumber.X64Mul(a.Get(0, r), b.Get(c, 0))+vpnumber.X64Mul(a.Get(1, r), b.Get(c, 1)))
		}
	}

	return &ret
}

// X64Mat2Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func X64Mat2Inv(mat *X64Mat2) *X64Mat2 {
	ret := X64Mat2{
		mat.Get(1, 1),
		-mat.Get(0, 1),
		-mat.Get(1, 0),
		mat.Get(0, 0),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
