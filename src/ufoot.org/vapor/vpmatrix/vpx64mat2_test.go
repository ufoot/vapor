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
	"math/rand"
	"testing"
	"ufoot.org/vapor/vpnumber"
)

func TestX64Mat2Basic(t *testing.T) {
	var x11 = vpnumber.F64ToX64(3.0)
	var x12 = vpnumber.F64ToX64(-43.0)
	var x21 = vpnumber.F64ToX64(31.0)
	var x22 = vpnumber.F64ToX64(-12.0)

	var x51 = vpnumber.F64ToX64(-4.5)
	var x52 = vpnumber.F64ToX64(60.0)
	var x61 = vpnumber.F64ToX64(1.1)
	var x62 = vpnumber.F64ToX64(-4.0)

	var xmul = vpnumber.F64ToX64(10.0)

	var m1, m2, m3, m4 *X64Mat2

	m1 = X64Mat2New(x11, x12, x21, x22)
	if !X64Mat2IsSimilar(m1, m1) {
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

	m2 = X64Mat2New(x51, x52, x61, x62)
	m3 = X64Mat2Add(m1, m2)
	m4 = X64Mat2New(x11+x51, x12+x52, x21+x61, x22+x62)
	if !X64Mat2IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = X64Mat2Sub(m1, m2)
	m4 = X64Mat2New(x11-x51, x12-x52, x21-x61, x22-x62)
	if !X64Mat2IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = X64Mat2MulScale(m1, xmul)
	m4 = X64Mat2New(vpnumber.X64Mul(x11, xmul), vpnumber.X64Mul(x12, xmul), vpnumber.X64Mul(x21, xmul), vpnumber.X64Mul(x22, xmul))
	if !X64Mat2IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = X64Mat2DivScale(m3, xmul)
	if !X64Mat2IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func invertableX64Mat2() *X64Mat2 {
	var ret X64Mat2

	for vpnumber.X64IsSimilar(ret.Det(), vpnumber.X64Const0) {
		for i := range ret {
			ret[i] = vpnumber.I64ToX64(rand.Int63n(10))
		}
	}

	return &ret
}

func TestX64Mat2Comp(t *testing.T) {
	m1 := invertableX64Mat2()
	m2 := X64Mat2Inv(m1)
	id := X64Mat2Identity()

	m2.MulComp(m1)
	if !X64Mat2IsSimilar(m2, id) {
		t.Error("multiplicating matrix by its inverse does not return identity")
	}
}

func BenchmarkX64Mat2Add(b *testing.B) {
	mat := X64Mat2New(vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkX64Mat2Inv(b *testing.B) {
	mat := invertableX64Mat2()

	for i := 0; i < b.N; i++ {
		_ = X64Mat2Inv(mat)
	}
}
