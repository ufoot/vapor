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
	"github.com/ufoot/vapor/vpnumber"
	"math/rand"
	"testing"
)

func TestX32Mat2Math(t *testing.T) {
	var x11 = vpnumber.F32ToX32(3.0)
	var x12 = vpnumber.F32ToX32(-11.0)
	var x21 = vpnumber.F32ToX32(9.0)
	var x22 = vpnumber.F32ToX32(-12.0)

	var x51 = vpnumber.F32ToX32(-4.5)
	var x52 = vpnumber.F32ToX32(6.0)
	var x61 = vpnumber.F32ToX32(1.1)
	var x62 = vpnumber.F32ToX32(-4.0)

	var xmul = vpnumber.F32ToX32(10.0)

	var m1, m2, m3, m4 *X32Mat2

	m1 = X32Mat2New(x11, x12, x21, x22)
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

	m2 = X32Mat2New(x51, x52, x61, x62)
	m3 = X32Mat2Add(m1, m2)
	m4 = X32Mat2New(x11+x51, x12+x52, x21+x61, x22+x62)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = X32Mat2Sub(m1, m2)
	m4 = X32Mat2New(x11-x51, x12-x52, x21-x61, x22-x62)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = X32Mat2MulScale(m1, xmul)
	m4 = X32Mat2New(vpnumber.X32Mul(x11, xmul), vpnumber.X32Mul(x12, xmul), vpnumber.X32Mul(x21, xmul), vpnumber.X32Mul(x22, xmul))
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = X32Mat2DivScale(m3, xmul)
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

func invertableX32Mat2() *X32Mat2 {
	var ret X32Mat2

	for vpnumber.X32Abs(ret.Det()) < vpnumber.X32Const1 {
		for i := range ret {
			ret[i] = vpnumber.I32ToX32(rand.Int31n(10)) >> 1
		}
	}

	return &ret
}

func TestX32Mat2Comp(t *testing.T) {
	m1 := invertableX32Mat2()
	m2 := X32Mat2Inv(m1)
	id := X32Mat2Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestX32Mat2JSON(t *testing.T) {
	m1 := invertableX32Mat2()
	m2 := X32Mat2Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for X32Mat2 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X32Mat2")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X32Mat2, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X32Mat2")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkX32Mat2Add(b *testing.B) {
	mat := X32Mat2New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkX32Mat2Inv(b *testing.B) {
	mat := invertableX32Mat2()

	for i := 0; i < b.N; i++ {
		_ = X32Mat2Inv(mat)
	}
}
