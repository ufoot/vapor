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

package vpmath

import (
	"testing"
)

func TestLookup(t *testing.T) {
	const s0 = 5
	const s1 = 1
	const s2 = 3
	const s3 = 7

	sizes := make([]int, 4)
	sizes[0] = s0
	sizes[1] = s1
	sizes[2] = s2
	sizes[3] = s3

	size := LookupSize(sizes)
	if size == s0*s1*s2*s3 {
		t.Logf("size of lookup table is %d", size)
	} else {
		t.Errorf("bad size for lookup table %d", size)
	}

	for id := 0; id < size; id++ {
		indexes := LookupSplit(sizes, id)
		t.Logf("indexes for id %d are [%d][%d][%d][%d]", id, indexes[0], indexes[1], indexes[2], indexes[3])
		id2 := LookupJoin(sizes, indexes)
		if id != id2 {
			t.Errorf("ID mismatch id=%d id2=%d", id, id2)
		}
	}
}
