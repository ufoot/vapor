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

func TestF32Vec3Math(t *testing.T) {
	const f1 = 3.0
	const f2 = -4.0
	const f3 = 1.0

	const f5 = -4.5
	const f6 = 6.0
	const f7 = 2.0

	const fmul = 10.0
	const fsumsq = 26.0
	const flength = 5.099

	var v1, v2, v3, v4 *F32Vec3
	var f float32

	v1 = F32Vec3New(f1, f2, f3)
	if !F32Vec3IsSimilar(v1, v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2 = F32Vec3New(f5, f6, f7)
	v3 = F32Vec3Add(v1, v2)
	v4 = F32Vec3New(f1+f5, f2+f6, f3+f7)
	if !F32Vec3IsSimilar(v3, v4) {
		t.Error("Add error")
	}

	v3 = F32Vec3Sub(v1, v2)
	v4 = F32Vec3New(f1-f5, f2-f6, f3-f7)
	if !F32Vec3IsSimilar(v3, v4) {
		t.Error("Sub error")
	}

	v3 = F32Vec3MulScale(v1, fmul)
	v4 = F32Vec3New(f1*fmul, f2*fmul, f3*fmul)
	if !F32Vec3IsSimilar(v3, v4) {
		t.Error("MulScale error")
	}

	v3 = F32Vec3DivScale(v3, fmul)
	if !F32Vec3IsSimilar(v3, v1) {
		t.Error("DivScale error")
	}

	f = F32Vec3SumSq(v1)
	if !vpnumber.F32IsSimilar(f, fsumsq) {
		t.Error("SumSq error", f, fsumsq)
	}

	f = F32Vec3Length(v1)
	if !vpnumber.F32IsSimilar(f, flength) {
		t.Error("Length error", f, flength)
	}

	v3 = F32Vec3Normalize(v1)
	f = F32Vec3Length(v3)
	if f != vpnumber.F32Const1 {
		t.Error("Normalize error", f)
	}
}

func BenchmarkF32Vec3Add(b *testing.B) {
	vec := F32Vec3New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkF32Vec3Normalize(b *testing.B) {
	vec := F32Vec3New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
