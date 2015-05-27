// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by32
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

var ax32 = vpnumber.F32ToX32(0.1)
var ay32 = vpnumber.F32ToX32(0.2)
var az32 = vpnumber.F32ToX32(0.3)
var bx32 = vpnumber.F32ToX32(1.1)
var by32 = vpnumber.F32ToX32(2.2)
var bz32 = vpnumber.F32ToX32(3.3)
var cx32 = vpnumber.F32ToX32(1.0)
var cy32 = vpnumber.F32ToX32(2.0)
var cz32 = vpnumber.F32ToX32(3.0)
var dx32 = vpnumber.F32ToX32(11)
var dy32 = vpnumber.F32ToX32(22)
var dz32 = vpnumber.F32ToX32(33)

func TestX32New(t *testing.T) {
	seg := X32NewSegment(vpvec3.X32New(ax32, ay32, az32), vpvec3.X32New(bx32, by32, bz32))
	t.Logf("Segment=%s", seg.String())
	tr := X32NewTriangle(vpvec3.X32New(ax32, ay32, az32), vpvec3.X32New(bx32, by32, bz32), vpvec3.X32New(cx32, cy32, cz32))
	t.Logf("Triangle=%s", tr.String())
	quad := X32NewQuad(vpvec3.X32New(ax32, ay32, az32), vpvec3.X32New(bx32, by32, bz32), vpvec3.X32New(cx32, cy32, cz32), vpvec3.X32New(dx32, dy32, dz32))
	t.Logf("Quad=%s", quad.String())
}

func TestX32Conv(t *testing.T) {
	l1 := X32NewSegment(vpvec3.X32New(ax32, ay32, az32), vpvec3.X32New(bx32, by32, bz32))
	if !l1.IsSimilar(l1) {
		t.Error("IsSimilar does not detect equality")
	}

	l2 := l1.ToX64().ToX32()
	if !l1.IsSimilar(l2) {
		t.Error("X64 conversion error")
	}

	l2 = l1.ToF32().ToX32()
	if !l1.IsSimilar(l2) {
		t.Error("F32 conversion error")
	}

	l2 = l1.ToF64().ToX32()
	if !l1.IsSimilar(l2) {
		t.Error("F64 conversion error")
	}
}

func TestX32JSON(t *testing.T) {
	line1 := X32NewTriangle(vpvec3.X32New(ax32, ay32, az32), vpvec3.X32New(bx32, by32, bz32), vpvec3.X32New(cx32, cy32, cz32))
	var line2 X32

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(line1)
	if err == nil {
		t.Logf("encoded JSON for X32 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X32")
	}
	err = json.Unmarshal([]byte("nawak"), &line2)
	if err == nil {
		t.Error("able to decode JSON for X32, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &line2)
	if err != nil {
		t.Error("unable to decode JSON for X32")
	}
	if !line1.IsSimilar(&line2) {
		t.Error("unmarshalled vector is different from original")
	}
}
