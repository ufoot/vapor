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

// X64Bas3 is a 3D space basis, composed of 4 points in a 3D space.
// It's defined along with 3x3 matrix code as manipulating such basis
// requires 3x3 code. X and Y are considered relative positions,
// with O as the origin.
type X64Bas3 struct {
	O vpmatrix3.X64Vec3
	X vpmatrix3.X64Vec3
	Y vpmatrix3.X64Vec3
	Z vpmatrix3.X64Vec3
}

// X64Bas3New creates a new 3D space basis.
func X64Bas3New(o, x, y, z *vpmatrix3.X64Vec3) *X64Bas3 {
	return &X64Bas3{*o, *x, *y, *z}
}

// X64Bas3Default creates a new 3D space basis, using default
// orthogonal settings (origin at 0,0,0 with vectors 1,0,0 0,1,0 and 0,0,1).
func X64Bas3Default() *X64Bas3 {
	return &X64Bas3{*vpmatrix3.X64Vec3New(vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const0), *vpmatrix3.X64Vec3New(vpnumber.X64Const1, vpnumber.X64Const0, vpnumber.X64Const0), *vpmatrix3.X64Vec3New(vpnumber.X64Const0, vpnumber.X64Const1, vpnumber.X64Const0), *vpmatrix3.X64Vec3New(vpnumber.X64Const0, vpnumber.X64Const0, vpnumber.X64Const1)}
}

// Normalize normalizes a 3D space basis, by normalizing
// all vectors in it.
// It modifies the basis, and returns a pointer on it.
func (bas *X64Bas3) Normalize() *X64Bas3 {
	*bas = *X64Bas3Normalize(bas)

	return bas
}

// Ortho makes a 3D space basis orthogonal,
// by using an Z * X as Y, and X * Y as Z.
// It modifies the basis, and returns a pointer on it.
func (bas *X64Bas3) Ortho() *X64Bas3 {
	*bas = *X64Bas3Ortho(bas)

	return bas
}

// X64Bas3Normalize normalizes a 3D space basis, by normalizing
// all vectors in it.
// Args is left untouched, a pointer on a new object is returned.
func X64Bas3Normalize(bas *X64Bas3) *X64Bas3 {
	return &X64Bas3{bas.O, *vpmatrix3.X64Vec3Normalize(&bas.X), *vpmatrix3.X64Vec3Normalize(&bas.Y), *vpmatrix3.X64Vec3Normalize(&bas.Z)}
}

// X64Bas3Ortho makes a 3D space basis orthogonal,
// by using an Z * X as Y, and X * Y as Z.
// Args is left untouched, a pointer on a new object is returned.
func X64Bas3Ortho(bas *X64Bas3) *X64Bas3 {
	y := vpmatrix3.X64Vec3Cross(&bas.Z, &bas.X)
	z := vpmatrix3.X64Vec3Cross(&bas.X, y)
	return &X64Bas3{bas.O, bas.X, *y, *z}
}
