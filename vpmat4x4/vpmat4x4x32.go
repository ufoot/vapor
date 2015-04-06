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

package vpmat4x4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
	"github.com/ufoot/vapor/vpvec3"
	"github.com/ufoot/vapor/vpvec4"
)

// X32 is a matrix containing 4x4 fixed point 32 bit values.
// Can be used in 3D matrix transformations.
type X32 [Size]vpnumber.X32

// X32New creates a new matrix containing 4x4 fixed point 32 bit values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func X32New(x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16 vpnumber.X32) *X32 {
	return &X32{x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16}
}

// X32Identity creates a new identity matrix.
func X32Identity() *X32 {
	return &X32{vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32Trans creates a new translation matrix.
func X32Trans(vec *vpvec3.X32) *X32 {
	return &X32{vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vec[0], vec[1], vec[2], vpnumber.X32Const1}
}

// X32Scale creates a new scale matrix.
func X32Scale(vec *vpvec3.X32) *X32 {
	return &X32{vec[0], vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vec[1], vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vec[2], vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32RotX creates a new rotation matrix.
// The rotation is done in 3D over the x (1st) axis.
// Angle is given in radians.
func X32RotX(r vpnumber.X32) *X32 {
	return &X32{vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpmath.X32Cos(r), vpmath.X32Sin(r), vpnumber.X32Const0, vpnumber.X32Const0, -vpmath.X32Sin(r), vpmath.X32Cos(r), vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32RotY creates a new rotation matrix.
// The rotation is done in 3D over the y (2nd) axis.
// Angle is given in radians.
func X32RotY(r vpnumber.X32) *X32 {
	return &X32{vpmath.X32Cos(r), vpnumber.X32Const0, -vpmath.X32Sin(r), vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpmath.X32Sin(r), vpnumber.X32Const0, vpmath.X32Cos(r), vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32RotZ creates a new rotation matrix.
// The rotation is done in 3D over the z (3rd) axis.
// Angle is given in radians.
func X32RotZ(r vpnumber.X32) *X32 {
	return &X32{vpmath.X32Cos(r), vpmath.X32Sin(r), vpnumber.X32Const0, vpnumber.X32Const0, -vpmath.X32Sin(r), vpmath.X32Cos(r), vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1}
}

// X32RebaseOXYZ creates a matrix that translates from the default
// O=(0,0,0), X=(1,0,0), Y=(0,1,0), Z=(0,0,1) basis to the given
// basis. It assumes f(a+b) equals f(a)+f(b).
func X32RebaseOXYZ(Origin, PosX, PosY, PosZ *vpvec3.X32) *X32 {
	return &X32{PosX[0] - Origin[0], PosX[1] - Origin[1], PosX[2] - Origin[2], vpnumber.X32Const0, PosY[0] - Origin[0], PosY[1] - Origin[1], PosY[2] - Origin[2], vpnumber.X32Const0, PosZ[0] - Origin[0], PosZ[1] - Origin[1], PosZ[2] - Origin[2], vpnumber.X32Const0, Origin[0], Origin[1], Origin[2], vpnumber.X32Const1}
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *X32) ToX64() *X64 {
	var ret X64

	for i, v := range mat {
		ret[i] = vpnumber.X32ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *X32) ToF32() *F32 {
	var ret F32

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF32(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *X32) ToF64() *F64 {
	var ret F64

	for i, v := range mat {
		ret[i] = vpnumber.X32ToF64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *X32) Set(col, row int, val vpnumber.X32) {
	mat[col*Height+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *X32) Get(col, row int) vpnumber.X32 {
	return mat[col*Height+row]
}

// SetCol sets a column to the values contained in a vector.
func (mat *X32) SetCol(col int, vec *vpvec4.X32) {
	for row, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetCol gets a column and returns it in a vector.
func (mat *X32) GetCol(col int) *vpvec4.X32 {
	var ret vpvec4.X32

	for row := range ret {
		ret[row] = mat[col*Height+row]
	}

	return &ret
}

// SetRow sets a row to the values contained in a vector.
func (mat *X32) SetRow(row int, vec *vpvec4.X32) {
	for col, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetRow gets a row and returns it in a vector.
func (mat *X32) GetRow(row int) *vpvec4.X32 {
	var ret vpvec4.X32

	for col := range ret {
		ret[col] = mat[col*Height+row]
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *X32) MarshalJSON() ([]byte, error) {
	var tmpArray [Width][Height]int32

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = int32(mat[col*Height+row])
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X32")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *X32) UnmarshalJSON(data []byte) error {
	var tmpArray [Width][Height]int32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X32")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*Height+row] = vpnumber.X32(tmpArray[col][row])
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *X32) String() string {
	buf, err := mat.ToF32().MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) Add(op *X32) *X32 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) Sub(op *X32) *X32 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) MulScale(factor vpnumber.X32) *X32 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Mul(v, factor)
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) DivScale(factor vpnumber.X32) *X32 {
	for i, v := range mat {
		mat[i] = vpnumber.X32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *X32) IsSimilar(op *X32) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.X32IsSimilar(v, op[i])
	}

	return ret
}

// Transpose inverts rows and columns (matrix transposition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) Transpose(op *X32) *X32 {
	*mat = *X32Transpose(op)

	return mat
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) MulComp(op *X32) *X32 {
	*mat = *X32MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *X32) Det() vpnumber.X32 {
	return vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row2], mat[Col3Row0]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row2], mat[Col3Row0]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row3], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row3], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row1], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row1], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row1], mat[Col3Row3]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row1], mat[Col3Row3]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row2], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row2], mat[Col3Row3])
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *X32) Inv() *X32 {
	*mat = *X32Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *X32) MulVec(vec *vpvec4.X32) *vpvec4.X32 {
	var ret vpvec4.X32

	for i := range vec {
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
func (mat *X32) MulVecPos(vec *vpvec3.X32) *vpvec3.X32 {
	var ret vpvec3.X32

	for i := range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + vpnumber.X32Mul(mat.Get(2, i), vec[2]) + mat.Get(3, i)
	}

	return ret.DivScale(vpnumber.X32Mul(mat[Col0Row3],vec[0])+vpnumber.X32Mul(mat[Col1Row3],vec[1])+vpnumber.X32Mul(mat[Col2Row3],vec[2])+mat[Col3Row3])
}

// MulVecDir performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *X32) MulVecDir(vec *vpvec3.X32) *vpvec3.X32 {
	var ret vpvec3.X32

	for i := range vec {
		ret[i] = vpnumber.X32Mul(mat.Get(0, i), vec[0]) + vpnumber.X32Mul(mat.Get(1, i), vec[1]) + vpnumber.X32Mul(mat.Get(2, i), vec[2])
	}

	return &ret
}

// X32Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func X32Add(mata, matb *X32) *X32 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// X32Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func X32Sub(mata, matb *X32) *X32 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// X32MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32MulScale(mat *X32, factor vpnumber.X32) *X32 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// X32DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func X32DivScale(mat *X32, factor vpnumber.X32) *X32 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// X32Transpose inverts rows and columns (matrix transposition).
// Args is left untouched, a pointer on a new object is returned.
func X32Transpose(mat *X32) *X32 {
	var ret X32

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, mat.Get(r, c))
		}
	}

	return &ret
}

// X32MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func X32MulComp(a, b *X32) *X32 {
	var ret X32

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, vpnumber.X32Mul(a.Get(0, r), b.Get(c, 0))+vpnumber.X32Mul(a.Get(1, r), b.Get(c, 1))+vpnumber.X32Mul(a.Get(2, r), b.Get(c, 2))+vpnumber.X32Mul(a.Get(3, r), b.Get(c, 3)))
		}
	}

	return &ret
}

// X32Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func X32Inv(mat *X32) *X32 {
	ret := X32{
		vpnumber.X32Muln(mat[Col1Row2], mat[Col2Row3], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col1Row3], mat[Col2Row2], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col1Row3], mat[Col2Row1], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col1Row1], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col1Row2], mat[Col2Row1], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col1Row1], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row3], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col2Row3], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col2Row1], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col2Row3], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col2Row1], mat[Col3Row3]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row3], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row2], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row1], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row3], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row1], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row2], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row1]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row1]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row2]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row2]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row3]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row3]),
		vpnumber.X32Muln(mat[Col1Row3], mat[Col2Row2], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col1Row2], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col1Row3], mat[Col2Row0], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col1Row0], mat[Col2Row3], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col1Row2], mat[Col2Row0], mat[Col3Row3]) - vpnumber.X32Muln(mat[Col1Row0], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row2], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col2Row2], mat[Col3Row0]) + vpnumber.X32Muln(mat[Col0Row3], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col2Row3], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col2Row2], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row2], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row0], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row3], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row0], mat[Col3Row3]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row2], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row3], mat[Col2Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row2], mat[Col2Row0]) + vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row2]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row2]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row3]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row3]),
		vpnumber.X32Muln(mat[Col1Row1], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col1Row3], mat[Col2Row1], mat[Col3Row0]) + vpnumber.X32Muln(mat[Col1Row3], mat[Col2Row0], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col1Row0], mat[Col2Row3], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col1Row1], mat[Col2Row0], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col1Row0], mat[Col2Row1], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row3], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col2Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col2Row3], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col2Row0], mat[Col3Row3]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col2Row1], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row3], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row1], mat[Col3Row0]) + vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row0], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row3], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row0], mat[Col3Row3]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row1], mat[Col3Row3]),
		vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row1], mat[Col2Row0]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row3], mat[Col2Row0]) - vpnumber.X32Muln(mat[Col0Row3], mat[Col1Row0], mat[Col2Row1]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row3], mat[Col2Row1]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row3]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row3]),
		vpnumber.X32Muln(mat[Col1Row2], mat[Col2Row1], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col1Row1], mat[Col2Row2], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col1Row2], mat[Col2Row0], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col1Row0], mat[Col2Row2], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col1Row1], mat[Col2Row0], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col1Row0], mat[Col2Row1], mat[Col3Row2]),
		vpnumber.X32Muln(mat[Col0Row1], mat[Col2Row2], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col2Row1], mat[Col3Row0]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col2Row0], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col2Row2], mat[Col3Row1]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col2Row0], mat[Col3Row2]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col2Row1], mat[Col3Row2]),
		vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row1], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row2], mat[Col3Row0]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row0], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row2], mat[Col3Row1]) + vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row0], mat[Col3Row2]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row1], mat[Col3Row2]),
		vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row2], mat[Col2Row0]) - vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row1], mat[Col2Row0]) + vpnumber.X32Muln(mat[Col0Row2], mat[Col1Row0], mat[Col2Row1]) - vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row2], mat[Col2Row1]) - vpnumber.X32Muln(mat[Col0Row1], mat[Col1Row0], mat[Col2Row2]) + vpnumber.X32Muln(mat[Col0Row0], mat[Col1Row1], mat[Col2Row2]),
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
