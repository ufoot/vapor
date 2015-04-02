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
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpvec3"
	"github.com/ufoot/vapor/vpvec4"
	"github.com/ufoot/vapor/vpnumber"
	"math/rand"
	"testing"
)

func TestX32Mat4x4Math(t *testing.T) {
	var x11 = vpnumber.F32ToX32(1.0)
	var x12 = vpnumber.F32ToX32(2.0)
	var x13 = vpnumber.F32ToX32(3.0)
	var x14 = vpnumber.F32ToX32(4.0)
	var x21 = vpnumber.F32ToX32(-1.0)
	var x22 = vpnumber.F32ToX32(-2.0)
	var x23 = vpnumber.F32ToX32(-3.0)
	var x24 = vpnumber.F32ToX32(-4.0)
	var x31 = vpnumber.F32ToX32(11.0)
	var x32 = vpnumber.F32ToX32(9.0)
	var x33 = vpnumber.F32ToX32(8)
	var x34 = vpnumber.F32ToX32(4.0)
	var x41 = vpnumber.F32ToX32(1.0)
	var x42 = vpnumber.F32ToX32(-2.0)
	var x43 = vpnumber.F32ToX32(3.0)
	var x44 = vpnumber.F32ToX32(4.0)

	var x51 = vpnumber.F32ToX32(-6.15)
	var x52 = vpnumber.F32ToX32(-7.25)
	var x53 = vpnumber.F32ToX32(-8.35)
	var x54 = vpnumber.F32ToX32(-9.45)
	var x61 = vpnumber.F32ToX32(6.4)
	var x62 = vpnumber.F32ToX32(7.3)
	var x63 = vpnumber.F32ToX32(8.2)
	var x64 = vpnumber.F32ToX32(6.1)
	var x71 = vpnumber.F32ToX32(2.4)
	var x72 = vpnumber.F32ToX32(2.3)
	var x73 = vpnumber.F32ToX32(8.2)
	var x74 = vpnumber.F32ToX32(9.1)
	var x81 = vpnumber.F32ToX32(-6.01)
	var x82 = vpnumber.F32ToX32(-7.02)
	var x83 = vpnumber.F32ToX32(-8.03)
	var x84 = vpnumber.F32ToX32(-9.04)

	var xmul = vpnumber.F32ToX32(10.0)

	var m1, m2, m3, m4 *X32Mat4x4

	m1 = X32Mat4x4New(x11, x12, x13, x14, x21, x22, x23, x24, x31, x32, x33, x34, x41, x42, x43, x44)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX64().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("F32 conversion error")
	}

	m2 = m1.ToF64().ToX32()
	if !m1.IsSimilar(m2) {
		t.Error("F64 conversion error")
	}

	m2 = X32Mat4x4New(x51, x52, x53, x54, x61, x62, x63, x64, x71, x72, x73, x74, x81, x82, x83, x84)
	m3 = X32Mat4x4Add(m1, m2)
	m4 = X32Mat4x4New(x11+x51, x12+x52, x13+x53, x14+x54, x21+x61, x22+x62, x23+x63, x24+x64, x31+x71, x32+x72, x33+x73, x34+x74, x41+x81, x42+x82, x43+x83, x44+x84)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = X32Mat4x4Sub(m1, m2)
	m4 = X32Mat4x4New(x11-x51, x12-x52, x13-x53, x14-x54, x21-x61, x22-x62, x23-x63, x24-x64, x31-x71, x32-x72, x33-x73, x34-x74, x41-x81, x42-x82, x43-x83, x44-x84)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = X32Mat4x4MulScale(m1, xmul)
	m4 = X32Mat4x4New(vpnumber.X32Mul(x11, xmul), vpnumber.X32Mul(x12, xmul), vpnumber.X32Mul(x13, xmul), vpnumber.X32Mul(x14, xmul), vpnumber.X32Mul(x21, xmul), vpnumber.X32Mul(x22, xmul), vpnumber.X32Mul(x23, xmul), vpnumber.X32Mul(x24, xmul), vpnumber.X32Mul(x31, xmul), vpnumber.X32Mul(x32, xmul), vpnumber.X32Mul(x33, xmul), vpnumber.X32Mul(x34, xmul), vpnumber.X32Mul(x41, xmul), vpnumber.X32Mul(x42, xmul), vpnumber.X32Mul(x43, xmul), vpnumber.X32Mul(x44, xmul))
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = X32Mat4x4DivScale(m3, xmul)
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

func invertableX32Mat4x4() *X32Mat4x4 {
	var ret X32Mat4x4

	for vpnumber.X32Abs(ret.Det()) < vpnumber.X32Const1 {
		for i := range ret {
			ret[i] = vpnumber.I32ToX32(rand.Int31n(10)) >> 3
		}
	}

	return &ret
}

func TestX32Mat4x4Comp(t *testing.T) {
	m1 := invertableX32Mat4x4()
	m2 := X32Mat4x4Inv(m1)
	id := X32Mat4x4Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestX32Mat4x4Aff(t *testing.T) {
	p1 := vpnumber.F32ToX32(3.0)
	p2 := vpnumber.F32ToX32(4.0)
	p3 := vpnumber.F32ToX32(5.0)
	t1 := vpnumber.F32ToX32(6.5)
	t2 := vpnumber.F32ToX32(8.5)
	t3 := vpnumber.F32ToX32(10.5)

	v1 := vpvec4.X32Vec4New(p1, p2, p3, vpnumber.X32Const1)
	vt := vpvec3.X32Vec3New(t1, t2, t3)
	mt := X32Mat4x4Trans(vt)
	t.Logf("translation mat4x4 for %s is %s", vt.String(), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec4.X32Vec4New(p1+t1, p2+t2, p3+t3, vpnumber.X32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(v1.ToVec3())
	v3pos := v1.ToVec3().Add(vt)
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir := mt.MulVecDir(v1.ToVec3())
	v3dir := v1.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr := X32Mat4x4RotX(vpmath.X32ConstPi2)
	t.Logf("rotation mat4x4 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec4.X32Vec4New(v1[0], -v1[2], v1[1], vpnumber.X32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 Z rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec3())
	v3pos = v3.ToVec3()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 Z rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec3())
	v3dir = v3.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 Z rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr = X32Mat4x4RotY(vpmath.X32ConstPi2)
	t.Logf("rotation mat4x4 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec4.X32Vec4New(v1[2], v1[1], -v1[0], vpnumber.X32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 Z rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec3())
	v3pos = v3.ToVec3()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 Z rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec3())
	v3dir = v3.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 Z rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}

	mr = X32Mat4x4RotZ(vpmath.X32ConstPi2)
	t.Logf("rotation mat4x4 for PI/2 is %s", mr.String())
	v2 = mr.MulVec(v1)
	t.Logf("mat4x4 MulVec %s * %s = %s", mr.String(), v1.String(), v2.String())
	v3 = vpvec4.X32Vec4New(-v1[1], v1[0], v1[2], vpnumber.X32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat4x4 Z rotation MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos = mr.MulVecPos(v1.ToVec3())
	v3pos = v3.ToVec3()
	if !v2pos.IsSimilar(v3pos) {
		t.Errorf("mat4x4 Z rotation MulVecPos error v2pos=%s v3pos=%s", v2pos.String(), v3pos.String())
	}
	v2dir = mr.MulVecDir(v1.ToVec3())
	v3dir = v3.ToVec3()
	if !v2dir.IsSimilar(v3dir) {
		t.Errorf("mat4x4 Z rotation MulVecDir error v2dir=%s v3dir=%s", v2dir.String(), v3dir.String())
	}
}

func TestX32Mat4x4JSON(t *testing.T) {
	m1 := invertableX32Mat4x4()
	m2 := X32Mat4x4Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for X32Mat4x4 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X32Mat4x4")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X32Mat4x4, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X32Mat4x4")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkX32Mat4x4Add(b *testing.B) {
	mat := X32Mat4x4New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkX32Mat4x4Inv(b *testing.B) {
	mat := invertableX32Mat4x4()

	for i := 0; i < b.N; i++ {
		_ = X32Mat4x4Inv(mat)
	}
}
