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

package vpmat3x3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmat2x2"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpsys"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
	"math"
)

// F64 is a matrix containing 3x3 float64 values.
// Can hold the values of a point in space.
type F64 [Size]float64

// F64New creates a new matrix containing 3x3 float64 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F64New(f1, f2, f3, f4, f5, f6, f7, f8, f9 float64) *F64 {
	return &F64{f1, f2, f3, f4, f5, f6, f7, f8, f9}
}

// F64Identity creates a new identity matrix.
func F64Identity() *F64 {
	return &F64{vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1}
}

// F64Trans creates a new translation matrix.
func F64Trans(vec *vpvec2.F64) *F64 {
	return &F64{vpnumber.F64Const1, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1, vpnumber.F64Const0, vec[0], vec[1], vpnumber.F64Const1}
}

// F64Scale creates a new scale matrix.
func F64Scale(vec *vpvec2.F64) *F64 {
	return &F64{vec[0], vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vec[1], vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1}
}

// F64Rot creates a new rotation matrix.
// The rotation is done in 2D over a virtual z axis, such as z = cross(x,y).
// Angle is given in radians.
func F64Rot(r float64) *F64 {
	cos := math.Cos(r)
	sin := math.Sin(r)

	return &F64{cos, sin, vpnumber.F64Const0, -sin, cos, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const0, vpnumber.F64Const1}
}

// F64RebaseOXY creates a matrix that translates from the default
// O=(0,0), X=(1,0), Y=(0,1) basis to the given
// basis. It assumes f(a+b) equals f(a)+f(b).
func F64RebaseOXY(Origin, PosX, PosY *vpvec2.F64) *F64 {
	return &F64{PosX[0] - Origin[0], PosX[1] - Origin[1], vpnumber.F64Const0, PosY[0] - Origin[0], PosY[1] - Origin[1], vpnumber.F64Const0, Origin[0], Origin[1], vpnumber.F64Const1}
}

// F64RebaseOXYP creates a matrix that translates from the default
// O=(0,0), X=(1,0), Y=(0,1), P=(1,1) basis to the given
// basis. Note that there can be a projection, so  f(a+b) is not f(a)+f(b).
func F64RebaseOXYP(Origin, PosX, PosY, PosP *vpvec2.F64) *F64 {
	var tmpMat vpmat2x2.F64
	projMat := F64Identity()

	dX := vpvec2.F64Sub(PosX, Origin)
	dY := vpvec2.F64Sub(PosY, Origin)
	dP := vpvec2.F64Sub(PosP, Origin)
	tmpMat.SetCol(0, vpvec2.F64Sub(dX, dP))
	tmpMat.SetCol(1, vpvec2.F64Sub(dY, dP))
	tmpMat.Inv()
	tmpVec := vpvec2.F64Sub(dP, vpvec2.F64Add(dX, dY))
	lastRow := tmpMat.MulVec(tmpVec)
	colX := vpvec2.F64MulScale(dX, vpnumber.F64Const1+lastRow[0])
	colY := vpvec2.F64MulScale(dY, vpnumber.F64Const1+lastRow[1])
	projMat.SetCol(0, vpvec3.F64FromVec2(colX, lastRow[0]))
	projMat.SetCol(1, vpvec3.F64FromVec2(colY, lastRow[1]))
	transMat := F64Trans(Origin)

	ret := F64MulComp(transMat, projMat)

	return ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F64) ToX32() *X32 {
	var ret X32

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F64) ToX64() *X64 {
	var ret X64

	for i, v := range mat {
		ret[i] = vpnumber.F64ToX64(v)
	}

	return &ret
}

// ToF32 converts the matrix to a float32 matrix.
func (mat *F64) ToF32() *F32 {
	var ret F32

	for i, v := range mat {
		ret[i] = float32(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F64) Set(col, row int, val float64) {
	mat[col*Height+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F64) Get(col, row int) float64 {
	return mat[col*Height+row]
}

// SetCol sets a column to the values contained in a vector.
func (mat *F64) SetCol(col int, vec *vpvec3.F64) {
	for row, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetCol gets a column and returns it in a vector.
func (mat *F64) GetCol(col int) *vpvec3.F64 {
	var ret vpvec3.F64

	for row := range ret {
		ret[row] = mat[col*Height+row]
	}

	return &ret
}

// SetRow sets a row to the values contained in a vector.
func (mat *F64) SetRow(row int, vec *vpvec3.F64) {
	for col, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetRow gets a row and returns it in a vector.
func (mat *F64) GetRow(row int) *vpvec3.F64 {
	var ret vpvec3.F64

	for col := range ret {
		ret[col] = mat[col*Height+row]
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F64) MarshalJSON() ([]byte, error) {
	var tmpArray [Width][Height]float64

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*Height+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal F64")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F64) UnmarshalJSON(data []byte) error {
	var tmpArray [Width][Height]float64

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal F64")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*Height+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F64) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) Add(op *F64) *F64 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) Sub(op *F64) *F64 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) MulScale(factor float64) *F64 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) DivScale(factor float64) *F64 {
	for i, v := range mat {
		mat[i] = vpnumber.F64Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F64) IsSimilar(op *F64) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F64IsSimilar(v, op[i])
	}

	return ret
}

// Transpose inverts rows and columns (matrix transposition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) Transpose(op *F64) *F64 {
	*mat = *F64Transpose(op)

	return mat
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) MulComp(op *F64) *F64 {
	*mat = *F64MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F64) Det() float64 {
	return mat[Col0Row0]*mat[Col1Row1]*mat[Col2Row2] + mat[Col0Row1]*mat[Col1Row2]*mat[Col2Row0] + mat[Col0Row2]*mat[Col1Row0]*mat[Col2Row1] - mat[Col0Row0]*mat[Col1Row2]*mat[Col2Row1] - mat[Col0Row1]*mat[Col1Row0]*mat[Col2Row2] - mat[Col0Row2]*mat[Col1Row1]*mat[Col2Row0]
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F64) Inv() *F64 {
	*mat = *F64Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *F64) MulVec(vec *vpvec3.F64) *vpvec3.F64 {
	var ret vpvec3.F64

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
func (mat *F64) MulVecPos(vec *vpvec2.F64) *vpvec2.F64 {
	var ret vpvec2.F64

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)
	}

	return ret.DivScale(mat[Col0Row2]*vec[0] + mat[Col1Row2]*vec[1] + mat[Col2Row2])
}

// MulVecDir performs a multiplication of a vector by a 3x3 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 2 (a point in a plane) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F64) MulVecDir(vec *vpvec2.F64) *vpvec2.F64 {
	var ret vpvec2.F64

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1]
	}

	return &ret
}

// F64Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F64Add(mata, matb *F64) *F64 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F64Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F64Sub(mata, matb *F64) *F64 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F64MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64MulScale(mat *F64, factor float64) *F64 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F64DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F64DivScale(mat *F64, factor float64) *F64 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F64Transpose inverts rows and columns (matrix transposition).
// Args is left untouched, a pointer on a new object is returned.
func F64Transpose(mat *F64) *F64 {
	var ret F64

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, mat.Get(r, c))
		}
	}

	return &ret
}

// F64MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F64MulComp(a, b *F64) *F64 {
	var ret F64

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1)+a.Get(2, r)*b.Get(c, 2))
		}
	}

	return &ret
}

// F64Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F64Inv(mat *F64) *F64 {
	ret := F64{
		mat[Col1Row1]*mat[Col2Row2] - mat[Col1Row2]*mat[Col2Row1],
		mat[Col0Row2]*mat[Col2Row1] - mat[Col0Row1]*mat[Col2Row2],
		mat[Col0Row1]*mat[Col1Row2] - mat[Col0Row2]*mat[Col1Row1],
		mat[Col1Row2]*mat[Col2Row0] - mat[Col1Row0]*mat[Col2Row2],
		mat[Col0Row0]*mat[Col2Row2] - mat[Col0Row2]*mat[Col2Row0],
		mat[Col0Row2]*mat[Col1Row0] - mat[Col0Row0]*mat[Col1Row2],
		mat[Col1Row0]*mat[Col2Row1] - mat[Col1Row1]*mat[Col2Row0],
		mat[Col0Row1]*mat[Col2Row0] - mat[Col0Row0]*mat[Col2Row1],
		mat[Col0Row0]*mat[Col1Row1] - mat[Col0Row1]*mat[Col1Row0],
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
