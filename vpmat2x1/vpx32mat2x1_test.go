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

package vpmat2x1

import (
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"math/rand"
	"testing"
)

func TestX32Mat2x1Math(t *testing.T) {
	var x11 = vpnumber.F32ToX32(3.0)
	var x12 = vpnumber.F32ToX32(-11.0)
	var x21 = vpnumber.F32ToX32(9.0)
	var x22 = vpnumber.F32ToX32(-12.0)

	var x51 = vpnumber.F32ToX32(-4.5)
	var x52 = vpnumber.F32ToX32(6.0)
	var x61 = vpnumber.F32ToX32(1.1)
	var x62 = vpnumber.F32ToX32(-4.0)

	var xmul = vpnumber.F32ToX32(10.0)

	var m1, m2, m3, m4 *X32Mat2x1

	m1 = X32Mat2x1New(x11, x12, x21, x22)
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

	m2 = X32Mat2x1New(x51, x52, x61, x62)
	m3 = X32Mat2x1Add(m1, m2)
	m4 = X32Mat2x1New(x11+x51, x12+x52, x21+x61, x22+x62)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = X32Mat2x1Sub(m1, m2)
	m4 = X32Mat2x1New(x11-x51, x12-x52, x21-x61, x22-x62)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = X32Mat2x1MulScale(m1, xmul)
	m4 = X32Mat2x1New(vpnumber.X32Mul(x11, xmul), vpnumber.X32Mul(x12, xmul), vpnumber.X32Mul(x21, xmul), vpnumber.X32Mul(x22, xmul))
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = X32Mat2x1DivScale(m3, xmul)
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

func invertableX32Mat2x1() *X32Mat2x1 {
	var ret X32Mat2x1

	for vpnumber.X32Abs(ret.Det()) < vpnumber.X32Const1 {
		for i := range ret {
			ret[i] = vpnumber.I32ToX32(rand.Int31n(10)) >> 1
		}
	}

	return &ret
}

func TestX32Mat2x1Comp(t *testing.T) {
	m1 := invertableX32Mat2x1()
	m2 := X32Mat2x1Inv(m1)
	id := X32Mat2x1Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestX32Mat2x1Aff(t *testing.T) {
	p1 := vpnumber.F32ToX32(3.0)
	t1 := vpnumber.F32ToX32(6.0)

	v1 := vpvec2.X32Vec2New(p1, vpnumber.X32Const1)
	mt := X32Mat2x1Trans(t1)
	t.Logf("translation mat2 for %f is %s", vpnumber.X32ToF32(p1), mt.String())
	v2 := mt.MulVec(v1)
	t.Logf("mat2 MulVec %s * %s = %s", mt.String(), v1.String(), v2.String())
	v3 := vpvec2.X32Vec2New(p1+t1, vpnumber.X32Const1)
	if !v2.IsSimilar(v3) {
		t.Errorf("mat2 MulVec error v2=%s v3=%s", v2.String(), v3.String())
	}
	v2pos := mt.MulVecPos(p1)
	v3pos := p1 + t1
	if !vpnumber.X32IsSimilar(v2pos, v3pos) {
		t.Errorf("mat2 MulVecPos error v2pos=%f v3pos=%f", vpnumber.X32ToF32(v2pos), vpnumber.X32ToF32(v3pos))
	}
	v2dir := mt.MulVecDir(p1)
	v3dir := p1
	if !vpnumber.X32IsSimilar(v2dir, v3dir) {
		t.Errorf("mat2 MulVecDir error v2dir=%f v3dir=%f", vpnumber.X32ToF32(v2dir), vpnumber.X32ToF32(v3dir))
	}
}

func TestX32Mat2x1JSON(t *testing.T) {
	m1 := invertableX32Mat2x1()
	m2 := X32Mat2x1Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for X32Mat2x1 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X32Mat2x1")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X32Mat2x1, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X32Mat2x1")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkX32Mat2x1Add(b *testing.B) {
	mat := X32Mat2x1New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkX32Mat2x1Inv(b *testing.B) {
	mat := invertableX32Mat2x1()

	for i := 0; i < b.N; i++ {
		_ = X32Mat2x1Inv(mat)
	}
}
