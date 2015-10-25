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

func TestCompatible(t *testing.T) {
	p1 := NewPackage("foo", "Foo", "foo@bar.com", "http://bar.com")
	p2 := NewPackage("foo", "Foo 2", "foo2@bar.com", "http://bar.com")
	p3 := NewPackage("bar", "Bar", "bar@foo.com", "http://foo.com")

	t.Logf("p1=%s", p1.String())
	t.Logf("p2=%s", p2.String())
	t.Logf("p3=%s", p3.String())

	if !Compatible(p1, p2) {
		t.Error("p1==p2 does not work")
	}
	if Compatible(p1, p3) {
		t.Error("p1!=p3 does not work")
	}
}
