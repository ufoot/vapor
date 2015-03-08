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

func TestX32Vec3Math(t *testing.T) {
	var x1 = vpnumber.F32ToX32(3.0)
	var x2 = vpnumber.F32ToX32(-4.0)
	var x3 = vpnumber.F32ToX32(1.0)

	var x5 = vpnumber.F32ToX32(-4.5)
	var x6 = vpnumber.F32ToX32(6.0)
	var x7 = vpnumber.F32ToX32(2.0)

	var xmul = vpnumber.F32ToX32(10.0)
	var xsumsq = vpnumber.F32ToX32(26.0)
	var xlength = vpnumber.F32ToX32(5.099)

	var v1, v2, v3, v4 *X32Vec3
	var x vpnumber.X32

	v1 = X32Vec3New(x1, x2, x3)
	if !X32Vec3IsSimilar(v1, v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2 = X32Vec3New(x5, x6, x7)
	v3 = X32Vec3Add(v1, v2)
	v4 = X32Vec3New(x1+x5, x2+x6, x3+x7)
	if !X32Vec3IsSimilar(v3, v4) {
		t.Error("Add error")
	}

	v3 = X32Vec3Sub(v1, v2)
	v4 = X32Vec3New(x1-x5, x2-x6, x3-x7)
	if !X32Vec3IsSimilar(v3, v4) {
		t.Error("Sub error")
	}

	v3 = X32Vec3MulScale(v1, xmul)
	v4 = X32Vec3New(vpnumber.X32Mul(x1, xmul), vpnumber.X32Mul(x2, xmul), vpnumber.X32Mul(x3, xmul))
	if !X32Vec3IsSimilar(v3, v4) {
		t.Error("MulScale error")
	}

	v3 = X32Vec3DivScale(v3, xmul)
	if !X32Vec3IsSimilar(v3, v1) {
		t.Error("DivScale error")
	}

	x = X32Vec3SumSq(v1)
	if !vpnumber.X32IsSimilar(x, xsumsq) {
		t.Error("SumSq error", x, xsumsq)
	}

	x = X32Vec3Length(v1)
	if !vpnumber.X32IsSimilar(x, xlength) {
		t.Error("Length error", x, xlength)
	}

	v3 = X32Vec3Normalize(v1)
	x = X32Vec3Length(v3)
	if !vpnumber.X32IsSimilar(x, vpnumber.X32Const1) {
		t.Error("Normalize error", x)
	}
}

func BenchmarkX32Vec3Add(b *testing.B) {
	vec := X32Vec3New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkX32Vec3Normalize(b *testing.B) {
	vec := X32Vec3New(vpnumber.X32Const1, vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
