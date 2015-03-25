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

package vpmatrix4

import (
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF32Vec4Math(t *testing.T) {
	const f1 = 3.0
	const f2 = -4.0
	const f3 = 1.0
	const f4 = 10.0

	const f5 = -4.5
	const f6 = 6.0
	const f7 = 2.0
	const f8 = -30.0

	const fmul = 10.0
	const fsqmag = 126.0
	const flength = 11.225

	var v1, v2, v3, v4 *F32Vec4
	var f float32

	v1 = F32Vec4New(f1, f2, f3, f4)
	if !F32Vec4IsSimilar(v1, v1) {
		t.Error("IsSimilar does not detect equality")
	}

	v2 = v1.ToI32().ToF32()
	if !v1.IsSimilar(v2) {
		t.Error("I32 conversion error")
	}

	v2 = v1.ToI64().ToF32()
	if !v1.IsSimilar(v2) {
		t.Error("I64 conversion error")
	}

	v2 = v1.ToX32().ToF32()
	if !v1.IsSimilar(v2) {
		t.Error("X32 conversion error")
	}

	v2 = v1.ToX64().ToF32()
	if !v1.IsSimilar(v2) {
		t.Error("X64 conversion error")
	}

	v2 = v1.ToF64().ToF32()
	if !v1.IsSimilar(v2) {
		t.Error("F64 conversion error")
	}

	v2 = F32Vec4New(f5, f6, f7, f8)
	v3 = F32Vec4Add(v1, v2)
	v4 = F32Vec4New(f1+f5, f2+f6, f3+f7, f4+f8)
	if !F32Vec4IsSimilar(v3, v4) {
		t.Error("Add error")
	}

	v3 = F32Vec4Sub(v1, v2)
	v4 = F32Vec4New(f1-f5, f2-f6, f3-f7, f4-f8)
	if !F32Vec4IsSimilar(v3, v4) {
		t.Error("Sub error")
	}

	v3 = F32Vec4Add(v1, F32Vec4Neg(v2))
	v4 = F32Vec4Sub(v1, v2)
	if !F32Vec4IsSimilar(v3, v4) {
		t.Error("Neg error")
	}

	v3 = F32Vec4MulScale(v1, fmul)
	v4 = F32Vec4New(f1*fmul, f2*fmul, f3*fmul, f4*fmul)
	if !F32Vec4IsSimilar(v3, v4) {
		t.Error("MulScale error")
	}

	v3 = F32Vec4DivScale(v3, fmul)
	if !F32Vec4IsSimilar(v3, v1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	v3.DivScale(0)

	f = F32Vec4SqMag(v1)
	if !vpnumber.F32IsSimilar(f, fsqmag) {
		t.Error("SqMag error", f, fsqmag)
	}

	f = F32Vec4Length(v1)
	if !vpnumber.F32IsSimilar(f, flength) {
		t.Error("Length error", f, flength)
	}

	v3 = F32Vec4Normalize(v1)
	f = F32Vec4Length(v3)
	if f != vpnumber.F32Const1 {
		t.Error("Normalize error", f)
	}

	v3 = F32Vec4Dot(v1, v2)
	v4 = F32Vec4New(f1*f5, f2*f6, f3*f7, f4*f8)
	if !F32Vec4IsSimilar(v3, v4) {
		t.Error("Dot error")
	}
}

func TestF32Vec4JSON(t *testing.T) {
	m1 := F32Vec4New(0.1, 0.2, 0.3, 0.4)
	m2 := F32Vec4New(1.0, 0.0, 0.0, 0.0)

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F32Vec4 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F32Vec4")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F32Vec4, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F32Vec4")
	}
	if !F32Vec4IsSimilar(m1, m2) {
		t.Error("unmarshalled vector is different from original")
	}
}

func BenchmarkF32Vec4Add(b *testing.B) {
	vec := F32Vec4New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkF32Vec4Normalize(b *testing.B) {
	vec := F32Vec4New(vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
