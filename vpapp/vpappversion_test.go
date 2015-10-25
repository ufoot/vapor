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

package vpapp

import (
	"testing"
)

func TestCompare(t *testing.T) {
	v1 := NewVersion(1, 2, "toto")
	v2 := NewVersion(2, 1, "titi")
	v3 := NewVersion(2, 3, "tata")

	t.Logf("v1=%s", v1.String())
	t.Logf("v2=%s", v2.String())
	t.Logf("v3=%s", v3.String())

	if Compare(v1, v1) != 0 {
		t.Error("v1==v1 does not work")
	}
	if Compare(v1, v2) != -1 {
		t.Error("v1<v2 does not work")
	}
	if Compare(v2, v1) != 1 {
		t.Error("v2>v1 does not work")
	}
	if Compare(v2, v3) != -1 {
		t.Error("v2<v3 does not work")
	}
	if Compare(v3, v2) != 1 {
		t.Error("v3>v2 does not work")
	}
}

func TestEqual(t *testing.T) {
	v1 := NewVersion(1, 2, "foo")
	v2 := NewVersion(2, 1, "bar")

	if !Equal(v1, v1) {
		t.Error("v1==v1 does not work")
	}
	if Equal(v1, v2) {
		t.Error("v1!=v2 does not work")
	}
}
