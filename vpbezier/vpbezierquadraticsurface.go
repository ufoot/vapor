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

package vpbezier

import (
	"github.com/ufoot/vapor/vpnumber"
	//"github.com/ufoot/vapor/vpvec2"
	//"github.com/ufoot/vapor/vpvec3"
)

// F32QuadraticSurface1d returns a quadratic Bezier surface between 9 points.
func F32QuadraticSurface1d(p [3][3]float32, u float32, v float32) (float32, float32, float32) {
	var retP, retDu, retDv float32

	switch {
	case u < vpnumber.F32Const0:
		retP, retDv = F32QuadraticCurve1d(p[0][0], p[0][1], p[0][2], v)
		return retP, vpnumber.F32Const0, retDv
	case u > vpnumber.F32Const1:
		retP, retDv = F32QuadraticCurve1d(p[2][0], p[2][1], p[2][2], v)
		return retP, vpnumber.F32Const0, retDv
	case v < vpnumber.F32Const0:
		retP, retDv = F32QuadraticCurve1d(p[0][0], p[1][0], p[2][0], u)
		return retP, retDu, vpnumber.F32Const0
	case v > vpnumber.F32Const1:
		retP, retDv = F32QuadraticCurve1d(p[0][2], p[1][2], p[2][2], u)
		return retP, retDu, vpnumber.F32Const0
	}

	//oneMinusU := vpnumber.F32Const1 - u
	//oneMinusV := vpnumber.F32Const1 - v

	// todo...

	return retP, retDu, retDv
}
