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

func TestU32Vec2Math(t *testing.T) {
	const u1 = 0
	const u2 = 4

	const u5 = 10
	const u6 = 1000

	var v1, v2, v3, v4 *U32Vec2

	v1 = U32Vec2New(u1, u2)
	v2 = U32Vec2New(u5, u6)
	v3 = U32Vec2Add(v1, v2)
	v4 = U32Vec2New(u1+u5, u2+u6)
	if *v3 != *v4 {
		t.Error("Add error")
	}
}
