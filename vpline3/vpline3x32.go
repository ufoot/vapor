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

package vpline3

import (
	"encoding/json"
	"github.com/ufoot/vapor/vpvec3"
)

// X32 is a line in a 3D space.
type X32 []vpvec3.X32

// X32NewSegment creates a new segment with 2 float64 vectors.
func X32NewSegment(a, b vpvec3.X32) *X32 {
	l := make([]vpvec3.X32, B+1)
	l[A] = a
	l[B] = b
	ret := X32(l)
	return &ret
}

// X32NewTriangle creates a new triangle with 3 float64 vectors.
func X32NewTriangle(a, b, c vpvec3.X32) *X32 {
	l := make([]vpvec3.X32, C+1)
	l[A] = a
	l[B] = b
	l[C] = c
	ret := X32(l)
	return &ret
}

// X32NewQuad creates a new line with 4 float64 vectors.
func X32NewQuad(a, b, c, d vpvec3.X32) *X32 {
	l := make([]vpvec3.X32, D+1)
	l[A] = a
	l[B] = b
	l[C] = c
	l[D] = d
	ret := X32(l)
	return &ret
}

// ToX64 converts the line to a fixed point number line on 64 bits.
func (line *X32) ToX64() *X64 {
	var ret X64

	for i, v := range *line {
		ret[i] = *v.ToX64()
	}

	return &ret
}

// ToF32 converts the line to a float32 line.
func (line *X32) ToF32() *F32 {
	var ret F32

	for i, v := range *line {
		ret[i] = *v.ToF32()
	}

	return &ret
}

// ToF64 converts the line to a float64 line.
func (line *X32) ToF64() *F64 {
	var ret F64

	for i, v := range *line {
		ret[i] = *v.ToF64()
	}

	return &ret
}

// String returns a readable form of the line.
func (line *X32) String() string {
	buf, err := json.Marshal(line)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}
