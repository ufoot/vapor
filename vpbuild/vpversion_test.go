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

package vpbuild

import (
	"testing"
	"unicode"
)

func TestVersionMajor(t *testing.T) {
	major := VERSION_MAJOR

	if major < 0 {
		t.Errorf("negative major version %d", major)
	}

	t.Logf("major=%d", major)
}

func TestVersionMinor(t *testing.T) {
	minor := VERSION_MINOR

	if minor <= 0 {
		t.Errorf("zero or negative minor version %d", minor)
	}

	t.Logf("minor=%d", minor)
}

func TestVersionStamp(t *testing.T) {
	stamp_string := VERSION_STAMP
	stamp_runes := []rune(stamp_string)
	stamp_len := len(stamp_runes)

	if stamp_len <= 0 {
		t.Errorf("stamp is too short \"%s\"", stamp_string)
	}
	for _, r := range stamp_runes {
		if !(unicode.IsDigit(r) || (unicode.IsLetter(r) && r < 128)) {
			t.Errorf("stamp contains non letter/digit ascii7 char \"%s\"", stamp_string)
		}
	}

	t.Logf("stamp=%s", stamp_string)
}
