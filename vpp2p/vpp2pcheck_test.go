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
	"strings"
	"testing"
)

var testID []byte

const testTitle = "Toto"
const testURL = "http://toto.bar/foo"

var testPubKey []byte
var testSig []byte

func init() {
	testID = []byte("abcdefghij")
	testPubKey = make([]byte, 300)
	testSig = make([]byte, 300)
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

func TestCheckTitle(t *testing.T) {
	b, err := CheckTitle(testTitle)
	if b != true || err != nil {
		t.Error("CheckTitle returned an error", err)
	}
	b, err = CheckTitle(strings.Repeat(" ", MinLenTitle))
	if b != true || err != nil {
		t.Error("CheckTitle returned an error on short Title", err)
	}
	b, err = CheckTitle(strings.Repeat(" ", MaxLenTitle))
	if b != true || err != nil {
		t.Error("CheckTitle returned an error on long Title", err)
	}
	b, err = CheckTitle(strings.Repeat(" ", MinLenTitle-1))
	if b == true || err == nil {
		t.Error("CheckTitle does not report an error on too shot Title")
	}
	b, err = CheckTitle(strings.Repeat(" ", MaxLenTitle+1))
	if b == true || err == nil {
		t.Error("CheckTitle does not report an error on too long Title")
	}
}

func TestCheckDescription(t *testing.T) {
	b, err := CheckDescription(testDescription)
	if b != true || err != nil {
		t.Error("CheckDescription returned an error", err)
	}
	b, err = CheckDescription(strings.Repeat(" ", MinLenDescription))
	if b != true || err != nil {
		t.Error("CheckDescription returned an error on short Description", err)
	}
	b, err = CheckDescription(strings.Repeat(" ", MaxLenDescription))
	if b != true || err != nil {
		t.Error("CheckDescription returned an error on long Description", err)
	}
	b, err = CheckDescription(strings.Repeat(" ", MinLenDescription-1))
	if b == true || err == nil {
		t.Error("CheckDescription does not report an error on too shot Description")
	}
	b, err = CheckDescription(strings.Repeat(" ", MaxLenDescription+1))
	if b == true || err == nil {
		t.Error("CheckDescription does not report an error on too long Description")
	}
}

func TestCheckURL(t *testing.T) {
	b, err := CheckURL(testURL)
	if b != true || err != nil {
		t.Error("CheckURL returned an error", err)
	}
	b, err = CheckURL(fmt.Sprintf("%s%s", testURL, strings.Repeat("x", MaxLenURL+1)))
	if b == true || err == nil {
		t.Error("CheckURL does not report an error on too long URL")
	}
}

func TestCheckPubKey(t *testing.T) {
	b, err := CheckPubKey(testPubKey)
	if b != true || err != nil {
		t.Error("CheckPubKey returned an error", err)
	}
	b, err = CheckPubKey(make([]byte, MinLenPubKey))
	if b != true || err != nil {
		t.Error("CheckPubKey returned an error on short PubKey", err)
	}
	b, err = CheckPubKey(make([]byte, MaxLenPubKey))
	if b != true || err != nil {
		t.Error("CheckPubKey returned an error on long PubKey", err)
	}
	b, err = CheckPubKey(make([]byte, MinLenPubKey-1))
	if b == true || err == nil {
		t.Error("CheckPubKey does not report an error on too shot PubKey")
	}
	b, err = CheckPubKey(make([]byte, MaxLenPubKey+1))
	if b == true || err == nil {
		t.Error("CheckPubKey does not report an error on too long PubKey")
	}
}

func TestCheckSig(t *testing.T) {
	b, err := CheckSig(testSig)
	if b != true || err != nil {
		t.Error("CheckSig returned an error", err)
	}
	b, err = CheckSig(make([]byte, MinLenSig))
	if b != true || err != nil {
		t.Error("CheckSig returned an error on short Sig", err)
	}
	b, err = CheckSig(make([]byte, MaxLenSig))
	if b != true || err != nil {
		t.Error("CheckSig returned an error on long Sig", err)
	}
	b, err = CheckSig(make([]byte, MaxLenSig+1))
	if b == true || err == nil {
		t.Error("CheckSig does not report an error on too long Sig")
	}
}
