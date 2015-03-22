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
	"encoding/json"
	"ufoot.org/vapor/vpnumber"
	"ufoot.org/vapor/vpsys"
)

// F64Mat2 is a matrix containing 2x2 float64 values.
// Can hold the values of a point in a plane.
type F64Mat2 [4]float64

// F64Mat2New creates a new matrix containing 2x2 float64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F64Mat2New(f1, f2, f3, f4 float64) *F64Mat2 {
	return &F64Mat2{f1, f2, f3, f4}
}

// F64Mat2Identity creates a new identity matrix.
func F64Mat2Identity() *F64Mat2 {
	return &F64Mat2{vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1}
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

// ToF32 converts the matrix to a float32 matrix.
func (mat *F64Mat2) ToF32() *F32Mat2 {
	var ret F32Mat2

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F64Mat2) Set(col, row int, val float64) {
	mat[col*2+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F64Mat2) Get(col, row int) float64 {
	return mat[col*2+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F64Mat2) MarshalJSON() ([]byte, error) {
	var tmpArray [2][2]float64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*2+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F64Mat2")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F64Mat2) UnmarshalJSON(data []byte) error {
	var tmpArray [2][2]float64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F64Mat2")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*2+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F64Mat2) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2) Add(op *F64Mat2) *F64Mat2 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2) Sub(op *F64Mat2) *F64Mat2 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2) MulScale(factor float64) *F64Mat2 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2) DivScale(factor float64) *F64Mat2 {
	for i, v := range mat {
		mat[i] = vpnumber.F64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F64Mat2) IsSimilar(op *F64Mat2) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2) MulComp(op *F64Mat2) *F64Mat2 {
	*mat = *F64Mat2MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F64Mat2) Det() float64 {
	return mat.Get(0, 0)*mat.Get(1, 1) - mat.Get(0, 1)*mat.Get(1, 0)
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64Mat2) Inv() *F64Mat2 {
	*mat = *F64Mat2Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 2x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *F64Mat2) MulVec(vec *F64Vec2) *F64Vec2 {
	var ret F64Vec2

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1]
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 2x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *F64Mat2) MulVecPos(vec float64) float64 {
	return mat.Get(0, 0)*vec + mat.Get(1, 0)
}

// MulVecDir performs a multiplication of a vector by a 2x2 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 1 (here, a scalar) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F64Mat2) MulVecDir(vec float64) float64 {
	return mat.Get(0, 0) * vec
}

// F64Mat2Add adds two matrices.
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

// F64Mat2IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func F64Mat2IsSimilar(mata, matb *F64Mat2) bool {
	return mata.IsSimilar(matb)
}

// F64Mat2MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F64Mat2MulComp(a, b *F64Mat2) *F64Mat2 {
	var ret F64Mat2

	for c := 0; c < 2; c++ {
		for r := 0; r < 2; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1))
		}
	}

	return &ret
}

// F64Mat2Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F64Mat2Inv(mat *F64Mat2) *F64Mat2 {
	ret := F64Mat2{
		mat.Get(1, 1),
		-mat.Get(0, 1),
		-mat.Get(1, 0),
		mat.Get(0, 0),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
