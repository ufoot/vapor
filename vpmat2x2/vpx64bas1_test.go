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

package vpmat2x2

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestX64Bas1Math(t *testing.T) {
	var o = vpnumber.F64ToX64(2.5)
	var x = vpnumber.F64ToX64(3.5)

	b := X64Bas1New(o, x)
	t.Logf("X64Bas1 b=%s", b.String())
	b.Normalize()
	t.Logf("X64Bas1 normalized b=%s", b.String())
	if !vpnumber.X64IsSimilar(b.O, o) {
		t.Error("X64Bas1 normalized origin changed")
	}
	if !vpnumber.X64IsSimilar(b.X, vpnumber.X64Const1) {
		t.Error("X64Bas1 normalized size is wrong")
	}
}

func TestX64Bas1JSON(t *testing.T) {
	var o vpnumber.X64
	var x vpnumber.X64
	b1 := X64Bas1Default()
	b2 := X64Bas1New(o, x)

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for X64Bas1 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64Bas1")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for X64Bas1, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for X64Bas1")
	}
	if *b1 != *b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
