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
	"testing"
)

var testID []byte

const testHostTitle = "Toto"
const testHostURL = "http://toto.bar/foo"

func init() {
	testID = []byte("abcdefghij")
}

func TestCheckID(t *testing.T) {
	b, err := CheckID(testID)
	if b != true || err != nil {
		t.Error("CheckID returned an error", err)
	}
	b, err = CheckID(make([]byte, MinLenID))
	if b != true || err != nil {
		t.Error("CheckID returned an error on short ID", err)
	}
	b, err = CheckID(make([]byte, MaxLenID))
	if b != true || err != nil {
		t.Error("CheckID returned an error on long ID", err)
	}
	b, err = CheckID(make([]byte, MinLenID-1))
	if b == true || err == nil {
		t.Error("CheckID does not report an error on too shot ID")
	}
	b, err = CheckID(make([]byte, MaxLenID+1))
	if b == true || err == nil {
		t.Error("CheckID does not report an error on too long ID")
	}
}
