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

package vpkeydx

import (
	"fmt"
	"math/big"
)

const n31 = 31
const n256 = 256
const nOffset1 = n256 - n31
const nOffset2 = n256 - 2*n31
const nOffset3 = n256 - 3*n31
const mask31 = 0x7fffffff
const mul31 = int64(0x80000000)

// Mod31 returns the value modulo 2<<32. Usefull to avoid
// out-of-range errors when manipulating coordinates.
func Mod31(i int32) int32 {
	return i & mask31
}

// Scale31 returns a value expressed on 31 bit, between [0,n).
// The scale is linear, that is, if i == n / x, then returned
// value is about 2<<31 / x.
func Scale31(i int32, n int32) int32 {
	if i < 0 {
		i = 0
	}
	if n < 0 {
		n = 1
	}
	if i >= n {
		i = n - 1
	}
	i64 := int64(i)
	n64 := int64(n)

	return Mod31(int32((i64 * mul31) / n64))
}

// Inc31 increases the value by 1, then performs a modulo 2<<32.
func Inc31(i int32) int32 {
	return Mod31(i + 1)
}

// Dec31 decreases the value by 1, then performs a modulo 2<<32.
func Dec31(i int32) int32 {
	return Mod31(i - 1)
}

func toBigInt31(i int32) (*big.Int, error) {
	var ret big.Int

	ret.SetUint64(uint64(Mod31(i)))
	if ret.Int64() != int64(i) {
		return nil, fmt.Errorf("unable to convert %d to int31, out of range", i)
	}

	return &ret, nil
}
