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
	"github.com/ufoot/vapor/vpmatrix3"
	"github.com/ufoot/vapor/vpnumber"
)

// X32Bas2 is a 2D space basis, composed of 3 points in a 2D space.
// It's defined along with 3x3 matrix code as manipulating such basis
// requires 3x3 code. X and Y are considered relative positions,
// with O as the origin.
type X32Bas2 struct {
	O vpmatrix3.X32Vec3
	X vpmatrix3.X32Vec3
	Y vpmatrix3.X32Vec3
	Z vpmatrix3.X32Vec3
}

// X32Bas2New creates a new 2D space basis.
func X32Bas2New(o, x, y, z *vpmatrix3.X32Vec3) *X32Bas2 {
	return &X32Bas2{*o, *x, *y, *z}
}

// X32Bas2Default creates a new 2D space basis, using default
// orthogonal settings (origin at 0,0 with vectors 1,0 and 0,1).
func X32Bas2Default() *X32Bas2 {
	return &X32Bas2{*vpmatrix3.X32Vec3New(vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const0), *vpmatrix3.X32Vec3New(vpnumber.X32Const1, vpnumber.X32Const0, vpnumber.X32Const0), *vpmatrix3.X32Vec3New(vpnumber.X32Const0, vpnumber.X32Const1, vpnumber.X32Const0), *vpmatrix3.X32Vec3New(vpnumber.X32Const0, vpnumber.X32Const0, vpnumber.X32Const1)}
}

// Normalize normalizes a 2D space basis, by normalizing
// all vectors in it.
// It modifies the basis, and returns a pointer on it.
func (bas *X32Bas2) Normalize() *X32Bas2 {
	*bas = *X32Bas2Normalize(bas)

	return bas
}

// Ortho makes a 3D space basis orthogonal,
// by using an Z * X as Y, and X * Y as Z.
// It modifies the basis, and returns a pointer on it.
func (bas *X32Bas2) Ortho() *X32Bas2 {
	*bas = *X32Bas2Ortho(bas)

	return bas
}

// X32Bas2Normalize normalizes a 3D space basis, by normalizing
// all vectors in it.
// Args is left untouched, a pointer on a new object is returned.
func X32Bas2Normalize(bas *X32Bas2) *X32Bas2 {
	return &X32Bas2{bas.O, *vpmatrix3.X32Vec3Normalize(&bas.X), *vpmatrix3.X32Vec3Normalize(&bas.Y), *vpmatrix3.X32Vec3Normalize(&bas.Z)}
}

// X32Bas2Ortho makes a 3D space basis orthogonal,
// by using an Z * X as Y, and X * Y as Z.
// Args is left untouched, a pointer on a new object is returned.
func X32Bas2Ortho(bas *X32Bas2) *X32Bas2 {
	y := vpmatrix3.X32Vec3Cross(&bas.Z, &bas.X)
	z := vpmatrix3.X32Vec3Cross(&bas.X, y)
	return &X32Bas2{bas.O, bas.X, *y, *z}
}
