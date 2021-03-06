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
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
	"github.com/ufoot/vapor/go/vpvec3"
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

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			retP += F32Bernstein(3, i, u) * F32Bernstein(3, j, v) * p[i][j]
			retDu += F32BernsteinDerivative(3, j, u) * F32Bernstein(3, j, v) * p[i][j]
			retDv += F32Bernstein(3, j, u) * F32BernsteinDerivative(3, j, v) * p[i][j]
		}
	}

	return retP, retDu, retDv
}

// F32QuadraticSurface2d returns a quadratic Bezier surface between 9 points.
func F32QuadraticSurface2d(p *[3][3]vpvec2.F32, u float32, v float32) (*vpvec2.F32, *vpvec2.F32, *vpvec2.F32) {
	var retP, retDu, retDv *vpvec2.F32

	switch {
	case u < vpnumber.F32Const0:
		retP, retDv = F32QuadraticCurve2d(&(p[0][0]), &(p[0][1]), &(p[0][2]), v)
		return retP, new(vpvec2.F32), retDv
	case u > vpnumber.F32Const1:
		retP, retDv = F32QuadraticCurve2d(&(p[2][0]), &(p[2][1]), &(p[2][2]), v)
		return retP, new(vpvec2.F32), retDv
	case v < vpnumber.F32Const0:
		retP, retDv = F32QuadraticCurve2d(&(p[0][0]), &(p[1][0]), &(p[2][0]), u)
		return retP, retDu, new(vpvec2.F32)
	case v > vpnumber.F32Const1:
		retP, retDv = F32QuadraticCurve2d(&(p[0][2]), &(p[1][2]), &(p[2][2]), u)
		return retP, retDu, new(vpvec2.F32)
	}

	retP = new(vpvec2.F32)
	retDu = new(vpvec2.F32)
	retDv = new(vpvec2.F32)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			retP.Add(vpvec2.F32MulScale(&(p[i][j]), F32Bernstein(3, i, u)*F32Bernstein(3, j, v)))
			retDu.Add(vpvec2.F32MulScale(&(p[i][j]), F32BernsteinDerivative(3, j, u)*F32Bernstein(3, j, v)))
			retDv.Add(vpvec2.F32MulScale(&(p[i][j]), F32Bernstein(3, j, u)*F32BernsteinDerivative(3, j, v)))
		}
	}

	return retP, retDu, retDv
}

// F32QuadraticSurface3d returns a quadratic Bezier surface between 9 points.
func F32QuadraticSurface3d(p *[3][3]vpvec3.F32, u float32, v float32) (*vpvec3.F32, *vpvec3.F32, *vpvec3.F32) {
	var retP, retDu, retDv *vpvec3.F32

	switch {
	case u < vpnumber.F32Const0:
		retP, retDv = F32QuadraticCurve3d(&(p[0][0]), &(p[0][1]), &(p[0][2]), v)
		return retP, new(vpvec3.F32), retDv
	case u > vpnumber.F32Const1:
		retP, retDv = F32QuadraticCurve3d(&(p[2][0]), &(p[2][1]), &(p[2][2]), v)
		return retP, new(vpvec3.F32), retDv
	case v < vpnumber.F32Const0:
		retP, retDv = F32QuadraticCurve3d(&(p[0][0]), &(p[1][0]), &(p[2][0]), u)
		return retP, retDu, new(vpvec3.F32)
	case v > vpnumber.F32Const1:
		retP, retDv = F32QuadraticCurve3d(&(p[0][2]), &(p[1][2]), &(p[2][2]), u)
		return retP, retDu, new(vpvec3.F32)
	}

	retP = new(vpvec3.F32)
	retDu = new(vpvec3.F32)
	retDv = new(vpvec3.F32)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			retP.Add(vpvec3.F32MulScale(&(p[i][j]), F32Bernstein(3, i, u)*F32Bernstein(3, j, v)))
			retDu.Add(vpvec3.F32MulScale(&(p[i][j]), F32BernsteinDerivative(3, j, u)*F32Bernstein(3, j, v)))
			retDv.Add(vpvec3.F32MulScale(&(p[i][j]), F32Bernstein(3, j, u)*F32BernsteinDerivative(3, j, v)))
		}
	}

	return retP, retDu, retDv
}
