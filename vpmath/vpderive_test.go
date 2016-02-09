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

package vpmath

import (
	"github.com/ufoot/vapor/vpnumber"
	"testing"
)

func TestF32Derivative(t *testing.T) {
	const t0 = vpnumber.F32Const0
	const t1 = vpnumber.F32Const1
	dt := (t1 - t0) / 100
	step := (t1 - t0) / 10

	f := func(t float32) float32 {
		return t * t
	}

	g := func(t float32) float32 {
		return 2.0 * t
	}

	h := func(t float32) float32 {
		return 2.0
	}

	fp := F32DerivativeFunc(f, dt)
	fpp := F32DerivativeFunc(fp, dt)

	for tp := t0; tp <= t1; tp += step {
		yf := f(tp)
		t.Logf("f(%f)=%f", tp, yf)

		yg := g(tp)
		yfp := F32Derivative(f, tp, dt)

		if vpnumber.F32IsSimilar(yg, yfp) {
			t.Logf("g(%f)=%f f'(%f)=%f", tp, yg, tp, yfp)
		} else {
			t.Errorf("1st order derive problem g(%f)=%f f'(%f)=%f", tp, yg, tp, yfp)
		}

		yh := h(tp)
		yfpp := fpp(tp)

		if vpnumber.F32IsSimilar(yh, yfpp) {
			t.Logf("h(%f)=%f f''(%f)=%f", tp, yh, tp, yfpp)
		} else {
			t.Errorf("2nd order derive problem h(%f)=%f f''(%f)=%f", tp, yh, tp, yfpp)
		}
	}
}

func TestF64Derivative(t *testing.T) {
	const t0 = vpnumber.F64Const0
	const t1 = vpnumber.F64Const1
	dt := (t1 - t0) / 10000
	step := (t1 - t0) / 10

	f := func(t float64) float64 {
		return t * t
	}

	g := func(t float64) float64 {
		return 2.0 * t
	}

	h := func(t float64) float64 {
		return 2.0
	}

	fp := F64DerivativeFunc(f, dt)
	fpp := F64DerivativeFunc(fp, dt)

	for tp := t0; tp <= t1; tp += step {
		yf := f(tp)
		t.Logf("f(%f)=%f", tp, yf)

		yg := g(tp)
		yfp := F64Derivative(f, tp, dt)

		if vpnumber.F64IsSimilar(yg, yfp) {
			t.Logf("g(%f)=%f f'(%f)=%f", tp, yg, tp, yfp)
		} else {
			t.Errorf("1st order derive problem g(%f)=%f f'(%f)=%f", tp, yg, tp, yfp)
		}

		yh := h(tp)
		yfpp := fpp(tp)

		if vpnumber.F64IsSimilar(yh, yfpp) {
			t.Logf("h(%f)=%f f''(%f)=%f", tp, yh, tp, yfpp)
		} else {
			t.Errorf("2nd order derive problem h(%f)=%f f''(%f)=%f", tp, yh, tp, yfpp)
		}
	}
}

func TestX32Derivative(t *testing.T) {
	const t0 = vpnumber.X32Const0
	const t1 = vpnumber.X32Const1
	dt := (t1 - t0) / 10
	step := (t1 - t0) / 10

	f := func(t vpnumber.X32) vpnumber.X32 {
		return t
	}

	g := func(t vpnumber.X32) vpnumber.X32 {
		return vpnumber.X32Const1
	}

	h := func(t vpnumber.X32) vpnumber.X32 {
		return vpnumber.X32Const0
	}

	fp := X32DerivativeFunc(f, dt)
	fpp := X32DerivativeFunc(fp, dt)

	for tp := t0; tp <= t1; tp += step {
		yf := f(tp)
		t.Logf("f(%s)=%s", tp.String(), yf.String())

		yg := g(tp)
		yfp := X32Derivative(f, tp, dt)

		if vpnumber.X32IsSimilar(yg, yfp) {
			t.Logf("g(%s)=%s f'(%s)=%s", tp.String(), yg.String(), tp.String(), yfp.String())
		} else {
			t.Errorf("1st order derive problem g(%s)=%s f'(%s)=%s", tp.String(), yg.String(), tp.String(), yfp.String())
		}

		yh := h(tp)
		yfpp := fpp(tp)

		if vpnumber.X32IsSimilar(yh, yfpp) {
			t.Logf("h(%s)=%s f''(%s)=%s", tp.String(), yh.String(), tp.String(), yfpp.String())
		} else {
			t.Errorf("2nd order derive problem h(%s)=%s f''(%s)=%s", tp.String(), yh.String(), tp.String(), yfpp.String())
		}
	}
}

func TestX64Derivative(t *testing.T) {
	const t0 = vpnumber.X64Const0
	const t1 = vpnumber.X64Const1
	dt := (t1 - t0) / 100
	step := (t1 - t0) / 10

	f := func(t vpnumber.X64) vpnumber.X64 {
		return t
	}

	g := func(t vpnumber.X64) vpnumber.X64 {
		return vpnumber.X64Const1
	}

	h := func(t vpnumber.X64) vpnumber.X64 {
		return vpnumber.X64Const0
	}

	fp := X64DerivativeFunc(f, dt)
	fpp := X64DerivativeFunc(fp, dt)

	for tp := t0; tp <= t1; tp += step {
		yf := f(tp)
		t.Logf("f(%s)=%s", tp.String(), yf.String())

		yg := g(tp)
		yfp := X64Derivative(f, tp, dt)

		if vpnumber.X64IsSimilar(yg, yfp) {
			t.Logf("g(%s)=%s f'(%s)=%s", tp.String(), yg.String(), tp.String(), yfp.String())
		} else {
			t.Errorf("1st order derive problem g(%s)=%s f'(%s)=%s", tp.String(), yg.String(), tp.String(), yfp.String())
		}

		yh := h(tp)
		yfpp := fpp(tp)

		if vpnumber.X64IsSimilar(yh, yfpp) {
			t.Logf("h(%s)=%s f''(%s)=%s", tp.String(), yh.String(), tp.String(), yfpp.String())
		} else {
			t.Errorf("2nd order derive problem h(%s)=%s f''(%s)=%s", tp.String(), yh.String(), tp.String(), yfpp.String())
		}
	}
}
