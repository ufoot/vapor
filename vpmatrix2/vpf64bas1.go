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
	"github.com/ufoot/vapor/vpnumber"
)

// F64Bas1 is a 1D space basis, composed of 2 points in a 1D space.
// It's defined along with 2x2 matrix code as manipulating such basis
// requires 2x2 code. X is considered a relative position,
// with O as the origin.
type F64Bas1 struct {
	O float64
	X float64
}

// F64Bas1New creates a new 1D space basis.
func F64Bas1New(o, x float64) *F64Bas1 {
	return &F64Bas1{o, x}
}

// F64Bas1Default creates a new 1D space basis, using default
// orthogonal settings (origin at 0 with vectors 1).
func F64Bas1Default() *F64Bas1 {
	return &F64Bas1{vpnumber.F64Const0, vpnumber.F64Const1}
}

// Normalize normalizes a 1D space basis, by normalizing
// all vectors in it.
// It modifies the basis, and returns a pointer on it.
func (bas *F64Bas1) Normalize() *F64Bas1 {
	*bas = *F64Bas1Normalize(bas)

	return bas
}

// F64Bas1Normalize normalizes a 1D space basis, by normalizing
// all vectors in it.
// Args is left untouched, a pointer on a new object is returned.
func F64Bas1Normalize(bas *F64Bas1) *F64Bas1 {
	return &F64Bas1{bas.O, vpnumber.F64Const1}
}
