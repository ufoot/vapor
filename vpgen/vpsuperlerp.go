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

package vpgen

import (
	"github.com/ufoot/vapor/vpmat4x4"
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec4"
)

var superLerpMat vpmat4x4.F32

func init() {
	// Here we build a matrix which represents (lines) :
	// a*x0^3 + b*x0^2 + c*x0 + d
	// 3*a*x0^2 + 2*b*x0 + c
	// a*x1^3 + b*x1^2 + c*x1 + d
	// 3*a*x1^2 + 2*b*x1 + c
	// with x0=0 and x1=1 (Lerp convention).
	// (a,b,c,d) or to be found by inversing the matrix
	// and multiplying by (f1[0], f1[1], f2[0], f2[1])
	mat := vpmat4x4.F32New(vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, 3*vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const0, vpnumber.F32Const1, 2*vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const0, vpnumber.F32Const1, vpnumber.F32Const0)

	superLerpMat = *vpmat4x4.F32Inv(mat)
}

// F32SuperLerp performs a linear interpolation between a and b,
// with a and b being two vectors containing a value and its derived
// value. Think of a point and the speed at this point.
func F32SuperLerp(f1, f2 *vpvec2.F32, beta float32) *vpvec2.F32 {
	switch {
	case beta < vpnumber.F32Const0:
		return vpvec2.F32New(f1[0], vpnumber.F32Const0)
	case beta > vpnumber.F32Const1:
		return vpvec2.F32New(f2[0], vpnumber.F32Const0)
	case vpnumber.F32IsSimilar(f1[0], f2[0]):
		return vpvec2.F32New(vpmath.F32Lerp(f1[0], f2[0], beta), vpmath.F32Lerp(f1[1], f2[1], beta))
	}

	vec := superLerpMat.MulVec(vpvec4.F32New(f1[0], f1[1], f2[0], f2[1]))

	return vpvec2.F32New(vec[0]*beta*beta*beta+vec[1]*beta*beta+vec[2]*beta+vec[3], 3*vec[0]*beta*beta+2*vec[1]*beta+vec[2])
}
