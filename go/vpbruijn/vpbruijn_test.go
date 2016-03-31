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
	"encoding/hex"
	"math/rand"
	"testing"
)

var mList = []int{2, 16, 31}
var nList = []int{3, 64, 113}

func TestBruijnSize(t *testing.T) {
	for _, m := range mList {
		for _, n := range nList {
			t.Logf("testing Bruijn graph walker for m=%d n=%d", m, n)
			walker, err := BruijnNew(m, n)
			if err != nil {
				t.Error("unable to create walker", err)
			}
			if walker.M() != m {
				t.Errorf("M()!=m M()=%d m=%d", walker.M(), m)
			}
			if walker.N() != n {
				t.Errorf("M()!=m M()=%d m=%d", walker.N(), n)
			}
			t.Logf("NbBits() -> %d", walker.NbBits())
			t.Logf("NbBytes() -> %d", walker.NbBytes())
		}
	}
}

func TestBruijnForward(t *testing.T) {
	for _, m := range mList {
		for _, n := range nList {
			t.Logf("testing Bruijn graph walker for m=%d n=%d", m, n)
			walker, err := BruijnNew(m, n)
			if err != nil {
				t.Error("unable to create walker", err)
			}
			zero := make([]byte, walker.NbBytes())
			max := walker.Decr(zero)
			forwardPath := walker.ForwardPath(zero, max)
			lastV := make([]byte, walker.NbBytes())
			for i, v := range forwardPath {
				if i > 0 {
					possibleV := walker.NextList(lastV)
					found := 0
					for _, w := range possibleV {
						if walker.Cmp(v, w) == 0 {
							t.Logf("found key %s as being next to %s", hex.EncodeToString(w), hex.EncodeToString(lastV))
							found++
						}
					}
					if found == 0 {
						t.Errorf("no matches for v=%s", hex.EncodeToString(v))
					}
					if found > 1 {
						t.Errorf("too many (%d) matches for v=%s", found, hex.EncodeToString(v))
					}
				}
				copy(lastV, v)
			}
		}
	}
}

func TestBruijnBackward(t *testing.T) {
	for _, m := range mList {
		for _, n := range nList {
			t.Logf("testing Bruijn graph walker for m=%d n=%d", m, n)
			walker, err := BruijnNew(m, n)
			if err != nil {
				t.Error("unable to create walker", err)
			}
			zero := make([]byte, walker.NbBytes())
			max := walker.Decr(zero)
			backwardPath := walker.BackwardPath(zero, max)
			lastV := make([]byte, walker.NbBytes())
			for i, v := range backwardPath {
				if i > 0 {
					possibleV := walker.PrevList(lastV)
					found := 0
					for _, w := range possibleV {
						if walker.Cmp(v, w) == 0 {
							t.Logf("found key %s as being prev to %s", hex.EncodeToString(w), hex.EncodeToString(lastV))
							found++
						}
					}
					if found == 0 {
						t.Errorf("no matches for v=%s", hex.EncodeToString(v))
					}
					if found > 1 {
						t.Errorf("too many (%d) matches for v=%s", found, hex.EncodeToString(v))
					}
				}
				copy(lastV, v)
			}
		}
	}
}

func TestBruijnMath(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	for _, m := range mList {
		for _, n := range nList {
			walker, err := BruijnNew(m, n)
			if err != nil {
				t.Error("unable to create walker", err)
			}
			randKey := walker.Rand(r)
			t.Logf("random key for m,n=%d,%d: %s", m, n, hex.EncodeToString(randKey))
			zeroKey := walker.Zero()
			t.Logf("zero key for m,n=%d,%d: %s", m, n, hex.EncodeToString(zeroKey))
			oneKey := walker.Incr(zeroKey)
			t.Logf("one key for m,n=%d,%d: %s", m, n, hex.EncodeToString(oneKey))
			zeroCheckKey := walker.Decr(oneKey)
			if walker.Cmp(zeroKey, zeroCheckKey) != 0 {
				t.Errorf("zero (check) key for m,n=%d,%d: %s but it should be zero", m, n, hex.EncodeToString(zeroCheckKey))
			}
			maxKey := walker.Decr(zeroKey)
			t.Logf("max key for m,n=%d,%d: %s", m, n, hex.EncodeToString(maxKey))
		}
	}
}
