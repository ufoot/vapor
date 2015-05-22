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
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
)

// F32QuadraticCurve1d returns the quadratic Bezier curve from p0 to p2,
// going through p1. Point p1 is typically not reached, but it influences
// the curve, as it heads to p1 from p0, and conversely from p2 to p1.
// Works on float32 scalars.
// First returned value is the position, second is the derivative.
func F32QuadraticCurve1d(p0, p1, p2 float32, t float32) (float32, float32) {
	switch {
	case t < vpnumber.F32Const0:
		return p0, vpnumber.F32Const0
	case t > vpnumber.F32Const1:
		return p1, vpnumber.F32Const0
	}

	oneMinusT := vpnumber.F32Const1 - t

	retP := p0*oneMinusT*oneMinusT + p1*2*oneMinusT*t + p2*t*t
	retDt := (p1-p0)*2*oneMinusT + (p2-p1)*2*t

	return retP, retDt
}

// F32QuadraticCurve2d returns the quadratic Bezier curve from p0 to p2,
// going through p1. Point p1 is typically not reached, but it influences
// the curve, as it heads to p1 from p0, and conversely from p2 to p1.
// Works on float32 2d vectors.
// First returned value is the position, second is the derivative.
func F32QuadraticCurve2d(p0, p1, p2 *vpvec2.F32, t float32) (*vpvec2.F32, *vpvec2.F32) {
	switch {
	case t < vpnumber.F32Const0:
		return p0, new(vpvec2.F32)
	case t > vpnumber.F32Const1:
		return p1, new(vpvec2.F32)
	}

	oneMinusT := vpnumber.F32Const1 - t

	retP := vpvec2.F32Add(vpvec2.F32MulScale(p0, oneMinusT*oneMinusT), vpvec2.F32MulScale(p1, 2*oneMinusT*t)).Add(vpvec2.F32MulScale(p2, t*t))
	retDt := vpvec2.F32Sub(p1, p0).MulScale(2 * oneMinusT).Add(vpvec2.F32Sub(p2, p1).MulScale(2 * t))

	return retP, retDt
}

// F32QuadraticCurve3d returns the quadratic Bezier curve from p0 to p2,
// going through p1. Point p1 is typically not reached, but it influences
// the curve, as it heads to p1 from p0, and conversely from p2 to p1.
// Works on float32 3d vectors.
// First returned value is the position, second is the derivative.
func F32QuadraticCurve3d(p0, p1, p2 *vpvec3.F32, t float32) (*vpvec3.F32, *vpvec3.F32) {
	switch {
	case t < vpnumber.F32Const0:
		return p0, new(vpvec3.F32)
	case t > vpnumber.F32Const1:
		return p1, new(vpvec3.F32)
	}

	oneMinusT := vpnumber.F32Const1 - t

	retP := vpvec3.F32Add(vpvec3.F32MulScale(p0, oneMinusT*oneMinusT), vpvec3.F32MulScale(p1, 2*oneMinusT*t)).Add(vpvec3.F32MulScale(p2, t*t))
	retDt := vpvec3.F32Sub(p1, p0).MulScale(2 * oneMinusT).Add(vpvec3.F32Sub(p2, p1).MulScale(2 * t))

	return retP, retDt
}
