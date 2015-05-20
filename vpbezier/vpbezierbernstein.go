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

package vpbezier

import (
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
)

// F32Bernstein returns a Berstein polynomial value.
func F32Bernstein(n, i int, t float32) float32 {
	b := vpmath.Binomial(n, i)
	ret := float32(b)
	var l int

	oneMinusT := vpnumber.F32Const1 - t

	for l = 0; l < n; l++ {
		if l < i {
			ret *= t
		} else {
			ret *= oneMinusT
		}
	}

	return ret
}
