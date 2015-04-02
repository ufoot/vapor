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

package vpmat3x3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"testing"
)

func TestX64Bas2Math(t *testing.T) {
	o := vpvec2.F64Vec2New(1.5, 2.5).ToX64()
	x := vpvec2.F64Vec2New(2.1, 2.1).ToX64()
	y := vpvec2.F64Vec2New(3.1, 3.1).ToX64()

	b := X64Bas2New(o, x, y)
	t.Logf("X64Bas2 b=%s", b.String())
	b.Normalize()
	t.Logf("X64Bas2 normalized b=%s", b.String())
	if !b.O.IsSimilar(o) {
		t.Error("X64Bas2 normalized origin changed")
	}
	if !vpnumber.X64IsSimilar(b.X.Length(), vpnumber.X64Const1) {
		t.Error("X64Bas2 X normalized size is wrong")
	}
	if !vpnumber.X64IsSimilar(b.Y.Length(), vpnumber.X64Const1) {
		t.Error("X64Bas2 Y normalized size is wrong")
	}
	b.Ortho()
	dot := b.X.Dot(&b.Y)
	t.Logf("X64Bas2 ortho'ed b=%s, dot=%f", b.String(), dot)
	if !vpnumber.X64IsSimilar(dot, vpnumber.X64Const0) {
		t.Error("X64Bas2 ortho'ed does not yield 0 dot product")
	}
}

func TestX64Bas2JSON(t *testing.T) {
	b1 := X64Bas2Default()
	var b2 X64Bas2

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for X64Bas2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64Bas2")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for X64Bas2, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for X64Bas2")
	}
	if *b1 != b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
