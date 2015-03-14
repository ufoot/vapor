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

package vpcrypto

import (
	"math/big"
)

func ZeroesInBigInt(i *big.Int) int {
	n := i.BitLen()
	var p int

	for p = 0; p < n && i.Bit(p) == 0; p++ {
	}

	return p
}

func ZeroesInBytes(data []byte) int {
	var i big.Int
	var p int

	i.SetBytes(data)
	n := i.BitLen()
	for p = 0; p < n && i.Bit(p) == 0; p++ {
	}

	return p
}

func ZeroesIn64(i uint64) int {
	var p int

	for p = 0; p < 64 && (i&1) == 0; p++ {
		i >>= 1
	}

	return p
}

func ZeroesIn32(i uint32) int {
	var p int

	for p = 0; p < 32 && (i&1) == 0; p++ {
		i >>= 1
	}

	return p
}
