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

package vpmatrix3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmatrix2"
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF64Bas2Math(t *testing.T) {
	o := vpmatrix2.F64Vec2New(1.5, 2.5)
	x := vpmatrix2.F64Vec2New(2.1, 2.1)
	y := vpmatrix2.F64Vec2New(3.1, 3.1)

	b := F64Bas2New(o, x, y)
	t.Logf("F64Bas2 b=%s", b.String())
	b.Normalize()
	t.Logf("F64Bas2 normalized b=%s", b.String())
	if !b.O.IsSimilar(o) {
		t.Error("F64Bas2 normalized origin changed")
	}
	if !vpnumber.F64IsSimilar(b.X.Length(), vpnumber.F64Const1) {
		t.Error("F64Bas2 X normalized size is wrong")
	}
	if !vpnumber.F64IsSimilar(b.Y.Length(), vpnumber.F64Const1) {
		t.Error("F64Bas2 Y normalized size is wrong")
	}
	b.Ortho()
	dot := b.X.Dot(&b.Y)
	t.Logf("F64Bas2 ortho'ed b=%s, dot=%f", b.String(), dot)
	if !vpnumber.F64IsSimilar(dot, vpnumber.F64Const0) {
		t.Error("F64Bas2 ortho'ed does not yield 0 dot product")
	}
}

func TestF64Bas2JSON(t *testing.T) {
	b1 := F64Bas2Default()
	var b2 F64Bas2

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for F64Bas2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64Bas2")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for F64Bas2, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for F64Bas2")
	}
	if *b1 != b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
