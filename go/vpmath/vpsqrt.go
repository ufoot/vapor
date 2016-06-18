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

// X32Sqrt is a lookup table based implementation of the square root
// function, working with fixed point numbers on 32 bits.
func X32Sqrt(x vpnumber.X32) vpnumber.X32 {
	if x <= 0 {
		return 0
	}
	exponent1 := vpnumber.X32Exponent(x)
	if (exponent1 & 0x01) != 0 {
		exponent1--
	}
	switch {
	case exponent1 < 0:
		x = x << uint(-exponent1)
	case exponent1 > 0:
		x = x >> uint(exponent1)
	}
	i := ((int(x) - int(vpnumber.X32Const1)) * x32TableSize) / (3 * int(vpnumber.X32Const1))
	switch {
	case i < 0:
		i = 0
	case i > x32TableSize:
		i = x32TableSize
	}
	mantis := x32SqrtTable[i]
	exponent2 := exponent1 >> 1
	if exponent2 < 0 {
		return mantis >> uint(-exponent2)
	}
	return mantis << uint(exponent2)
}

// X64Sqrt is a lookup table based implementation of the square root
// function, working with fixed point numbers on 64 bits.
func X64Sqrt(x vpnumber.X64) vpnumber.X64 {
	if x <= 0 {
		return 0
	}
	exponent1 := vpnumber.X64Exponent(x)
	if (exponent1 & 0x01) != 0 {
		exponent1--
	}
	switch {
	case exponent1 < 0:
		x = x << uint(-exponent1)
	case exponent1 > 0:
		x = x >> uint(exponent1)
	}
	i := ((int(x) - int(vpnumber.X64Const1)) * x64TableSize) / (3 * int(vpnumber.X64Const1))
	switch {
	case i < 0:
		i = 0
	case i > x64TableSize:
		i = x64TableSize
	}
	mantis := x64SqrtTable[i]
	exponent2 := exponent1 >> 1
	if exponent2 < 0 {
		return mantis >> uint(-exponent2)
	}
	return mantis << uint(exponent2)
}
