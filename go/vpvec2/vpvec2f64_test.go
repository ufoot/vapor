// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpvec2

import (
	"encoding/json"
	"github.com/ufoot/vapor/go/vpnumber"
	"testing"
)

func TestF64Math(t *testing.T) {
	const f1 = 3.0
	const f2 = -4.0

	const f5 = -4.5
	const f6 = 6.0

	const fmul = 10.0
	const fsqmag = 25.0
	const flength = 5.0

	var v1, v2, v3, v4 *F64
	var f float64

	v1 = F64New(f1, f2)
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

	v2 = F64New(f5, f6)
	v3 = F64Add(v1, v2)
	v4 = F64New(f1+f5, f2+f6)
	if !v3.IsSimilar(v4) {
		t.Error("Add error")
	}

	v3 = F64Sub(v1, v2)
	v4 = F64New(f1-f5, f2-f6)
	if !v3.IsSimilar(v4) {
		t.Error("Sub error")
	}

	v3 = F64Add(v1, F64Neg(v2))
	v4 = F64Sub(v1, v2)
	if !v3.IsSimilar(v4) {
		t.Error("Neg error")
	}

	v3 = F64MulScale(v1, fmul)
	v4 = F64New(f1*fmul, f2*fmul)
	if !v3.IsSimilar(v4) {
		t.Error("MulScale error")
	}

	v3 = F64DivScale(v3, fmul)
	if !v3.IsSimilar(v1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	v3.DivScale(0)

	f = v1.SqMag()
	if !vpnumber.F64IsSimilar(f, fsqmag) {
		t.Error("SqMag error", f, fsqmag)
	}

	f = v1.Length()
	if !vpnumber.F64IsSimilar(f, flength) {
		t.Error("Length error", f, flength)
	}

	v3 = F64Normalize(v1)
	f = v3.Length()
	if f != vpnumber.F64Const1 {
		t.Error("Normalize error", f)
	}

	v3 = F64Homogeneous(v1)
	f = v3[Size-1]
	if f != vpnumber.F64Const1 {
		t.Error("Homogeneous error", f)
	}

	f = 0.3
	v3 = F64Lerp(v1, v2, f)
	v4 = F64Add(F64MulScale(v1, vpnumber.F64Const1-f), F64MulScale(v2, f))
	if !v3.IsSimilar(v4) {
		t.Errorf("Lerp error v3=%s v4=%s", v3.String(), v4.String())
	}

	dot1 := v1.Dot(v2)
	dot2 := f1*f5 + f2*f6
	if !vpnumber.F64IsSimilar(dot1, dot2) {
		t.Error("Dot error")
	}
}

func TestF64JSON(t *testing.T) {
	m1 := F64New(0.1, 0.2)
	m2 := F64AxisX()

	var err error
	var jsonBuf []byte

	jsonBuf, err = json.Marshal(m1)
	if err == nil {
		t.Logf("encoded JSON for F64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64")
	}
	err = json.Unmarshal([]byte("nawak"), m2)
	if err == nil {
		t.Error("able to decode JSON for F64, but json is not correct")
	}
	err = json.Unmarshal(jsonBuf, m2)
	if err != nil {
		t.Error("unable to decode JSON for F64")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled vector is different from original")
	}
}

func BenchmarkF64Add(b *testing.B) {
	vec := F64New(vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkF64Normalize(b *testing.B) {
	vec := F64New(vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
