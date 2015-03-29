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
	"github.com/ufoot/vapor/vpmatrix3"
	"github.com/ufoot/vapor/vpnumber"
	"math"
	"math/rand"
	"testing"
)

func TestF32Mat4Math(t *testing.T) {
	const f11 = 13.0
	const f12 = 23.0
	const f13 = 33.0
	const f14 = 43.0
	const f21 = -14.0
	const f22 = -24.0
	const f23 = -34.0
	const f24 = -44.0
	const f31 = 11.0
	const f32 = 21.0
	const f33 = 31.0
	const f34 = 41.0
	const f41 = 110.0
	const f42 = 210.0
	const f43 = 310.0
	const f44 = 410.0

	const f51 = -64.15
	const f52 = -74.25
	const f53 = -84.35
	const f54 = -94.45
	const f61 = 66.4
	const f62 = 76.3
	const f63 = 86.2
	const f64 = 96.1
	const f71 = 62.4
	const f72 = 72.3
	const f73 = 82.2
	const f74 = 92.1
	const f81 = -630.01
	const f82 = -730.02
	const f83 = -830.03
	const f84 = -930.04

	const fmul = 10.0

	var m1, m2, m3, m4 *F32Mat4

	m1 = F32Mat4New(f11, f12, f13, f14, f21, f22, f23, f24, f31, f32, f33, f34, f41, f42, f43, f44)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX32().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("F64 conversion error")
	}

	m2 = F32Mat4New(f51, f52, f53, f54, f61, f62, f63, f64, f71, f72, f73, f74, f81, f82, f83, f84)
	m3 = F32Mat4Add(m1, m2)
	m4 = F32Mat4New(f11+f51, f12+f52, f13+f53, f14+f54, f21+f61, f22+f62, f23+f63, f24+f64, f31+f71, f32+f72, f33+f73, f34+f74, f41+f81, f42+f82, f43+f83, f44+f84)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = F32Mat4Sub(m1, m2)
	m4 = F32Mat4New(f11-f51, f12-f52, f13-f53, f14-f54, f21-f61, f22-f62, f23-f63, f24-f64, f31-f71, f32-f72, f33-f73, f34-f74, f41-f81, f42-f82, f43-f83, f44-f84)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = F32Mat4MulScale(m1, fmul)
	m4 = F32Mat4New(f11*fmul, f12*fmul, f13*fmul, f14*fmul, f21*fmul, f22*fmul, f23*fmul, f24*fmul, f31*fmul, f32*fmul, f33*fmul, f34*fmul, f41*fmul, f42*fmul, f43*fmul, f44*fmul)
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = F32Mat4DivScale(m3, fmul)
	if !m3.IsSimilar(m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func invertableF32Mat4() *F32Mat4 {
	var ret F32Mat4

	for math.Abs(float64(ret.Det())) < 0.5 {
		for i := range ret {
			ret[i] = rand.Float32()
		}
	}

	return &ret
}

func TestF32Mat4Comp(t *testing.T) {
	m1 := invertableF32Mat4()
	m2 := F32Mat4Inv(m1)
	id := F32Mat4Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF32Mat4Aff(t *testing.T) {
	const p1 = 3.0
	const p2 = 4.0
	const p3 = 5.0
	const t1 = 6.5
	const t2 = 8.5
	const t3 = 10.5

	v1 := F32Vec4New(p1, p2, p3, vpnumber.F32Const1)
	vt := vpmatrix3.F32Vec3New(t1, t2, t3)
	mt := F32Mat4Trans(vt)
	t.Logf("translation mat4 for %s is %s", vt.String(), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat4 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := F32Vec4New(p1+t1, p2+t2, p3+t3, vpnumber.F32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4 MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(v1.ToVec3())
	v3pos := v1.ToVec3().Add(vt)
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4 MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir := mt.MulVecDir(v1.ToVec3())
	v3dir := v1.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4 MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}
}

func TestF32Mat4JSON(t *testing.T) {
	m1 := invertableF32Mat4()
	m2 := F32Mat4Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F32Mat4 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32Mat4")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F32Mat4, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F32Mat4")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF32Mat4Add(b *testing.B) {
	mat := F32Mat4New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF32Mat4Inv(b *testing.B) {
	mat := invertableF32Mat4()

	for i := 0; i < b.N; i++ {
		_ = F32Mat4Inv(mat)
	}
}
