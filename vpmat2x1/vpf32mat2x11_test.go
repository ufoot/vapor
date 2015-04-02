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

package vpmat2x1

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF32Bas1Math(t *testing.T) {
	const o = 2.5
	const x = 3.5

	b := F32Bas1New(o, x)
	t.Logf("F32Bas1 b=%s", b.String())
	b.Normalize()
	t.Logf("F32Bas1 normalized b=%s", b.String())
	if !vpnumber.F32IsSimilar(b.O, o) {
		t.Error("F32Bas1 normalized origin changed")
	}
	if !vpnumber.F32IsSimilar(b.X, vpnumber.F32Const1) {
		t.Error("F32Bas1 normalized size is wrong")
	}
}

func TestF32Bas1JSON(t *testing.T) {
	var o float32
	var x float32
	b1 := F32Bas1Default()
	b2 := F32Bas1New(o, x)

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for F32Bas1 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32Bas1")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for F32Bas1, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for F32Bas1")
	}
	if *b1 != *b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
