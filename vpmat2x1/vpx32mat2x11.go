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

package vpmat2x1

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
)

// X32Bas1 is a 1D space basis, composed of 2 points in a 1D space.
// It's defined along with 2x1 matrix code as manipulating such basis
// requires 2x1 code. X is considered a relative position,
// with O as the origin.
type X32Bas1 struct {
	O vpnumber.X32
	X vpnumber.X32
}

// X32Bas1New creates a new 1D space basis.
func X32Bas1New(o, x vpnumber.X32) *X32Bas1 {
	return &X32Bas1{o, x}
}

// X32Bas1Default creates a new 1D space basis, using default
// orthogonal settings (origin at 0 with vectors 1).
func X32Bas1Default() *X32Bas1 {
	return &X32Bas1{vpnumber.X32Const0, vpnumber.X32Const1}
}

// String returns a readable form of the basis.
func (bas *X32Bas1) String() string {
	buf, err := json.Marshal(bas)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// Normalize normalizes a 1D space basis, by normalizing
// all vectors in it.
// It modifies the basis, and returns a pointer on it.
func (bas *X32Bas1) Normalize() *X32Bas1 {
	*bas = *X32Bas1Normalize(bas)

	return bas
}

// X32Bas1Normalize normalizes a 1D space basis, by normalizing
// all vectors in it.
// Args is left untouched, a pointer on a new object is returned.
func X32Bas1Normalize(bas *X32Bas1) *X32Bas1 {
	return &X32Bas1{bas.O, vpnumber.X32Const1}
}
