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
	"github.com/ufoot/vapor/vpnumber"
	"math/rand"
	"testing"
)

func TestX64Mat3Math(t *testing.T) {
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

	var m1, m2, m3, m4 *X64Mat3

	m1 = X64Mat3New(x11, x12, x13, x21, x22, x23, x31, x32, x33)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToX64()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToX64()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
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

	m2 = X64Mat3New(x51, x52, x53, x61, x62, x63, x71, x72, x73)
	m3 = X64Mat3Add(m1, m2)
	m4 = X64Mat3New(x11+x51, x12+x52, x13+x53, x21+x61, x22+x62, x23+x63, x31+x71, x32+x72, x33+x73)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = X64Mat3Sub(m1, m2)
	m4 = X64Mat3New(x11-x51, x12-x52, x13-x53, x21-x61, x22-x62, x23-x63, x31-x71, x32-x72, x33-x73)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = X64Mat3MulScale(m1, xmul)
	m4 = X64Mat3New(vpnumber.X64Mul(x11, xmul), vpnumber.X64Mul(x12, xmul), vpnumber.X64Mul(x13, xmul), vpnumber.X64Mul(x21, xmul), vpnumber.X64Mul(x22, xmul), vpnumber.X64Mul(x23, xmul), vpnumber.X64Mul(x31, xmul), vpnumber.X64Mul(x32, xmul), vpnumber.X64Mul(x33, xmul))
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = X64Mat3DivScale(m3, xmul)
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

func invertableX64Mat3() *X64Mat3 {
	var ret X64Mat3

	for vpnumber.X64Abs(ret.Det()) < vpnumber.X64Const1 {
		for i := range ret {
			ret[i] = vpnumber.I64ToX64(rand.Int63n(10)) >> 2
		}
	}

	return &ret
}

func TestX64Mat3Comp(t *testing.T) {
	m1 := invertableX64Mat3()
	m2 := X64Mat3Inv(m1)
	id := X64Mat3Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestX64Mat3JSON(t *testing.T) {
	m1 := invertableX64Mat3()
	m2 := X64Mat3Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for X64Mat3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64Mat3")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X64Mat3, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X64Mat3")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkX64Mat3Add(b *testing.B) {
	mat := X64Mat3New(vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkX64Mat3Inv(b *testing.B) {
	mat := invertableX64Mat3()

	for i := 0; i < b.N; i++ {
		_ = X64Mat3Inv(mat)
	}
}
