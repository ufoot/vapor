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

func TestX64Mat4Math(t *testing.T) {
	var x1 = vpnumber.F64ToX64(3.0)
	var x2 = vpnumber.F64ToX64(-4.0)
	var x3 = vpnumber.F64ToX64(1.0)
	var x4 = vpnumber.F64ToX64(10.0)

	var x5 = vpnumber.F64ToX64(-4.5)
	var x6 = vpnumber.F64ToX64(6.0)
	var x7 = vpnumber.F64ToX64(2.0)
	var x8 = vpnumber.F64ToX64(-30.0)

	var xmul = vpnumber.F64ToX64(10.0)

	var m1, m2, m3, m4 *X64Mat4

	m1 = X64Mat4New(x1, x2, x3, x4)
	if !X64Mat4IsSimilar(m1, m1) {
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

	m2 = X64Mat4New(x5, x6, x7, x8)
	m3 = X64Mat4Add(m1, m2)
	m4 = X64Mat4New(x1+x5, x2+x6, x3+x7, x4+x8)
	if !X64Mat4IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = X64Mat4Sub(m1, m2)
	m4 = X64Mat4New(x1-x5, x2-x6, x3-x7, x4-x8)
	if !X64Mat4IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = X64Mat4MulScale(m1, xmul)
	m4 = X64Mat4New(vpnumber.X64Mul(x1, xmul), vpnumber.X64Mul(x2, xmul), vpnumber.X64Mul(x3, xmul), vpnumber.X64Mul(x4, xmul))
	if !X64Mat4IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = X64Mat4DivScale(m3, xmul)
	if !X64Mat4IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func BenchmarkX64Mat4Add(b *testing.B) {
	mat := X64Mat4New(vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}
