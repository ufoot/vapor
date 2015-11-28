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

func TestBruijnGenericNext(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {0, 1}, {2, 3}, {4, 5}, {6, 7}}
	impl := bruijnGenericNew(m, n)

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		nf, err := impl.NextFirst(bigToBytes(bi, impl.NbBytes()))
		if err != nil {
			t.Error("unable to call NextFirst:", err)
		}
		if bytesToBig(nf).Int64() != int64(v[0]) {
			t.Errorf("bad nf for i=%d, got %s", i, nf)
		}
		nl, err := impl.NextLast(bigToBytes(bi, impl.NbBytes()))
		if err != nil {
			t.Error("unable to call NextLast:", err)
		}
		if bytesToBig(nl).Int64() != int64(v[1]) {
			t.Errorf("bad nl for i=%d, got %s", i, nl)
		}
		nList, err := impl.NextList(bigToBytes(bi, impl.NbBytes()))
		if err != nil {
			t.Error("unable to call NextList:", err)
		}
		for j, w := range nList {
			if bytesToBig(w).Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnGenericPrev(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 4}, {0, 4}, {1, 5}, {1, 5}, {2, 6}, {2, 6}, {3, 7}, {3, 7}}
	impl := bruijnGenericNew(m, n)

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		pf, err := impl.PrevFirst(bigToBytes(bi, impl.NbBytes()))
		if err != nil {
			t.Error("unable to call PrevFirst:", err)
		}
		if bytesToBig(pf).Int64() != int64(v[0]) {
			t.Errorf("bad pf for i=%d, got %s", i, pf)
		}
		pl, err := impl.PrevLast(bigToBytes(bi, impl.NbBytes()))
		if err != nil {
			t.Error("unable to call PrevLast:", err)
		}
		if bytesToBig(pl).Int64() != int64(v[1]) {
			t.Errorf("bad pl for i=%d, got %s", i, pl)
		}
		pList, err := impl.PrevList(bigToBytes(bi, impl.NbBytes()))
		if err != nil {
			t.Error("unable to call PrevList:", err)
		}
		for j, w := range pList {
			if bytesToBig(w).Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnGenericForwardPath(t *testing.T) {
	const m = 10
	const n = 6
	const from = 234567
	const to = 987654
	impl := bruijnGenericNew(m, n)

	path, err := impl.ForwardPath(bigToBytes(big.NewInt(from), impl.NbBytes()), bigToBytes(big.NewInt(to), impl.NbBytes()))
	if err != nil {
		t.Error("unable to call ForwardPath:", err)
	}
	for i, v := range path {
		t.Logf("path[%d]=%d", i, bytesToBig(v).Int64())
		if i > 0 {
			found := false
			nList, err := impl.NextList(path[i-1])
			if err != nil {
				t.Error("unable to call NextList:", err)
			}
			for j, w := range nList {
				if bytesToBig(v).Cmp(bytesToBig(w)) == 0 {
					t.Logf("path[%d]: successor %d of %d found, was in position %d", i, bytesToBig(w).Int64(), bytesToBig(path[i-1]).Int64(), j)
					found = true
				} else {
					t.Logf("path[%d]: successor %d of %d in position %d, not what we search", i, bytesToBig(w).Int64(), bytesToBig(path[i-1]).Int64(), j)
				}
			}
			if !found {
				t.Errorf("v[%d]=%d not found in successors of v[%d]", i, bytesToBig(v).Int64(), i-1)
			}
		}
		w, err := impl.ForwardElem(bigToBytes(big.NewInt(from), impl.NbBytes()), bigToBytes(big.NewInt(to), impl.NbBytes()), i)
		if err != nil {
			t.Error(vpsys.ErrorChain(err, "unable to call ForwardElem"))
		}
		if bytesToBig(v).Cmp(bytesToBig(w)) != 0 {
			t.Errorf("values for ForwardPath=%d and ForwardElem=%d differ, i=%d", bytesToBig(v).Int64(), bytesToBig(w).Int64(), i)
		}
	}
}

func TestBruijnGenericBackwardPath(t *testing.T) {
	const m = 10
	const n = 4
	const from = 1234
	const to = 9876
	impl := bruijnGenericNew(m, n)

	path, err := impl.BackwardPath(bigToBytes(big.NewInt(from), impl.NbBytes()), bigToBytes(big.NewInt(to), impl.NbBytes()))
	if err != nil {
		t.Error("unable to call BackwardPath:", err)
	}
	for i, v := range path {
		t.Logf("path[%d]=%d", i, bytesToBig(v).Int64())
		if i > 0 {
			found := false
			nList, err := impl.PrevList(path[i-1])
			if err != nil {
				t.Error("unable to call PrevList:", err)
			}
			for j, w := range nList {
				if bytesToBig(v).Cmp(bytesToBig(w)) == 0 {
					t.Logf("path[%d]: predecessor %d of %d found, was in position %d", i, bytesToBig(w).Int64(), bytesToBig(path[i-1]).Int64(), j)
					found = true
				} else {
					t.Logf("path[%d]: predecessor %d of %d in position %d, not what we search", i, bytesToBig(w).Int64(), bytesToBig(path[i-1]).Int64(), j)
				}
			}
			if !found {
				t.Errorf("v[%d]=%d not found in predecessors of v[%d]", i, bytesToBig(v).Int64(), i-1)
			}
		}
		w, err := impl.BackwardElem(bigToBytes(big.NewInt(from), impl.NbBytes()), bigToBytes(big.NewInt(to), impl.NbBytes()), i)
		if err != nil {
			t.Error(vpsys.ErrorChain(err, "unable to call BackwardElem"))
		}
		if bytesToBig(v).Cmp(bytesToBig(w)) != 0 {
			t.Errorf("values for BackwardPath=%d and BackwardElem=%d differ, i=%d", bytesToBig(v).Int64(), bytesToBig(w).Int64(), i)
		}
	}
}

func BrenchmarkBruijnGenericNext_2_32(b *testing.B) {
	impl := bruijnGenericNew(2, 32)
	v := bigToBytes(big.NewInt(1000), impl.NbBytes())
	for i := 0; i < b.N; i++ {
		impl.NextFirst(v)
	}
}

func BrenchmarkBruijnGenericPrev_2_32(b *testing.B) {
	impl := bruijnGenericNew(2, 32)
	v := bigToBytes(big.NewInt(1000), impl.NbBytes())
	for i := 0; i < b.N; i++ {
		impl.PrevFirst(v)
	}
}

func BrenchmarkBruijnGenericNext_16_64(b *testing.B) {
	impl := bruijnGenericNew(16, 64)
	v := bigToBytes(big.NewInt(1000), impl.NbBytes())
	for i := 0; i < b.N; i++ {
		impl.NextFirst(v)
	}
}

func BrenchmarkBruijnGenericPrev_16_64(b *testing.B) {
	impl := bruijnGenericNew(16, 64)
	v := bigToBytes(big.NewInt(1000), impl.NbBytes())
	for i := 0; i < b.N; i++ {
		impl.PrevFirst(v)
	}
}

func BrenchmarkBruijnGenericNext_7_100(b *testing.B) {
	impl := bruijnGenericNew(7, 100)
	v := bigToBytes(big.NewInt(1000), impl.NbBytes())
	for i := 0; i < b.N; i++ {
		impl.NextFirst(v)
	}
}

func BrenchmarkBruijnGenericPrev_7_100(b *testing.B) {
	impl := bruijnGenericNew(7, 100)
	v := bigToBytes(big.NewInt(1000), impl.NbBytes())
	for i := 0; i < b.N; i++ {
		impl.PrevFirst(v)
	}
}
