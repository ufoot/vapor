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

// F64 is a line in a 3D space.
type F64 []vpvec3.F64

// F64NewSegment creates a new segment with 2 float64 vectors.
func F64NewSegment(a, b vpvec3.F64) *F64 {
	l := make([]vpvec3.F64, B+1)
	l[A] = a
	l[B] = b
	ret := F64(l)
	return &ret
}

// F64NewTriangle creates a new triangle with 3 float64 vectors.
func F64NewTriangle(a, b, c vpvec3.F64) *F64 {
	l := make([]vpvec3.F64, C+1)
	l[A] = a
	l[B] = b
	l[C] = c
	ret := F64(l)
	return &ret
}

// F64NewQuad creates a new line with 4 float64 vectors.
func F64NewQuad(a, b, c, d vpvec3.F64) *F64 {
	l := make([]vpvec3.F64, D+1)
	l[A] = a
	l[B] = b
	l[C] = c
	l[D] = d
	ret := F64(l)
	return &ret
}

// ToX32 converts the line to a fixed point number line on 32 bits.
func (line *F64) ToX32() *X32 {
	var ret X32

	for i, v := range *line {
		ret[i] = *v.ToX32()
	}

	return &ret
}

// ToX64 converts the line to a fixed point number line on 64 bits.
func (line *F64) ToX64() *X64 {
	var ret X64

	for i, v := range *line {
		ret[i] = *v.ToX64()
	}

	return &ret
}

// ToF32 converts the line to a float32 line.
func (line *F64) ToF32() *F32 {
	var ret F32

	for i, v := range *line {
		ret[i] = *v.ToF32()
	}

	return &ret
}

// String returns a readable form of the line.
func (line *F64) String() string {
	buf, err := json.Marshal(line)

	if err != nil {
		// Catching & ignoring error
		return ""
	}

	return string(buf)
}
