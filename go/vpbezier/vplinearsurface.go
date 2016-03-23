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

package vpbezier

import (
	"github.com/ufoot/vapor/go/vpmath"
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
	"github.com/ufoot/vapor/go/vpvec3"
)

func scalarFromF32Vec2(p *[2][2]vpvec2.F32, i int) [2][2]float32 {
	return [2][2]float32{{p[0][0][i], p[0][1][i]}, {p[1][0][i], p[1][1][i]}}
}

func scalarFromF32Vec3(p *[2][2]vpvec3.F32, i int) [2][2]float32 {
	return [2][2]float32{{p[0][0][i], p[0][1][i]}, {p[1][0][i], p[1][1][i]}}
}

// F32LinearSurface1d returns the linear Bezier surface between 4 points.
func F32LinearSurface1d(p [2][2]float32, u, v float32) (float32, float32, float32) {
	switch {
	case u < vpnumber.F32Const0:
		return vpmath.F32Lerp(p[0][0], p[0][1], v), vpnumber.F32Const0, p[0][1] - p[0][0]
	case u > vpnumber.F32Const1:
		return vpmath.F32Lerp(p[1][0], p[1][1], v), vpnumber.F32Const0, p[1][1] - p[1][0]
	case v < vpnumber.F32Const0:
		return vpmath.F32Lerp(p[0][0], p[1][0], u), p[1][0] - p[0][0], vpnumber.F32Const0
	case v > vpnumber.F32Const1:
		return vpmath.F32Lerp(p[0][1], p[1][1], u), p[1][1] - p[0][1], vpnumber.F32Const0
	}

	oneMinusU := vpnumber.F32Const1 - u
	oneMinusV := vpnumber.F32Const1 - v

	retP := vpmath.F32Lerp2(p, u, v)
	retDu := p[1][0]*oneMinusV + p[1][1]*v - p[0][0]*oneMinusV - p[0][1]*v
	retDv := p[0][1]*oneMinusU + p[1][1]*u - p[0][0]*oneMinusU - p[1][0]*u

	return retP, retDu, retDv
}

// F32LinearSurface2d returns the linear Bezier surface between 4 points.
func F32LinearSurface2d(p *[2][2]vpvec2.F32, u, v float32) (*vpvec2.F32, *vpvec2.F32, *vpvec2.F32) {
	switch {
	case u < vpnumber.F32Const0:
		return vpvec2.F32Lerp(&(p[0][0]), &(p[0][1]), v), new(vpvec2.F32), vpvec2.F32Sub(&(p[0][1]), &(p[0][0]))
	case u > vpnumber.F32Const1:
		return vpvec2.F32Lerp(&(p[1][0]), &(p[1][1]), v), new(vpvec2.F32), vpvec2.F32Sub(&(p[1][1]), &(p[1][0]))
	case v < vpnumber.F32Const0:
		return vpvec2.F32Lerp(&(p[0][0]), &(p[1][0]), u), vpvec2.F32Sub(&(p[1][0]), &(p[0][0])), new(vpvec2.F32)
	case v > vpnumber.F32Const1:
		return vpvec2.F32Lerp(&(p[0][1]), &(p[1][1]), u), vpvec2.F32Sub(&(p[1][1]), &(p[0][1])), new(vpvec2.F32)
	}

	oneMinusU := vpnumber.F32Const1 - u
	oneMinusV := vpnumber.F32Const1 - v

	retP := vpvec2.F32New(vpmath.F32Lerp2(scalarFromF32Vec2(p, 0), u, v), vpmath.F32Lerp2(scalarFromF32Vec2(p, 1), u, v))
	retDu := vpvec2.F32MulScale(&(p[1][0]), oneMinusV).Add(vpvec2.F32MulScale(&(p[1][1]), v)).Sub(vpvec2.F32MulScale(&(p[0][0]), oneMinusV)).Sub(vpvec2.F32MulScale(&(p[0][1]), v))
	retDv := vpvec2.F32MulScale(&(p[0][1]), oneMinusU).Add(vpvec2.F32MulScale(&(p[1][1]), u)).Sub(vpvec2.F32MulScale(&(p[0][0]), oneMinusU)).Sub(vpvec2.F32MulScale(&(p[1][0]), u))

	return retP, retDu, retDv
}

// F32LinearSurface3d returns the linear Bezier surface between 4 points.
func F32LinearSurface3d(p *[2][2]vpvec3.F32, u, v float32) (*vpvec3.F32, *vpvec3.F32, *vpvec3.F32) {
	switch {
	case u < vpnumber.F32Const0:
		return vpvec3.F32Lerp(&(p[0][0]), &(p[0][1]), v), new(vpvec3.F32), vpvec3.F32Sub(&(p[0][1]), &(p[0][0]))
	case u > vpnumber.F32Const1:
		return vpvec3.F32Lerp(&(p[1][0]), &(p[1][1]), v), new(vpvec3.F32), vpvec3.F32Sub(&(p[1][1]), &(p[1][0]))
	case v < vpnumber.F32Const0:
		return vpvec3.F32Lerp(&(p[0][0]), &(p[1][0]), u), vpvec3.F32Sub(&(p[1][0]), &(p[0][0])), new(vpvec3.F32)
	case v > vpnumber.F32Const1:
		return vpvec3.F32Lerp(&(p[0][1]), &(p[1][1]), u), vpvec3.F32Sub(&(p[1][1]), &(p[0][1])), new(vpvec3.F32)
	}

	oneMinusU := vpnumber.F32Const1 - u
	oneMinusV := vpnumber.F32Const1 - v

	retP := vpvec3.F32New(vpmath.F32Lerp2(scalarFromF32Vec3(p, 0), u, v), vpmath.F32Lerp2(scalarFromF32Vec3(p, 1), u, v), vpmath.F32Lerp2(scalarFromF32Vec3(p, 2), u, v))
	retDu := vpvec3.F32MulScale(&(p[1][0]), oneMinusV).Add(vpvec3.F32MulScale(&(p[1][1]), v)).Sub(vpvec3.F32MulScale(&(p[0][0]), oneMinusV)).Sub(vpvec3.F32MulScale(&(p[0][1]), v))
	retDv := vpvec3.F32MulScale(&(p[0][1]), oneMinusU).Add(vpvec3.F32MulScale(&(p[1][1]), u)).Sub(vpvec3.F32MulScale(&(p[0][0]), oneMinusU)).Sub(vpvec3.F32MulScale(&(p[1][0]), u))

	return retP, retDu, retDv
}
