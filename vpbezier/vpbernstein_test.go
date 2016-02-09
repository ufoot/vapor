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
	"testing"
)

func TestF32OneMinusT(t *testing.T) {
	x := vpnumber.F32Const1 / 10.0
	y := x * x * x * (vpnumber.F32Const1 - x) * (vpnumber.F32Const1 - x)
	z := f32TOneMinusT(3, 2, x)
	if !vpnumber.F32IsSimilar(y, z) {
		t.Errorf("OneMinusT problem %f != %f", y, z)
	}
}

func TestF32OneMinusTDerivative(t *testing.T) {
	const x0 = vpnumber.F32Const0
	const x1 = vpnumber.F32Const1
	xStep := (x1 - x0) / 10
	dX := (x1 - x0) / 1000

	for n := 1; n <= 4; n++ {
		for i := 0; i <= n; i++ {
			t.Logf("(n,i)=(%d,%d)", n, i)
			f := func(x float32) float32 { return f32TOneMinusT(i, n-i, x) }
			g := func(x float32) float32 { return f32TOneMinusTDerivative(i, n-i, x) }
			h := vpmath.F32DerivativeFunc(f, dX)
			for x := x0 + xStep; x <= x1-xStep; x += xStep {
				y := g(x)
				z := h(x)
				if vpnumber.F32IsSimilar(y, z) {
					t.Logf("x=%f : y=%f == z=%f", x, y, z)
				} else {
					t.Errorf("Derivative mismatch x=%f : y=%f != z=%f", x, y, z)
				}
			}
		}
	}
}

func TestBerstein(t *testing.T) {
	const x0 = vpnumber.F32Const0
	const x1 = vpnumber.F32Const1
	xStep := (x1 - x0) / 10

	for n := 1; n <= 4; n++ {
		for x := x0 + xStep; x <= x1-xStep; x += xStep {
			t.Logf("n%d x=%f", n, x)
			z := vpnumber.F32Const0
			for i := 0; i <= n; i++ {
				y := F32Bernstein(n, i, x)
				z += y
			}
			if !vpnumber.F32IsSimilar(z, vpnumber.F32Const1) {
				t.Errorf("Bernstein values should sum up to 1, got %f", z)
			}
		}
	}
}
