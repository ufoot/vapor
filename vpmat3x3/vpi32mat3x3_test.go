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
	"testing"
)

func TestI32Mat3x3Math(t *testing.T) {
	const i11 = 3
	const i12 = 333
	const i13 = 31
	const i21 = -4
	const i22 = -24
	const i23 = -4
	const i31 = 1
	const i32 = 11
	const i33 = 7

	const i51 = 9
	const i52 = 12
	const i53 = -4
	const i61 = -123
	const i62 = 12
	const i63 = 3
	const i71 = 2
	const i72 = 2
	const i73 = 1

	var m1, m2, m3, m4 *I32Mat3x3

	m1 = I32Mat3x3New(i11, i12, i13, i21, i22, i23, i31, i32, i33)

	m2 = m1.ToI64().ToI32()
	if *m1 != *m2 {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX32().ToI32()
	if *m1 != *m2 {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToI32()
	if *m1 != *m2 {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToI32()
	if *m1 != *m2 {
		t.Error("F32 conversion error")
	}

	m2 = m1.ToF64().ToI32()
	if *m1 != *m2 {
		t.Error("F64 conversion error")
	}

	m2 = I32Mat3x3New(i51, i52, i53, i61, i62, i63, i71, i72, i73)
	m3 = I32Mat3x3Add(m1, m2)
	m4 = I32Mat3x3New(i11+i51, i12+i52, i13+i53, i21+i61, i22+i62, i23+i63, i31+i71, i32+i72, i33+i73)
	if *m3 != *m4 {
		t.Error("Add error")
	}

	m3 = I32Mat3x3Sub(m1, m2)
	m4 = I32Mat3x3New(i11-i51, i12-i52, i13-i53, i21-i61, i22-i62, i23-i63, i31-i71, i32-i72, i33-i73)
	if *m3 != *m4 {
		t.Error("Sub error")
	}
}

func TestI32Mat3x3JSON(t *testing.T) {
	m1 := I32Mat3x3Identity()
	var m2 I32Mat3x3

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for I32Mat3x3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for I32Mat3x3")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for I32Mat3x3, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for I32Mat3x3")
	}
	if *m1 != m2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
