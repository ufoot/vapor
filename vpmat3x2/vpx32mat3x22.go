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
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
)

// X32Bas2 is a 2D space basis, composed of 3 points in a 2D space.
// It's defined along with 3x3 matrix code as manipulating such basis
// requires 3x3 code. X and Y are considered relative positions,
// with O as the origin.
type X32Bas2 struct {
	O vpvec2.X32Vec2
	X vpvec2.X32Vec2
	Y vpvec2.X32Vec2
}

// X32Bas2New creates a new 2D space basis.
func X32Bas2New(o, x, y *vpvec2.X32Vec2) *X32Bas2 {
	return &X32Bas2{*o, *x, *y}
}

// X32Bas2Default creates a new 2D space basis, using default
// orthogonal settings (origin at 0,0 with vectors 1,0 and 0,1).
func X32Bas2Default() *X32Bas2 {
	return &X32Bas2{*vpvec2.X32Vec2New(vpnumber.X32Const0, vpnumber.X32Const0), *vpvec2.X32Vec2New(vpnumber.X32Const1, vpnumber.X32Const0), *vpvec2.X32Vec2New(vpnumber.X32Const0, vpnumber.X32Const1)}
}

// String returns a readable form of the basis.
func (bas *X32Bas2) String() string {
	buf, err := json.Marshal(bas)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Normalize normalizes a 2D space basis, by normalizing
// all vectors in it.
// It modifies the basis, and returns a pointer on it.
func (bas *X32Bas2) Normalize() *X32Bas2 {
	*bas = *X32Bas2Normalize(bas)

	return bas
}

// Ortho makes a 2D space basis orthogonal,
// by using an orthogonal vector to X as the Y vector.
// It modifies the basis, and returns a pointer on it.
func (bas *X32Bas2) Ortho() *X32Bas2 {
	*bas = *X32Bas2Ortho(bas)

	return bas
}

// X32Bas2Normalize normalizes a 2D space basis, by normalizing
// all vectors in it.
// Args is left untouched, a pointer on a new object is returned.
func X32Bas2Normalize(bas *X32Bas2) *X32Bas2 {
	return &X32Bas2{bas.O, *vpvec2.X32Vec2Normalize(&bas.X), *vpvec2.X32Vec2Normalize(&bas.Y)}
}

// X32Bas2Ortho makes a 2D space basis orthogonal,
// by using an orthogonal vector to X as the Y vector.
// Args is left untouched, a pointer on a new object is returned.
func X32Bas2Ortho(bas *X32Bas2) *X32Bas2 {
	return &X32Bas2{bas.O, bas.X, *vpvec2.X32Vec2New(-bas.X[1], bas.X[0])}
}
