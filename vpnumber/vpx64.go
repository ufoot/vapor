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
	"encoding/json"
	"github.com/ufoot/vapor/vpsys"
	"math/big"
)

// X64 is as fixed point number on 64 bits, the idea is to arbitrary
// consider that 2^32 is 1, the 32 strongest bits being the
// integer part, and the 32 weakest bits being the fractionnal
// part. This might sound useless as all modern systems do have
// floats but in some cases it's convenient to have accurate
// integer calculations representing decimal numbers.
type X64 int64

// X64Const0 contains 0 as a fixed point number on 64 bits.
const X64Const0 X64 = 0

// X64Const1 contains 1 as a fixed point number on 64 bits.
const X64Const1 X64 = 0x100000000

// X64Shift tells how many bits must be shifted to convert a
// fixed point number on 64 bits to its int64 corresponding value.
const X64Shift = 32

const x64Const2 X64 = 0x200000000
const x64Const3 X64 = 0x300000000
const x64Const4 X64 = 0x400000000
const x64Const5 X64 = 0x500000000
const x64Const6 X64 = 0x600000000
const x64Const7 X64 = 0x700000000
const x64Const8 X64 = 0x800000000
const x64Const9 X64 = 0x900000000

const x64Half = 0x080000000
const x64ShiftHalf = 16
const x64SimilarDiff = 0x01000000
const x64SimilarScale = 0x00800000
const x64IpMask = 0xffffffff00000000

// X64ToI32 converts a fixed point number on 64 bits to an int32.
func X64ToI32(x X64) int32 {
	return int32(int64(x) >> X64Shift)
}

// X64ToI64 converts a fixed point number on 64 bits to an int64.
func X64ToI64(x X64) int64 {
	return int64(x) >> X64Shift
}

// X64ToX32 converts a fixed point number on 64 bits to a fixed point number on 32 bits.
func X64ToX32(x X64) X32 {
	return X32(int32(int64(x) >> (X64Shift - X32Shift)))
}

// X64ToF32 converts a fixed point number on 64 bits to a float32.
func X64ToF32(x X64) float32 {
	return float32(x) / float32(X64Const1)
}

// X64ToF64 converts a fixed point number on 64 bits to a float64.
func X64ToF64(x X64) float64 {
	return float64(x) / float64(X64Const1)
}

// MarshalJSON implements the json.Marshaler interface.
func (x *X64) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal(int64(*x))
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to marshal X64")
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (x *X64) UnmarshalJSON(data []byte) error {
	var tmp int64

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return vpsys.ErrorChain(err, "unable to unmarshal X64")
	}

	*x = X64(tmp)

	return nil
}

// String returns a readable form of the matrix.
func (x *X64) String() string {
	buf, err := json.Marshal(X64ToF64(*x))

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}

// X64Min returns the minimum value bewteen 2 fixed point numbers on 64 bits.
func X64Min(x1 X64, x2 X64) X64 {
	if x1 < x2 {
		return x1
	}
	return x2
}

// X64Max returns the maximum value bewteen 2 fixed point numbers on 64 bits.
func X64Max(x1 X64, x2 X64) X64 {
	if x1 < x2 {
		return x2
	}
	return x1
}

// X64Abs returns the absolute value of a fixed point number on 64 bits.
func X64Abs(x X64) X64 {
	if x < 0 {
		return -x
	}
	return x
}

// X64Mul multiplicates 2 fixed point numbers on 64 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big.
// Beware of rounding errors, might be very approximative on limit cases.
func X64Mul(x1 X64, x2 X64) X64 {
	return X64((int64(x1) >> x64ShiftHalf) * (int64(x2) >> x64ShiftHalf))
}

// X64Muln multiplicates several fixed point numbers on 64 bits.
// Calls the 2 args mul function recursively.
// Beware of rounding errors, might be very approximative on limit cases.
func X64Muln(x X64, xn ...X64) X64 {
	if xn == nil || len(xn) <= 0 {
		return x
	}
	x = X64Mul(x, xn[0])
	return X64Muln(x, xn[1:]...)
}

// X64Div divides 2 fixed point numbers on 64 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big. Does not raise any error on
// division by zero. Instead it will divide by the smallest value
// available, the results will probably be inconsistent but at least
// no panic around, program flow is not interrupted.
// Beware of rounding errors, might be very approximative on limit cases.
func X64Div(x1 X64, x2 X64) X64 {
	var d = int64(x2) >> x64ShiftHalf
	if d == 0 {
		return X64((int64(x1) << X64Shift))
	}
	return X64((int64(x1) << x64ShiftHalf) / d)
}

// X64Mulp multiplicates 2 fixed point numbers on 64 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big.
// Internally, this function uses 128 bits integers so that results
// have minimal rounding errors. As a consequence, it is quite slow.
func X64Mulp(x1 X64, x2 X64) X64 {
	bx1 := big.NewInt(int64(x1))
	bx2 := big.NewInt(int64(x2))

	var bt big.Int
	var br big.Int

	bt.Mul(bx1, bx2)
	br.Rsh(&bt, X64Shift)

	return X64(br.Int64())
}

// X64Divp divides 2 fixed point numbers on 64 bits. This is not
// equivalent to the standard * operator as values need be shifted so
// that they do not get artificially big. Does not raise any error on
// division by zero. Instead it will divide by the smallest value
// available, the results will probably be inconsistent but at least
// no panic around, program flow is not interrupted.
// Internally, this function uses 128 bits integers so that results
// have minimal rounding errors. As a consequence, it is quite slow.
func X64Divp(x1 X64, x2 X64) X64 {
	bx1 := big.NewInt(int64(x1))
	if x2 == 0 {
		return X64((int64(x1) << X64Shift))
	}
	bx2 := big.NewInt(int64(x2))

	var bt big.Int
	var br big.Int

	bt.Lsh(bx1, X64Shift)
	br.Div(&bt, bx2)

	return X64(br.Int64())
}

// X64Exponent returns the exponent of a fixed point number on 64 bits.
// This means a logical value of 1 has an exponent of 0, and the smallest
// strictly positive integer has an exponent of -32. 0 returns 0.
func X64Exponent(x X64) int {
	var ret int
	var val uint64

	switch {
	case x < 0:
		val = uint64(-x)
	case x > 0:
		val = uint64(x)
	default:
		return 0
	}

	if (val & x64IpMask) != 0 {
		for (val&x64IpMask) != 0 && ret <= X64Shift {
			val >>= 1
			ret++
		}
		ret--
	} else {
		for (val&x64IpMask) == 0 && ret > -X64Shift {
			val <<= 1
			ret--
		}
	}

	return ret
}

// X64Mantis returns the mantis of a fixed point number on 64 bits.
// Note that on numbers greater than 2 (exponent greater than 1) there
// is a loss of information as bits are shifted to the right.
func X64Mantis(x X64) X64 {
	exponent := X64Exponent(x)
	if exponent < 0 {
		return x << uint(-exponent)
	} else if exponent > 0 {
		return x >> uint(exponent)
	}

	return x
}

// X64Round rounds a fixed point number on 64 bits to the
// nearest integer value.
// Casted as an integer the result has the 32 lowest bits set to 0.
func X64Round(x X64) X64 {
	return X64(uint64(x+x64Half) & x64IpMask)
}

// X64Floor rounds a fixed point number on 64 bits to the
// integer value just below.
// Casted as an integer the result has the 32 lowest bits set to 0.
func X64Floor(x X64) X64 {
	return X64(uint64(x) & x64IpMask)
}

// X64Ceil rounds a fixed point number on 64 bits to the nearest
// integer value just above.
// Casted as an integer the result has the 32 lowest bits set to 0.
func X64Ceil(x X64) X64 {
	return X64((uint64(x)-1)&x64IpMask) + X64Const1
}

// X64IsSimilar returns true if args are approximatively the same.
// This is a workarround to ignore rounding errors.
func X64IsSimilar(x1 X64, x2 X64) bool {
	diff := X64Abs(x1 - x2)
	scale := X64Max(X64Abs(x1), X64Abs(x2))

	if scale <= x64SimilarScale {
		return true
	}

	if X64Div(diff, scale) <= x64SimilarDiff {
		return true
	}

	return false
}

// X64Lerp performs a linear interpolation between a and b.
func X64Lerp(x1 X64, x2 X64, beta X64) X64 {
	switch {
	case beta <= X64Const0:
		return x1
	case beta >= X64Const1:
		return x2
	}

	return X64Mul(x1, (X64Const1-beta)) + X64Mul(x2, beta)
}
