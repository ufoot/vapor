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
	"math"
)

// F64Const0 contains 0 as a float64.
const F64Const0 float64 = 0.0

// F64Const1 contains 1 as a float64.
const F64Const1 float64 = 1.0

const f64Const2 float64 = 2.0
const f64Const3 float64 = 3.0
const f64Const4 float64 = 4.0
const f64Const5 float64 = 5.0
const f64Const6 float64 = 6.0
const f64Const7 float64 = 7.0
const f64Const8 float64 = 8.0
const f64Const9 float64 = 9.0

const f64SimilarDiff float64 = 1.0e-5
const f64SimilarScale float64 = 1.0e-15

// F64ToX32 converts a float64 to a fixed point number on 32 bits.
func F64ToX32(f float64) X32 {
	return X32(int32(F64Round(f * float64(X32Const1))))
}

// F64ToX64 converts a float64 to a fixed point number on 64 bits.
func F64ToX64(f float64) X64 {
	return X64(int64(F64Round(f * float64(X64Const1))))
}

// F64Div does a simple division, without raising any error on
// division by zero. Instead it will divide by the smallest value
// available, the results will probably be inconsistent but at least
// no panic around, program flow is not interrupted.
func F64Div(f1 float64, f2 float64) float64 {
	if f2 == 0.0 {
		f2 = math.SmallestNonzeroFloat64
	}
	return f1 / f2
}

// F64Round rounds to the nearest integer, while still returning
// a floating point value.
func F64Round(f float64) float64 {
	return math.Floor(f + 0.5)
}

// F64IsSimilar returns true if args are approximatively the same.
// This is a workarround to ignore rounding errors.
func F64IsSimilar(f1 float64, f2 float64) bool {
	diff := math.Abs(f1 - f2)
	scale := math.Max(math.Abs(f1), math.Abs(f2))

	if scale < f64SimilarScale {
		return true
	}

	if (diff / scale) <= f64SimilarDiff {
		return true
	}

	return false
}

// F64Lerp performs a linear interpolation between a and b.
func F64Lerp(f1 float64, f2 float64, beta float64) float64 {
	switch {
	case beta <= F64Const0:
		return f1
	case beta >= F64Const1:
		return f2
	}

	return f1*(F64Const1-beta) + f2*beta
}
