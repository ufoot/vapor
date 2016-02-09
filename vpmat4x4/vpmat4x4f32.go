// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
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
	"github.com/ufoot/vapor/vperror"
	"github.com/ufoot/vapor/vpmat3x3"
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
	"github.com/ufoot/vapor/vpvec4"
	"math"
)

// F32 is a matrix containing 4x4 float32 values.
// Can be used in 3D matrix transformations.
type F32 [Size]float32

// F32New creates a new matrix containing 4x4 float32 values.
// The column-major (OpenGL notation) mode is used,
// first elements fill first column.
func F32New(f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16 float32) *F32 {
	return &F32{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16}
}

// F32Identity creates a new identity matrix.
func F32Identity() *F32 {
	return &F32{vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1}
}

// F32Translation creates a new translation matrix.
func F32Translation(vec *vpvec3.F32) *F32 {
	return &F32{vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vec[0], vec[1], vec[2], vpnumber.F32Const1}
}

// F32Scale creates a new scale matrix.
func F32Scale(vec *vpvec3.F32) *F32 {
	return &F32{vec[0], vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vec[1], vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vec[2], vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1}
}

// F32RotX creates a new rotation matrix.
// The rotation is done in 3D over the x (1st) axis.
// Angle is given in radians.
func F32RotX(r float32) *F32 {
	cos := float32(math.Cos(float64(r)))
	sin := float32(math.Sin(float64(r)))

	return &F32{vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, cos, sin, vpnumber.F32Const0, vpnumber.F32Const0, -sin, cos, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1}
}

// F32RotY creates a new rotation matrix.
// The rotation is done in 3D over the y (2nd) axis.
// Angle is given in radians.
func F32RotY(r float32) *F32 {
	cos := float32(math.Cos(float64(r)))
	sin := float32(math.Sin(float64(r)))

	return &F32{cos, vpnumber.F32Const0, -sin, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, sin, vpnumber.F32Const0, cos, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1}
}

// F32RotZ creates a new rotation matrix.
// The rotation is done in 3D over the z (3rd) axis.
// Angle is given in radians.
func F32RotZ(r float32) *F32 {
	cos := float32(math.Cos(float64(r)))
	sin := float32(math.Sin(float64(r)))

	return &F32{cos, sin, vpnumber.F32Const0, vpnumber.F32Const0, -sin, cos, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1}
}

// F32RebaseOXYZ creates a matrix that translates from the default
// O=(0,0,0), X=(1,0,0), Y=(0,1,0), Z=(0,0,1) basis to the given
// basis. It assumes f(a+b) equals f(a)+f(b).
func F32RebaseOXYZ(Origin, PosX, PosY, PosZ *vpvec3.F32) *F32 {
	return &F32{PosX[0] - Origin[0], PosX[1] - Origin[1], PosX[2] - Origin[2], vpnumber.F32Const0, PosY[0] - Origin[0], PosY[1] - Origin[1], PosY[2] - Origin[2], vpnumber.F32Const0, PosZ[0] - Origin[0], PosZ[1] - Origin[1], PosZ[2] - Origin[2], vpnumber.F32Const0, Origin[0], Origin[1], Origin[2], vpnumber.F32Const1}
}

// F32RebaseOXYZP creates a matrix that translates from the default
// O=(0,0,0), X=(1,0,0), Y=(0,1,0), Z=(0,0,1), P=(1,1,1) basis to the given
// basis. Note that there can be a projection, so  f(a+b) is not f(a)+f(b).
func F32RebaseOXYZP(Origin, PosX, PosY, PosZ, PosP *vpvec3.F32) *F32 {
	var tmpMat vpmat3x3.F32
	projMat := F32Identity()

	dX := vpvec3.F32Sub(PosX, Origin)
	dY := vpvec3.F32Sub(PosY, Origin)
	dZ := vpvec3.F32Sub(PosZ, Origin)
	dP := vpvec3.F32Sub(PosP, Origin)
	tmpMat.SetCol(0, vpvec3.F32Sub(dX, dP))
	tmpMat.SetCol(1, vpvec3.F32Sub(dY, dP))
	tmpMat.SetCol(2, vpvec3.F32Sub(dZ, dP))
	tmpMat.Inv()
	tmpVec := vpvec3.F32Sub(dP, vpvec3.F32Add(dX, vpvec3.F32Add(dY, dZ)))
	lastRow := tmpMat.MulVec(tmpVec)
	colX := vpvec3.F32MulScale(dX, vpnumber.F32Const1+lastRow[0])
	colY := vpvec3.F32MulScale(dY, vpnumber.F32Const1+lastRow[1])
	colZ := vpvec3.F32MulScale(dZ, vpnumber.F32Const1+lastRow[2])
	projMat.SetCol(0, vpvec4.F32FromVec3(colX, lastRow[0]))
	projMat.SetCol(1, vpvec4.F32FromVec3(colY, lastRow[1]))
	projMat.SetCol(2, vpvec4.F32FromVec3(colZ, lastRow[2]))
	transMat := F32Translation(Origin)

	ret := F32MulComp(transMat, projMat)

	return ret
}

// F32Ortho creates a projection matrix the way the standard OpenGL glOrtho
// would (see https://www.opengl.org/sdk/docs/man2/xhtml/glOrtho.xml).
// Note: use -nearVal and -farVal to initialize.
// It's a little akward, if you expect to pass vectors with positions
// ranging from nearVal to farVal then you need to pass -nearVal and
// -farVal to this function. This is probably due to the fact that
// with a right-handed basis and X,Y set up "as usual", then Z is negative
// when going farther and farther. This tweak allows farVal to yield
// +1 and nearVal -1. We keep this function as is here, as this is the
// way OpenGL functions seem to work.
func F32Ortho(left, right, bottom, top, nearVal, farVal float32) *F32 {
	var ret F32

	ret[Col0Row0] = vpnumber.F32Div(2.0, right-left)
	ret[Col1Row1] = vpnumber.F32Div(2.0, top-bottom)
	ret[Col2Row2] = vpnumber.F32Div(-2.0, farVal-nearVal)
	ret[Col3Row0] = -vpnumber.F32Div(right+left, right-left)
	ret[Col3Row1] = -vpnumber.F32Div(top+bottom, top-bottom)
	ret[Col3Row2] = -vpnumber.F32Div(farVal+nearVal, farVal-nearVal)
	ret[Col3Row3] = vpnumber.F32Const1

	return &ret
}

// F32Perspective creates a projection matrix the way the standard GLU
// gluPerspective function would (see
// https://www.opengl.org/sdk/docs/man2/xhtml/gluPerspective.xml).
// Beware, fovy is in degrees, not radians.
func F32Perspective(fovy, aspect, zNear, zFar float32) *F32 {
	var ret F32

	radFovy2 := float64(vpmath.F32DegToRad(vpmath.F32DegMod(fovy) / 2.0))
	f := vpnumber.F32Div(float32(math.Cos(radFovy2)), float32(math.Sin(radFovy2)))

	ret[Col0Row0] = vpnumber.F32Div(f, aspect)
	ret[Col1Row1] = f
	ret[Col2Row2] = vpnumber.F32Div(zFar+zNear, zNear-zFar)
	ret[Col2Row3] = -vpnumber.F32Const1
	ret[Col3Row2] = vpnumber.F32Div(2.0*zFar*zNear, zNear-zFar)

	return &ret
}

// ToX32 converts the matrix to a fixed point number matrix on 32 bits.
func (mat *F32) ToX32() *X32 {
	var ret X32

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX32(v)
	}

	return &ret
}

// ToX64 converts the matrix to a fixed point number matrix on 64 bits.
func (mat *F32) ToX64() *X64 {
	var ret X64

	for i, v := range mat {
		ret[i] = vpnumber.F32ToX64(v)
	}

	return &ret
}

// ToF64 converts the matrix to a float64 matrix.
func (mat *F32) ToF64() *F64 {
	var ret F64

	for i, v := range mat {
		ret[i] = float64(v)
	}

	return &ret
}

// Set sets the value of the matrix for a given column and row.
func (mat *F32) Set(col, row int, val float32) {
	mat[col*Height+row] = val
}

// Get gets the value of the matrix for a given column and row.
func (mat *F32) Get(col, row int) float32 {
	return mat[col*Height+row]
}

// SetCol sets a column to the values contained in a vector.
func (mat *F32) SetCol(col int, vec *vpvec4.F32) {
	for row, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetCol gets a column and returns it in a vector.
func (mat *F32) GetCol(col int) *vpvec4.F32 {
	var ret vpvec4.F32

	for row := range ret {
		ret[row] = mat[col*Height+row]
	}

	return &ret
}

// SetRow sets a row to the values contained in a vector.
func (mat *F32) SetRow(row int, vec *vpvec4.F32) {
	for col, val := range vec {
		mat[col*Height+row] = val
	}
}

// GetRow gets a row and returns it in a vector.
func (mat *F32) GetRow(row int) *vpvec4.F32 {
	var ret vpvec4.F32

	for col := range ret {
		ret[col] = mat[col*Height+row]
	}

	return &ret
}

// MarshalJSON implements the json.Marshaler interface.
func (mat *F32) MarshalJSON() ([]byte, error) {
	var tmpArray [Width][Height]float32

	for col := range tmpArray {
		for row := range tmpArray[col] {
			tmpArray[col][row] = mat[col*Height+row]
		}
	}

	ret, err := json.Marshal(tmpArray)
	if err != nil {
		return nil, vperror.Chain(err, "unable to marshal F32")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mat *F32) UnmarshalJSON(data []byte) error {
	var tmpArray [Width][Height]float32

	err := json.Unmarshal(data, &tmpArray)
	if err != nil {
		return vperror.Chain(err, "unable to unmarshal F32")
	}

	for col := range tmpArray {
		for row := range tmpArray[col] {
			mat[col*Height+row] = tmpArray[col][row]
		}
	}

	return nil
}

// String returns a readable form of the matrix.
func (mat *F32) String() string {
	buf, err := mat.MarshalJSON()

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Add adds operand to the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Add(op *F32) *F32 {
	for i, v := range op {
		mat[i] += v
	}

	return mat
}

// Sub substracts operand from the matrix.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Sub(op *F32) *F32 {
	for i, v := range op {
		mat[i] -= v
	}

	return mat
}

// MulScale multiplies all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) MulScale(factor float32) *F32 {
	for i, v := range mat {
		mat[i] = v * factor
	}

	return mat
}

// DivScale divides all values of the matrix by factor.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) DivScale(factor float32) *F32 {
	for i, v := range mat {
		mat[i] = vpnumber.F32Div(v, factor)
	}

	return mat
}

// IsSimilar returns true if matrices are approximatively the same.
// This is a workarround to ignore rounding errors.
func (mat *F32) IsSimilar(op *F32) bool {
	ret := true
	for i, v := range mat {
		ret = ret && vpnumber.F32IsSimilar(v, op[i])
	}

	return ret
}

// Transpose inverts rows and columns (matrix transposition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Transpose(op *F32) *F32 {
	*mat = *F32Transpose(op)

	return mat
}

// MulComp multiplies the matrix by another matrix (composition).
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) MulComp(op *F32) *F32 {
	*mat = *F32MulComp(mat, op)

	return mat
}

// Det returns the matrix determinant.
func (mat *F32) Det() float32 {
	return mat[Col0Row3]*mat[Col1Row2]*mat[Col2Row1]*mat[Col3Row0] - mat[Col0Row2]*mat[Col1Row3]*mat[Col2Row1]*mat[Col3Row0] - mat[Col0Row3]*mat[Col1Row1]*mat[Col2Row2]*mat[Col3Row0] + mat[Col0Row1]*mat[Col1Row3]*mat[Col2Row2]*mat[Col3Row0] + mat[Col0Row2]*mat[Col1Row1]*mat[Col2Row3]*mat[Col3Row0] - mat[Col0Row1]*mat[Col1Row2]*mat[Col2Row3]*mat[Col3Row0] - mat[Col0Row3]*mat[Col1Row2]*mat[Col2Row0]*mat[Col3Row1] + mat[Col0Row2]*mat[Col1Row3]*mat[Col2Row0]*mat[Col3Row1] + mat[Col0Row3]*mat[Col1Row0]*mat[Col2Row2]*mat[Col3Row1] - mat[Col0Row0]*mat[Col1Row3]*mat[Col2Row2]*mat[Col3Row1] - mat[Col0Row2]*mat[Col1Row0]*mat[Col2Row3]*mat[Col3Row1] + mat[Col0Row0]*mat[Col1Row2]*mat[Col2Row3]*mat[Col3Row1] + mat[Col0Row3]*mat[Col1Row1]*mat[Col2Row0]*mat[Col3Row2] - mat[Col0Row1]*mat[Col1Row3]*mat[Col2Row0]*mat[Col3Row2] - mat[Col0Row3]*mat[Col1Row0]*mat[Col2Row1]*mat[Col3Row2] + mat[Col0Row0]*mat[Col1Row3]*mat[Col2Row1]*mat[Col3Row2] + mat[Col0Row1]*mat[Col1Row0]*mat[Col2Row3]*mat[Col3Row2] - mat[Col0Row0]*mat[Col1Row1]*mat[Col2Row3]*mat[Col3Row2] - mat[Col0Row2]*mat[Col1Row1]*mat[Col2Row0]*mat[Col3Row3] + mat[Col0Row1]*mat[Col1Row2]*mat[Col2Row0]*mat[Col3Row3] + mat[Col0Row2]*mat[Col1Row0]*mat[Col2Row1]*mat[Col3Row3] - mat[Col0Row0]*mat[Col1Row2]*mat[Col2Row1]*mat[Col3Row3] - mat[Col0Row1]*mat[Col1Row0]*mat[Col2Row2]*mat[Col3Row3] + mat[Col0Row0]*mat[Col1Row1]*mat[Col2Row2]*mat[Col3Row3]
}

// Inv inverts the matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// It modifies the matrix, and returns a pointer on it.
func (mat *F32) Inv() *F32 {
	*mat = *F32Inv(mat)

	return mat
}

// MulVec performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
func (mat *F32) MulVec(vec *vpvec4.F32) *vpvec4.F32 {
	var ret vpvec4.F32

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2] + mat.Get(3, i)*vec[3]
	}

	return &ret
}

// MulVecPos performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 1, so in practice a
// position vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations and translations to be accumulated
// within the matrix and then performed at once.
func (mat *F32) MulVecPos(vec *vpvec3.F32) *vpvec3.F32 {
	var ret vpvec3.F32

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2] + mat.Get(3, i)
	}

	return ret.DivScale(mat[Col0Row3]*vec[0] + mat[Col1Row3]*vec[1] + mat[Col2Row3]*vec[2] + mat[Col3Row3])
}

// MulVecDir performs a multiplication of a vector by a 4x4 matrix,
// considering the vector is a column vector (matrix left, vector right).
// The last member of the vector is assumed to be 0, so in practice a
// direction vector of length 3 (a point in space) is passed. This allow geometric
// transformations such as rotations to be accumulated
// within the matrix and then performed at once.
func (mat *F32) MulVecDir(vec *vpvec3.F32) *vpvec3.F32 {
	var ret vpvec3.F32

	for i := range vec {
		ret[i] = mat.Get(0, i)*vec[0] + mat.Get(1, i)*vec[1] + mat.Get(2, i)*vec[2]
	}

	return &ret
}

// F32Add adds two matrices.
// Args are left untouched, a pointer on a new object is returned.
func F32Add(mata, matb *F32) *F32 {
	var ret = *mata

	_ = ret.Add(matb)

	return &ret
}

// F32Sub substracts matrix b from matrix a.
// Args are left untouched, a pointer on a new object is returned.
func F32Sub(mata, matb *F32) *F32 {
	var ret = *mata

	_ = ret.Sub(matb)

	return &ret
}

// F32MulScale multiplies all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32MulScale(mat *F32, factor float32) *F32 {
	var ret = *mat

	_ = ret.MulScale(factor)

	return &ret
}

// F32DivScale divides all values of a matrix by a scalar.
// Args are left untouched, a pointer on a new object is returned.
func F32DivScale(mat *F32, factor float32) *F32 {
	var ret = *mat

	_ = ret.DivScale(factor)

	return &ret
}

// F32Transpose inverts rows and columns (matrix transposition).
// Args is left untouched, a pointer on a new object is returned.
func F32Transpose(mat *F32) *F32 {
	var ret F32

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, mat.Get(r, c))
		}
	}

	return &ret
}

// F32MulComp multiplies two matrices (composition).
// Args are left untouched, a pointer on a new object is returned.
func F32MulComp(a, b *F32) *F32 {
	var ret F32

	for c := 0; c < Width; c++ {
		for r := 0; r < Height; r++ {
			ret.Set(c, r, a.Get(0, r)*b.Get(c, 0)+a.Get(1, r)*b.Get(c, 1)+a.Get(2, r)*b.Get(c, 2)+a.Get(3, r)*b.Get(c, 3))
		}
	}

	return &ret
}

// F32Inv inverts a matrix.
// Never fails (no division by zero error, never) but if the
// matrix can't be inverted, result does not make sense.
// Args is left untouched, a pointer on a new object is returned.
func F32Inv(mat *F32) *F32 {
	ret := F32{
		mat[Col1Row2]*mat[Col2Row3]*mat[Col3Row1] - mat[Col1Row3]*mat[Col2Row2]*mat[Col3Row1] + mat[Col1Row3]*mat[Col2Row1]*mat[Col3Row2] - mat[Col1Row1]*mat[Col2Row3]*mat[Col3Row2] - mat[Col1Row2]*mat[Col2Row1]*mat[Col3Row3] + mat[Col1Row1]*mat[Col2Row2]*mat[Col3Row3],
		mat[Col0Row3]*mat[Col2Row2]*mat[Col3Row1] - mat[Col0Row2]*mat[Col2Row3]*mat[Col3Row1] - mat[Col0Row3]*mat[Col2Row1]*mat[Col3Row2] + mat[Col0Row1]*mat[Col2Row3]*mat[Col3Row2] + mat[Col0Row2]*mat[Col2Row1]*mat[Col3Row3] - mat[Col0Row1]*mat[Col2Row2]*mat[Col3Row3],
		mat[Col0Row2]*mat[Col1Row3]*mat[Col3Row1] - mat[Col0Row3]*mat[Col1Row2]*mat[Col3Row1] + mat[Col0Row3]*mat[Col1Row1]*mat[Col3Row2] - mat[Col0Row1]*mat[Col1Row3]*mat[Col3Row2] - mat[Col0Row2]*mat[Col1Row1]*mat[Col3Row3] + mat[Col0Row1]*mat[Col1Row2]*mat[Col3Row3],
		mat[Col0Row3]*mat[Col1Row2]*mat[Col2Row1] - mat[Col0Row2]*mat[Col1Row3]*mat[Col2Row1] - mat[Col0Row3]*mat[Col1Row1]*mat[Col2Row2] + mat[Col0Row1]*mat[Col1Row3]*mat[Col2Row2] + mat[Col0Row2]*mat[Col1Row1]*mat[Col2Row3] - mat[Col0Row1]*mat[Col1Row2]*mat[Col2Row3],
		mat[Col1Row3]*mat[Col2Row2]*mat[Col3Row0] - mat[Col1Row2]*mat[Col2Row3]*mat[Col3Row0] - mat[Col1Row3]*mat[Col2Row0]*mat[Col3Row2] + mat[Col1Row0]*mat[Col2Row3]*mat[Col3Row2] + mat[Col1Row2]*mat[Col2Row0]*mat[Col3Row3] - mat[Col1Row0]*mat[Col2Row2]*mat[Col3Row3],
		mat[Col0Row2]*mat[Col2Row3]*mat[Col3Row0] - mat[Col0Row3]*mat[Col2Row2]*mat[Col3Row0] + mat[Col0Row3]*mat[Col2Row0]*mat[Col3Row2] - mat[Col0Row0]*mat[Col2Row3]*mat[Col3Row2] - mat[Col0Row2]*mat[Col2Row0]*mat[Col3Row3] + mat[Col0Row0]*mat[Col2Row2]*mat[Col3Row3],
		mat[Col0Row3]*mat[Col1Row2]*mat[Col3Row0] - mat[Col0Row2]*mat[Col1Row3]*mat[Col3Row0] - mat[Col0Row3]*mat[Col1Row0]*mat[Col3Row2] + mat[Col0Row0]*mat[Col1Row3]*mat[Col3Row2] + mat[Col0Row2]*mat[Col1Row0]*mat[Col3Row3] - mat[Col0Row0]*mat[Col1Row2]*mat[Col3Row3],
		mat[Col0Row2]*mat[Col1Row3]*mat[Col2Row0] - mat[Col0Row3]*mat[Col1Row2]*mat[Col2Row0] + mat[Col0Row3]*mat[Col1Row0]*mat[Col2Row2] - mat[Col0Row0]*mat[Col1Row3]*mat[Col2Row2] - mat[Col0Row2]*mat[Col1Row0]*mat[Col2Row3] + mat[Col0Row0]*mat[Col1Row2]*mat[Col2Row3],
		mat[Col1Row1]*mat[Col2Row3]*mat[Col3Row0] - mat[Col1Row3]*mat[Col2Row1]*mat[Col3Row0] + mat[Col1Row3]*mat[Col2Row0]*mat[Col3Row1] - mat[Col1Row0]*mat[Col2Row3]*mat[Col3Row1] - mat[Col1Row1]*mat[Col2Row0]*mat[Col3Row3] + mat[Col1Row0]*mat[Col2Row1]*mat[Col3Row3],
		mat[Col0Row3]*mat[Col2Row1]*mat[Col3Row0] - mat[Col0Row1]*mat[Col2Row3]*mat[Col3Row0] - mat[Col0Row3]*mat[Col2Row0]*mat[Col3Row1] + mat[Col0Row0]*mat[Col2Row3]*mat[Col3Row1] + mat[Col0Row1]*mat[Col2Row0]*mat[Col3Row3] - mat[Col0Row0]*mat[Col2Row1]*mat[Col3Row3],
		mat[Col0Row1]*mat[Col1Row3]*mat[Col3Row0] - mat[Col0Row3]*mat[Col1Row1]*mat[Col3Row0] + mat[Col0Row3]*mat[Col1Row0]*mat[Col3Row1] - mat[Col0Row0]*mat[Col1Row3]*mat[Col3Row1] - mat[Col0Row1]*mat[Col1Row0]*mat[Col3Row3] + mat[Col0Row0]*mat[Col1Row1]*mat[Col3Row3],
		mat[Col0Row3]*mat[Col1Row1]*mat[Col2Row0] - mat[Col0Row1]*mat[Col1Row3]*mat[Col2Row0] - mat[Col0Row3]*mat[Col1Row0]*mat[Col2Row1] + mat[Col0Row0]*mat[Col1Row3]*mat[Col2Row1] + mat[Col0Row1]*mat[Col1Row0]*mat[Col2Row3] - mat[Col0Row0]*mat[Col1Row1]*mat[Col2Row3],
		mat[Col1Row2]*mat[Col2Row1]*mat[Col3Row0] - mat[Col1Row1]*mat[Col2Row2]*mat[Col3Row0] - mat[Col1Row2]*mat[Col2Row0]*mat[Col3Row1] + mat[Col1Row0]*mat[Col2Row2]*mat[Col3Row1] + mat[Col1Row1]*mat[Col2Row0]*mat[Col3Row2] - mat[Col1Row0]*mat[Col2Row1]*mat[Col3Row2],
		mat[Col0Row1]*mat[Col2Row2]*mat[Col3Row0] - mat[Col0Row2]*mat[Col2Row1]*mat[Col3Row0] + mat[Col0Row2]*mat[Col2Row0]*mat[Col3Row1] - mat[Col0Row0]*mat[Col2Row2]*mat[Col3Row1] - mat[Col0Row1]*mat[Col2Row0]*mat[Col3Row2] + mat[Col0Row0]*mat[Col2Row1]*mat[Col3Row2],
		mat[Col0Row2]*mat[Col1Row1]*mat[Col3Row0] - mat[Col0Row1]*mat[Col1Row2]*mat[Col3Row0] - mat[Col0Row2]*mat[Col1Row0]*mat[Col3Row1] + mat[Col0Row0]*mat[Col1Row2]*mat[Col3Row1] + mat[Col0Row1]*mat[Col1Row0]*mat[Col3Row2] - mat[Col0Row0]*mat[Col1Row1]*mat[Col3Row2],
		mat[Col0Row1]*mat[Col1Row2]*mat[Col2Row0] - mat[Col0Row2]*mat[Col1Row1]*mat[Col2Row0] + mat[Col0Row2]*mat[Col1Row0]*mat[Col2Row1] - mat[Col0Row0]*mat[Col1Row2]*mat[Col2Row1] - mat[Col0Row1]*mat[Col1Row0]*mat[Col2Row2] + mat[Col0Row0]*mat[Col1Row1]*mat[Col2Row2],
	}

	det := mat.Det()
	ret.DivScale(det)

	return &ret
}
