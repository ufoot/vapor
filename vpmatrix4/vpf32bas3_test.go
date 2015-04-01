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

package vpmatrix4

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpmatrix3"
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF32Bas3Math(t *testing.T) {
	o := vpmatrix3.F32Vec3New(1.5, 2.5, 3.5)
	x := vpmatrix3.F32Vec3New(1.1, 2.1, 3.1)
	y := vpmatrix3.F32Vec3New(1.3, 2.3, 3.3)
	z := vpmatrix3.F32Vec3New(1.7, 2.7, 3.7)

	b := F32Bas3New(o, x, y, z)
	t.Logf("F32Bas3 b=%s", b.String())
	b.Normalize()
	t.Logf("F32Bas3 normalized b=%s", b.String())
	if !b.O.IsSimilar(o) {
		t.Error("F32Bas3 normalized origin changed")
	}
	if !vpnumber.F32IsSimilar(b.X.Length(), vpnumber.F32Const1) {
		t.Error("F32Bas3 X normalized size is wrong")
	}
	if !vpnumber.F32IsSimilar(b.Y.Length(), vpnumber.F32Const1) {
		t.Error("F32Bas3 Y normalized size is wrong")
	}
	if !vpnumber.F32IsSimilar(b.Z.Length(), vpnumber.F32Const1) {
		t.Error("F32Bas3 Z normalized size is wrong")
	}
	b.Ortho()
	dotXY := b.X.Dot(&b.Y)
	dotYZ := b.Y.Dot(&b.Z)
	dotZX := b.Z.Dot(&b.X)
	t.Logf("F32Bas3 ortho'ed b=%s, dotXY=%f dotYZ=%d dotZX=%f", b.String(), dotXY, dotYZ, dotZX)
	if !vpnumber.F32IsSimilar(dotXY, vpnumber.F32Const0) || !vpnumber.F32IsSimilar(dotYZ, vpnumber.F32Const0) || !vpnumber.F32IsSimilar(dotZX, vpnumber.F32Const0) {
		t.Error("F32Bas3 ortho'ed does not yield 0 dot products")
	}
}

func TestF32Bas3JSON(t *testing.T) {
	b1 := F32Bas3Default()
	var b2 F32Bas3

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(b1)
	if err == nil {
		t.Logf("encoded JSON for F32Bas3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32Bas3")
	}
	err = json.Unmarshal([]byte("nawak"), &b2)
	if err == nil {
		t.Error("able to decode JSON for F32Bas3, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &b2)
	if err != nil {
		t.Error("unable to decode JSON for F32Bas3")
	}
	if *b1 != b2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
