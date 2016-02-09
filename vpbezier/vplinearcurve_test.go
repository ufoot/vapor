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

package vpbezier

import (
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
	"testing"
)

const p0x float32 = 2
const p0y float32 = 4
const p0z float32 = 8
const p1x float32 = 3
const p1y float32 = -1
const p1z float32 = 11

func TestF32LinearCurve1d(t *testing.T) {
	alpha0 := vpnumber.F32Const0
	alpha1 := vpnumber.F32Const1
	step := (alpha1 - alpha0) / 5.0
	dt := (alpha1 - alpha0) / 1000.0
	p0 := p0x
	p1 := p1x

	f := func(a float32) float32 {
		ret, _ := F32LinearCurve1d(p0, p1, a)
		return ret
	}
	fp := vpmath.F32DerivativeFunc(f, dt)

	for alpha := alpha0; alpha <= alpha1; alpha += step {
		v, dv := F32LinearCurve1d(p0, p1, alpha)
		t.Logf("alpha=%f v=%f dv=%f", alpha, v, dv)
		dv2 := fp(alpha)
		if alpha-alpha0 >= step && alpha1-alpha >= step {
			if !vpnumber.F32IsSimilar(dv, dv2) {
				t.Errorf("derivative mismatch alpha=%f dv=%f dv2=%f", alpha, dv, dv2)
			}
		}
	}
}

func TestF32LinearCurve2d(t *testing.T) {
	alpha0 := vpnumber.F32Const0
	alpha1 := vpnumber.F32Const1
	step := (alpha1 - alpha0) / 5.0
	dt := (alpha1 - alpha0) / 1000.0
	p0 := vpvec2.F32New(p0x, p0y)
	p1 := vpvec2.F32New(p1x, p1y)

	f := func(a float32) float32 {
		ret, _ := F32LinearCurve2d(p0, p1, a)
		return ret[1]
	}
	fp := vpmath.F32DerivativeFunc(f, dt)

	for alpha := alpha0; alpha <= alpha1; alpha += step {
		v, dv := F32LinearCurve2d(p0, p1, alpha)
		t.Logf("alpha=%f v=%s dv=%s", alpha, v.String(), dv.String())
		dv2 := fp(alpha)
		if alpha-alpha0 >= step && alpha1-alpha >= step {
			if !vpnumber.F32IsSimilar(dv[1], dv2) {
				t.Errorf("derivative mismatch alpha=%f dv=%f dv2=%f", alpha, dv[1], dv2)
			}
		}
	}
}

func TestF32LinearCurve3d(t *testing.T) {
	alpha0 := vpnumber.F32Const0
	alpha1 := vpnumber.F32Const1
	step := (alpha1 - alpha0) / 5.0
	dt := (alpha1 - alpha0) / 1000.0
	p0 := vpvec3.F32New(p0x, p0y, p0z)
	p1 := vpvec3.F32New(p1x, p1y, p1z)

	f := func(a float32) float32 {
		ret, _ := F32LinearCurve3d(p0, p1, a)
		return ret[2]
	}
	fp := vpmath.F32DerivativeFunc(f, dt)

	for alpha := alpha0; alpha <= alpha1; alpha += step {
		v, dv := F32LinearCurve3d(p0, p1, alpha)
		t.Logf("alpha=%f v=%s dv=%s", alpha, v.String(), dv.String())
		dv2 := fp(alpha)
		if alpha-alpha0 >= step && alpha1-alpha >= step {
			if !vpnumber.F32IsSimilar(dv[2], dv2) {
				t.Errorf("derivative mismatch alpha=%f dv=%f dv2=%f", alpha, dv[2], dv2)
			}
		}
	}
}
