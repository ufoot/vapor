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

package vpnumber

import (
	"math"
)

// F32Const0 contains 0 as a float32.
const F32Const0 float32 = 0.0

// F32Const1 contains 1 as a float32.
const F32Const1 float32 = 1.0

const f32Const2 float32 = 2.0
const f32Const3 float32 = 3.0
const f32Const4 float32 = 4.0
const f32Const5 float32 = 5.0
const f32Const6 float32 = 6.0
const f32Const7 float32 = 7.0
const f32Const8 float32 = 8.0
const f32Const9 float32 = 9.0

const f32SimilarDiff float32 = 1.0e-3
const f32SimilarScale float32 = 1.0e-4

// F32ToX32 converts a float32 to a fixed point number on 32 bits.
func F32ToX32(f float32) X32 {
	return X32(int32(F32Round(f * float32(X32Const1))))
}

// F32ToX64 cnverts a float32 to a fixed point number on 64 bits.
func F32ToX64(f float32) X64 {
	return X64(int64(F32Round(f * float32(X64Const1))))
}

// F32Div does a simple division, without raising any error on
// division by zero. Instead it will divide by the smallest value
// available, the results will probably be inconsistent but at least
// no panic around, program flow is not interrupted.
func F32Div(f1 float32, f2 float32) float32 {
	if f2 == 0.0 {
		f2 = math.SmallestNonzeroFloat32
	}
	return f1 / f2
}

// F32Round rounds to the nearest integer, while still returning
// a floating point value.
func F32Round(f float32) float32 {
	return float32(math.Floor(float64(f) + 0.5))
}

// F32IsSimilar returns true if args are approximatively the same.
// This is a workarround to ignore rounding errors.
func F32IsSimilar(f1 float32, f2 float32) bool {
	diff := float32(math.Abs(float64(f1 - f2)))
	scale := float32(math.Max(math.Abs(float64(f1)), math.Abs(float64(f2))))

	if scale <= f32SimilarScale {
		return true
	}

	if (diff / scale) <= f32SimilarDiff {
		return true
	}

	return false
}
