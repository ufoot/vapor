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

func TestF32Lerp2(t *testing.T) {
	p := [2][2]float32{{-1, 10}, {1, 20}}
	const u float32 = 0.3
	const v float32 = 0.7
	const lerp float32 = 8.98
	var f float32
	var fCheck float32

	f = F32Lerp2(p, u, v)
	if !vpnumber.F32IsSimilar(f, lerp) {
		t.Errorf("bad lerp2, got %f should be %f", f, lerp)
	}
	f = F32Lerp2(p, -vpnumber.F32Const1, v)
	fCheck = F32Lerp(p[0][0], p[0][1], v)
	if !vpnumber.F32IsSimilar(f, fCheck) {
		t.Errorf("bad lerp2 on negative u, got %f should be %f", f, fCheck)
	}
	f = F32Lerp2(p, vpnumber.F32Const1+vpnumber.F32Const1, v)
	fCheck = F32Lerp(p[1][0], p[1][1], v)
	if !vpnumber.F32IsSimilar(f, fCheck) {
		t.Errorf("bad lerp2 on u>1, got %f should be %f", f, fCheck)
	}
}

func TestF64Lerp2(t *testing.T) {
	p := [2][2]float64{{-1, 10}, {1, 20}}
	const u float64 = 0.3
	const v float64 = 0.7
	const lerp float64 = 8.98
	var f float64
	var fCheck float64

	f = F64Lerp2(p, u, v)
	if !vpnumber.F64IsSimilar(f, lerp) {
		t.Errorf("bad lerp2, got %f should be %f", f, lerp)
	}
	f = F64Lerp2(p, -vpnumber.F64Const1, v)
	fCheck = F64Lerp(p[0][0], p[0][1], v)
	if !vpnumber.F64IsSimilar(f, fCheck) {
		t.Errorf("bad lerp2 on negative u, got %f should be %f", f, fCheck)
	}
	f = F64Lerp2(p, vpnumber.F64Const1+vpnumber.F64Const1, v)
	fCheck = F64Lerp(p[1][0], p[1][1], v)
	if !vpnumber.F64IsSimilar(f, fCheck) {
		t.Errorf("bad lerp2 on u>1, got %f should be %f", f, fCheck)
	}
}

func TestX32Lerp2(t *testing.T) {
	p := [2][2]vpnumber.X32{{vpnumber.I32ToX32(-1), vpnumber.I32ToX32(1)}, {vpnumber.I32ToX32(4), vpnumber.I32ToX32(10)}}
	u := vpnumber.F32ToX32(0.25)
	v := vpnumber.F32ToX32(0.75)
	lerp := vpnumber.F32ToX32(2.5)
	var x vpnumber.X32
	var xCheck vpnumber.X32

	x = X32Lerp2(p, u, v)
	if !vpnumber.X32IsSimilar(x, lerp) {
		t.Errorf("bad lerp2, got %s should be %s", x.String(), lerp.String())
	}
	x = X32Lerp2(p, -vpnumber.X32Const1, v)
	xCheck = X32Lerp(p[0][0], p[0][1], v)
	if !vpnumber.X32IsSimilar(x, xCheck) {
		t.Errorf("bad lerp2 on negative u, got %s should be %s", x.String(), xCheck.String())
	}
	x = X32Lerp2(p, vpnumber.X32Const1+vpnumber.X32Const1, v)
	xCheck = X32Lerp(p[1][0], p[1][1], v)
	if !vpnumber.X32IsSimilar(x, xCheck) {
		t.Errorf("bad lerp2 on u>1, got %s should be %s", x.String(), xCheck.String())
	}
}

func TestX64Lerp2(t *testing.T) {
	p := [2][2]vpnumber.X64{{vpnumber.I64ToX64(-1), vpnumber.I64ToX64(1)}, {vpnumber.I64ToX64(4), vpnumber.I64ToX64(10)}}
	u := vpnumber.F64ToX64(0.25)
	v := vpnumber.F64ToX64(0.75)
	lerp := vpnumber.F64ToX64(2.5)
	var x vpnumber.X64
	var xCheck vpnumber.X64

	x = X64Lerp2(p, u, v)
	if !vpnumber.X64IsSimilar(x, lerp) {
		t.Errorf("bad lerp2, got %s should be %s", x.String(), lerp.String())
	}
	x = X64Lerp2(p, -vpnumber.X64Const1, v)
	xCheck = X64Lerp(p[0][0], p[0][1], v)
	if !vpnumber.X64IsSimilar(x, xCheck) {
		t.Errorf("bad lerp2 on negative u, got %s should be %s", x.String(), xCheck.String())
	}
	x = X64Lerp2(p, vpnumber.X64Const1+vpnumber.X64Const1, v)
	xCheck = X64Lerp(p[1][0], p[1][1], v)
	if !vpnumber.X64IsSimilar(x, xCheck) {
		t.Errorf("bad lerp2 on u>1, got %s should be %s", x.String(), xCheck.String())
	}
}
