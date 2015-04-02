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
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
	"github.com/ufoot/vapor/vpnumber"
	"math"
	"math/rand"
	"testing"
)

func TestF64Mat3x3Math(t *testing.T) {
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

	var m1, m2, m3, m4 *F64Mat3x3

	m1 = F64Mat3x3New(f11, f12, f13, f21, f22, f23, f31, f32, f33)
	if !m1.IsSimilar(m1) {
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

	m2 = F64Mat3x3New(f51, f52, f53, f61, f62, f63, f71, f72, f73)
	m3 = F64Mat3x3Add(m1, m2)
	m4 = F64Mat3x3New(f11+f51, f12+f52, f13+f53, f21+f61, f22+f62, f23+f63, f31+f71, f32+f72, f33+f73)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = F64Mat3x3Sub(m1, m2)
	m4 = F64Mat3x3New(f11-f51, f12-f52, f13-f53, f21-f61, f22-f62, f23-f63, f31-f71, f32-f72, f33-f73)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = F64Mat3x3MulScale(m1, fmul)
	m4 = F64Mat3x3New(f11*fmul, f12*fmul, f13*fmul, f21*fmul, f22*fmul, f23*fmul, f31*fmul, f32*fmul, f33*fmul)
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = F64Mat3x3DivScale(m3, fmul)
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

func invertableF64Mat3x3() *F64Mat3x3 {
	var ret F64Mat3x3

	for math.Abs(float64(ret.Det())) < 0.25 {
		for i := range ret {
			ret[i] = rand.Float64()
		}
	}

	return &ret
}

func TestF64Mat3x3Comp(t *testing.T) {
	m1 := invertableF64Mat3x3()
	m2 := F64Mat3x3Inv(m1)
	id := F64Mat3x3Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF64Mat3x3Aff(t *testing.T) {
	const p1 = 3.0
	const p2 = 4.0
	const t1 = 6.5
	const t2 = 8.5

	v1 := vpvec3.F64Vec3New(p1, p2, vpnumber.F64Const1)
	vt := vpvec2.F64Vec2New(t1, t2)
	mt := F64Mat3x3Trans(vt)
	t.Logf("translation mat3x3x3 for %s is %s", vt.String(), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat3x3x3 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec3.F64Vec3New(p1+t1, p2+t2, vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat3x3x3 translation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(v1.ToVec2())
	v3pos := v1.ToVec2().Add(vt)
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat3x3x3 translation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir := mt.MulVecDir(v1.ToVec2())
	v3dir := v1.ToVec2()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat3x3x3 translation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr := F64Mat3x3Rot(math.Pi / 2)
	t.Logf("rotation mat3x3x3 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat3x3x3 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec3.F64Vec3New(-v1[1], v1[0], vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat3x3x3 rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec2())
	v3pos = v3.ToVec2()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat3x3x3 rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec2())
	v3dir = v3.ToVec2()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat3x3x3 rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}
}

func TestF64Mat3x3JSON(t *testing.T) {
	m1 := invertableF64Mat3x3()
	m2 := F64Mat3x3Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F64Mat3x3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64Mat3x3")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F64Mat3x3, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F64Mat3x3")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF64Mat3x3Add(b *testing.B) {
	mat := F64Mat3x3New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF64Mat3x3Inv(b *testing.B) {
	mat := invertableF64Mat3x3()

	for i := 0; i < b.N; i++ {
		_ = F64Mat3x3Inv(mat)
	}
}
