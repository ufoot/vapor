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
	"testing"
)

func TestBruijnNext(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {0, 1}, {2, 3}, {4, 5}, {6, 7}}

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		n0, err := BruijnNext0(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnNext0:", err)
		}
		if n0.Int64() != int64(v[0]) {
			t.Errorf("bad n0 for i=%d, got %s", i, n0)
		}
		nList, err := BruijnNextList(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnNextList:", err)
		}
		for j, w := range nList {
			if w.Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnPrev(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 4}, {0, 4}, {1, 5}, {1, 5}, {2, 6}, {2, 6}, {3, 7}, {3, 7}}

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		n0, err := BruijnPrev0(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnPrev0:", err)
		}
		if n0.Int64() != int64(v[0]) {
			t.Errorf("bad n0 for i=%d, got %s", i, n0)
		}
		nList, err := BruijnPrevList(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnPrevList:", err)
		}
		for j, w := range nList {
			if w.Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func BenchmarkNext_2_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruijnNextList(2, 32, big.NewInt(int64(i)))
	}
}

func BenchmarkPrev_2_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruijnPrevList(2, 32, big.NewInt(int64(i)))
	}
}

func BenchmarkNext_4_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruijnNextList(4, 512, big.NewInt(int64(i)))
	}
}

func BenchmarkPrev_4_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruijnPrevList(4, 512, big.NewInt(int64(i)))
	}
}

func BenchmarkNext_7_100(b *testing.B) {
	// 7**100 -> approx 280 bits
	for i := 0; i < b.N; i++ {
		BruijnNextList(7, 100, big.NewInt(int64(i)))
	}
}

func BenchmarkPrev_7_100(b *testing.B) {
	// 7**100 -> approx 280 bits
	for i := 0; i < b.N; i++ {
		BruijnPrevList(7, 100, big.NewInt(int64(i)))
	}
}
