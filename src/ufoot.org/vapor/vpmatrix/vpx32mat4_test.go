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

package vpmatrix

import (
	"testing"
	"ufoot.org/vapor/vpnumber"
)

func TestX32Mat4Math(t *testing.T) {
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

	var m1, m2, m3, m4 *X32Mat4

	m1 = X32Mat4New(x11, x12, x13, x14, x21, x22, x23, x24, x31, x32, x33, x34, x41, x42, x43, x44)
	if !X32Mat4IsSimilar(m1, m1) {
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

	m2 = X32Mat4New(x51, x52, x53, x54, x61, x62, x63, x64, x71, x72, x73, x74, x81, x82, x83, x84)
	m3 = X32Mat4Add(m1, m2)
	m4 = X32Mat4New(x11+x51, x12+x52, x13+x53, x14+x54, x21+x61, x22+x62, x23+x63, x24+x64, x31+x71, x32+x72, x33+x73, x34+x74, x41+x81, x42+x82, x43+x83, x44+x84)
	if !X32Mat4IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = X32Mat4Sub(m1, m2)
	m4 = X32Mat4New(x11-x51, x12-x52, x13-x53, x14-x54, x21-x61, x22-x62, x23-x63, x24-x64, x31-x71, x32-x72, x33-x73, x34-x74, x41-x81, x42-x82, x43-x83, x44-x84)
	if !X32Mat4IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = X32Mat4MulScale(m1, xmul)
	m4 = X32Mat4New(vpnumber.X32Mul(x11, xmul), vpnumber.X32Mul(x12, xmul), vpnumber.X32Mul(x13, xmul), vpnumber.X32Mul(x14, xmul), vpnumber.X32Mul(x21, xmul), vpnumber.X32Mul(x22, xmul), vpnumber.X32Mul(x23, xmul), vpnumber.X32Mul(x24, xmul), vpnumber.X32Mul(x31, xmul), vpnumber.X32Mul(x32, xmul), vpnumber.X32Mul(x33, xmul), vpnumber.X32Mul(x34, xmul), vpnumber.X32Mul(x41, xmul), vpnumber.X32Mul(x42, xmul), vpnumber.X32Mul(x43, xmul), vpnumber.X32Mul(x44, xmul))
	if !X32Mat4IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = X32Mat4DivScale(m3, xmul)
	if !X32Mat4IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func BenchmarkX32Mat4Add(b *testing.B) {
	mat := X32Mat4New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}
