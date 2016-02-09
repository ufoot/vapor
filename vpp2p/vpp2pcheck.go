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
)

// CheckID checks that an ID has the right format.
func CheckID(ID []byte) (bool, error) {
	lenID := len(ID)

	if lenID < MinLenID {
		return false, fmt.Errorf("ID is too short len=%d  min=%d", lenID, MinLenID)
	}
	if lenID > MaxLenID {
		return false, fmt.Errorf("ID is too long len=%d max=%d", lenID, MaxLenID)
	}

	return true, nil
}

// CheckTitle checks that a title is correct
func CheckTitle(title string) (bool, error) {
	lenTitle := len(title)

	if lenTitle < MinLenTitle {
		return false, fmt.Errorf("Title is too short len=%d  min=%d", lenTitle, MinLenTitle)
	}
	if lenTitle > MaxLenTitle {
		return false, fmt.Errorf("Title is too long len=%d max=%d", lenTitle, MaxLenTitle)
	}
	if !utf8.ValidString(title) {
		return false, fmt.Errorf("Title is not a valid UTF-8 string, len=%d", lenTitle)
	}
	for i, r := range title {
		if int(r) < 32 {
			return false, fmt.Errorf("Title contains invalid char %d at pos %d", int(r), i)
		}
	}

	return true, nil
}
