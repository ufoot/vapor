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
	"testing"
)

const init_i = 24
const expected_z = 3

func TestZeroesInBigInt(t *testing.T) {
	i := big.NewInt(init_i)

	z := ZeroesInBigInt(i)
	if z != expected_z {
		t.Errorf("bad zeroes in BigInt for %d, got %d, expected %d", init_i, z, expected_z)
	} else {
		t.Logf("zeroes in BigInt for %d is %d", init_i, expected_z)
	}
	i.SetUint64(0)
	z = ZeroesInBigInt(i)
	if z != 0 {
		t.Errorf("bad zeroes in BigInt for 0, got %d, expected 0", 0, z)
	} else {
		t.Logf("zeroes in BigInt for 0 is 0")
	}
}

func TestZeroesInBytes(t *testing.T) {
	b := make([]byte, 10)

	b[9] = byte(init_i)
	z := ZeroesInBytes(b)
	if z != expected_z {
		t.Errorf("bad zeroes in Bytes for %d, got %d, expected %d", init_i, z, expected_z)
	} else {
		t.Logf("zeroes in Bytes for %d is %d", init_i, expected_z)
	}
	b = make([]byte, 0)
	z = ZeroesInBytes(b)
	if z != 0 {
		t.Errorf("bad zeroes in Bytes for empty Bytes, got %d, expected 0", z)
	} else {
		t.Logf("zeroes in Bytes for empty Bytes is 0")
	}
	b = make([]byte, 1)
	z = ZeroesInBytes(b)
	if z != 0 {
		t.Errorf("bad zeroes in Bytes for [1]Bytes, got %d, expected 8", z)
	} else {
		t.Logf("zeroes in Bytes for [1]Bytes is 8")
	}
}

func TestZeroesIn64(t *testing.T) {
	i := uint64(init_i)

	z := ZeroesIn64(i)
	if z != expected_z {
		t.Errorf("bad zeroes in uint64 for %d, got %d, expected %d", init_i, z, expected_z)
	} else {
		t.Logf("zeroes in uint64 for %d is %d", init_i, expected_z)
	}
	z = ZeroesIn64(0)
	if z != 64 {
		t.Errorf("bad zeroes in uint64 for 0, got %d, expected 64", z)
	} else {
		t.Logf("zeroes in uint64 for 0 is 64")
	}
}

func TestZeroesIn32(t *testing.T) {
	i := uint32(init_i)

	z := ZeroesIn32(i)
	if z != expected_z {
		t.Errorf("bad zeroes in uint32 for %d, got %d, expected %d", init_i, z, expected_z)
	} else {
		t.Logf("zeroes in uint32 for %d is %d", init_i, expected_z)
	}
	z = ZeroesIn32(0)
	if z != 32 {
		t.Errorf("bad zeroes in uint32 for 0, got %d, expected 32", z)
	} else {
		t.Logf("zeroes in uint32 for 0 is 64")
	}
}
