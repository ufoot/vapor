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

func TestRand512(t *testing.T) {
	const l = 1000000
	r := NewRand()
	n := big.NewInt(l)

	i := Rand512(r, n).Int64()
	if i >= 0 && i < l {
		t.Logf("rand512 n is %d", i)
	} else {
		t.Errorf("rand512 out of range, is %d, should be between 0 and %d", i, l)
	}
	i = Rand512(r, nil).Int64()
	t.Logf("rand512.Int64 n is %016x", i)
}

func TestRand256(t *testing.T) {
	const l = 1000000
	r := NewRand()
	n := big.NewInt(l)

	i := Rand256(r, n).Int64()
	if i >= 0 && i < l {
		t.Logf("rand256 n is %d", i)
	} else {
		t.Errorf("rand256 out of range, is %d, should be between 0 and %d", i, l)
	}
	i = Rand256(r, nil).Int64()
	t.Logf("rand256.Int64 n is %016x", i)
}

func TestRand128(t *testing.T) {
	const l = 1000000
	r := NewRand()
	n := big.NewInt(l)

	i := Rand128(r, n).Int64()
	if i >= 0 && i < l {
		t.Logf("rand128 n is %d", i)
	} else {
		t.Errorf("rand128 out of range, is %d, should be between 0 and %d", i, l)
	}
	i = Rand128(r, nil).Int64()
	t.Logf("rand128.Int64 n is %016x", i)
}

func TestRand64(t *testing.T) {
	const l = 1000000
	r := NewRand()

	i := int64(Rand64(r, l))
	if i >= 0 && i < l {
		t.Logf("rand64 n is %d", i)
	} else {
		t.Errorf("rand64 out of range, is %d, should be between 0 and %d", i, l)
	}
}

func TestRand32(t *testing.T) {
	const l = 1000000
	r := NewRand()

	i := int32(Rand32(r, l))
	if i >= 0 && i < l {
		t.Logf("rand32 n is %d", i)
	} else {
		t.Errorf("rand32 out of range, is %d, should be between 0 and %d", i, l)
	}
}

func BenchmarkNewRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewRand()
	}
}

func BenchmarkRand512(b *testing.B) {
	r := NewRand()
	n := big.NewInt(10000)

	for i := 0; i < b.N; i++ {
		_ = Rand512(r, n)
	}
}

func BenchmarkRand256(b *testing.B) {
	r := NewRand()
	n := big.NewInt(10000)

	for i := 0; i < b.N; i++ {
		_ = Rand256(r, n)
	}
}

func BenchmarkRand128(b *testing.B) {
	r := NewRand()
	n := big.NewInt(10000)

	for i := 0; i < b.N; i++ {
		_ = Rand128(r, n)
	}
}

func BenchmarkRand64(b *testing.B) {
	r := NewRand()
	for i := 0; i < b.N; i++ {
		_ = Rand64(r, 1000)
	}
}

func BenchmarkRand32(b *testing.B) {
	r := NewRand()
	for i := 0; i < b.N; i++ {
		_ = Rand32(r, 1000)
	}
}
