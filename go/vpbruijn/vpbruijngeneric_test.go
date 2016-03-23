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

package vpbruijn

import (
	"bytes"
	"math/big"
	"testing"
)

func TestBruijnGenericNext(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {0, 1}, {2, 3}, {4, 5}, {6, 7}}
	implGeneric := bruijnGenericNew(m, n)

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		nf := implGeneric.NextFirst(bigToBytes(bi, implGeneric.NbBytes()))
		if bytesToBig(nf, implGeneric.bigMax).Int64() != int64(v[0]) {
			t.Errorf("bad nf for i=%d, got %s", i, nf)
		}
		nl := implGeneric.NextLast(bigToBytes(bi, implGeneric.NbBytes()))
		if bytesToBig(nl, implGeneric.bigMax).Int64() != int64(v[1]) {
			t.Errorf("bad nl for i=%d, got %s", i, nl)
		}
		nList := implGeneric.NextList(bigToBytes(bi, implGeneric.NbBytes()))
		for j, w := range nList {
			if bytesToBig(w, implGeneric.bigMax).Int64() != int64(v[j]) {
				t.Errorf("bad value for i,j=%d,%d got %s", i, j, w)
			}
		}
	}
}

func TestBruijnGenericPrev(t *testing.T) {
	const m = 2
	const n = 3
	bruijn23 := [m * m * m][n]int{{0, 4}, {0, 4}, {1, 5}, {1, 5}, {2, 6}, {2, 6}, {3, 7}, {3, 7}}
	implGeneric := bruijnGenericNew(m, n)

	for i, v := range bruijn23 {
		bi := big.NewInt(int64(i))
		pf := implGeneric.PrevFirst(bigToBytes(bi, implGeneric.NbBytes()))
		if bytesToBig(pf, implGeneric.bigMax).Int64() != int64(v[0]) {
			t.Errorf("bad pf for i=%d, got %s", i, pf)
		}
		pl := implGeneric.PrevLast(bigToBytes(bi, implGeneric.NbBytes()))
		if bytesToBig(pl, implGeneric.bigMax).Int64() != int64(v[1]) {
			t.Errorf("bad pl for i=%d, got %s", i, pl)
		}
		pList := implGeneric.PrevList(bigToBytes(bi, implGeneric.NbBytes()))
		for j, w := range pList {
			if bytesToBig(w, implGeneric.bigMax).Int64() != int64(v[j]) {
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
	implGeneric := bruijnGenericNew(m, n)

	path := implGeneric.ForwardPath(bigToBytes(big.NewInt(from), implGeneric.NbBytes()), bigToBytes(big.NewInt(to), implGeneric.NbBytes()))
	for i, v := range path {
		t.Logf("path[%d]=%d", i, bytesToBig(v, implGeneric.bigMax).Int64())
		if i > 0 {
			found := 0
			nList := implGeneric.NextList(path[i-1])
			for j, w := range nList {
				if bytesToBig(v, implGeneric.bigMax).Cmp(bytesToBig(w, implGeneric.bigMax)) == 0 {
					t.Logf("path[%d]: successor %d of %d found, was in position %d", i, bytesToBig(w, implGeneric.bigMax).Int64(), bytesToBig(path[i-1], implGeneric.bigMax).Int64(), j)
					found++
				} else {
					t.Logf("path[%d]: successor %d of %d in position %d, not what we search", i, bytesToBig(w, implGeneric.bigMax).Int64(), bytesToBig(path[i-1], implGeneric.bigMax).Int64(), j)
				}
			}
			if found <= 0 {
				t.Errorf("v[%d]=%d not found in successors of v[%d]", i, bytesToBig(v, implGeneric.bigMax).Int64(), i-1)
			}
			if found > 1 {
				t.Errorf("v[%d]=%d found too many times (%d) in successors of v[%d]", i, bytesToBig(v, implGeneric.bigMax).Int64(), found, i-1)
			}
		}
		w := implGeneric.ForwardElem(bigToBytes(big.NewInt(from), implGeneric.NbBytes()), bigToBytes(big.NewInt(to), implGeneric.NbBytes()), i)
		if bytesToBig(v, implGeneric.bigMax).Cmp(bytesToBig(w, implGeneric.bigMax)) != 0 {
			t.Errorf("values for ForwardPath=%d and ForwardElem=%d differ, i=%d", bytesToBig(v, implGeneric.bigMax).Int64(), bytesToBig(w, implGeneric.bigMax).Int64(), i)
		}
	}
}

func TestBruijnGenericBackwardPath(t *testing.T) {
	const m = 10
	const n = 4
	const from = 1234
	const to = 9876
	implGeneric := bruijnGenericNew(m, n)

	path := implGeneric.BackwardPath(bigToBytes(big.NewInt(from), implGeneric.NbBytes()), bigToBytes(big.NewInt(to), implGeneric.NbBytes()))
	for i, v := range path {
		t.Logf("path[%d]=%d", i, bytesToBig(v, implGeneric.bigMax).Int64())
		if i > 0 {
			found := 0
			nList := implGeneric.PrevList(path[i-1])
			for j, w := range nList {
				if bytesToBig(v, implGeneric.bigMax).Cmp(bytesToBig(w, implGeneric.bigMax)) == 0 {
					t.Logf("path[%d]: predecessor %d of %d found, was in position %d", i, bytesToBig(w, implGeneric.bigMax).Int64(), bytesToBig(path[i-1], implGeneric.bigMax).Int64(), j)
					found++
				} else {
					t.Logf("path[%d]: predecessor %d of %d in position %d, not what we search", i, bytesToBig(w, implGeneric.bigMax).Int64(), bytesToBig(path[i-1], implGeneric.bigMax).Int64(), j)
				}
			}
			if found <= 0 {
				t.Errorf("v[%d]=%d not found in predecessors of v[%d]", i, bytesToBig(v, implGeneric.bigMax).Int64(), i-1)
			}
			if found > 1 {
				t.Errorf("v[%d]=%d found too many times (%d) in predecessors of v[%d]", i, bytesToBig(v, implGeneric.bigMax).Int64(), found, i-1)
			}
		}
		w := implGeneric.BackwardElem(bigToBytes(big.NewInt(from), implGeneric.NbBytes()), bigToBytes(big.NewInt(to), implGeneric.NbBytes()), i)
		if bytesToBig(v, implGeneric.bigMax).Cmp(bytesToBig(w, implGeneric.bigMax)) != 0 {
			t.Errorf("values for BackwardPath=%d and BackwardElem=%d differ, i=%d", bytesToBig(v, implGeneric.bigMax).Int64(), bytesToBig(w, implGeneric.bigMax).Int64(), i)
		}
	}
}

func TestBruijnGenericIntArray(t *testing.T) {
	const m = 10
	const n = 9
	const x = 123654789

	implGeneric := bruijnGenericNew(m, n)
	bytesX := bigToBytes(big.NewInt(x), implGeneric.NbBytes())

	intArrayX := implGeneric.BytesToIntArray(bytesX)
	t.Log("int array for", x, "is", intArrayX)
	checkBytesX := implGeneric.IntArrayToBytes(intArrayX)
	intX := int(bytesToBig(bytesX, implGeneric.bigMax).Int64())
	checkIntX := int(bytesToBig(checkBytesX, implGeneric.bigMax).Int64())

	if !bytes.Equal(bytesX, checkBytesX) {
		t.Error("conversion of bytes to int array to bytes failed", intX, checkIntX, bytesX, checkBytesX)
	}
}

func BenchmarkBruijnGenericNext2_32(b *testing.B) {
	implGeneric := bruijnGenericNew(2, 32)
	v := bigToBytes(big.NewInt(1000), implGeneric.NbBytes())
	for i := 0; i < b.N; i++ {
		implGeneric.NextFirst(v)
	}
}

func BenchmarkBruijnGenericPrev2_32(b *testing.B) {
	implGeneric := bruijnGenericNew(2, 32)
	v := bigToBytes(big.NewInt(1000), implGeneric.NbBytes())
	for i := 0; i < b.N; i++ {
		implGeneric.PrevFirst(v)
	}
}

func BenchmarkBruijnGenericNext16_64(b *testing.B) {
	implGeneric := bruijnGenericNew(16, 64)
	v := bigToBytes(big.NewInt(1000), implGeneric.NbBytes())
	for i := 0; i < b.N; i++ {
		implGeneric.NextFirst(v)
	}
}

func BenchmarkBruijnGenericPrev16_64(b *testing.B) {
	implGeneric := bruijnGenericNew(16, 64)
	v := bigToBytes(big.NewInt(1000), implGeneric.NbBytes())
	for i := 0; i < b.N; i++ {
		implGeneric.PrevFirst(v)
	}
}

func BenchmarkBruijnGenericNext7_100(b *testing.B) {
	implGeneric := bruijnGenericNew(7, 100)
	v := bigToBytes(big.NewInt(1000), implGeneric.NbBytes())
	for i := 0; i < b.N; i++ {
		implGeneric.NextFirst(v)
	}
}

func BenchmarkBruijnGenericPrev7_100(b *testing.B) {
	implGeneric := bruijnGenericNew(7, 100)
	v := bigToBytes(big.NewInt(1000), implGeneric.NbBytes())
	for i := 0; i < b.N; i++ {
		implGeneric.PrevFirst(v)
	}
}
