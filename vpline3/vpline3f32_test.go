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

package vpline3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpvec3"
	"testing"
)

func TestF32New(t *testing.T) {
	seg := F32NewSegment(vpvec3.F32New(0.1, 0.2, 0.3), vpvec3.F32New(1.1, 2.2, 3.3))
	t.Logf("Segment=%s", seg.String())
	tr := F32NewTriangle(vpvec3.F32New(0.1, 0.2, 0.3), vpvec3.F32New(1.1, 2.2, 3.3), vpvec3.F32New(1.0, 2.0, 3.0))
	t.Logf("Triangle=%s", tr.String())
	quad := F32NewQuad(vpvec3.F32New(0.1, 0.2, 0.3), vpvec3.F32New(1.1, 2.2, 3.3), vpvec3.F32New(1.0, 2.0, 3.0), vpvec3.F32New(11, 22, 33))
	t.Logf("Quad=%s", quad.String())
}

func TestF32Conv(t *testing.T) {
	l1 := F32NewSegment(vpvec3.F32New(0.1, 0.2, 0.3), vpvec3.F32New(1.1, 2.2, 3.3))
	if !l1.IsSimilar(l1) {
		t.Error("IsSimilar does not detect equality")
	}

	l2 := l1.ToX32().ToF32()
	if !l1.IsSimilar(l2) {
		t.Error("X32 conversion error")
	}

	l2 = l1.ToX64().ToF32()
	if !l1.IsSimilar(l2) {
		t.Error("X64 conversion error")
	}

	l2 = l1.ToF64().ToF32()
	if !l1.IsSimilar(l2) {
		t.Error("F64 conversion error")
	}
}

func TestF32JSON(t *testing.T) {
	line1 := F32NewTriangle(vpvec3.F32New(0.1, 0.2, 0.3), vpvec3.F32New(1.1, 2.2, 3.3), vpvec3.F32New(1.0, 2.0, 3.0))
	var line2 F32

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(line1)
	if err == nil {
		t.Logf("encoded JSON for F32 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32")
	}
	err = json.Unmarshal([]byte("nawak"), &line2)
	if err == nil {
		t.Error("able to decode JSON for F32, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, &line2)
	if err != nil {
		t.Error("unable to decode JSON for F32")
	}
	if !line1.IsSimilar(&line2) {
		t.Error("unmarshalled vector is different from original")
	}
}
