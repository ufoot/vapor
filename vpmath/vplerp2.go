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

package vpmath

import (
	"github.com/ufoot/vapor/vpnumber"
)

// F32Lerp2 performs a linear interpolation between 4 points.
func F32Lerp2(p [2][2]float32, u, v float32) float32 {
	switch {
	case u <= vpnumber.F32Const0:
		return F32Lerp(p[0][0], p[0][1], v)
	case u >= vpnumber.F32Const1:
		return F32Lerp(p[1][0], p[1][1], v)
	case v <= vpnumber.F32Const0:
		return F32Lerp(p[0][0], p[1][0], u)
	case v >= vpnumber.F32Const1:
		return F32Lerp(p[0][1], p[1][1], u)
	}

	return p[0][0]*(vpnumber.F32Const1-u)*(vpnumber.F32Const1-v) +
		p[0][1]*(vpnumber.F32Const1-u)*v +
		p[1][0]*u*(vpnumber.F32Const1-v) +
		p[1][1]*u*v
}

// F64Lerp2 performs a linear interpolation between 4 points.
func F64Lerp2(p [2][2]float64, u, v float64) float64 {
	switch {
	case u <= vpnumber.F64Const0:
		return F64Lerp(p[0][0], p[0][1], v)
	case u >= vpnumber.F64Const1:
		return F64Lerp(p[1][0], p[1][1], v)
	case v <= vpnumber.F64Const0:
		return F64Lerp(p[0][0], p[1][0], u)
	case v >= vpnumber.F64Const1:
		return F64Lerp(p[0][1], p[1][1], u)
	}

	return p[0][0]*(vpnumber.F64Const1-u)*(vpnumber.F64Const1-v) +
		p[0][1]*(vpnumber.F64Const1-u)*v +
		p[1][0]*u*(vpnumber.F64Const1-v) +
		p[1][1]*u*v
}

// X32Lerp2 performs a linear interpolation between 4 points.
func X32Lerp2(p [2][2]vpnumber.X32, u, v vpnumber.X32) vpnumber.X32 {
	switch {
	case u <= vpnumber.X32Const0:
		return X32Lerp(p[0][0], p[0][1], v)
	case u >= vpnumber.X32Const1:
		return X32Lerp(p[1][0], p[1][1], v)
	case v <= vpnumber.X32Const0:
		return X32Lerp(p[0][0], p[1][0], u)
	case v >= vpnumber.X32Const1:
		return X32Lerp(p[0][1], p[1][1], u)
	}

	return vpnumber.X32Muln(p[0][0], (vpnumber.X32Const1-u), (vpnumber.X32Const1-v)) +
		vpnumber.X32Muln(p[0][1], (vpnumber.X32Const1-u), v) +
		vpnumber.X32Muln(p[1][0], u, (vpnumber.X32Const1-v)) +
		vpnumber.X32Muln(p[1][1], u, v)
}

// X64Lerp2 performs a linear interpolation between 4 points.
func X64Lerp2(p [2][2]vpnumber.X64, u, v vpnumber.X64) vpnumber.X64 {
	switch {
	case u <= vpnumber.X64Const0:
		return X64Lerp(p[0][0], p[0][1], v)
	case u >= vpnumber.X64Const1:
		return X64Lerp(p[1][0], p[1][1], v)
	case v <= vpnumber.X64Const0:
		return X64Lerp(p[0][0], p[1][0], u)
	case v >= vpnumber.X64Const1:
		return X64Lerp(p[0][1], p[1][1], u)
	}

	return vpnumber.X64Muln(p[0][0], (vpnumber.X64Const1-u), (vpnumber.X64Const1-v)) +
		vpnumber.X64Muln(p[0][1], (vpnumber.X64Const1-u), v) +
		vpnumber.X64Muln(p[1][0], u, (vpnumber.X64Const1-v)) +
		vpnumber.X64Muln(p[1][1], u, v)
}
