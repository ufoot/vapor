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

package vprand

import (
	"github.com/ufoot/vapor/go/vpsum"
	"math/big"
	"math/rand"
	"time"
)

var bigZero *big.Int
var bigOne512 *big.Int
var bigOne256 *big.Int
var bigOne128 *big.Int

func init() {
	bigZero = big.NewInt(0)
	bigOne512 = big.NewInt(1)
	bigOne512.Lsh(bigOne512, 512)
	bigOne256 = big.NewInt(1)
	bigOne256.Lsh(bigOne256, 256)
	bigOne128 = big.NewInt(1)
	bigOne128.Lsh(bigOne128, 128)
}
func newSource() rand.Source {
	var seed uint64
	var source rand.Source
	var now time.Time

	now = time.Now()
	seed = vpsum.PseudoRand64(uint64(now.Unix()<<32)^uint64(now.Nanosecond()), 0)
	source = rand.NewSource(int64(seed))

	return source
}

// NewRand returns a new random number generator.
// It is seeded with current time, to give it a little entropy.
func NewRand() *rand.Rand {
	return rand.New(newSource())
}

// Rand512 returns a random number on 512 bits.
// If n is nil, 0, or negative, or greater than 2^512, then
// everything is done as if n was 2^512.
func Rand512(r *rand.Rand, n *big.Int) *big.Int {
	ret := big.NewInt(0)

	if n == nil || n.Cmp(bigZero) <= 0 || n.Cmp(bigOne512) > 0 {
		n = bigOne512
	}
	if r == nil {
		r = NewRand()
	}

	return ret.Rand(r, n)
}

// Rand256 returns a random number on 256 bits.
// If n is nil, 0, or negative, or greater than 2^256, then
// everything is done as if n was 2^256.
func Rand256(r *rand.Rand, n *big.Int) *big.Int {
	ret := big.NewInt(0)

	if n == nil || n.Cmp(bigZero) <= 0 || n.Cmp(bigOne256) > 0 {
		n = bigOne256
	}
	if r == nil {
		r = NewRand()
	}

	return ret.Rand(r, n)
}

// Rand128 returns a random number on 128 bits.
// If n is nil, 0, or negative, or greater than 2^128, then
// everything is done as if n was 2^128.
func Rand128(r *rand.Rand, n *big.Int) *big.Int {
	ret := big.NewInt(0)

	if n == nil || n.Cmp(bigZero) <= 0 || n.Cmp(bigOne128) > 0 {
		n = bigOne128
	}
	if r == nil {
		r = NewRand()
	}

	return ret.Rand(r, n)
}

// Rand64 returns a random number on 64 bits.
// If n 0 then everything is done as if n was 2^64.
func Rand64(r *rand.Rand, n uint64) uint64 {
	if r == nil {
		r = NewRand()
	}

	ret := uint64(r.Int63()<<1) ^ uint64(r.Int31())

	if n > 0 {
		ret %= n
	}

	return ret
}

// Rand32 returns a random number on 64 bits.
// If n 0 then everything is done as if n was 2^32.
func Rand32(r *rand.Rand, n uint32) uint32 {
	if r == nil {
		r = NewRand()
	}

	ret := uint32(r.Int63())

	if n > 0 {
		ret %= n
	}

	return ret
}
