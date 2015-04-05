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

package vpnumber

// I64Const0 contains 0 as an int64.
const I64Const0 int64 = 0

// I64Const1 contains 1 as an int64.
const I64Const1 int64 = 1

const i64Const2 int64 = 2
const i64Const3 int64 = 3
const i64Const4 int64 = 4
const i64Const5 int64 = 5
const i64Const6 int64 = 6
const i64Const7 int64 = 7
const i64Const8 int64 = 8
const i64Const9 int64 = 9

// I64ToX32 converts an int64 to a fixed point number on 32 bits.
func I64ToX32(i int64) X32 {
	return X32(int32(i) << X32Shift)
}

// I64ToX64 converts an int64 to a fixed point number on 64 bits.
func I64ToX64(i int64) X64 {
	return X64(i << X64Shift)
}
