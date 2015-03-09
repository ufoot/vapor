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

func TestF64Vec2Math(t *testing.T) {
	const f1 = 3.0
	const f2 = -4.0

	const f5 = -4.5
	const f6 = 6.0

	const fmul = 10.0
	const fsumsq = 25.0
	const flength = 5.0

	var v1, v2, v3, v4 *F64Vec2
	var f float64

	v1 = F64Vec2New(f1, f2)
	if !F64Vec2IsSimilar(v1, v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2=v1.ToI32().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("I32 conversion error")
	}	

	v2=v1.ToI64().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("I64 conversion error")
	}	

	v2=v1.ToX32().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("X32 conversion error")
	}	

	v2=v1.ToX64().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("X64 conversion error")
	}	

	v2=v1.ToF32().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("F32 conversion error")
	}	
	
	v2 = F64Vec2New(f5, f6)
	v3 = F64Vec2Add(v1, v2)
	v4 = F64Vec2New(f1+f5, f2+f6)
	if !F64Vec2IsSimilar(v3, v4) {
		t.Error("Add error")
	}

	v3 = F64Vec2Sub(v1, v2)
	v4 = F64Vec2New(f1-f5, f2-f6)
	if !F64Vec2IsSimilar(v3, v4) {
		t.Error("Sub error")
	}

	v3 = F64Vec2MulScale(v1, fmul)
	v4 = F64Vec2New(f1*fmul, f2*fmul)
	if !F64Vec2IsSimilar(v3, v4) {
		t.Error("MulScale error")
	}

	v3 = F64Vec2DivScale(v3, fmul)
	if !F64Vec2IsSimilar(v3, v1) {
		t.Error("DivScale error")
	}

	f = F64Vec2SumSq(v1)
	if !vpnumber.F64IsSimilar(f, fsumsq) {
		t.Error("SumSq error", f, fsumsq)
	}

	f = F64Vec2Length(v1)
	if !vpnumber.F64IsSimilar(f, flength) {
		t.Error("Length error", f, flength)
	}

	v3 = F64Vec2Normalize(v1)
	f = F64Vec2Length(v3)
	if f != vpnumber.F64Const1 {
		t.Error("Normalize error", f)
	}
}

func BenchmarkF64Vec2Add(b *testing.B) {
	vec := F64Vec2New(vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkF64Vec2Normalize(b *testing.B) {
	vec := F64Vec2New(vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
