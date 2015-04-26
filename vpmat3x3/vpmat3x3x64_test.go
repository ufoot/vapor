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
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
	"math/rand"
	"testing"
)

func TestX64Math(t *testing.T) {
	var x11 = vpnumber.F64ToX64(3.0)
	var x12 = vpnumber.F64ToX64(333.0)
	var x13 = vpnumber.F64ToX64(31.0)
	var x21 = vpnumber.F64ToX64(-4.0)
	var x22 = vpnumber.F64ToX64(-24.0)
	var x23 = vpnumber.F64ToX64(-4.0)
	var x31 = vpnumber.F64ToX64(1.0)
	var x32 = vpnumber.F64ToX64(11.0)
	var x33 = vpnumber.F64ToX64(1.0)

	var x51 = vpnumber.F64ToX64(-4.5)
	var x52 = vpnumber.F64ToX64(-4.2)
	var x53 = vpnumber.F64ToX64(-4.5)
	var x61 = vpnumber.F64ToX64(6.0)
	var x62 = vpnumber.F64ToX64(4.0)
	var x63 = vpnumber.F64ToX64(3.0)
	var x71 = vpnumber.F64ToX64(2.0)
	var x72 = vpnumber.F64ToX64(2.0)
	var x73 = vpnumber.F64ToX64(1.0)

	var xmul = vpnumber.F64ToX64(10.0)

	var m1, m2, m3, m4 *X64

	m1 = X64New(x11, x12, x13, x21, x22, x23, x31, x32, x33)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToX32().ToX64()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToF32().ToX64()
	if !m1.IsSimilar(m2) {
		t.Error("F32 conversion error")
	}

	m2 = m1.ToF64().ToX64()
	if !m1.IsSimilar(m2) {
		t.Error("F64 conversion error")
	}

	m2 = X64New(x51, x52, x53, x61, x62, x63, x71, x72, x73)
	m3 = X64Add(m1, m2)
	m4 = X64New(x11+x51, x12+x52, x13+x53, x21+x61, x22+x62, x23+x63, x31+x71, x32+x72, x33+x73)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = X64Sub(m1, m2)
	m4 = X64New(x11-x51, x12-x52, x13-x53, x21-x61, x22-x62, x23-x63, x31-x71, x32-x72, x33-x73)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = X64MulScale(m1, xmul)
	m4 = X64New(vpnumber.X64Mul(x11, xmul), vpnumber.X64Mul(x12, xmul), vpnumber.X64Mul(x13, xmul), vpnumber.X64Mul(x21, xmul), vpnumber.X64Mul(x22, xmul), vpnumber.X64Mul(x23, xmul), vpnumber.X64Mul(x31, xmul), vpnumber.X64Mul(x32, xmul), vpnumber.X64Mul(x33, xmul))
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = X64DivScale(m3, xmul)
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

func invertableX64() *X64 {
	var ret X64

	for vpnumber.X64Abs(ret.Det()) < vpnumber.X64Const1 {
		for i := range ret {
			ret[i] = vpnumber.I64ToX64(rand.Int63n(10)) >> 2
		}
	}

	return &ret
}

func randomVecX64() *vpvec2.X64 {
	var ret vpvec2.X64

	for i := range ret {
		ret[i] = vpnumber.I64ToX64(rand.Int63n(10)) >> 2
	}

	return &ret
}

func TestX64Comp(t *testing.T) {
	m1 := invertableX64()
	m2 := X64Inv(m1)
	id := X64Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestX64Aff(t *testing.T) {
	p1 := vpnumber.F64ToX64(3.0)
	p2 := vpnumber.F64ToX64(4.0)
	t1 := vpnumber.F64ToX64(6.5)
	t2 := vpnumber.F64ToX64(8.5)

	v1 := vpvec3.X64New(p1, p2, vpnumber.X64Const1)
	vt := vpvec2.X64New(t1, t2)
	mt := X64Translation(vt)
	t.Logf("translation mat3x3 for %s is %s", vt.String(), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat3x3 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec3.X64New(p1+t1, p2+t2, vpnumber.X64Const1)
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

	mr := X64Rot(vpmath.X64ConstPi2)
	t.Logf("rotation mat3x3 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat3x3 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec3.X64New(-v1[1], v1[0], vpnumber.X64Const1)
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

func TestX64Rebase(t *testing.T) {
	cx := vpnumber.F64ToX64(1.1)
	cy := vpnumber.F64ToX64(1.2)
	m1 := invertableX64()
	var vo1 vpvec2.X64
	vx1 := vpvec2.X64AxisX()
	vy1 := vpvec2.X64AxisY()
	vp1 := vpvec2.X64Add(vx1, vy1)
	vo2 := randomVecX64()
	vx2 := vpvec2.X64Add(m1.GetCol(0).ToVec2(), vo2)
	vy2 := vpvec2.X64Add(m1.GetCol(1).ToVec2(), vo2)
	vp2 := vpvec2.X64Add(vo2, vpvec2.X64Add(vpvec2.X64MulScale(vpvec2.X64Sub(vx2, vo2), cx), vpvec2.X64MulScale(vpvec2.X64Sub(vy2, vo2), cy)))

	m2 := X64RebaseOXY(vo2, vx2, vy2)
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

	m2 = X64RebaseOXYP(vo2, vx2, vy2, vp2)
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

func TestX64JSON(t *testing.T) {
	m1 := invertableX64()
	m2 := X64Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for X64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X64, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X64")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkX64Add(b *testing.B) {
	mat := X64New(vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkX64Inv(b *testing.B) {
	mat := invertableX64()

	for i := 0; i < b.N; i++ {
		_ = X64Inv(mat)
	}
}
