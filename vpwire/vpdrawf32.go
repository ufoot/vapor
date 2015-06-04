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

package vpwire

import (
	"github.com/ufoot/vapor/vpline3"
	"github.com/ufoot/vapor/vpmat4x4"
	"github.com/ufoot/vapor/vpvec3"
	"image"
	"image/color"
)

// F32Proj calculates the projection for a line, so that all points in the
// line are visible, given a direction for the center "ray".
func F32Proj(img *image.RGBA, line *vpline3.F32, dir *vpvec3.F32) *vpmat4x4.F32 {
	var ret vpmat4x4.F32

	// todo...

	return &ret
}

// F32Draw draws a line on a image,
// line are visible, given a direction for the center "ray".
func F32Draw(img *image.RGBA, proj *vpmat4x4.F32, line *vpline3.F32, col color.RGBA) *image.RGBA {
	var ret image.RGBA

	// todo...

	return &ret
}
