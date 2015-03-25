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

package vpmatrix3

import (
	"math"
	"math/rand"
	"testing"
	"ufoot.org/vapor/vpnumber"
)

func TestF64Mat3Math(t *testing.T) {
	const f11 = 3.0
	const f12 = 333.0
	const f13 = 31.0
	const f21 = -4.0
	const f22 = -24.0
	const f23 = -4.0
	const f31 = 1.0
	const f32 = 11.0
	const f33 = 1.0

	const f51 = -4.5
	const f52 = -4.2
	const f53 = -4.5
	const f61 = 6.0
	const f62 = 4.0
	const f63 = 3.0
	const f71 = 2.0
	const f72 = 2.0
	const f73 = 1.0

	const fmul = 10.0

	var m1, m2, m3, m4 *F64Mat3

	m1 = F64Mat3New(f11, f12, f13, f21, f22, f23, f31, f32, f33)
	if !F64Mat3IsSimilar(m1, m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("F32 conversion error")
	}

	m2 = F64Mat3New(f51, f52, f53, f61, f62, f63, f71, f72, f73)
	m3 = F64Mat3Add(m1, m2)
	m4 = F64Mat3New(f11+f51, f12+f52, f13+f53, f21+f61, f22+f62, f23+f63, f31+f71, f32+f72, f33+f73)
	if !F64Mat3IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = F64Mat3Sub(m1, m2)
	m4 = F64Mat3New(f11-f51, f12-f52, f13-f53, f21-f61, f22-f62, f23-f63, f31-f71, f32-f72, f33-f73)
	if !F64Mat3IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = F64Mat3MulScale(m1, fmul)
	m4 = F64Mat3New(f11*fmul, f12*fmul, f13*fmul, f21*fmul, f22*fmul, f23*fmul, f31*fmul, f32*fmul, f33*fmul)
	if !F64Mat3IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = F64Mat3DivScale(m3, fmul)
	if !F64Mat3IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func invertableF64Mat3() *F64Mat3 {
	var ret F64Mat3

	for math.Abs(float64(ret.Det())) < 0.25 {
		for i := range ret {
			ret[i] = rand.Float64()
		}
	}

	return &ret
}

func TestF64Mat3Comp(t *testing.T) {
	m1 := invertableF64Mat3()
	m2 := F64Mat3Inv(m1)
	id := F64Mat3Identity()

	m2.MulComp(m1)
	if F64Mat3IsSimilar(m2, id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF64Mat3JSON(t *testing.T) {
	m1 := invertableF64Mat3()
	m2 := F64Mat3Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F64Mat3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64Mat3")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F64Mat3, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F64Mat3")
	}
	if !F64Mat3IsSimilar(m1, m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF64Mat3Add(b *testing.B) {
	mat := F64Mat3New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}
