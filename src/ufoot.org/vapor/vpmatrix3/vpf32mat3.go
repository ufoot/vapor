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

package vpmatrix3

import (
	"encoding/json"
	"ufoot.org/vapor/vpnumber"
	"ufoot.org/vapor/vpsys"
	"ufoot.org/vapor/vpmatrix2"
)

// F32Mat3 is a matrix containing 3x3 float32 values.
// Can hold the values of a point in space.
type F32Mat3 [9]float32

// F32Mat3New creates a new matrix containing 3x3 float32 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F32Mat3New(f1, f2, f3, f4, f5, f6, f7, f8, f9 float32) *F32Mat3 {
	return &F32Mat3{f1, f2, f3, f4, f5, f6, f7, f8, f9}
}

// F32Mat3Identity creates a new identity matrix.
func F32Mat3Identity() *F32Mat3 {
	return &F32Mat3{vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1}
}

// ToI32 converts the matrix to an int32 matrix.
func (mat *F32Mat3) ToI32() *I32Mat3 {
	var ret I32Mat3

	for i, v := range mat {
		ret[i] = int32(v)
	}

	return &ret
}

// ToI64 converts the matrix to an int64 matrix.
func (mat *F32Mat3) ToI64() *I64Mat3 {
	var ret I64Mat3

	for i, v := range mat {
		ret[i] = int64(v)
	}

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F32Mat3) ToX32() *X32Mat3 {
	var ret X32Mat3

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F32Mat3) ToX64() *X64Mat3 {
	var ret X64Mat3

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *F32Mat3) ToF64() *F64Mat3 {
	var ret F64Mat3

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F32Mat3) Set(col, row int, val float32) {
	mat[col*3+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F32Mat3) Get(col, row int) float32 {
	return mat[col*3+row]
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F32Mat3) MarshalJSON() ([]byte, error) {
	var tmpArray [3][3]float32

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*3+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F32Mat3")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F32Mat3) UnmarshalJSON(data []byte) error {
	var tmpArray [3][3]float32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F32Mat3")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*3+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F32Mat3) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat3) Add(op *F32Mat3) *F32Mat3 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat3) Sub(op *F32Mat3) *F32Mat3 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat3) MulScale(factor float32) *F32Mat3 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat3) DivScale(factor float32) *F32Mat3 {
	for i, v := range mat {
		mat[i] = vpnumber.F32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F32Mat3) IsSimilar(op *F32Mat3) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat3) MulComp(op *F32Mat3) *F32Mat3 {
	*mat = *F32Mat3MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F32Mat3) Det() float32 {
	return mat.Get(0, 0)*mat.Get(1, 1)*mat.Get(2, 2) + mat.Get(0, 1)*mat.Get(1, 2)*mat.Get(2, 0) + mat.Get(0, 2)*mat.Get(1, 0)*mat.Get(2, 1) - mat.Get(0, 0)*mat.Get(1, 2)*mat.Get(2, 1) - mat.Get(0, 1)*mat.Get(1, 0)*mat.Get(2, 2) - mat.Get(0, 2)*mat.Get(1, 1)*mat.Get(2, 0)
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32Mat3) Inv() *F32Mat3 {
	*mat = *F32Mat3Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *F32Mat3) MulVec(vec *F32Vec3) *F32Vec3 {
	var ret F32Vec3

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2]
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *F32Mat3) MulVecPos(vec *vpmatrix2.F32Vec2) *vpmatrix2.F32Vec2 {
	var ret vpmatrix2.F32Vec2

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)
	}

	return &ret
}

// MulVecDir performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F32Mat3) MulVecDir(vec *vpmatrix2.F32Vec2) *vpmatrix2.F32Vec2 {
	var ret vpmatrix2.F32Vec2

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1]
	}

	return &ret
}

// F32Mat3Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat3Add(mata, matb *F32Mat3) *F32Mat3 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F32Mat3Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat3Sub(mata, matb *F32Mat3) *F32Mat3 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F32Mat3MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat3MulScale(mat *F32Mat3, factor float32) *F32Mat3 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F32Mat3DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32Mat3DivScale(mat *F32Mat3, factor float32) *F32Mat3 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F32Mat3IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func F32Mat3IsSimilar(mata, matb *F32Mat3) bool {
	return mata.IsSimilar(matb)
}

// F32Mat3MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F32Mat3MulComp(a, b *F32Mat3) *F32Mat3 {
	var ret F32Mat3

	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1)+a.Get(2, r)*b.Get(c, 2))
		}
	}

	return &ret
}

// F32Mat3Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F32Mat3Inv(mat *F32Mat3) *F32Mat3 {
	ret := F32Mat3{
		mat.Get(1, 1)*mat.Get(2, 2) - mat.Get(1, 2)*mat.Get(2, 1),
		mat.Get(0, 2)*mat.Get(2, 1) - mat.Get(0, 1)*mat.Get(2, 2),
		mat.Get(0, 1)*mat.Get(1, 2) - mat.Get(0, 2)*mat.Get(1, 1),
		mat.Get(1, 2)*mat.Get(2, 0) - mat.Get(1, 0)*mat.Get(2, 2),
		mat.Get(0, 0)*mat.Get(2, 2) - mat.Get(0, 2)*mat.Get(2, 0),
		mat.Get(0, 2)*mat.Get(1, 0) - mat.Get(0, 0)*mat.Get(1, 2),
		mat.Get(1, 0)*mat.Get(2, 1) - mat.Get(1, 1)*mat.Get(2, 0),
		mat.Get(0, 1)*mat.Get(2, 0) - mat.Get(0, 0)*mat.Get(2, 1),
		mat.Get(0, 0)*mat.Get(1, 1) - mat.Get(0, 1)*mat.Get(1, 0),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}