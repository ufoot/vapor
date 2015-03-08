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

package vpnumber

import (
	"testing"
)

func TestI32Convert(t *testing.T) {
	var x32 X32
	var x64 X64

	x32 = I32ToX32(I32Const1)
	if x32 != X32Const1 {
		t.Error("can't convert positive I32 to X32", x32)
	}
	x32 = I32ToX32(-I32Const1)
	if x32 != -X32Const1 {
		t.Error("can't convert negtive I32 to X32", x32)
	}
	x64 = I32ToX64(I32Const1)
	if x64 != X64Const1 {
		t.Error("can't convert positive I32 to X64", x64)
	}
	x64 = I32ToX64(-I32Const1)
	if x64 != -X64Const1 {
		t.Error("can't convert negtive I32 to X64", x64)
	}
}
