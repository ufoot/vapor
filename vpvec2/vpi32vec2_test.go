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

package vpvec2

import (
	"testing"
)

func TestI32Vec2Math(t *testing.T) {
	const i1 = 0
	const i2 = -4

	const i5 = -10
	const i6 = 1000

	var v1, v2, v3, v4 *I32Vec2

	v1 = I32Vec2New(i1, i2)

	v2 = v1.ToI64().ToI32()
	if *v1 != *v2 {
		t.Error("I64 conversion error")
	}

	v2 = v1.ToX32().ToI32()
	if *v1 != *v2 {
		t.Error("X32 conversion error")
	}

	v2 = v1.ToX64().ToI32()
	if *v1 != *v2 {
		t.Error("X64 conversion error")
	}

	v2 = v1.ToF32().ToI32()
	if *v1 != *v2 {
		t.Error("F32 conversion error")
	}

	v2 = v1.ToF64().ToI32()
	if *v1 != *v2 {
		t.Error("F64 conversion error")
	}

	v2 = I32Vec2New(i5, i6)
	v3 = I32Vec2Add(v1, v2)
	v4 = I32Vec2New(i1+i5, i2+i6)
	if *v3 != *v4 {
		t.Error("Add error")
	}

	v3 = I32Vec2Sub(v1, v2)
	v4 = I32Vec2New(i1-i5, i2-i6)
	if *v3 != *v4 {
		t.Error("Sub error")
	}

	v3 = I32Vec2Add(v1, I32Vec2Neg(v2))
	v4 = I32Vec2Sub(v1, v2)
	if *v3 != *v4 {
		t.Errorf("Neg error")
	}
}

func TestI32Vec2JSON(t *testing.T) {
	m1 := I32Vec2New(10, 20)
	m2 := I32Vec2New(1, 0)

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for I32Vec2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for I32Vec2")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for I32Vec2, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for I32Vec2")
	}
	if *m1 != *m2 {
		t.Error("unmarshalled vector is different from original")
	}
}
