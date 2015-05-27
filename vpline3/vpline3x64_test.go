// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by64
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

package vpline3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
	"testing"
)

var ax64 = vpnumber.F64ToX64(0.1)
var ay64 = vpnumber.F64ToX64(0.2)
var az64 = vpnumber.F64ToX64(0.3)
var bx64 = vpnumber.F64ToX64(1.1)
var by64 = vpnumber.F64ToX64(2.2)
var bz64 = vpnumber.F64ToX64(3.3)
var cx64 = vpnumber.F64ToX64(1.0)
var cy64 = vpnumber.F64ToX64(2.0)
var cz64 = vpnumber.F64ToX64(3.0)
var dx64 = vpnumber.F64ToX64(11)
var dy64 = vpnumber.F64ToX64(22)
var dz64 = vpnumber.F64ToX64(33)

func TestX64New(t *testing.T) {
	seg := X64NewSegment(vpvec3.X64New(ax64, ay64, az64), vpvec3.X64New(bx64, by64, bz64))
	t.Logf("Segment=%s", seg.String())
	tr := X64NewTriangle(vpvec3.X64New(ax64, ay64, az64), vpvec3.X64New(bx64, by64, bz64), vpvec3.X64New(cx64, cy64, cz64))
	t.Logf("Triangle=%s", tr.String())
	quad := X64NewQuad(vpvec3.X64New(ax64, ay64, az64), vpvec3.X64New(bx64, by64, bz64), vpvec3.X64New(cx64, cy64, cz64), vpvec3.X64New(dx64, dy64, dz64))
	t.Logf("Quad=%s", quad.String())
}

func TestX64Conv(t *testing.T) {
	l1 := X64NewSegment(vpvec3.X64New(ax64, ay64, az64), vpvec3.X64New(bx64, by64, bz64))
	if !l1.IsSimilar(l1) {
		t.Error("IsSimilar does not detect equality")
	}

	l2 := l1.ToX32().ToX64()
	if !l1.IsSimilar(l2) {
		t.Error("X32 conversion error")
	}

	l2 = l1.ToF32().ToX64()
	if !l1.IsSimilar(l2) {
		t.Error("F32 conversion error")
	}

	l2 = l1.ToF64().ToX64()
	if !l1.IsSimilar(l2) {
		t.Error("F64 conversion error")
	}
}

func TestX64JSON(t *testing.T) {
	line1 := X64NewTriangle(vpvec3.X64New(ax64, ay64, az64), vpvec3.X64New(bx64, by64, bz64), vpvec3.X64New(cx64, cy64, cz64))
	var line2 X64

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(line1)
	if err == nil {
		t.Logf("encoded JSON for X64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64")
	}
	err = json.Unmarshal([]byte("nawak"), &line2)
	if err == nil {
		t.Error("able to decode JSON for X64, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &line2)
	if err != nil {
		t.Error("unable to decode JSON for X64")
	}
	if !line1.IsSimilar(&line2) {
		t.Error("unmarshalled vector is different from original")
	}
}
