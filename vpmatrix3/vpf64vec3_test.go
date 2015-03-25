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

package vpmatrix3

import (
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF64Vec3Math(t *testing.T) {
	const f1 = 3.0
	const f2 = -4.0
	const f3 = 1.0

	const f5 = -4.5
	const f6 = 6.0
	const f7 = 2.0

	const fmul = 10.0
	const fsqmag = 26.0
	const flength = 5.099019

	var v1, v2, v3, v4 *F64Vec3
	var f float64

	v1 = F64Vec3New(f1, f2, f3)
	if !v1.IsSimilar(v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2 = v1.ToI32().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("I32 conversion error")
	}

	v2 = v1.ToI64().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("I64 conversion error")
	}

	v2 = v1.ToX32().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("X32 conversion error")
	}

	v2 = v1.ToX64().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("X64 conversion error")
	}

	v2 = v1.ToF32().ToF64()
	if !v1.IsSimilar(v2) {
		t.Error("F32 conversion error")
	}

	v2 = F64Vec3New(f5, f6, f7)
	v3 = F64Vec3Add(v1, v2)
	v4 = F64Vec3New(f1+f5, f2+f6, f3+f7)
	if !v3.IsSimilar(v4) {
		t.Error("Add error")
	}

	v3 = F64Vec3Sub(v1, v2)
	v4 = F64Vec3New(f1-f5, f2-f6, f3-f7)
	if !v3.IsSimilar(v4) {
		t.Error("Sub error")
	}

	v3 = F64Vec3Add(v1, F64Vec3Neg(v2))
	v4 = F64Vec3Sub(v1, v2)
	if !v3.IsSimilar(v4) {
		t.Error("Neg error")
	}

	v3 = F64Vec3MulScale(v1, fmul)
	v4 = F64Vec3New(f1*fmul, f2*fmul, f3*fmul)
	if !v3.IsSimilar(v4) {
		t.Error("MulScale error")
	}

	v3 = F64Vec3DivScale(v3, fmul)
	if !v3.IsSimilar(v1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	v3.DivScale(0)

	f = F64Vec3SqMag(v1)
	if !vpnumber.F64IsSimilar(f, fsqmag) {
		t.Error("SqMag error", f, fsqmag)
	}

	f = F64Vec3Length(v1)
	if !vpnumber.F64IsSimilar(f, flength) {
		t.Error("Length error", f, flength)
	}

	v3 = F64Vec3Normalize(v1)
	f = F64Vec3Length(v3)
	if f != vpnumber.F64Const1 {
		t.Error("Normalize error", f)
	}

	v3 = F64Vec3Dot(v1, v2)
	v4 = F64Vec3New(f1*f5, f2*f6, f3*f7)
	if !v3.IsSimilar(v4) {
		t.Error("Dot error")
	}
}

func TestF64Vec3JSON(t *testing.T) {
	m1 := F64Vec3New(0.1, 0.2, 0.3)
	m2 := F64Vec3New(1.0, 0.0, 0.0)

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F64Vec3 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64Vec3")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F64Vec3, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F64Vec3")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled vector is different from original")
	}
}

func BenchmarkF64Vec3Add(b *testing.B) {
	vec := F64Vec3New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkF64Vec3Normalize(b *testing.B) {
	vec := F64Vec3New(vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
