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
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpmath

import (
	"ufoot.org/vapor/vpnumber"
)

// X32Cos is a lookup table based implementation of the cosinus
// function, working with fixed point numbers on 32 bits.
func X32Cos(x vpnumber.X32) vpnumber.X32 {
	return X32Sin(x + X32ConstPi2)
}

// X64Cos is a lookup table based implementation of the cosinus
// function, working with fixed point numbers on 64 bits.
func X64Cos(x vpnumber.X64) vpnumber.X64 {
	return X64Sin(x + X64ConstPi2)
}
