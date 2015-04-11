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
)

const f32Const2Pi = float32(2 * math.Pi)
const f64Const2Pi = 2 * math.Pi

// F32Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func F32Anglelerp(f1, f2, beta float32) float32 {
	switch {
	case beta <= vpnumber.F32Const0:
		return f1
	case beta >= vpnumber.F32Const1:
		return f2
	}

	f1 = float32(math.Mod(float64(f1), float64(f32Const2Pi)))
	f2 = float32(math.Mod(float64(f2), float64(f32Const2Pi)))
	if f1 < 0 {
		f1 += f32Const2Pi
	}
	if f2 < 0 {
		f2 += f32Const2Pi
	}

	switch {
	case f1 < f2:
		if f2-f1 < math.Pi {
			return F32Lerp(f1, f2, beta)
		}
		return F32Lerp(f1+f32Const2Pi, f2, beta)
	case f2 < f1:
		if f1-f2 < math.Pi {
			return F32Lerp(f1, f2, beta)
		}
		return F32Lerp(f1, f2+f32Const2Pi, beta)
	}

	return f1
}

// F64Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func F64Anglelerp(f1, f2, beta float64) float64 {
	switch {
	case beta <= vpnumber.F64Const0:
		return f1
	case beta >= vpnumber.F64Const1:
		return f2
	}

	f1 = math.Mod(f1, f64Const2Pi)
	f2 = math.Mod(f2, f64Const2Pi)
	if f1 < 0 {
		f1 += f64Const2Pi
	}
	if f2 < 0 {
		f2 += f64Const2Pi
	}

	switch {
	case f1 < f2:
		if f2-f1 < math.Pi {
			return F64Lerp(f1, f2, beta)
		}
		return F64Lerp(f1+f64Const2Pi, f2, beta)
	case f2 < f1:
		if f1-f2 < math.Pi {
			return F64Lerp(f1, f2, beta)
		}
		return F64Lerp(f1, f2+f64Const2Pi, beta)
	}

	return f1
}

// X32Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func X32Anglelerp(x1, x2, beta vpnumber.X32) vpnumber.X32 {
	switch {
	case beta <= vpnumber.X32Const0:
		return x1
	case beta >= vpnumber.X32Const1:
		return x2
	}

	x1 = x1 % X32Const2Pi
	x2 = x2 % X32Const2Pi
	if x1 < 0 {
		x1 += X32Const2Pi
	}
	if x2 < 0 {
		x2 += X32Const2Pi
	}

	switch {
	case x1 < x2:
		if x2-x1 < X32ConstPi {
			return X32Lerp(x1, x2, beta)
		}
		return X32Lerp(x1+X32Const2Pi, x2, beta)
	case x2 < x1:
		if x1-x2 < X32ConstPi {
			return X32Lerp(x1, x2, beta)
		}
		return X32Lerp(x1, x2+X32Const2Pi, beta)
	}

	return x1
}

// X64Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func X64Anglelerp(x1, x2, beta vpnumber.X64) vpnumber.X64 {
	switch {
	case beta <= vpnumber.X64Const0:
		return x1
	case beta >= vpnumber.X64Const1:
		return x2
	}

	x1 = x1 % X64Const2Pi
	x2 = x2 % X64Const2Pi
	if x1 < 0 {
		x1 += X64Const2Pi
	}
	if x2 < 0 {
		x2 += X64Const2Pi
	}

	switch {
	case x1 < x2:
		if x2-x1 < X64ConstPi {
			return X64Lerp(x1, x2, beta)
		}
		return X64Lerp(x1+X64Const2Pi, x2, beta)
	case x2 < x1:
		if x1-x2 < X64ConstPi {
			return X64Lerp(x1, x2, beta)
		}
		return X64Lerp(x1, x2+X64Const2Pi, beta)
	}

	return x1
}
