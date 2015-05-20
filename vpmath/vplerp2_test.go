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

// X32 X64 tests to be written...
