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

package vpmatrix

import (
	"testing"
)

func TestU32Vec4Math(t *testing.T) {
	const u1 = 0
	const u2 = 4
	const u3 = 42
	const u4 = 100000

	const u5 = 10
	const u6 = 1000
	const u7 = 222
	const u8 = 1

	var v1, v2, v3, v4 *U32Vec4

	v1 = U32Vec4New(u1, u2, u3, u4)
	v2 = U32Vec4New(u5, u6, u7, u8)
	v3 = U32Vec4Add(v1, v2)
	v4 = U32Vec4New(u1+u5, u2+u6, u3+u7, u4+u8)
	if *v3 != *v4 {
		t.Error("Add error")
	}
}
