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

//go:generate bash ./stamp.sh
	
// I32Const0 contains 0 as an int32.
const I32Const0 int32 = 0

// I32Const1 contains 1 as an int32.
const I32Const1 int32 = 1

const i32Const2 int32 = 2
const i32Const3 int32 = 3
const i32Const4 int32 = 4
const i32Const5 int32 = 5
const i32Const6 int32 = 6
const i32Const7 int32 = 7
const i32Const8 int32 = 8
const i32Const9 int32 = 9

// I32ToX32 converts an int32 to a fixed point number on 32 bits.
func I32ToX32(i int32) X32 {
	return X32(i << X32Shift)
}

// I32ToX64 converts an int32 to a fixed point number on 64 bits.
func I32ToX64(i int32) X64 {
	return X64(int64(i) << X64Shift)
}
