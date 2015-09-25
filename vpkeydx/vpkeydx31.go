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

func toBigInt31(i int32) (*big.Int, error) {
	var ret big.Int

	ret.SetUint64(uint64(i) & mask31)
	if ret.Int64() != int64(i) {
		return nil, fmt.Errorf("unable to convert %d to int31, out of range", i)
	}

	return &ret, nil
}
