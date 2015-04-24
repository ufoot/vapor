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

// F32Const90 stores 90 degrees as a float32.
const F32Const90 float32 = 90.0

// F32Const180 stores 180 degrees as a float32.
const F32Const180 float32 = 180.0

// F32Const360 stores 360 degrees as a float32.
const F32Const360 float32 = 360.0

// F64Const90 stores 90 degrees as a float64.
const F64Const90 float64 = 90.0

// F64Const180 stores 180 degrees as a float64.
const F64Const180 float64 = 180.0

// F64Const360 stores 360 degrees as a float64.
const F64Const360 float64 = 360.0

// X32Const90 stores 90 degrees as a 32-bit fixed point number..
const X32Const90 vpnumber.X32 = 90 * vpnumber.X32Const1

// X32Const180 stores 180 degrees as a 32-bit fixed point number..
const X32Const180 vpnumber.X32 = 180 * vpnumber.X32Const1

// X32Const360 stores 360 degrees as a 32-bit fixed point number..
const X32Const360 vpnumber.X32 = 360 * vpnumber.X32Const1

// X64Const90 stores 90 degrees as a 64-bit fixed point number..
const X64Const90 vpnumber.X64 = 90 * vpnumber.X64Const1

// X64Const180 stores 180 degrees as a 64-bit fixed point number..
const X64Const180 vpnumber.X64 = 180 * vpnumber.X64Const1

// X64Const360 stores 360 degrees as a 64-bit fixed point number..
const X64Const360 vpnumber.X64 = 360 * vpnumber.X64Const1

const f32Const2Pi = float32(2 * math.Pi)
const f64Const2Pi = 2 * math.Pi

const f32FactorDegToRad float32 = float32(math.Pi) / F32Const180
const f64FactorDegToRad float64 = math.Pi / F64Const180
const f32FactorRadToDeg float32 = F32Const180 / float32(math.Pi)
const f64FactorRadToDeg float64 = F64Const180 / math.Pi

// F32DegMod returns the angle modulo 2*Pi, and always as a positive value.
func F32DegMod(deg float32) float32 {
	deg = float32(math.Mod(float64(deg), float64(F32Const360)))
	if deg < 0 {
		deg += F32Const360
	}
	return deg
}

// F64DegMod returns the angle modulo 2*Pi, and always as a positive value.
func F64DegMod(deg float64) float64 {
	deg = math.Mod(deg, F64Const360)
	if deg < 0 {
		deg += F64Const360
	}
	return deg
}

// X32DegMod returns the angle modulo 2*Pi, and always as a positive value.
func X32DegMod(deg vpnumber.X32) vpnumber.X32 {
	deg = deg % X32Const360
	if deg < 0 {
		deg += X32Const360
	}
	return deg
}

// X64DegMod returns the angle modulo 2*Pi, and always as a positive value.
func X64DegMod(deg vpnumber.X64) vpnumber.X64 {
	deg = deg % X64Const360
	if deg < 0 {
		deg += X64Const360
	}
	return deg
}

// F32RadMod returns the angle modulo 2*Pi, and always as a positive value.
func F32RadMod(rad float32) float32 {
	rad = float32(math.Mod(float64(rad), float64(f32Const2Pi)))
	if rad < 0 {
		rad += f32Const2Pi
	}
	return rad
}

// F64RadMod returns the angle modulo 2*Pi, and always as a positive value.
func F64RadMod(rad float64) float64 {
	rad = math.Mod(rad, float64(f64Const2Pi))
	if rad < 0 {
		rad += f64Const2Pi
	}
	return rad
}

// X32RadMod returns the angle modulo 2*Pi, and always as a positive value.
func X32RadMod(rad vpnumber.X32) vpnumber.X32 {
	rad = rad % X32Const2Pi
	if rad < 0 {
		rad += X32Const2Pi
	}
	return rad
}

// X64RadMod returns the angle modulo 2*Pi, and always as a positive value.
func X64RadMod(rad vpnumber.X64) vpnumber.X64 {
	rad = rad % X64Const2Pi
	if rad < 0 {
		rad += X64Const2Pi
	}
	return rad
}

// F32DegToRad converts an angle from degrees to radians.
func F32DegToRad(deg float32) float32 {
	return deg * f32FactorDegToRad
}

// F64DegToRad converts an angle from degrees to radians.
func F64DegToRad(deg float64) float64 {
	return deg * f64FactorDegToRad
}

// X32DegToRad converts an angle from degrees to radians.
func X32DegToRad(deg vpnumber.X32) vpnumber.X32 {
	return vpnumber.X32Div(vpnumber.X32Mul(deg, X32ConstPi>>2), X32Const180>>2)
}

// X64DegToRad converts an angle from degrees to radians.
func X64DegToRad(deg vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Div(vpnumber.X64Mul(deg, X64ConstPi>>2), X64Const180>>2)
}

// F32RadToDeg converts an angle from radians to degrees.
func F32RadToDeg(rad float32) float32 {
	return rad * f32FactorRadToDeg
}

// F64RadToDeg converts an angle from radians to degrees.
func F64RadToDeg(rad float64) float64 {
	return rad * f64FactorRadToDeg
}

// X32RadToDeg converts an angle from radians to degrees.
func X32RadToDeg(rad vpnumber.X32) vpnumber.X32 {
	return vpnumber.X32Div(vpnumber.X32Mul(rad, X32Const180>>2), X32ConstPi>>2)
}

// X64RadToDeg converts an angle from radians to degrees.
func X64RadToDeg(rad vpnumber.X64) vpnumber.X64 {
	return vpnumber.X64Div(vpnumber.X64Mul(rad, X64Const180>>2), X64ConstPi>>2)
}
