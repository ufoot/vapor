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

func TestF64Mat2Math(t *testing.T) {
	const f1 = 3.0
	const f2 = -4.0

	const f5 = -4.5
	const f6 = 6.0

	const fmul = 10.0

	var m1, m2, m3, m4 *F64Mat2

	m1 = F64Mat2New(f1, f2)
	if !F64Mat2IsSimilar(m1, m1) {
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

	m2 = F64Mat2New(f5, f6)
	m3 = F64Mat2Add(m1, m2)
	m4 = F64Mat2New(f1+f5, f2+f6)
	if !F64Mat2IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = F64Mat2Sub(m1, m2)
	m4 = F64Mat2New(f1-f5, f2-f6)
	if !F64Mat2IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = F64Mat2MulScale(m1, fmul)
	m4 = F64Mat2New(f1*fmul, f2*fmul)
	if !F64Mat2IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = F64Mat2DivScale(m3, fmul)
	if !F64Mat2IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func BenchmarkF64Mat2Add(b *testing.B) {
	mat := F64Mat2New(vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

