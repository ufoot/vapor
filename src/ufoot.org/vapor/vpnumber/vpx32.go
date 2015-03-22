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

// X32 is a fixed point number on 32 bits. The idea is to arbitrary
// consider that 2^16 is 1, the 16 strongest bits being the
// integer part, and the 16 weakest bits being the fractionnal
// part. This might sound useless as all modern systems do have
// floats but in some cases it's convenient to have accurate
// integer calculations representing decimal numbers.
type X32 int32

// Constant containing 0 as a fixed point number on 32 bits.
const X32Const0 X32 = 0

// Constant containing 1 as a fixed point number on 32 bits.
const X32Const1 X32 = 0x10000

// X32Shift tells how many bits must be shifted to convert a
// fixed point number on 32 bits to its int32 corresponding value.
const X32Shift = 16

const x32Const2 X32 = 0x20000
const x32Const3 X32 = 0x30000
const x32Const4 X32 = 0x40000
const x32Const5 X32 = 0x50000
const x32Const6 X32 = 0x60000
const x32Const7 X32 = 0x70000
const x32Const8 X32 = 0x80000
const x32Const9 X32 = 0x90000

const x32Half = 0x08000
const x32ShiftHalf = 8
const x32SimilarDiff = 0x0800
const x32SimilarScale = 0x0400
const x32IpMask = 0xffff0000

// X32ToI32 converts a fixed point number on 32 bits to an int32.
func X32ToI32(x X32) int32 {
	return int32(x) >> X32Shift
}

// X32ToI64 converts a fixed point number on 32 bits to an int64.
func X32ToI64(x X32) int64 {
	return int64(x) >> X32Shift
}

// X32ToX64 converts a fixed point number on 32 bits to a fixed point number on 64 bits.
func X32ToX64(x X32) X64 {
	return X64(int64(x) << (X64Shift - X32Shift))
}

// X32ToF32 converts a fixed point number on 32 bits to a float32.
func X32ToF32(x X32) float32 {
	return float32(x) / float32(X32Const1)
}

// X32ToF64 converts a fixed point number on 32 bits to a float64.
func X32ToF64(x X32) float64 {
	return float64(x) / float64(X32Const1)
}

// X32Min returns the minimum value bewteen 2 fixed point numbers on 32 bits.
func X32Min(x1 X32, x2 X32) X32 {
	if x1 < x2 {
		return x1
	}
	return x2
}

// X32Max returns the maximum value bewteen 2 fixed point numbers on 32 bits.
func X32Max(x1 X32, x2 X32) X32 {
	if x1 < x2 {
		return x2
	}
	return x1
}

// X32Abs returns the absolute value of a fixed point number on 32 bits.
func X32Abs(x X32) X32 {
	if x < 0 {
		return -x
	}
	return x
}

// X32Mul multiplicates 2 fixed point numbers on 32 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big.
// Beware of rounding errors, might be very approximative on limit cases.
func X32Mul(x1 X32, x2 X32) X32 {
	return X32((int32(x1) >> x32ShiftHalf) * (int32(x2) >> x32ShiftHalf))
}

// X32Muln multiplicates several fixed point numbers on 32 bits.
// Calls the 2 args mul function recursively.
// Beware of rounding errors, might be very approximative on limit cases.
func X32Muln(x X32, xn ...X32) X32 {
	switch len(xn) {
	case 0:
		return x
	case 1:
		return X32Mul(x, xn[0])
	}
	x = X32Mul(x, xn[0])
	return X32Muln(x, xn[1:]...)
}

// X32Div divides 2 fixed point numbers on 32 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big. Does not raise any error on
// division by zero. Instead it will divide by the smallest value
// available, the results will probably be inconsistent but at least
// no panic around, program flow is not interrupted.
// Beware of rounding errors, might be very approximative on limit cases.
func X32Div(x1 X32, x2 X32) X32 {
	var d = int32(x2) >> x32ShiftHalf
	if d == 0 {
		return X32((int32(x1) << X32Shift))
	}
	return X32((int32(x1) << x32ShiftHalf) / d)
}

// X32Mulp multiplicates 2 fixed point numbers on 32 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big.
// Internally, this function uses 64 bits integers so that results
// have minimal rounding errors.
func X32Mulp(x1 X32, x2 X32) X32 {
	return X32((int64(x1) * int64(x2)) >> X32Shift)
}

// X32Divp divides 2 fixed point numbers on 32 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big. Does not raise any error on
// division by zero. Instead it will divide by the smallest value
// available, the results will probably be inconsistent but at least
// no panic around, program flow is not interrupted.
// Internally, this function uses 64 bits integers so that results
// have minimal rounding errors.
func X32Divp(x1 X32, x2 X32) X32 {
	if x2 == 0 {
		return X32((int32(x1) << X32Shift))
	}
	return X32((int64(x1) << X32Shift) / int64(x2))
}

// X32Exponent returns the exponent of a fixed point number on 32 bits.
// This means a logical value of 1 has an exponent of 0, and the smallest
// strictly positive integer has an exponent of -16. 0 returns 0.
func X32Exponent(x X32) int {
	var ret int
	var val uint32

	switch {
	case x < 0:
		val = uint32(-x)
	case x > 0:
		val = uint32(x)
	default:
		return 0
	}

	if (val & x32IpMask) != 0 {
		for (val&x32IpMask) != 0 && ret <= X32Shift {
			val >>= 1
			ret++
		}
		ret--
	} else {
		for (val&x32IpMask) == 0 && ret > -X32Shift {
			val <<= 1
			ret--
		}
	}

	return ret
}

// X32Mantis returns the mantis of a fixed point number on 32 bits.
// Note that on numbers greater than 2 (exponent greater than 1) there
// is a loss of information as bits are shifted to the right.
func X32Mantis(x X32) X32 {
	exponent := X32Exponent(x)
	if exponent < 0 {
		return x << uint(-exponent)
	} else if exponent > 0 {
		return x >> uint(exponent)
	}

	return x
}

// X32Round rounds a fixed point number on 32 bits to the
// nearest integer value.
// Casted as an integer the result has the 16 lowest bits set to 0.
func X32Round(x X32) X32 {
	return X32(uint32(x+x32Half) & x32IpMask)
}

// X32Floor rounds a fixed point number on 32 bits to the
// integer value just below.
// Casted as an integer the result has the 16 lowest bits set to 0.
func X32Floor(x X32) X32 {
	return X32(uint32(x) & x32IpMask)
}

// X32Ceil rounds a fixed point number on 32 bits to the nearest
// integer value just above.
// Casted as an integer the result has the 16 lowest bits set to 0.
func X32Ceil(x X32) X32 {
	return X32((uint32(x)-1)&x32IpMask) + X32Const1
}

// X32IsSimilar returns true if args are approximatively the same.
// This is a workarround to ignore rounding errors.
func X32IsSimilar(x1 X32, x2 X32) bool {
	diff := X32Abs(x1 - x2)
	scale := X32Max(X32Abs(x1), X32Abs(x2))

	if scale < x32SimilarScale {
		return true
	}

	if X32Div(diff, scale) <= x32SimilarDiff {
		return true
	}

	return false
}

// X32Lerp performs a linear interpolation between a and b.
func X32Lerp(x1 X32, x2 X32, beta X32) X32 {
	switch {
	case beta <= X32Const0:
		return x1
	case beta >= X32Const1:
		return x2
	}

	return X32Mul(x1, (X32Const1-beta)) + X32Mul(x2, beta)
}
