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
	"testing"
)

func TestI64Mat4x4x4Math(t *testing.T) {
	const i11 = 130
	const i12 = 230
	const i13 = 330
	const i14 = 430
	const i21 = -140
	const i22 = -240
	const i23 = -340
	const i24 = -440
	const i31 = 110
	const i32 = 210
	const i33 = 310
	const i34 = 410
	const i41 = 1100
	const i42 = 2100
	const i43 = 3100
	const i44 = 4100

	const i51 = -6415
	const i52 = -7425
	const i53 = -8435
	const i54 = -9445
	const i61 = 664
	const i62 = 763
	const i63 = 862
	const i64 = 961
	const i71 = 624
	const i72 = 723
	const i73 = 822
	const i74 = 921
	const i81 = -63001
	const i82 = -73002
	const i83 = -83003
	const i84 = -93004

	var m1, m2, m3, m4 *I64Mat4x4x4

	m1 = I64Mat4x4x4New(i11, i12, i13, i14, i21, i22, i23, i24, i31, i32, i33, i34, i41, i42, i43, i44)

	m2 = m1.ToI32().ToI64()
	if *m1 != *m2 {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToX32().ToI64()
	if *m1 != *m2 {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToI64()
	if *m1 != *m2 {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToI64()
	if *m1 != *m2 {
		t.Error("F32 conversion error")
	}

	m2 = m1.ToF64().ToI64()
	if *m1 != *m2 {
		t.Error("F64 conversion error")
	}

	m2 = I64Mat4x4x4New(i51, i52, i53, i54, i61, i62, i63, i64, i71, i72, i73, i74, i81, i82, i83, i84)
	m3 = I64Mat4x4x4Add(m1, m2)
	m4 = I64Mat4x4x4New(i11+i51, i12+i52, i13+i53, i14+i54, i21+i61, i22+i62, i23+i63, i24+i64, i31+i71, i32+i72, i33+i73, i34+i74, i41+i81, i42+i82, i43+i83, i44+i84)
	if *m3 != *m4 {
		t.Error("Add error")
	}

	m3 = I64Mat4x4x4Sub(m1, m2)
	m4 = I64Mat4x4x4New(i11-i51, i12-i52, i13-i53, i14-i54, i21-i61, i22-i62, i23-i63, i24-i64, i31-i71, i32-i72, i33-i73, i34-i74, i41-i81, i42-i82, i43-i83, i44-i84)
	if *m3 != *m4 {
		t.Error("Sub error")
	}
}

func TestI64Mat4x4x4JSON(t *testing.T) {
	m1 := I64Mat4x4x4Identity()
	var m2 I64Mat4x4x4

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for I64Mat4x4x4 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for I64Mat4x4x4")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for I64Mat4x4x4, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for I64Mat4x4x4")
	}
	if *m1 != m2 {
		t.Error("unmarshalled matrix is different from original")
	}
}
