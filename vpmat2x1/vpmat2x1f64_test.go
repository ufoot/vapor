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

package vpmat2x1

import (
	"github.com/ufoot/vapor/vpnumber"
	"math"
	"math/rand"
	"testing"
)

func TestF64Math(t *testing.T) {
	const f11 = 3.0
	const f21 = -40.0

	const f51 = -4.5
	const f61 = 6.0

	const fmul = 10.0

	var m1, m2, m3, m4 *F64

	m1 = F64New(f11, f21)
	if !m1.IsSimilar(m1) {
		t.Error("IsSimilar does not detect equality")
	}

	m2 = m1.ToX32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("X32 conversion error")
	}

	m2 = m1.ToX64().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("X64 conversion error")
	}

	m2 = m1.ToF32().ToF64()
	if !m1.IsSimilar(m2) {
		t.Error("F32 conversion error")
	}

	m2 = F64New(f51, f61)
	m3 = F64Add(m1, m2)
	m4 = F64New(f11+f51, f21+f61)
	if !m3.IsSimilar(m4) {
		t.Error("Add error")
	}

	m3 = F64Sub(m1, m2)
	m4 = F64New(f11-f51, f21-f61)
	if !m3.IsSimilar(m4) {
		t.Error("Sub error")
	}

	m3 = F64MulScale(m1, fmul)
	m4 = F64New(f11*fmul, f21*fmul)
	if !m3.IsSimilar(m4) {
		t.Error("MulScale error")
	}

	m3 = F64DivScale(m3, fmul)
	if !m3.IsSimilar(m1) {
		t.Error("DivScale error")
	}

	// Yes, div by 0 is valid, it should raise no error.
	// The results are inconsistent, but no big deal.
	// While it should theorically raise an error, the consequence
	// it 3d math is usually : glitch in display. This is less
	// disastrous than a floating point exception.
	m3.DivScale(0)
}

func invertableF64() *F64 {
	var ret F64

	for math.Abs(ret.Det()) < 0.25 {
		for i := range ret {
			ret[i] = rand.Float64()
		}
	}

	return &ret
}

func TestF64Comp(t *testing.T) {
	m1 := invertableF64()
	m2 := F64Inv(m1)
	id := F64Identity()

	m2.MulComp(m1)
	if m2.IsSimilar(id) {
		t.Logf("multiplicating matrix by its inverse return something similar to identity m2=%s", m2.String())
	} else {
		t.Errorf("multiplicating matrix by its inverse does not return identity m1=%s m2=%s", m1.String(), m2.String())
	}
}

func TestF64Aff(t *testing.T) {
	const p1 = 3.0
	const t1 = 6.0

	mt := F64Trans(t1)
	t.Logf("translation mat2 for %f is %s", p1, mt.String())
	v2pos := mt.MulVecPos(p1)
	v3pos := float64(p1 + t1)
	if !vpnumber.F64IsSimilar(v2pos, v3pos) {
		t.Errorf("mat2 MulVecPos error v2pos=%f v3pos=%f", v2pos, v3pos)
	}
	v2dir := mt.MulVecDir(p1)
	v3dir := float64(p1)
	if !vpnumber.F64IsSimilar(v2dir, v3dir) {
		t.Errorf("mat2 MulVecDir error v2dir=%f v3dir=%f", v2dir, v3dir)
	}
}

func TestF64JSON(t *testing.T) {
	m1 := invertableF64()
	m2 := F64Identity()

	var err error
	var jsonBuf []byte

	jsonBuf, err = m1.MarshalJSON()
	if err == nil {
		t.Logf("encoded JSON for F64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for F64")
	}
	err = m2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for F64, but json is not correct")
	}
	err = m2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for F64")
	}
	if !m1.IsSimilar(m2) {
		t.Error("unmarshalled matrix is different from original")
	}
}

func BenchmarkF64Add(b *testing.B) {
	mat := F64New(vpnumber.F64Const1, vpnumber.F64Const1)

	for i := 0; i < b.N; i++ {
		_ = mat.Add(mat)
	}
}

func BenchmarkF64Inv(b *testing.B) {
	mat := invertableF64()

	for i := 0; i < b.N; i++ {
		_ = F64Inv(mat)
	}
}
