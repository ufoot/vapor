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

import (
	"github.com/ufoot/vapor/go/vpnumber"
)

// X32Atan is a lookup table based implementation of the arctangent
// function, working with fixed point numbers on 32 bits.
func X32Atan(x vpnumber.X32) vpnumber.X32 {
	var i int
	var s vpnumber.X32 = 1
	var inv = false

	if x < 0 {
		x = -x
		s = -1
	}

	if x > vpnumber.X32Const1 {
		x = vpnumber.X32Div(vpnumber.X32Const1, x)
		inv = true
	}

	i = (int(x) * x32TableSize) / int(vpnumber.X32Const1)

	switch {
	case i < 0:
		i = 0
	case i > x32TableSize:
		i = x32TableSize
	}

	if inv {
		return (X32ConstPi2 - x32AtanTable[i]) * s
	}
	return x32AtanTable[i] * s
}

// X64Atan is a lookup table based implementation of the arctangent
// function, working with fixed point numbers on 64 bits.
func X64Atan(x vpnumber.X64) vpnumber.X64 {
	var i int
	var s vpnumber.X64 = 1
	var inv = false

	if x < 0 {
		x = -x
		s = -1
	}

	if x > vpnumber.X64Const1 {
		x = vpnumber.X64Div(vpnumber.X64Const1, x)
		inv = true
	}

	i = (int(x) * x64TableSize) / int(vpnumber.X64Const1)

	switch {
	case i < 0:
		i = 0
	case i > x64TableSize:
		i = x64TableSize
	}

	if inv {
		return (X64ConstPi2 - x64AtanTable[i]) * s
	}
	return x64AtanTable[i] * s
}
