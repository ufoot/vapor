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

package vpmat4x4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
	"testing"
)

func TestX64Bas3Math(t *testing.T) {
	o := vpvec3.F64Vec3New(1.5, 2.5, 3.5).ToX64()
	x := vpvec3.F64Vec3New(1.1, 2.1, 3.1).ToX64()
	y := vpvec3.F64Vec3New(1.3, 2.3, 3.3).ToX64()
	z := vpvec3.F64Vec3New(1.7, 2.7, 3.7).ToX64()

	b := X64Bas3New(o, x, y, z)
	t.Logf("X64Bas3 b=%s", b.String())
	b.Normalize()
	t.Logf("X64Bas3 normalized b=%s", b.String())
	if !b.O.IsSimilar(o) {
		t.Error("X64Bas3 normalized origin changed")
	}
	if !vpnumber.X64IsSimilar(b.X.Length(), vpnumber.X64Const1) {
		t.Error("X64Bas3 X normalized size is wrong")
	}
	if !vpnumber.X64IsSimilar(b.Y.Length(), vpnumber.X64Const1) {
		t.Error("X64Bas3 Y normalized size is wrong")
	}
	if !vpnumber.X64IsSimilar(b.Z.Length(), vpnumber.X64Const1) {
		t.Error("X64Bas3 Z normalized size is wrong")
	}
	b.Ortho()
	dotXY := b.X.Dot(&b.Y)
	dotYZ := b.Y.Dot(&b.Z)
	dotZX := b.Z.Dot(&b.X)
	t.Logf("X64Bas3 ortho'ed b=%s, dotXY=%f dotYZ=%d dotZX=%f", b.String(), dotXY, dotYZ, dotZX)
	if !vpnumber.X64IsSimilar(dotXY, vpnumber.X64Const0) || !vpnumber.X64IsSimilar(dotYZ, vpnumber.X64Const0) || !vpnumber.X64IsSimilar(dotZX, vpnumber.X64Const0) {
		t.Error("X64Bas3 ortho'ed does not yield 0 dot products")
	}
}

func TestX64Bas3JSON(t *testing.T) {
	b1 := X64Bas3Default()
	var b2 X64Bas3

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for X64Bas3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64Bas3")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for X64Bas3, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for X64Bas3")
	}
	if *b1 != b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
