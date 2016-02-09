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

package vpnumber

import (
	"testing"
)

func TestF64Convert(t *testing.T) {
	var x32 X32
	var x64 X64

	x32 = F64ToX32(F64Const1)
	if x32 != X32Const1 {
		t.Error("can't convert positive F64 to X32", x32)
	}
	x32 = F64ToX32(-F64Const1)
	if x32 != -X32Const1 {
		t.Error("can't convert negtive F64 to X32", x32)
	}
	x64 = F64ToX64(F64Const1)
	if x64 != X64Const1 {
		t.Error("can't convert positive F64 to X64", x64)
	}
	x64 = F64ToX64(-F64Const1)
	if x64 != -X64Const1 {
		t.Error("can't convert negtive F64 to X64", x64)
	}
}

func TestF64Float(t *testing.T) {
	var f float64

	f = F64Round(5.0 / 4.0)
	if f != F64Const1 {
		t.Error("Round problem on positive numbers", f)
	}
	f = F64Round(-5.0 / 4.0)
	if f != -F64Const1 {
		t.Error("Round problem on negative numbers", f)
	}
}

func TestF64Similar(t *testing.T) {
	var f float64

	if !F64IsSimilar(F64Const1, F64Const1) {
		t.Error("Can't figure out same float64 is similar")
	}
	if F64IsSimilar(F64Const1, -F64Const1) {
		t.Error("Can't figure out different float64 are not similar")
	}
	f = float64(F64Const1) * 1.000000001
	if f == float64(F64Const1) {
		t.Error("Similar values should not be the same", float64(F64Const1), f)
	}
	if !F64IsSimilar(F64Const1, f) {
		t.Error("Can't figure out float64 is similar to 1", f)
	}
}
