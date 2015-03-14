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
	"math/rand"
	"time"
)

func NewSource() rand.Source {
	var seed uint64
	var source rand.Source
	var now time.Time

	now = time.Now()
	seed = PseudoRand64(uint64(now.Year())*uint64(now.Month())*uint64(now.Day())*uint64(now.Hour())*uint64(now.Minute())*uint64(now.Second())*uint64(now.Nanosecond()), 0)
	source = rand.NewSource(int64(seed))

	return source
}

func NewRand() *rand.Rand {
	return rand.New(NewSource())
}

func Rand512(r *rand.Rand, n *big.Int) *big.Int {
	ret := big.NewInt(0)

	if n == nil || n.Cmp(big1) <= 0 {
		n = big.NewInt(1)
		n.Lsh(n, 512)
	}

	return ret.Rand(r, n)
}

func Rand256(r *rand.Rand, n *big.Int) *big.Int {
	ret := big.NewInt(0)

	if n == nil || n.Cmp(big1) <= 0 {
		n = big.NewInt(1)
		n.Lsh(n, 256)
	}

	return ret.Rand(r, n)
}

func Rand128(r *rand.Rand, n *big.Int) *big.Int {
	ret := big.NewInt(0)

	if n == nil || n.Cmp(big1) <= 0 {
		n = big.NewInt(1)
		n.Lsh(n, 128)
	}

	return ret.Rand(r, n)
}

func Rand64(r *rand.Rand, n uint64) uint64 {
	ret := uint64(r.Int63()<<1) ^ uint64(r.Int31())
	if n > 0 {
		ret %= n
	}

	return ret
}

func Rand32(r *rand.Rand, n uint32) uint32 {
	ret := uint32(r.Int63())
	if n > 0 {
		ret %= n
	}

	return ret
}
