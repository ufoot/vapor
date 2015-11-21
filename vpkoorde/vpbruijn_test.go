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
	"github.com/ufoot/vapor/vpsys"
	"math/big"
	"testing"
)

func TestBruijnNext(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {0, 1}, {2, 3}, {4, 5}, {6, 7}}

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		nf, err := BruijnNextFirst(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnNextFirst:", err)
		}
		if nf.Int64() != int64(v[0]) {
			t.Errorf("bad nf for i=%d, got %s", i, nf)
		}
		nl, err := BruijnNextLast(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnNextLast:", err)
		}
		if nl.Int64() != int64(v[1]) {
			t.Errorf("bad nl for i=%d, got %s", i, nl)
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
		pf, err := BruijnPrevFirst(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnPrevFirst:", err)
		}
		if pf.Int64() != int64(v[0]) {
			t.Errorf("bad pf for i=%d, got %s", i, pf)
		}
		pl, err := BruijnPrevLast(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnPrevLast:", err)
		}
		if pl.Int64() != int64(v[1]) {
			t.Errorf("bad pl for i=%d, got %s", i, pl)
		}
		pList, err := BruijnPrevList(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnPrevList:", err)
		}
		for j, w := range pList {
			if w.Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnForwardPath(t *testing.T) {
	const m = 10
	const n = 6
	const from = 234567
	const to = 987654

	path, err := BruijnForwardPath(m, n, big.NewInt(from), big.NewInt(to))
	if err != nil {
		t.Error("unable to call BruijnForwardPath:", err)
	}
	for i, v := range path {
		t.Logf("path[%d]=%d", i, v.Int64())
		if i > 0 {
			found := false
			nList, err := BruijnNextList(m, n, path[i-1])
			if err != nil {
				t.Error("unable to call BruijnNextList:", err)
			}
			for j, w := range nList {
				if v.Cmp(w) == 0 {
					t.Logf("path[%d]: successor %d of %d found, was in position %d", i, w.Int64(), path[i-1].Int64(), j)
					found = true
				} else {
					t.Logf("path[%d]: successor %d of %d in position %d, not what we search", i, w.Int64(), path[i-1].Int64(), j)
				}
			}
			if !found {
				t.Errorf("v[%d]=%d not found in successors of v[%d]", i, v.Int64(), i-1)
			}
		}
		w, err := BruijnForwardElem(m, n, big.NewInt(from), big.NewInt(to), i)
		if err != nil {
			t.Error(vpsys.ErrorChain(err, "unable to call BruijnForwardElem"))
		}
		if v.Cmp(w) != 0 {
			t.Errorf("values for BruijnForwardPath=%s and BruijnForwardElem=%s differ, i=%d", v.String(), w.String(), i)
		}
	}
}

func TestBruijnBackwardPath(t *testing.T) {
	const m = 10
	const n = 4
	const from = 1234
	const to = 9876

	path, err := BruijnBackwardPath(m, n, big.NewInt(from), big.NewInt(to))
	if err != nil {
		t.Error("unable to call BruijnBackwardPath:", err)
	}
	for i, v := range path {
		t.Logf("path[%d]=%d", i, v.Int64())
		if i > 0 {
			found := false
			nList, err := BruijnPrevList(m, n, path[i-1])
			if err != nil {
				t.Error("unable to call BruijnPrevList:", err)
			}
			for j, w := range nList {
				if v.Cmp(w) == 0 {
					t.Logf("path[%d]: predecessor %d of %d found, was in position %d", i, w.Int64(), path[i-1].Int64(), j)
					found = true
				} else {
					t.Logf("path[%d]: predecessor %d of %d in position %d, not what we search", i, w.Int64(), path[i-1].Int64(), j)
				}
			}
			if !found {
				t.Errorf("v[%d]=%d not found in predecessors of v[%d]", i, v.Int64(), i-1)
			}
		}
		w, err := BruijnBackwardElem(m, n, big.NewInt(from), big.NewInt(to), i)
		if err != nil {
			t.Error(vpsys.ErrorChain(err, "unable to call BruijnBackwardElem"))
		}
		if v.Cmp(w) != 0 {
			t.Errorf("values for BruijnBackwardPath=%s and BruijnBackwardElem=%s differ, i=%d", v.String(), w.String(), i)
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
