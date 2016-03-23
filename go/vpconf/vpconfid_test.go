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

package vpconf

import (
	"testing"
)

func TestConfIdSplit(t *testing.T) {
	var split []string
	checkFooBarStrings := []string{"foo.bar", "FOO.BAR", "..Foo..Bar.."}
	var id string
	var l int

	for _, id = range checkFooBarStrings {
		split = ConfIDSplit(id)
		t.Log("split", id, "->", split)
		l = len(split)
		if l != 2 {
			t.Error("bad length", l)
		} else {
			if split[0] != "foo" {
				t.Error("bad foo", split[0])
			}
			if split[1] != "bar" {
				t.Error("bad bar", split[1])
			}
		}
	}
}
