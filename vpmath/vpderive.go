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
)

// F32DeriveFunc returns a numerical derivative function.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func F32DeriveFunc(f func(float32) float32, dt float32) func(float32) float32 {
	return func(t float32) float32 {
		return vpnumber.F32Div(f(t+dt/2.0)-f(t-dt/2.0), dt)
	}
}

// F32Derive returns a numerical derivative value at a given point.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func F32Derive(f func(float32) float32, t, dt float32) float32 {
	return F32DeriveFunc(f, dt)(t)
}

// F64DeriveFunc returns a numerical derivative function.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func F64DeriveFunc(f func(float64) float64, dt float64) func(float64) float64 {
	return func(t float64) float64 {
		return vpnumber.F64Div(f(t+dt/2.0)-f(t-dt/2.0), dt)
	}
}

// F64Derive returns a numerical derivative value at a given point.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func F64Derive(f func(float64) float64, t, dt float64) float64 {
	return F64DeriveFunc(f, dt)(t)
}

// X32DeriveFunc returns a numerical derivative function.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func X32DeriveFunc(f func(vpnumber.X32) vpnumber.X32, dt vpnumber.X32) func(vpnumber.X32) vpnumber.X32 {
	return func(t vpnumber.X32) vpnumber.X32 {
		return vpnumber.X32Div(f(t+(dt>>1))-f(t-(dt>>1)), dt)
	}
}

// X32Derive returns a numerical derivative value at a given point.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func X32Derive(f func(vpnumber.X32) vpnumber.X32, t, dt vpnumber.X32) vpnumber.X32 {
	return X32DeriveFunc(f, dt)(t)
}

// X64DeriveFunc returns a numerical derivative function.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func X64DeriveFunc(f func(vpnumber.X64) vpnumber.X64, dt vpnumber.X64) func(vpnumber.X64) vpnumber.X64 {
	return func(t vpnumber.X64) vpnumber.X64 {
		return vpnumber.X64Div(f(t+(dt>>1))-f(t-(dt>>1)), dt)
	}
}

// X64Derive returns a numerical derivative value at a given point.
// The dt arg is here to give the interval on which to test.
// Function will never fail unless the called function itself fails.
// But... it only returns an approximation.
func X64Derive(f func(vpnumber.X64) vpnumber.X64, t, dt vpnumber.X64) vpnumber.X64 {
	return X64DeriveFunc(f, dt)(t)
}
