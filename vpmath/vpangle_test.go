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
	"math"
	"testing"
)

const f32DegIn float32 = -11.0 * F32Const90
const f32DegOut float32 = 5.0 * F32Const90
const f64DegIn float64 = -11.0 * F64Const90
const f64DegOut float64 = 5.0 * F64Const90
const x32DegIn vpnumber.X32 = -11 * X32Const90
const x32DegOut vpnumber.X32 = 5 * X32Const90
const x64DegIn vpnumber.X64 = -11 * X64Const90
const x64DegOut vpnumber.X64 = 5 * X64Const90
const f32RadIn float32 = -11.0 * math.Pi / 4.0
const f32RadOut float32 = 5.0 * math.Pi / 4.0
const f64RadIn float64 = -11.0 * math.Pi / 4.0
const f64RadOut float64 = 5.0 * math.Pi / 4.0
const x32RadIn vpnumber.X32 = -11 * X32ConstPi4
const x32RadOut vpnumber.X32 = 5 * X32ConstPi4
const x64RadIn vpnumber.X64 = -11 * X64ConstPi4
const x64RadOut vpnumber.X64 = 5 * X64ConstPi4

func testDegMod(t *testing.T) {
	f32 := F32DegMod(f32DegIn)
	if !vpnumber.F32IsSimilar(f32, f32DegOut) {
		t.Error("error doing modulo on degrees (F32)")
	}
	f64 := F64DegMod(f64DegIn)
	if !vpnumber.F64IsSimilar(f64, f64DegOut) {
		t.Error("error doing modulo on degrees (F64)")
	}
	x32 := X32DegMod(x32DegIn)
	if !vpnumber.X32IsSimilar(x32, x32DegOut) {
		t.Error("error doing modulo on degrees (X32)")
	}
	x64 := X64DegMod(x64DegIn)
	if !vpnumber.X64IsSimilar(x64, x64DegOut) {
		t.Error("error doing modulo on degrees (X64)")
	}
}

func testRadMod(t *testing.T) {
	f32 := F32RadMod(f32RadIn)
	if !vpnumber.F32IsSimilar(f32, f32RadOut) {
		t.Error("error doing modulo on radians (F32)")
	}
	f64 := F64RadMod(f64RadIn)
	if !vpnumber.F64IsSimilar(f64, f64RadOut) {
		t.Error("error doing modulo on radians (F64)")
	}
	x32 := X32RadMod(x32RadIn)
	if !vpnumber.X32IsSimilar(x32, x32RadOut) {
		t.Error("error doing modulo on radians (X32)")
	}
	x64 := X64RadMod(x64RadIn)
	if !vpnumber.X64IsSimilar(x64, x64RadOut) {
		t.Error("error doing modulo on radians (X64)")
	}
}

func TestDegToRad(t *testing.T) {
	f32 := F32DegToRad(F32Const90)
	if !vpnumber.F32IsSimilar(f32, float32(math.Pi/2.0)) {
		t.Error("error converting from degrees to radians (F32)")
	}
	f64 := F64DegToRad(F64Const90)
	if !vpnumber.F64IsSimilar(f64, math.Pi/2.0) {
		t.Error("error converting from degrees to radians (F64)")
	}
	x32 := X32DegToRad(X32Const90)
	if !vpnumber.X32IsSimilar(x32, X32ConstPi2) {
		t.Error("error converting from degrees to radians (X32)")
	}
	x64 := X64DegToRad(X64Const90)
	if !vpnumber.X64IsSimilar(x64, X64ConstPi2) {
		t.Error("error converting from degrees to radians (X64)")
	}
}

func TestRadToDeg(t *testing.T) {
	f32 := F32RadToDeg(float32(math.Pi / 2.0))
	if !vpnumber.F32IsSimilar(f32, F32Const90) {
		t.Error("error converting from degrees to radians (F32)")
	}
	f64 := F64RadToDeg(math.Pi / 2.0)
	if !vpnumber.F64IsSimilar(f64, F64Const90) {
		t.Error("error converting from degrees to radians (F64)")
	}
	x32 := X32RadToDeg(X32ConstPi2)
	if !vpnumber.X32IsSimilar(x32, X32Const90) {
		t.Error("error converting from degrees to radians (X32)")
	}
	x64 := X64RadToDeg(X64ConstPi2)
	if !vpnumber.X64IsSimilar(x64, X64Const90) {
		t.Error("error converting from degrees to radians (X64)")
	}
}
