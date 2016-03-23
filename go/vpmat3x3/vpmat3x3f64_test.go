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

package vpmat3x3

import (
	"github.com/ufoot/vapor/go/vpnumber"
	"github.com/ufoot/vapor/go/vpvec2"
	"github.com/ufoot/vapor/go/vpvec3"
	"math"
	"math/rand"
	"testing"
)

func TestF64Math(t *testing.T) {
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

	var m1, m2, m3, m4 *F64

	m1 = F64New(f11, f12, f13, f21, f22, f23, f31, f32, f33)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
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

	m2 = F64New(f51, f52, f53, f61, f62, f63, f71, f72, f73)
	m3 = F64Add(m1, m2)
	m4 = F64New(f11+f51, f12+f52, f13+f53, f21+f61, f22+f62, f23+f63, f31+f71, f32+f72, f33+f73)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = F64Sub(m1, m2)
	m4 = F64New(f11-f51, f12-f52, f13-f53, f21-f61, f22-f62, f23-f63, f31-f71, f32-f72, f33-f73)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = F64MulScale(m1, fmul)
	m4 = F64New(f11*fmul, f12*fmul, f13*fmul, f21*fmul, f22*fmul, f23*fmul, f31*fmul, f32*fmul, f33*fmul)
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = F64DivScale(m3, fmul)
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

func invertableF64() *F64 {
	var ret F64

	for math.Abs(float64(ret.Det())) < 0.25 {
		for i := range ret {
			ret[i] = rand.Float64()
		}
	}

	return &ret
}

func randomVecF64() *vpvec2.F64 {
	var ret vpvec2.F64

	for i := range ret {
		ret[i] = rand.Float64()
	}

	return &ret
}

func TestF64Comp(t *testing.T) {
	m1 := invertableF64()
	m2 := F64Inv(m1)
	id := F64Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF64Aff(t *testing.T) {
	const p1 = 3.0
	const p2 = 4.0
	const t1 = 6.5
	const t2 = 8.5

	v1 := vpvec3.F64New(p1, p2, vpnumber.F64Const1)
	vt := vpvec2.F64New(t1, t2)
	mt := F64Translation(vt)
	t.Logf("translation mat3x3 for %s is %s", vt.String(), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat3x3 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec3.F64New(p1+t1, p2+t2, vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat3x3 translation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(v1.ToVec2())
	v3pos := v1.ToVec2().Add(vt)
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat3x3 translation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir := mt.MulVecDir(v1.ToVec2())
	v3dir := v1.ToVec2()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat3x3 translation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr := F64Rot(math.Pi / 2)
	t.Logf("rotation mat3x3 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat3x3 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec3.F64New(-v1[1], v1[0], vpnumber.F64Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat3x3 rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec2())
	v3pos = v3.ToVec2()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat3x3 rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec2())
	v3dir = v3.ToVec2()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat3x3 rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}
}

func TestF64Rebase(t *testing.T) {
	const cx = 1.5
	const cy = 2.0
	m1 := invertableF64()
	var vo1 vpvec2.F64
	vx1 := vpvec2.F64AxisX()
	vy1 := vpvec2.F64AxisY()
	vp1 := vpvec2.F64Add(vx1, vy1)
	vo2 := randomVecF64()
	vx2 := vpvec2.F64Add(m1.GetCol(0).ToVec2(), vo2)
	vy2 := vpvec2.F64Add(m1.GetCol(1).ToVec2(), vo2)
	vp2 := vpvec2.F64Add(vo2, vpvec2.F64Add(vpvec2.F64MulScale(vpvec2.F64Sub(vx2, vo2), cx), vpvec2.F64MulScale(vpvec2.F64Sub(vy2, vo2), cy)))

	m2 := F64RebaseOXY(vo2, vx2, vy2)
	t.Logf("transformation matrix for O=%s X=%s Y=%s is M=%s", vo2.String(), vx2.String(), vy2.String(), m2.String())

	vo3 := m2.MulVecPos(&vo1)
	if !vo3.IsSimilar(vo2) {
		t.Errorf("vo1 -> vo2 error vo1=%s vo2=%s vo3=%s", vo1.String(), vo2.String(), vo3.String())
	}
	vx3 := m2.MulVecPos(vx1)
	if !vx3.IsSimilar(vx2) {
		t.Errorf("vx1 -> vx2 error vx1=%s vx2=%s vx3=%s", vx1.String(), vx2.String(), vx3.String())
	}
	vy3 := m2.MulVecPos(vy1)
	if !vy3.IsSimilar(vy2) {
		t.Errorf("vy1 -> vy2 error vy1=%s vy2=%s vy3=%s", vy1.String(), vy2.String(), vy3.String())
	}

	m2 = F64RebaseOXYP(vo2, vx2, vy2, vp2)
	t.Logf("transformation matrix for O=%s X=%s Y=%s P=%s is M=%s", vo2.String(), vx2.String(), vy2.String(), vp2.String(), m2.String())
	vo3 = m2.MulVecPos(&vo1)
	if !vo3.IsSimilar(vo2) {
		t.Errorf("vo1 -> vo2 error vo1=%s vo2=%s vo3=%s", vo1.String(), vo2.String(), vo3.String())
	}
	vx3 = m2.MulVecPos(vx1)
	if !vx3.IsSimilar(vx2) {
		t.Errorf("vx1 -> vx2 error vx1=%s vx2=%s vx3=%s", vx1.String(), vx2.String(), vx3.String())
	}
	vy3 = m2.MulVecPos(vy1)
	if !vy3.IsSimilar(vy2) {
		t.Errorf("vy1 -> vy2 error vy1=%s vy2=%s vy3=%s", vy1.String(), vy2.String(), vy3.String())
	}
	vp3 := m2.MulVecPos(vp1)
	if !vp3.IsSimilar(vp2) {
		t.Errorf("vp1 -> vp2 error vp1=%s vp2=%s vp3=%s", vp1.String(), vp2.String(), vp3.String())
	}
}

func TestF64JSON(t *testing.T) {
	m1 := invertableF64()
	m2 := F64Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F64, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F64")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF64Add(b *testing.B) {
	mat := F64New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF64Inv(b *testing.B) {
	mat := invertableF64()

	for i := 0; i < b.N; i++ {
		_ = F64Inv(mat)
	}
}
