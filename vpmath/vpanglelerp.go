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

package vpmath

import (
	"github.com/ufoot/vapor/vpnumber"
	"math"
)

// F32Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func F32Anglelerp(rad1, rad2, beta float32) float32 {
	switch {
	case beta <= vpnumber.F32Const0:
		return rad1
	case beta >= vpnumber.F32Const1:
		return rad2
	}

	rad1 = F32RadMod(rad1)
	rad2 = F32RadMod(rad2)

	switch {
	case rad1 < rad2:
		if rad2-rad1 < math.Pi {
			return F32Lerp(rad1, rad2, beta)
		}
		return F32Lerp(rad1, rad2-f32Const2Pi, beta)
	case rad2 < rad1:
		if rad1-rad2 < math.Pi {
			return F32Lerp(rad1, rad2, beta)
		}
		return F32Lerp(rad1-f32Const2Pi, rad2, beta)
	}

	return rad1
}

// F64Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func F64Anglelerp(rad1, rad2, beta float64) float64 {
	switch {
	case beta <= vpnumber.F64Const0:
		return rad1
	case beta >= vpnumber.F64Const1:
		return rad2
	}

	rad1 = F64RadMod(rad1)
	rad2 = F64RadMod(rad2)

	switch {
	case rad1 < rad2:
		if rad2-rad1 < math.Pi {
			return F64Lerp(rad1, rad2, beta)
		}
		return F64Lerp(rad1, rad2-f64Const2Pi, beta)
	case rad2 < rad1:
		if rad1-rad2 < math.Pi {
			return F64Lerp(rad1, rad2, beta)
		}
		return F64Lerp(rad1-f64Const2Pi, rad2, beta)
	}

	return rad1
}

// X32Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func X32Anglelerp(rad1, rad2, beta vpnumber.X32) vpnumber.X32 {
	switch {
	case beta <= vpnumber.X32Const0:
		return rad1
	case beta >= vpnumber.X32Const1:
		return rad2
	}

	rad1 = X32RadMod(rad1)
	rad2 = X32RadMod(rad2)

	switch {
	case rad1 < rad2:
		if rad2-rad1 < X32ConstPi {
			return X32Lerp(rad1, rad2, beta)
		}
		return X32Lerp(rad1, rad2-X32Const2Pi, beta)
	case rad2 < rad1:
		if rad1-rad2 < X32ConstPi {
			return X32Lerp(rad1, rad2, beta)
		}
		return X32Lerp(rad1-X32Const2Pi, rad2, beta)
	}

	return rad1
}

// X64Anglelerp performs a linear interpolation between two angles.
// Angles need a special lerp as they wrap arround.
func X64Anglelerp(rad1, rad2, beta vpnumber.X64) vpnumber.X64 {
	switch {
	case beta <= vpnumber.X64Const0:
		return rad1
	case beta >= vpnumber.X64Const1:
		return rad2
	}

	rad1 = X64RadMod(rad1)
	rad2 = X64RadMod(rad2)

	switch {
	case rad1 < rad2:
		if rad2-rad1 < X64ConstPi {
			return X64Lerp(rad1, rad2, beta)
		}
		return X64Lerp(rad1, rad2-X64Const2Pi, beta)
	case rad2 < rad1:
		if rad1-rad2 < X64ConstPi {
			return X64Lerp(rad1, rad2, beta)
		}
		return X64Lerp(rad1-X64Const2Pi, rad2, beta)
	}

	return rad1
}
