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
	"math/rand"
	"time"
)

const ID_RAND_MASK_OR uint64 = 0x0400000000000000
const ID_RAND_MASK_AND uint64 = 0x7fffffffffffffff

func NewIdRandSource() rand.Source {
	var source_seed uint64
	var source rand.Source
	var now time.Time

	now = time.Now()
	source_seed = PredictableRandomU64(int64(now.Year()) * int64(now.Month()) * int64(now.Day()) * int64(now.Hour()) * int64(now.Minute()) * int64(now.Second()) * int64(now.Nanosecond()))
	source = rand.NewSource(int64(source_seed))

	return source
}

func NewIdRandSeed(source rand.Source) uint64 {
	var seed uint64

	seed = uint64(source.Int63())
	seed = seed | ID_RAND_MASK_OR
	seed = seed & ID_RAND_MASK_AND

	return seed
}

func IdRandMask(raw uint64) uint64 {
	return (raw | ID_RAND_MASK_OR) & ID_RAND_MASK_AND
}
