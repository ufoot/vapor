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
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
	"image"
	"image/color"
)

func f32ProjCenterAndScale(line *vpline3.F32) *vpmat4x4.F32 {
	ret := vpmat4x4.F32Identity()

	min := line.Reduce(vpvec3.F32Min)
	max := line.Reduce(vpvec3.F32Max)
	avg := line.Reduce(vpvec3.F32Add).DivScale(float32(len(*line)))

	var diff float32
	for i, v := range avg {
		if v-min[i] > diff {
			diff = v - min[i]
		}
		if max[i]-v > diff {
			diff = max[i] - v
		}
	}

	ret.MulComp(vpmat4x4.F32Translation(avg))

	ret.MulComp(vpmat4x4.F32Scale(vpvec3.F32New(diff, diff, diff)))

	return ret
}

func f32ProjCameraRotate(dir *vpvec3.F32) *vpmat4x4.F32 {
	cameraDirX := vpvec3.F32Normalize(dir)
	cameraDirZ := vpvec3.F32Cross(cameraDirX, vpvec3.F32AxisY()).Normalize()
	cameraDirY := vpvec3.F32Cross(cameraDirZ, cameraDirX).Normalize()

	ret := vpmat4x4.F32RebaseOXYZ(new(vpvec3.F32), cameraDirX, cameraDirY, cameraDirZ)

	return ret
}

func f32ProjPerspectiveToWorld() *vpmat4x4.F32 {
	ret := vpmat4x4.F32RebaseOXYZ(vpvec3.F32New(-vpnumber.F32Const1, -vpnumber.F32Const1, vpnumber.F32Const1), vpvec3.F32New(vpnumber.F32Const1, -vpnumber.F32Const1, vpnumber.F32Const1), vpvec3.F32New(-vpnumber.F32Const1, vpnumber.F32Const1, vpnumber.F32Const1), vpvec3.F32New(-vpnumber.F32Const1, -vpnumber.F32Const1, -vpnumber.F32Const1))

	return ret
}

func f32ProjPerspective(img *image.RGBA) *vpmat4x4.F32 {
	bounds := img.Bounds()
	var topLeftCorner image.Point
	var bottomRightCorner image.Point
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	if width > height {
		topLeftCorner.X = (width - height) >> 1
		bottomRightCorner.X = topLeftCorner.X + width
	} else {
		topLeftCorner.Y = (height - width) >> 1
		bottomRightCorner.Y = topLeftCorner.Y + height
	}

	perspLerp := float32(0.25)
	ret := vpmat4x4.F32RebaseOXYZP(vpvec3.F32New(float32(topLeftCorner.X), float32(bottomRightCorner.Y), vpnumber.F32Const0),
		vpvec3.F32New(float32(topLeftCorner.X), float32(bottomRightCorner.Y), vpnumber.F32Const0),
		vpvec3.F32New(float32(topLeftCorner.X), vpnumber.F32Const0, vpnumber.F32Const0),
		vpvec3.F32New(vpmath.F32Lerp(float32(topLeftCorner.X), float32(bottomRightCorner.X), perspLerp), vpmath.F32Lerp(float32(bottomRightCorner.Y), float32(topLeftCorner.Y), perspLerp), vpnumber.F32Const1),
		vpvec3.F32New(vpmath.F32Lerp(float32(bottomRightCorner.X), float32(topLeftCorner.X), perspLerp), vpmath.F32Lerp(float32(topLeftCorner.Y), float32(bottomRightCorner.Y), perspLerp), vpnumber.F32Const1))

	return ret
}

// F32Proj calculates the projection for a line, so that all points in the
// line are visible, given a direction for the center "ray".
func F32Proj(img *image.RGBA, line *vpline3.F32, dir *vpvec3.F32) *vpmat4x4.F32 {
	ret := vpmat4x4.F32Identity()

	ret.MulComp(f32ProjCenterAndScale(line))
	ret.MulComp(f32ProjCameraRotate(dir))
	ret.MulComp(f32ProjPerspectiveToWorld())

	ret.Inv()

	ret = vpmat4x4.F32MulComp(f32ProjPerspective(img), ret)

	return ret
}

// F32Draw draws a line on a image,
// line are visible, given a direction for the center "ray".
func F32Draw(img *image.RGBA, proj *vpmat4x4.F32, line *vpline3.F32, col color.RGBA) *image.RGBA {
	var ret image.RGBA

	// todo...

	return &ret
}
