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

func TestF32Lerp(t *testing.T) {
	var f float32
	const f1 float32 = -2.0
	const f2 float32 = 8.0
	const beta float32 = 0.7
	const lerp float32 = 5.0

	f = F32Lerp(f1, f2, beta)
	if !vpnumber.F32IsSimilar(f, lerp) {
		t.Errorf("bad lerp, got %f should be %f", f, lerp)
	}
	f = F32Lerp(f1, f2, -vpnumber.F32Const1)
	if f != f1 {
		t.Errorf("bad lerp on negative beta, got %f should be %f", f, f1)
	}
	f = F32Lerp(f1, f2, vpnumber.F32Const1+vpnumber.F32Const1)
	if f != f2 {
		t.Errorf("bad lerp on beta>1, got %f should be %f", f, f2)
	}
}

func TestF64Lerp(t *testing.T) {
	var f float64
	const f1 float64 = -2.0
	const f2 float64 = 8.0
	const beta float64 = 0.7
	const lerp float64 = 5.0

	f = F64Lerp(f1, f2, beta)
	if !vpnumber.F64IsSimilar(f, lerp) {
		t.Errorf("bad lerp, got %f should be %f", f, lerp)
	}
	f = F64Lerp(f1, f2, -vpnumber.F64Const1)
	if f != f1 {
		t.Errorf("bad lerp on negative beta, got %f should be %f", f, f1)
	}
	f = F64Lerp(f1, f2, vpnumber.F64Const1+vpnumber.F64Const1)
	if f != f2 {
		t.Errorf("bad lerp on beta>1, got %f should be %f", f, f2)
	}
}

func TestX32Lerp(t *testing.T) {
	var x vpnumber.X32
	var x1 = vpnumber.F32ToX32(-2.0)
	var x2 = vpnumber.F32ToX32(8.0)
	var beta = vpnumber.F32ToX32(0.7)
	var lerp = vpnumber.F32ToX32(5.0)

	x = X32Lerp(x1, x2, beta)
	if !vpnumber.X32IsSimilar(x, lerp) {
		t.Errorf("bad lerp, got %s should be %s", x.String(), lerp.String())
	}
	x = X32Lerp(x1, x2, -vpnumber.X32Const1)
	if x != x1 {
		t.Errorf("bad lerp on negative beta, got %x should be %x", x, x1)
	}
	x = X32Lerp(x1, x2, vpnumber.X32Const1+vpnumber.X32Const1)
	if x != x2 {
		t.Errorf("bad lerp on beta>1, got %x should be %x", x, x2)
	}
}

func TestX64Lerp(t *testing.T) {
	var x vpnumber.X64
	var x1 = vpnumber.F64ToX64(-2.0)
	var x2 = vpnumber.F64ToX64(8.0)
	var beta = vpnumber.F64ToX64(0.7)
	var lerp = vpnumber.F64ToX64(5.0)

	x = X64Lerp(x1, x2, beta)
	if !vpnumber.X64IsSimilar(x, lerp) {
		t.Errorf("bad lerp, got %s should be %s", x.String(), lerp.String())
	}
	x = X64Lerp(x1, x2, -vpnumber.X64Const1)
	if x != x1 {
		t.Errorf("bad lerp on negative beta, got %x should be %x", x, x1)
	}
	x = X64Lerp(x1, x2, vpnumber.X64Const1+vpnumber.X64Const1)
	if x != x2 {
		t.Errorf("bad lerp on beta>1, got %x should be %x", x, x2)
	}
}
