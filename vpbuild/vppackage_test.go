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

package vpbuild

import (
	"testing"
	"unicode"
)

func TestPackageTarname(t *testing.T) {
	tarnameString := PackageTarname
	tarnameRunes := []rune(tarnameString)
	tarnameLen := len(tarnameRunes)

	if tarnameLen <= 0 {
		t.Errorf("package tarname is too short \"%s\"", tarnameString)
	}
	for _, r := range tarnameRunes {
		if !(unicode.IsDigit(r) || (unicode.IsLetter(r) && unicode.IsLower(r) && r < 128)) {
			t.Errorf("package tarname contains non letter/digit ascii7 char \"%s\"", tarnameString)
		}
	}

	t.Logf("tarname=%s", tarnameString)
}

func TestPackageName(t *testing.T) {
	nameString := PackageName
	nameRunes := []rune(nameString)
	nameLen := len(nameRunes)

	if nameLen <= 0 {
		t.Errorf("package name is too short \"%s\"", nameString)
	}
	for _, r := range nameRunes {
		if !(r >= 32 && r < 128) {
			t.Errorf("package name contains non ascii7 char \"%s\"", nameString)
		}
	}

	t.Logf("name=%s", nameString)
}

func TestPackageEmail(t *testing.T) {
	emailString := PackageEmail
	emailRunes := []rune(emailString)
	emailLen := len(emailRunes)

	if emailLen <= 0 {
		t.Errorf("package email is too short \"%s\"", emailString)
	}
	for _, r := range emailRunes {
		if !(r >= 32 && r < 128) {
			t.Errorf("package email contains non ascii7 char \"%s\"", emailString)
		}
	}

	t.Logf("email=%s", emailString)
}

func TestPackageURL(t *testing.T) {
	urlString := PackageURL
	urlRunes := []rune(urlString)
	urlLen := len(urlRunes)

	if urlLen <= 0 {
		t.Errorf("package url is too short \"%s\"", urlString)
	}
	for _, r := range urlRunes {
		if !(r >= 32 && r < 128) {
			t.Errorf("package url contains non ascii7 char \"%s\"", urlString)
		}
	}

	t.Logf("url=%s", urlString)
}
