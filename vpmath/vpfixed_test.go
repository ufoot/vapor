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
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpmath

import (
	"math"
	"testing"
	"github.com/ufoot/vapor/vpnumber"
)

func TestFixedX32Consts(t *testing.T) {
	if X32Const2Pi < vpnumber.I32ToX32(6) || X32Const2Pi > vpnumber.I32ToX32(7) {
		t.Error("inconsistent X32Const2Pi value", X32Const2Pi)
	}
	if X32ConstPi < vpnumber.I32ToX32(3) || X32ConstPi > vpnumber.I32ToX32(4) {
		t.Error("inconsistent X32ConstPi value", X32ConstPi)
	}
	if X32ConstPi2 < vpnumber.I32ToX32(1) || X32ConstPi2 > vpnumber.I32ToX32(2) {
		t.Error("inconsistent X32ConstPi2 value", X32ConstPi2)
	}
	if X32ConstPi4 < vpnumber.I32ToX32(0) || X32ConstPi4 > vpnumber.I32ToX32(1) {
		t.Error("inconsistent X32ConstPi4 value", X32ConstPi4)
	}
}

func TestFixedX64Consts(t *testing.T) {
	if X64Const2Pi < vpnumber.I64ToX64(6) || X64Const2Pi > vpnumber.I64ToX64(7) {
		t.Error("inconsistent X64Const2Pi value", X64Const2Pi)
	}
	if X64ConstPi < vpnumber.I64ToX64(3) || X64ConstPi > vpnumber.I64ToX64(4) {
		t.Error("inconsistent X64ConstPi value", X64ConstPi)
	}
	if X64ConstPi2 < vpnumber.I64ToX64(1) || X64ConstPi2 > vpnumber.I64ToX64(2) {
		t.Error("inconsistent X64ConstPi2 value", X64ConstPi2)
	}
	if X64ConstPi4 < vpnumber.I64ToX64(0) || X64ConstPi4 > vpnumber.I64ToX64(1) {
		t.Error("inconsistent X64ConstPi4 value", X64ConstPi4)
	}
}

func TestFixedX32SqrtTable(t *testing.T) {
	var x vpnumber.X32

	x = x32SqrtTable[0]
	if x != vpnumber.X32Const1 {
		t.Error("bad X32 table value for sqrt(1)", x)
	}
	x = x32SqrtTable[x32TableSize]
	if x != vpnumber.X32Const1<<1 {
		t.Error("bad X32 table value for sqrt(4)", x)
	}
}

func TestFixedX64SqrtTable(t *testing.T) {
	var x vpnumber.X64

	x = x64SqrtTable[0]
	if x != vpnumber.X64Const1 {
		t.Error("bad X64 table value for sqrt(1)", x)
	}
	x = x64SqrtTable[x64TableSize]
	if x != vpnumber.X64Const1<<1 {
		t.Error("bad X64 table value for sqrt(4)", x)
	}
}

func TestFixedX32AtanTable(t *testing.T) {
	var x vpnumber.X32

	x = x32AtanTable[0]
	if x != vpnumber.X32Const0 {
		t.Error("bad X32 table value for atan(0)", x)
	}
	x = x32AtanTable[x32TableSize]
	if !vpnumber.X32IsSimilar(x, vpnumber.F64ToX32(math.Atan(1.0))) {
		t.Error("bad X32 table value for atan(1)", x)
	}
}

func TestFixedX64AtanTable(t *testing.T) {
	var x vpnumber.X64

	x = x64AtanTable[0]
	if x != vpnumber.X64Const0 {
		t.Error("bad X64 table value for atan(0)", x)
	}
	x = x64AtanTable[x64TableSize]
	if !vpnumber.X64IsSimilar(x, vpnumber.F64ToX64(math.Atan(1.0))) {
		t.Error("bad X64 table value for atan(1)", x)
	}
}

func TestFixedX32SinTable(t *testing.T) {
	var x vpnumber.X32

	x = x32SinTable[0]
	if x != vpnumber.X32Const0 {
		t.Error("bad X32 table value for sin(0)", x)
	}
	x = x32SinTable[x32TableSize]
	if !vpnumber.X32IsSimilar(x, vpnumber.F64ToX32(math.Sin(math.Pi/2.0))) {
		t.Error("bad X32 table value for sin(pi/2)", x)
	}
}

func TestFixedX64SinTable(t *testing.T) {
	var x vpnumber.X64

	x = x64SinTable[0]
	if x != vpnumber.X64Const0 {
		t.Error("bad X64 table value for sin(0)", x)
	}
	x = x64SinTable[x64TableSize]
	if !vpnumber.X64IsSimilar(x, vpnumber.F64ToX64(math.Sin(math.Pi/2.0))) {
		t.Error("bad X64 table value for sin(pi/2)", x)
	}
}
