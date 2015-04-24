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

// X32Tan is a lookup table based implementation of the tangent
// function, working with fixed point numbers on 32 bits.
func X32Tan(x vpnumber.X32) vpnumber.X32 {
	return vpnumber.X32Div(X32Sin(x), X32Cos(x))
}

// X64Tan is a lookup table based implementation of the tangent
// function, working with fixed point numbers on 64 bits.
func X64Tan(x vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Div(X64Sin(x), X64Cos(x))
}
