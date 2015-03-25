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

func TestPackageTarname(t *testing.T) {
	tarname_string := PACKAGE_TARNAME
	tarname_runes := []rune(tarname_string)
	tarname_len := len(tarname_runes)

	if tarname_len <= 0 {
		t.Errorf("package tarname is too short \"%s\"", tarname_string)
	}
	for _, r := range tarname_runes {
		if !(unicode.IsDigit(r) || (unicode.IsLetter(r) && unicode.IsLower(r) && r < 128)) {
			t.Errorf("package tarname contains non letter/digit ascii7 char \"%s\"", tarname_string)
		}
	}

	t.Logf("tarname=%s", tarname_string)
}

func TestPackageName(t *testing.T) {
	name_string := PACKAGE_NAME
	name_runes := []rune(name_string)
	name_len := len(name_runes)

	if name_len <= 0 {
		t.Errorf("package name is too short \"%s\"", name_string)
	}
	for _, r := range name_runes {
		if !(r >= 32 && r < 128) {
			t.Errorf("package name contains non ascii7 char \"%s\"", name_string)
		}
	}

	t.Logf("name=%s", name_string)
}

func TestPackageEmail(t *testing.T) {
	email_string := PACKAGE_EMAIL
	email_runes := []rune(email_string)
	email_len := len(email_runes)

	if email_len <= 0 {
		t.Errorf("package email is too short \"%s\"", email_string)
	}
	for _, r := range email_runes {
		if !(r >= 32 && r < 128) {
			t.Errorf("package email contains non ascii7 char \"%s\"", email_string)
		}
	}

	t.Logf("email=%s", email_string)
}

func TestPackageUrl(t *testing.T) {
	url_string := PACKAGE_URL
	url_runes := []rune(url_string)
	url_len := len(url_runes)

	if url_len <= 0 {
		t.Errorf("package url is too short \"%s\"", url_string)
	}
	for _, r := range url_runes {
		if !(r >= 32 && r < 128) {
			t.Errorf("package url contains non ascii7 char \"%s\"", url_string)
		}
	}

	t.Logf("url=%s", url_string)
}
