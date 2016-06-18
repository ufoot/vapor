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

// F32Lerp performs a linear interpolation between a and b.
func F32Lerp(f1, f2, beta float32) float32 {
	switch {
	case beta <= vpnumber.F32Const0:
		return f1
	case beta >= vpnumber.F32Const1:
		return f2
	}

	return f1*(vpnumber.F32Const1-beta) + f2*beta
}

// F64Lerp performs a linear interpolation between a and b.
func F64Lerp(f1, f2, beta float64) float64 {
	switch {
	case beta <= vpnumber.F64Const0:
		return f1
	case beta >= vpnumber.F64Const1:
		return f2
	}

	return f1*(vpnumber.F64Const1-beta) + f2*beta
}

// X32Lerp performs a linear interpolation between a and b.
func X32Lerp(x1, x2, beta vpnumber.X32) vpnumber.X32 {
	switch {
	case beta <= vpnumber.X32Const0:
		return x1
	case beta >= vpnumber.X32Const1:
		return x2
	}

	return vpnumber.X32Mul(x1, (vpnumber.X32Const1-beta)) + vpnumber.X32Mul(x2, beta)
}

// X64Lerp performs a linear interpolation between a and b.
func X64Lerp(x1, x2, beta vpnumber.X64) vpnumber.X64 {
	switch {
	case beta <= vpnumber.X64Const0:
		return x1
	case beta >= vpnumber.X64Const1:
		return x2
	}

	return vpnumber.X64Mul(x1, (vpnumber.X64Const1-beta)) + vpnumber.X64Mul(x2, beta)
}
