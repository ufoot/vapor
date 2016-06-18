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

package vpmath

//go:generate bash ./stamp.sh
	
import (
	"github.com/ufoot/vapor/go/vpnumber"
)

// F32Unlerp performs the inverse of a linear interpolation between a and b.
func F32Unlerp(f1, f2, f float32) float32 {
	switch {
	case f1 < f2:
		switch {
		case f <= f1:
			return vpnumber.F32Const0
		case f >= f2:
			return vpnumber.F32Const1
		}
	case f1 > f2:
		switch {
		case f >= f1:
			return vpnumber.F32Const0
		case f <= f2:
			return vpnumber.F32Const1
		}
	default:
		return vpnumber.F32Const0
	}

	return (f - f1) / (f2 - f1)
}

// F64Unlerp performs the inverse of a linear interpolation between a and b.
func F64Unlerp(f1, f2, f float64) float64 {
	switch {
	case f1 < f2:
		switch {
		case f <= f1:
			return vpnumber.F64Const0
		case f >= f2:
			return vpnumber.F64Const1
		}
	case f1 > f2:
		switch {
		case f >= f1:
			return vpnumber.F64Const0
		case f <= f2:
			return vpnumber.F64Const1
		}
	default:
		return vpnumber.F64Const0
	}

	return (f - f1) / (f2 - f1)
}

// X32Unlerp performs the inverse of a linear interpolation between a and b.
func X32Unlerp(x1, x2, x vpnumber.X32) vpnumber.X32 {
	switch {
	case x1 < x2:
		switch {
		case x <= x1:
			return vpnumber.X32Const0
		case x >= x2:
			return vpnumber.X32Const1
		}
	case x1 > x2:
		switch {
		case x >= x1:
			return vpnumber.X32Const0
		case x <= x2:
			return vpnumber.X32Const1
		}
	default:
		return vpnumber.X32Const0
	}

	return vpnumber.X32Div(x-x1, x2-x1)
}

// X64Unlerp performs the inverse of a linear interpolation between a and b.
func X64Unlerp(x1, x2, x vpnumber.X64) vpnumber.X64 {
	switch {
	case x1 < x2:
		switch {
		case x <= x1:
			return vpnumber.X64Const0
		case x >= x2:
			return vpnumber.X64Const1
		}
	case x1 > x2:
		switch {
		case x >= x1:
			return vpnumber.X64Const0
		case x <= x2:
			return vpnumber.X64Const1
		}
	default:
		return vpnumber.X64Const0
	}

	return vpnumber.X64Div(x-x1, x2-x1)
}
