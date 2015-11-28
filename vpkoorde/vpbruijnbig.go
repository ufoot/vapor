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

package vpkoorde

import (
	"math/big"
)

func bigToBytes(i *big.Int, nbBytes int) []byte {
	raw := i.Bytes()
	l := len(raw)
	if l == nbBytes {
		return raw
	}
	if l > nbBytes {
		return raw[l-nbBytes : l]
	}
	return append(make([]byte, nbBytes-l), raw...)
}

func bytesToBig(b []byte) *big.Int {
	i := big.NewInt(0)
	i.SetBytes(b)
	return i
}

func nextBigFirst(x, bm, max *big.Int) *big.Int {
	nf := big.NewInt(0)
	nf.Mul(x, bm)
	return nf.Mod(nf, max)
}

func prevBigFirst(x, bm *big.Int) *big.Int {
	pf := big.NewInt(0)
	// no need to do a modulo here : it *is* smaller than m**n
	return pf.Div(x, bm)
}

func composeBig(x, y, bm, max *big.Int, m, n, i int) *big.Int {
	c := big.NewInt(0)

	if i <= 0 {
		return c.Set(x)
	}
	if i >= n {
		return c.Set(y)
	}

	tX := big.NewInt(0)
	tX.Exp(bm, big.NewInt(int64(i)), max)
	tX.Mul(tX, x)

	tY := big.NewInt(0)
	tY.Exp(bm, big.NewInt(int64(n-i)), max)
	tY.Div(y, tY)

	c.Add(c, tX)
	c.Add(c, tY)
	c.Mod(c, max)

	return c
}
