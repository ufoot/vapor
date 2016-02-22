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

package vpp2p

import (
	"fmt"
	"net/url"
	"unicode/utf8"
)

const (
	// MinLenID is the minimum length for ID fields
	MinLenID = 4
	// MaxLenID is the maximum length for ID fields
	MaxLenID = 1000
	// MinLenTitle is the minimum length for Title fields
	MinLenTitle = 1
	// MaxLenTitle is the maximum length for Title fields
	MaxLenTitle = 100
	// MinLenURL is the minimum length for URL fields
	MinLenURL = 3
	// MaxLenURL is the maximum length for URL fields
	MaxLenURL = 1000
	// MinLenPubKey is the minimum length for PubKey fields
	MinLenPubKey = 100
	// MaxLenPubKey is the maximum length for PubKey fields
	MaxLenPubKey = 10000
)

func checkLenByte(fieldName string, content []byte, minLen, maxLen int) (bool, error) {
	l := len(content)

	if l < minLen {
		return false, fmt.Errorf("%s is too short len=%d min=%d", fieldName, l, minLen)
	}
	if l > maxLen {
		return false, fmt.Errorf("%s is too long len=%d max=%d", fieldName, l, maxLen)
	}

	return true, nil
}

func checkLenString(fieldName, content string, minLen, maxLen int) (bool, error) {
	return checkLenByte(fieldName, []byte(content), minLen, maxLen)
}

func checkUTF8(fieldName, content string) (bool, error) {
	if !utf8.ValidString(content) {
		return false, fmt.Errorf("%s is not a valid UTF-8 string", fieldName)
	}
	for i, r := range content {
		if int(r) < 32 {
			return false, fmt.Errorf("%s contains invalid char %d at pos %d", fieldName, int(r), i)
		}
	}

	return true, nil
}

func checkASCII(fieldName, content string) (bool, error) {
	if !utf8.ValidString(content) {
		return false, fmt.Errorf("%s is not a valid UTF-8 string", fieldName)
	}
	for i, r := range content {
		if int(r) < 32 || int(r) > 127 {
			return false, fmt.Errorf("%s contains invalid char %d at pos %d", fieldName, int(r), i)
		}
	}

	return true, nil
}

// CheckID checks that an ID has the right format.
func CheckID(ID []byte) (bool, error) {
	return checkLenByte("ID", ID, MinLenID, MaxLenID)
}

// CheckTitle checks that a title is correct
func CheckTitle(title string) (bool, error) {
	b, err := checkLenString("Title", title, MinLenTitle, MaxLenTitle)
	if b != true || err != nil {
		return false, err
	}
	b, err = checkUTF8("title", title)
	if b != true || err != nil {
		return false, err
	}

	return true, nil
}

// CheckURL checks that a a URL is correct
func CheckURL(u string) (bool, error) {
	var parsedURL *url.URL

	b, err := checkLenString("URL", u, MinLenURL, MaxLenURL)
	if b != true || err != nil {
		return false, err
	}
	b, err = checkUTF8("URL", u)
	if b != true || err != nil {
		return false, err
	}
	parsedURL, err = url.Parse(u)
	if err != nil {
		return false, err
	}
	if !parsedURL.IsAbs() {
		return false, fmt.Errorf("URL \"%s\" is not  absolute", u)
	}

	return true, nil
}

// CheckPubKey checks that a public key is correct
func CheckPubKey(pubKey []byte) (bool, error) {
	b, err := checkLenByte("PubKey", pubKey, MinLenPubKey, MaxLenPubKey)
	if b != true || err != nil {
		return false, err
	}
	b, err = checkASCII("PubKey", pubKey)
	if b != true || err != nil {
		return false, err
	}

	return true, nil
}
