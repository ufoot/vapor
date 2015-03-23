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
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpmatrix2

import (
	"testing"
)

func TestI64Mat2Math(t *testing.T) {
	const i11 = 3
	const i12 = 4
	const i21 = -5
	const i22 = -6

	const i51 = -12
	const i52 = -13
	const i61 = 25
	const i62 = 121

	var m1, m2, m3, m4 *I64Mat2

	m1 = I64Mat2New(i11, i12, i21, i22)

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

	m2 = I64Mat2New(i51, i52, i61, i62)
	m3 = I64Mat2Add(m1, m2)
	m4 = I64Mat2New(i11+i51, i12+i52, i21+i61, i22+i62)
	if *m3 != *m4 {
		t.Error("Add error")
	}

	m3 = I64Mat2Sub(m1, m2)
	m4 = I64Mat2New(i11-i51, i12-i52, i21-i61, i22-i62)
	if *m3 != *m4 {
		t.Error("Sub error")
	}
}

func TestI64Mat2JSON(t *testing.T) {
	m1 := I64Mat2Identity()
	var m2 I64Mat2
	
	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for I64Mat2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for I64Mat2")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for I64Mat2, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for I64Mat2")
	}
	if *m1!= m2 {
		t.Error("unmarshalled matrix is different from original")
	}
}

