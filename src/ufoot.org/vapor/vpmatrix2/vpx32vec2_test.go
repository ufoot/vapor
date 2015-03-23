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
	"testing"
	"ufoot.org/vapor/vpnumber"
)

func TestX32Vec2Math(t *testing.T) {
	var x1 = vpnumber.F32ToX32(3.0)
	var x2 = vpnumber.F32ToX32(-4.0)

	var x5 = vpnumber.F32ToX32(-4.5)
	var x6 = vpnumber.F32ToX32(6.0)

	var xmul = vpnumber.F32ToX32(10.0)
	var xsqmag = vpnumber.F32ToX32(25.0)
	var xlength = vpnumber.F32ToX32(5.0)

	var v1, v2, v3, v4 *X32Vec2
	var x vpnumber.X32

	v1 = X32Vec2New(x1, x2)
	if !X32Vec2IsSimilar(v1, v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2 = v1.ToI32().ToX32()
	if !v1.IsSimilar(v2) {
		t.Error("I32 conversion error")
	}

	v2 = v1.ToI64().ToX32()
	if !v1.IsSimilar(v2) {
		t.Error("I64 conversion error")
	}

	v2 = v1.ToX64().ToX32()
	if !v1.IsSimilar(v2) {
		t.Error("X64 conversion error")
	}

	v2 = v1.ToF32().ToX32()
	if !v1.IsSimilar(v2) {
		t.Error("F32 conversion error")
	}

	v2 = v1.ToF64().ToX32()
	if !v1.IsSimilar(v2) {
		t.Error("F64 conversion error")
	}

	v2 = X32Vec2New(x5, x6)
	v3 = X32Vec2Add(v1, v2)
	v4 = X32Vec2New(x1+x5, x2+x6)
	if !X32Vec2IsSimilar(v3, v4) {
		t.Error("Add error")
	}

	v3 = X32Vec2Sub(v1, v2)
	v4 = X32Vec2New(x1-x5, x2-x6)
	if !X32Vec2IsSimilar(v3, v4) {
		t.Error("Sub error")
	}

	v3 = X32Vec2Add(v1, X32Vec2Neg(v2))
	v4 = X32Vec2Sub(v1, v2)
	if !X32Vec2IsSimilar(v3, v4) {
		t.Error("Neg error")
	}

	v3 = X32Vec2MulScale(v1, xmul)
	v4 = X32Vec2New(vpnumber.X32Mul(x1, xmul), vpnumber.X32Mul(x2, xmul))
	if !X32Vec2IsSimilar(v3, v4) {
		t.Error("MulScale error")
	}

	v3 = X32Vec2DivScale(v3, xmul)
	if !X32Vec2IsSimilar(v3, v1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	v3.DivScale(0)

	x = X32Vec2SqMag(v1)
	if !vpnumber.X32IsSimilar(x, xsqmag) {
		t.Error("SqMag error", x, xsqmag)
	}

	x = X32Vec2Length(v1)
	if !vpnumber.X32IsSimilar(x, xlength) {
		t.Error("Length error", x, xlength)
	}

	v3 = X32Vec2Normalize(v1)
	x = X32Vec2Length(v3)
	if !vpnumber.X32IsSimilar(x, vpnumber.X32Const1) {
		t.Error("Normalize error", x)
	}
}

func BenchmarkX32Vec2Add(b *testing.B) {
	vec := X32Vec2New(vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkX32Vec2Normalize(b *testing.B) {
	vec := X32Vec2New(vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
