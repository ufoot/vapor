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

func TestX64Vec4Math(t *testing.T) {
	var x1 = vpnumber.F64ToX64(3.0)
	var x2 = vpnumber.F64ToX64(-4.0)
	var x3 = vpnumber.F64ToX64(1.0)
	var x4 = vpnumber.F64ToX64(10.0)

	var x5 = vpnumber.F64ToX64(-4.5)
	var x6 = vpnumber.F64ToX64(6.0)
	var x7 = vpnumber.F64ToX64(2.0)
	var x8 = vpnumber.F64ToX64(-30.0)

	var xmul = vpnumber.F64ToX64(10.0)
	var xsqmag = vpnumber.F64ToX64(126.0)
	var xlength = vpnumber.F64ToX64(11.224972)

	var v1, v2, v3, v4 *X64Vec4
	var x vpnumber.X64

	v1 = X64Vec4New(x1, x2, x3, x4)
	if !X64Vec4IsSimilar(v1, v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2 = v1.ToI32().ToX64()
	if !v1.IsSimilar(v2) {
		t.Error("I32 conversion error")
	}

	v2 = v1.ToI64().ToX64()
	if !v1.IsSimilar(v2) {
		t.Error("I64 conversion error")
	}

	v2 = v1.ToX32().ToX64()
	if !v1.IsSimilar(v2) {
		t.Error("X32 conversion error")
	}

	v2 = v1.ToF32().ToX64()
	if !v1.IsSimilar(v2) {
		t.Error("F32 conversion error")
	}

	v2 = v1.ToF64().ToX64()
	if !v1.IsSimilar(v2) {
		t.Error("F64 conversion error")
	}

	v2 = X64Vec4New(x5, x6, x7, x8)
	v3 = X64Vec4Add(v1, v2)
	v4 = X64Vec4New(x1+x5, x2+x6, x3+x7, x4+x8)
	if !X64Vec4IsSimilar(v3, v4) {
		t.Error("Add error")
	}

	v3 = X64Vec4Sub(v1, v2)
	v4 = X64Vec4New(x1-x5, x2-x6, x3-x7, x4-x8)
	if !X64Vec4IsSimilar(v3, v4) {
		t.Error("Sub error")
	}

	v3 = X64Vec4Add(v1, X64Vec4Neg(v2))
	v4 = X64Vec4Sub(v1, v2)
	if !X64Vec4IsSimilar(v3, v4) {
		t.Error("Neg error")
	}

	v3 = X64Vec4MulScale(v1, xmul)
	v4 = X64Vec4New(vpnumber.X64Mul(x1, xmul), vpnumber.X64Mul(x2, xmul), vpnumber.X64Mul(x3, xmul), vpnumber.X64Mul(x4, xmul))
	if !X64Vec4IsSimilar(v3, v4) {
		t.Error("MulScale error")
	}

	v3 = X64Vec4DivScale(v3, xmul)
	if !X64Vec4IsSimilar(v3, v1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	v3.DivScale(0)

	x = X64Vec4SqMag(v1)
	if !vpnumber.X64IsSimilar(x, xsqmag) {
		t.Error("SqMag error", x, xsqmag)
	}

	x = X64Vec4Length(v1)
	if !vpnumber.X64IsSimilar(x, xlength) {
		t.Error("Length error", x, xlength)
	}

	v3 = X64Vec4Normalize(v1)
	x = X64Vec4Length(v3)
	if !vpnumber.X64IsSimilar(x, vpnumber.X64Const1) {
		t.Error("Normalize error", x)
	}
}

func BenchmarkX64Vec4Add(b *testing.B) {
	vec := X64Vec4New(vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkX64Vec4Normalize(b *testing.B) {
	vec := X64Vec4New(vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1, vpnumber.X64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
