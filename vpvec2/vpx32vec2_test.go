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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpvec2

import (
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestX32Math(t *testing.T) {
	var x1 = vpnumber.F32ToX32(3.0)
	var x2 = vpnumber.F32ToX32(-4.0)

	var x5 = vpnumber.F32ToX32(-4.5)
	var x6 = vpnumber.F32ToX32(6.0)

	var xmul = vpnumber.F32ToX32(10.0)
	var xsqmag = vpnumber.F32ToX32(25.0)
	var xlength = vpnumber.F32ToX32(5.0)

	var v1, v2, v3, v4 *X32
	var x vpnumber.X32

	v1 = X32New(x1, x2)
	if !v1.IsSimilar(v1) {
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

	v2 = X32New(x5, x6)
	v3 = X32Add(v1, v2)
	v4 = X32New(x1+x5, x2+x6)
	if !v3.IsSimilar(v4) {
		t.Error("Add error")
	}

	v3 = X32Sub(v1, v2)
	v4 = X32New(x1-x5, x2-x6)
	if !v3.IsSimilar(v4) {
		t.Error("Sub error")
	}

	v3 = X32Add(v1, X32Neg(v2))
	v4 = X32Sub(v1, v2)
	if !v3.IsSimilar(v4) {
		t.Error("Neg error")
	}

	v3 = X32MulScale(v1, xmul)
	v4 = X32New(vpnumber.X32Mul(x1, xmul), vpnumber.X32Mul(x2, xmul))
	if !v3.IsSimilar(v4) {
		t.Error("MulScale error")
	}

	v3 = X32DivScale(v3, xmul)
	if !v3.IsSimilar(v1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	v3.DivScale(0)

	x = v1.SqMag()
	if !vpnumber.X32IsSimilar(x, xsqmag) {
		t.Error("SqMag error", x, xsqmag)
	}

	x = v1.Length()
	if !vpnumber.X32IsSimilar(x, xlength) {
		t.Error("Length error", x, xlength)
	}

	v3 = X32Normalize(v1)
	x = v3.Length()
	if !vpnumber.X32IsSimilar(x, vpnumber.X32Const1) {
		t.Error("Normalize error", x)
	}

	dot1 := v1.Dot(v2)
	dot2 := vpnumber.X32Mul(x1, x5) + vpnumber.X32Mul(x2, x6)
	if !vpnumber.X32IsSimilar(dot1, dot2) {
		t.Error("Dot error")
	}
}

func TestX32JSON(t *testing.T) {
	m1 := X32New(vpnumber.I32ToX32(10), vpnumber.I32ToX32(20))
	m2 := X32New(vpnumber.X32Const1, vpnumber.X32Const0)

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for X32 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X32")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X32, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X32")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled vector is different from original")
	}
}

func BenchmarkX32Add(b *testing.B) {
	vec := X32New(vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Add(vec)
	}
}

func BenchmarkX32Normalize(b *testing.B) {
	vec := X32New(vpnumber.X32Const1, vpnumber.X32Const1)

	for i := 0; i < b.N; i++ {
		_ = vec.Normalize()
	}
}
