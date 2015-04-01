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

package vpmatrix4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmatrix3"
	"github.com/ufoot/vapor/vpnumber"
)

// F32Bas3 is a 3D space basis, composed of 4 points in a 4D space.
// It's defined along with 4x4 matrix code as manipulating such basis
// requires 4x4 code. X, Y and Z are considered relative positions,
// with O as the origin.
type F32Bas3 struct {
	O vpmatrix3.F32Vec3
	X vpmatrix3.F32Vec3
	Y vpmatrix3.F32Vec3
	Z vpmatrix3.F32Vec3
}

// F32Bas3New creates a new 3D space basis.
func F32Bas3New(o, x, y, z *vpmatrix3.F32Vec3) *F32Bas3 {
	return &F32Bas3{*o, *x, *y, *z}
}

// F32Bas3Default creates a new 3D space basis, using default
// orthogonal settings (origin at 0,0,0 with vectors 1,0,0 0,1,0 and 0,0,1).
func F32Bas3Default() *F32Bas3 {
	return &F32Bas3{*vpmatrix3.F32Vec3New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const0), *vpmatrix3.F32Vec3New(vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0), *vpmatrix3.F32Vec3New(vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0), *vpmatrix3.F32Vec3New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1)}
}

// String returns a readable form of the basis.
func (bas *F32Bas3) String() string {
	buf, err := json.Marshal(bas)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Normalize normalizes a 3D space basis, by normalizing
// all vectors in it.
// It modifies the basis, and returns a pointer on it.
func (bas *F32Bas3) Normalize() *F32Bas3 {
	*bas = *F32Bas3Normalize(bas)

	return bas
}

// Ortho makes a 3D space basis orthogonal,
// by using an Z * X as Y, and X * Y as Z.
// It modifies the basis, and returns a pointer on it.
func (bas *F32Bas3) Ortho() *F32Bas3 {
	*bas = *F32Bas3Ortho(bas)

	return bas
}

// F32Bas3Normalize normalizes a 3D space basis, by normalizing
// all vectors in it.
// Args is left untouched, a pointer on a new object is returned.
func F32Bas3Normalize(bas *F32Bas3) *F32Bas3 {
	return &F32Bas3{bas.O, *vpmatrix3.F32Vec3Normalize(&bas.X), *vpmatrix3.F32Vec3Normalize(&bas.Y), *vpmatrix3.F32Vec3Normalize(&bas.Z)}
}

// F32Bas3Ortho makes a 3D space basis orthogonal,
// by using an Z * X as Y, and X * Y as Z.
// Args is left untouched, a pointer on a new object is returned.
func F32Bas3Ortho(bas *F32Bas3) *F32Bas3 {
	y := vpmatrix3.F32Vec3Cross(&bas.Z, &bas.X)
	z := vpmatrix3.F32Vec3Cross(&bas.X, y)
	return &F32Bas3{bas.O, bas.X, *y, *z}
}
