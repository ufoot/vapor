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

//go:generate bash ./stamp.sh
	
import (
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
	"github.com/ufoot/vapor/go/vpvec3"
)

// F32CubicCurve1d returns the cubic Bezier curve from p0 to p3,
// going through p1 and p2. Points p1 and p2 are typically not reached,
// but they influence the curve, as it heads to p1 from p0, and conversely
// from p3 to p2. Works on float32 scalars.
// First returned value is the position, second is the derivative.
func F32CubicCurve1d(p0, p1, p2, p3 float32, t float32) (float32, float32) {
	switch {
	case t < vpnumber.F32Const0:
		return p0, vpnumber.F32Const0
	case t > vpnumber.F32Const1:
		return p1, vpnumber.F32Const0
	}

	oneMinusT := vpnumber.F32Const1 - t

	retP := p0*oneMinusT*oneMinusT*oneMinusT + p1*3*oneMinusT*oneMinusT*t + p2*3*oneMinusT*t*t + p3*t*t*t
	retDt := (p1-p0)*3*oneMinusT*oneMinusT + (p2-p1)*6*oneMinusT*t + (p3-p2)*3*t*t

	return retP, retDt
}

// F32CubicCurve2d returns the cubic Bezier curve from p0 to p3,
// going through p1 and p2. Points p1 and p2 are typically not reached,
// but they influence the curve, as it heads to p1 from p0, and conversely
// from p3 to p2. Works on float32 2d vectors.
// First returned value is the position, second is the derivative.
func F32CubicCurve2d(p0, p1, p2, p3 *vpvec2.F32, t float32) (*vpvec2.F32, *vpvec2.F32) {
	switch {
	case t < vpnumber.F32Const0:
		return p0, new(vpvec2.F32)
	case t > vpnumber.F32Const1:
		return p1, new(vpvec2.F32)
	}

	oneMinusT := vpnumber.F32Const1 - t

	retP := vpvec2.F32Add(vpvec2.F32MulScale(p0, oneMinusT*oneMinusT*oneMinusT), vpvec2.F32MulScale(p1, 3*oneMinusT*oneMinusT*t)).Add(vpvec2.F32MulScale(p2, 3*oneMinusT*t*t)).Add(vpvec2.F32MulScale(p3, t*t*t))
	retDt := vpvec2.F32Sub(p1, p0).MulScale(3 * oneMinusT * oneMinusT).Add(vpvec2.F32Sub(p2, p1).MulScale(6 * oneMinusT * t)).Add(vpvec2.F32Sub(p3, p2).MulScale(3 * t * t))

	return retP, retDt
}

// F32CubicCurve3d returns the cubic Bezier curve from p0 to p3,
// going through p1 and p2. Points p1 and p2 are typically not reached,
// but they influence the curve, as it heads to p1 from p0, and conversely
// from p3 to p2. Works on float32 3d vectors.
// First returned value is the position, second is the derivative.
func F32CubicCurve3d(p0, p1, p2, p3 *vpvec3.F32, t float32) (*vpvec3.F32, *vpvec3.F32) {
	switch {
	case t < vpnumber.F32Const0:
		return p0, new(vpvec3.F32)
	case t > vpnumber.F32Const1:
		return p1, new(vpvec3.F32)
	}

	oneMinusT := vpnumber.F32Const1 - t

	retP := vpvec3.F32Add(vpvec3.F32MulScale(p0, oneMinusT*oneMinusT*oneMinusT), vpvec3.F32MulScale(p1, 3*oneMinusT*oneMinusT*t)).Add(vpvec3.F32MulScale(p2, 3*oneMinusT*t*t)).Add(vpvec3.F32MulScale(p3, t*t*t))
	retDt := vpvec3.F32Sub(p1, p0).MulScale(3 * oneMinusT * oneMinusT).Add(vpvec3.F32Sub(p2, p1).MulScale(6 * oneMinusT * t)).Add(vpvec3.F32Sub(p3, p2).MulScale(3 * t * t))

	return retP, retDt
}
