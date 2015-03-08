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

// X32Sin is a lookup table based implementation of the sinus
// function, working with fixed point numbers on 32 bits.
func X32Sin(x vpnumber.X32) vpnumber.X32 {
	var i int
	var d vpnumber.X32
	var s vpnumber.X32 = 1

	for i = 10; i >= 0; i-- {
		d = X32Const2Pi << uint(i)
		for x > d {
			x -= d
		}
		for x < 0 {
			x += d
		}
	}

	if x > X32ConstPi {
		x -= X32ConstPi
		s = -1
	}
	if x > X32ConstPi2 {
		x = X32ConstPi - x
	}

	i = (int(x) * x32TableSize) / int(X32ConstPi2)

	switch {
	case i < 0:
		i = 0
	case i > x32TableSize:
		i = x32TableSize
	}

	return x32SinTable[i] * s
}

// X64Sin is a lookup table based implementation of the sinus
// function, working with fixed point numbers on 64 bits.
func X64Sin(x vpnumber.X64) vpnumber.X64 {
	var i int
	var d vpnumber.X64
	var s vpnumber.X64 = 1

	for i = 10; i >= 0; i-- {
		d = X64Const2Pi << uint(i)
		for x > d {
			x -= d
		}
		for x < 0 {
			x += d
		}
	}

	if x > X64ConstPi {
		x -= X64ConstPi
		s = -1
	}
	if x > X64ConstPi2 {
		x = X64ConstPi - x
	}

	i = (int(x) * x64TableSize) / int(X64ConstPi2)

	switch {
	case i < 0:
		i = 0
	case i > x64TableSize:
		i = x64TableSize
	}

	return x64SinTable[i] * s
}
