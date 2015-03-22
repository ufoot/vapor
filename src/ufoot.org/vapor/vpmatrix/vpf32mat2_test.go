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

func TestF32Mat2Basic(t *testing.T) {
	const f11 = 3.0
	const f12 = 3.0
	const f21 = -8.0
	const f22 = -4.0

	const f51 = -4.5
	const f52 = -4.1
	const f61 = 6.0
	const f62 = 6.6

	const fmul = 10.0

	var m1, m2, m3, m4 *F32Mat2

	m1 = F32Mat2New(f11, f12, f21, f22)
	if !F32Mat2IsSimilar(m1, m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToI32().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("I32 conversion error")
	}

	m2 = m1.ToI64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("I64 conversion error")
	}

	m2 = m1.ToX32().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF64().ToF32()
	if !m1.IsSimilar(m2) {
		t.Error("F64 conversion error")
	}

	m2 = F32Mat2New(f51, f52, f61, f62)
	m3 = F32Mat2Add(m1, m2)
	m4 = F32Mat2New(f11+f51, f12+f52, f21+f61, f22+f62)
	if !F32Mat2IsSimilar(m3, m4) {
		t.Error("Add error")
	}

	m3 = F32Mat2Sub(m1, m2)
	m4 = F32Mat2New(f11-f51, f12-f52, f21-f61, f22-f62)
	if !F32Mat2IsSimilar(m3, m4) {
		t.Error("Sub error")
	}

	m3 = F32Mat2MulScale(m1, fmul)
	m4 = F32Mat2New(f11*fmul, f12*fmul, f21*fmul, f22*fmul)
	if !F32Mat2IsSimilar(m3, m4) {
		t.Error("MulScale error")
	}

	m3 = F32Mat2DivScale(m3, fmul)
	if !F32Mat2IsSimilar(m3, m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func invertableF32Mat2() *F32Mat2 {
	var ret F32Mat2

	for vpnumber.F32IsSimilar(ret.Det(), vpnumber.F32Const0) {
		for i := range ret {
			ret[i] = rand.Float32()
		}
	}

	return &ret
}

func TestF32Mat2Comp(t *testing.T) {
	m1 := invertableF32Mat2()
	m2 := F32Mat2Inv(m1)
	id := F32Mat2Identity()

	m2.MulComp(m1)
	if !F32Mat2IsSimilar(m2, id) {
		t.Error("multiplicating matrix by its inverse does not return identity")
	}
}

func BenchmarkF32Mat2Add(b *testing.B) {
	mat := F32Mat2New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF32Mat2Inv(b *testing.B) {
	mat := invertableF32Mat2()

	for i := 0; i < b.N; i++ {
		_ = F32Mat2Inv(mat)
	}
}
