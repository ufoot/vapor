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
	"github.com/ufoot/vapor/go/vpnumber"
	"testing"
)

func TestF32Unlerp(t *testing.T) {
	var f float32
	const f1 float32 = -2.0
	const f2 float32 = 8.0
	const beta float32 = 0.7
	const lerp float32 = 5.0

	f = F32Unlerp(f1, f2, lerp)
	if !vpnumber.F32IsSimilar(f, beta) {
		t.Errorf("bad unlerp, got %f should be %f", f, beta)
	}
	f = F32Unlerp(f1, f2, f1-vpnumber.F32Const1)
	if f != vpnumber.F32Const0 {
		t.Errorf("bad unlerp on negative beta, got %f should be %f", f, vpnumber.F32Const0)
	}
	f = F32Unlerp(f1, f2, f2+vpnumber.F32Const1+vpnumber.F32Const1)
	if f != vpnumber.F32Const1 {
		t.Errorf("bad unlerp on beta>1, got %f should be %f", f, vpnumber.F32Const1)
	}
	f = F32Unlerp(-f1, -f2, -f1+vpnumber.F32Const1)
	if f != vpnumber.F32Const0 {
		t.Errorf("bad unlerp on negative beta, got %f should be %f", f, vpnumber.F32Const0)
	}
	f = F32Unlerp(-f1, -f2, -f2-vpnumber.F32Const1+vpnumber.F32Const1)
	if f != vpnumber.F32Const1 {
		t.Errorf("bad unlerp on beta>1, got %f should be %f", f, vpnumber.F32Const1)
	}
}

func TestF64Unlerp(t *testing.T) {
	var f float64
	const f1 float64 = -2.0
	const f2 float64 = 8.0
	const beta float64 = 0.7
	const lerp float64 = 5.0

	f = F64Unlerp(f1, f2, lerp)
	if !vpnumber.F64IsSimilar(f, beta) {
		t.Errorf("bad unlerp, got %f should be %f", f, beta)
	}
	f = F64Unlerp(f1, f2, f1-vpnumber.F64Const1)
	if f != vpnumber.F64Const0 {
		t.Errorf("bad unlerp on negative beta, got %f should be %f", f, vpnumber.F64Const0)
	}
	f = F64Unlerp(f1, f2, f2+vpnumber.F64Const1+vpnumber.F64Const1)
	if f != vpnumber.F64Const1 {
		t.Errorf("bad unlerp on beta>1, got %f should be %f", f, vpnumber.F64Const0)
	}
	f = F64Unlerp(-f1, -f2, -f1+vpnumber.F64Const1)
	if f != vpnumber.F64Const0 {
		t.Errorf("bad unlerp on negative beta, got %f should be %f", f, vpnumber.F64Const0)
	}
	f = F64Unlerp(-f1, -f2, -f2-vpnumber.F64Const1+vpnumber.F64Const1)
	if f != vpnumber.F64Const1 {
		t.Errorf("bad unlerp on beta>1, got %f should be %f", f, vpnumber.F64Const0)
	}
}

func TestX32Unlerp(t *testing.T) {
	var x vpnumber.X32
	var x1 = vpnumber.F32ToX32(-2.0)
	var x2 = vpnumber.F32ToX32(8.0)
	var beta = vpnumber.F32ToX32(0.7)
	var lerp = vpnumber.F32ToX32(5.0)

	x = X32Unlerp(x1, x2, lerp)
	if !vpnumber.X32IsSimilar(x, beta) {
		t.Errorf("bad unlerp, got %x should be %x", x, lerp)
	}
	x = X32Unlerp(x1, x2, x1-vpnumber.X32Const1)
	if x != vpnumber.X32Const0 {
		t.Errorf("bad unlerp on negative beta, got %x should be %x", x, vpnumber.X32Const0)
	}
	x = X32Unlerp(x1, x2, x2+vpnumber.X32Const1+vpnumber.X32Const1)
	if x != vpnumber.X32Const1 {
		t.Errorf("bad unlerp on beta>1, got %x should be %x", x, vpnumber.X32Const1)
	}
	x = X32Unlerp(-x1, -x2, -x1+vpnumber.X32Const1)
	if x != vpnumber.X32Const0 {
		t.Errorf("bad unlerp on negative beta, got %x should be %x", x, vpnumber.X32Const0)
	}
	x = X32Unlerp(-x1, -x2, -x2-vpnumber.X32Const1+vpnumber.X32Const1)
	if x != vpnumber.X32Const1 {
		t.Errorf("bad unlerp on beta>1, got %x should be %x", x, vpnumber.X32Const1)
	}
}

func TestX64Unlerp(t *testing.T) {
	var x vpnumber.X64
	var x1 = vpnumber.F64ToX64(-2.0)
	var x2 = vpnumber.F64ToX64(8.0)
	var beta = vpnumber.F64ToX64(0.7)
	var lerp = vpnumber.F64ToX64(5.0)

	x = X64Unlerp(x1, x2, lerp)
	if !vpnumber.X64IsSimilar(x, beta) {
		t.Errorf("bad unlerp, got %x should be %x", x, lerp)
	}
	x = X64Unlerp(x1, x2, x1-vpnumber.X64Const1)
	if x != vpnumber.X64Const0 {
		t.Errorf("bad unlerp on negative beta, got %x should be %x", x, vpnumber.X64Const0)
	}
	x = X64Unlerp(x1, x2, x2+vpnumber.X64Const1+vpnumber.X64Const1)
	if x != vpnumber.X64Const1 {
		t.Errorf("bad unlerp on beta>1, got %x should be %x", x, vpnumber.X64Const1)
	}
	x = X64Unlerp(-x1, -x2, -x1+vpnumber.X64Const1)
	if x != vpnumber.X64Const0 {
		t.Errorf("bad unlerp on negative beta, got %x should be %x", x, vpnumber.X64Const0)
	}
	x = X64Unlerp(-x1, -x2, -x2-vpnumber.X64Const1+vpnumber.X64Const1)
	if x != vpnumber.X64Const1 {
		t.Errorf("bad unlerp on beta>1, got %x should be %x", x, vpnumber.X64Const1)
	}
}
