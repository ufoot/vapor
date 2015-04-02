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
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF32Bas2Math(t *testing.T) {
	o := vpvec2.F32Vec2New(1.5, 2.5)
	x := vpvec2.F32Vec2New(2.1, 2.1)
	y := vpvec2.F32Vec2New(3.1, 3.1)

	b := F32Bas2New(o, x, y)
	t.Logf("F32Bas2 b=%s", b.String())
	b.Normalize()
	t.Logf("F32Bas2 normalized b=%s", b.String())
	if !b.O.IsSimilar(o) {
		t.Error("F32Bas2 normalized origin changed")
	}
	if !vpnumber.F32IsSimilar(b.X.Length(), vpnumber.F32Const1) {
		t.Error("F32Bas2 X normalized size is wrong")
	}
	if !vpnumber.F32IsSimilar(b.Y.Length(), vpnumber.F32Const1) {
		t.Error("F32Bas2 Y normalized size is wrong")
	}
	b.Ortho()
	dot := b.X.Dot(&b.Y)
	t.Logf("F32Bas2 ortho'ed b=%s, dot=%f", b.String(), dot)
	if !vpnumber.F32IsSimilar(dot, vpnumber.F32Const0) {
		t.Error("F32Bas2 ortho'ed does not yield 0 dot product")
	}
}

func TestF32Bas2JSON(t *testing.T) {
	b1 := F32Bas2Default()
	var b2 F32Bas2

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for F32Bas2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32Bas2")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for F32Bas2, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for F32Bas2")
	}
	if *b1 != b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
