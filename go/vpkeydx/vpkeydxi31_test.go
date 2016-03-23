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

package vpkeydx

import (
	"testing"
)

func TestMod31(t *testing.T) {
	minus1 := Mod31(-1)
	if minus1 != 0x7fffffff {
		t.Errorf("-1 not treated correctly %d", minus1)
	}
	zero := Mod31(0)
	if zero != 0 {
		t.Errorf("zero not treated correctly %d", zero)
	}
	ten := Mod31(10)
	if ten != 10 {
		t.Errorf("10 not treated correctly %d", ten)
	}
}

func TestScale31(t *testing.T) {
	n := int32(100)
	input := []int32{0, 10, 33, 50, 99}
	output := []int32{0, 214748364, 708669603, 1073741824, 2126008811}

	for i, v := range input {
		w1 := output[i]
		w2 := Scale31(v, n)
		if w1 != w2 {
			t.Errorf("Scale31 returned wrong value, expecting %d got %d", w1, w2)
		}
	}
}

func TestInc31(t *testing.T) {
	a := Inc31(0)
	if a != 1 {
		t.Errorf("Inc31 does not work on 0, got %d", a)
	}
	b := Inc31(mask31)
	if b != 0 {
		t.Errorf("Inc31 does not wraparround, got %d", b)
	}
}

func TestDec31(t *testing.T) {
	a := Dec31(1)
	if a != 0 {
		t.Errorf("Dec31 does not work on 1, got %d", a)
	}
	b := Dec31(0)
	if b != mask31 {
		t.Errorf("Dec31 does not wraparround, got %d", b)
	}
}
