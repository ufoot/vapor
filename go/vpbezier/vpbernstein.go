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

package vpbezier

import (
	"github.com/ufoot/vapor/go/vpmath"
	"github.com/ufoot/vapor/go/vpnumber"
)

func f32TOneMinusT(a, b int, t float32) float32 {
	ret := vpnumber.F32Const1
	oneMinusT := vpnumber.F32Const1 - t
	var l int

	for l = 0; l < a; l++ {
		ret *= t
	}
	for l = 0; l < b; l++ {
		ret *= oneMinusT
	}

	return ret
}

// F32Bernstein returns a Berstein polynomial value.
func F32Bernstein(n, i int, t float32) float32 {
	b := vpmath.Binomial(n, i)

	return float32(b) * f32TOneMinusT(i, n-i, t)
}

func f32TOneMinusTDerivative(a, b int, t float32) float32 {
	oneMinusT := vpnumber.F32Const1 - t
	switch {
	case a >= 1:
		return t*f32TOneMinusTDerivative(a-1, b, t) + f32TOneMinusT(a-1, b, t)
	case b >= 1:
		return oneMinusT*f32TOneMinusTDerivative(a, b-1, t) - f32TOneMinusT(a, b-1, t)
	}

	return vpnumber.F32Const0
}

// F32BernsteinDerivative returns the derivative of the
// the Berstein polynomial func at a given point.
func F32BernsteinDerivative(n, i int, t float32) float32 {
	b := vpmath.Binomial(n, i)

	return float32(b) * f32TOneMinusTDerivative(i, n-i, t)
}
