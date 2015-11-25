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

func TestBruijnBigNext(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {0, 1}, {2, 3}, {4, 5}, {6, 7}}

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		nf, err := BruijnBigNextFirst(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnBigNextFirst:", err)
		}
		if nf.Int64() != int64(v[0]) {
			t.Errorf("bad nf for i=%d, got %s", i, nf)
		}
		nl, err := BruijnBigNextLast(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnBigNextLast:", err)
		}
		if nl.Int64() != int64(v[1]) {
			t.Errorf("bad nl for i=%d, got %s", i, nl)
		}
		nList, err := BruijnBigNextList(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnBigNextList:", err)
		}
		for j, w := range nList {
			if w.Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnBigPrev(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 4}, {0, 4}, {1, 5}, {1, 5}, {2, 6}, {2, 6}, {3, 7}, {3, 7}}

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		pf, err := BruijnBigPrevFirst(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnBigPrevFirst:", err)
		}
		if pf.Int64() != int64(v[0]) {
			t.Errorf("bad pf for i=%d, got %s", i, pf)
		}
		pl, err := BruijnBigPrevLast(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnBigPrevLast:", err)
		}
		if pl.Int64() != int64(v[1]) {
			t.Errorf("bad pl for i=%d, got %s", i, pl)
		}
		pList, err := BruijnBigPrevList(m, n, bi)
		if err != nil {
			t.Error("unable to call BruijnBigPrevList:", err)
		}
		for j, w := range pList {
			if w.Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnBigForwardPath(t *testing.T) {
	const m = 10
	const n = 6
	const from = 234567
	const to = 987654

	path, err := BruijnBigForwardPath(m, n, big.NewInt(from), big.NewInt(to))
	if err != nil {
		t.Error("unable to call BruijnBigForwardPath:", err)
	}
	for i, v := range path {
		t.Logf("path[%d]=%d", i, v.Int64())
		if i > 0 {
			found := false
			nList, err := BruijnBigNextList(m, n, path[i-1])
			if err != nil {
				t.Error("unable to call BruijnBigNextList:", err)
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
		w, err := BruijnBigForwardElem(m, n, big.NewInt(from), big.NewInt(to), i)
		if err != nil {
			t.Error(vpsys.ErrorChain(err, "unable to call BruijnBigForwardElem"))
		}
		if v.Cmp(w) != 0 {
			t.Errorf("values for BruijnBigForwardPath=%s and BruijnBigForwardElem=%s differ, i=%d", v.String(), w.String(), i)
		}
	}
}

func TestBruijnBigBackwardPath(t *testing.T) {
	const m = 10
	const n = 4
	const from = 1234
	const to = 9876

	path, err := BruijnBigBackwardPath(m, n, big.NewInt(from), big.NewInt(to))
	if err != nil {
		t.Error("unable to call BruijnBigBackwardPath:", err)
	}
	for i, v := range path {
		t.Logf("path[%d]=%d", i, v.Int64())
		if i > 0 {
			found := false
			nList, err := BruijnBigPrevList(m, n, path[i-1])
			if err != nil {
				t.Error("unable to call BruijnBigPrevList:", err)
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
		w, err := BruijnBigBackwardElem(m, n, big.NewInt(from), big.NewInt(to), i)
		if err != nil {
			t.Error(vpsys.ErrorChain(err, "unable to call BruijnBigBackwardElem"))
		}
		if v.Cmp(w) != 0 {
			t.Errorf("values for BruijnBigBackwardPath=%s and BruijnBigBackwardElem=%s differ, i=%d", v.String(), w.String(), i)
		}
	}
}

func BenchmarkBruijnBigNext_2_32(b *testing.B) {
	v := big.NewInt(1000)
	for i := 0; i < b.N; i++ {
		BruijnBigNextList(2, 32, v)
	}
}

func BenchmarkBruijnBigPrev_2_32(b *testing.B) {
	v := big.NewInt(1000)
	for i := 0; i < b.N; i++ {
		BruijnBigPrevList(2, 32, v)
	}
}

func BenchmarkBruijnBigNext_16_64(b *testing.B) {
	v := big.NewInt(1000)
	for i := 0; i < b.N; i++ {
		BruijnBigNextList(16, 64, v)
	}
}

func BenchmarkBruijnBigPrev_16_64(b *testing.B) {
	v := big.NewInt(1000)
	for i := 0; i < b.N; i++ {
		BruijnBigPrevList(16, 64, v)
	}
}

func BenchmarkBruijnBigNext_7_100(b *testing.B) {
	v := big.NewInt(1000)
	for i := 0; i < b.N; i++ {
		BruijnBigNextList(7, 100, v)
	}
}

func BenchmarkBruijnBigPrev_7_100(b *testing.B) {
	v := big.NewInt(1000)
	for i := 0; i < b.N; i++ {
		BruijnBigPrevList(7, 100, v)
	}
}
