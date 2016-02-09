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

package vpwire

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/ufoot/vapor/vpline3"
	"github.com/ufoot/vapor/vpmat4x4"
	"github.com/ufoot/vapor/vpmath"
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec3"
	"image"
	"image/color"
)

func f64ProjCenterAndScale(line *vpline3.F64) *vpmat4x4.F64 {
	ret := vpmat4x4.F64Identity()

	min := line.Reduce(vpvec3.F64Min)
	max := line.Reduce(vpvec3.F64Max)
	avg := line.Reduce(vpvec3.F64Add).DivScale(float64(len(*line)))

	var diff float64
	for i, v := range avg {
		if v-min[i] > diff {
			diff = v - min[i]
		}
		if max[i]-v > diff {
			diff = max[i] - v
		}
	}

	ret.MulComp(vpmat4x4.F64Translation(avg))

	ret.MulComp(vpmat4x4.F64Scale(vpvec3.F64New(diff, diff, diff)))

	return ret
}

func f64ProjCameraRotate(dir *vpvec3.F64) *vpmat4x4.F64 {
	cameraDirX := vpvec3.F64Normalize(dir)
	cameraDirZ := vpvec3.F64Cross(cameraDirX, vpvec3.F64AxisY()).Normalize()
	cameraDirY := vpvec3.F64Cross(cameraDirZ, cameraDirX).Normalize()

	ret := vpmat4x4.F64RebaseOXYZ(new(vpvec3.F64), cameraDirX, cameraDirY, cameraDirZ)

	return ret
}

func f64ProjPerspectiveToWorld() *vpmat4x4.F64 {
	ret := vpmat4x4.F64RebaseOXYZ(vpvec3.F64New(-vpnumber.F64Const1, -vpnumber.F64Const1, vpnumber.F64Const1), vpvec3.F64New(vpnumber.F64Const1, -vpnumber.F64Const1, vpnumber.F64Const1), vpvec3.F64New(-vpnumber.F64Const1, vpnumber.F64Const1, vpnumber.F64Const1), vpvec3.F64New(-vpnumber.F64Const1, -vpnumber.F64Const1, -vpnumber.F64Const1))

	return ret
}

func f64ProjPerspective(img *image.RGBA) *vpmat4x4.F64 {
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

	perspLerp := float64(0.25)
	ret := vpmat4x4.F64RebaseOXYZP(vpvec3.F64New(float64(topLeftCorner.X), float64(bottomRightCorner.Y), vpnumber.F64Const0),
		vpvec3.F64New(float64(topLeftCorner.X), float64(bottomRightCorner.Y), vpnumber.F64Const0),
		vpvec3.F64New(float64(topLeftCorner.X), vpnumber.F64Const0, vpnumber.F64Const0),
		vpvec3.F64New(vpmath.F64Lerp(float64(topLeftCorner.X), float64(bottomRightCorner.X), perspLerp), vpmath.F64Lerp(float64(bottomRightCorner.Y), float64(topLeftCorner.Y), perspLerp), vpnumber.F64Const1),
		vpvec3.F64New(vpmath.F64Lerp(float64(bottomRightCorner.X), float64(topLeftCorner.X), perspLerp), vpmath.F64Lerp(float64(topLeftCorner.Y), float64(bottomRightCorner.Y), perspLerp), vpnumber.F64Const1))

	return ret
}

// F64Proj calculates the projection for a line, so that all points in the
// line are visible, given a direction for the center "ray".
func F64Proj(line *vpline3.F64, img *image.RGBA, dir *vpvec3.F64) *vpmat4x4.F64 {
	ret := vpmat4x4.F64Identity()

	ret.MulComp(f64ProjCenterAndScale(line))
	ret.MulComp(f64ProjCameraRotate(dir))
	ret.MulComp(f64ProjPerspectiveToWorld())

	ret.Inv()

	ret = vpmat4x4.F64MulComp(f64ProjPerspective(img), ret)

	return ret
}

// F64Draw draws a line on an image, the image is modified in-place, and
// returned modified.
func F64Draw(img *image.RGBA, line *vpline3.F64, proj *vpmat4x4.F64, mode DrawMode, col color.Color) *image.RGBA {
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetStrokeColor(col)

	switch mode {
	case DrawPoints:
		for _, vecWorld := range *line {
			vecDraw := proj.MulVecPos(&vecWorld)
			gc.MoveTo(vecDraw[vpvec3.X], vecDraw[vpvec3.Y])
			gc.LineTo(vecDraw[vpvec3.X], vecDraw[vpvec3.Y])
			gc.Stroke()
		}
	case DrawLines:
		for i := 0; i < len(*line)-1; i += 2 {
			vecWorld := (*line)[i]
			vecDraw := proj.MulVecPos(&vecWorld)
			gc.MoveTo(vecDraw[vpvec3.X], vecDraw[vpvec3.Y])
			vecDraw = proj.MulVecPos(&vecWorld)
			vecWorld = (*line)[i]
			gc.LineTo(vecDraw[vpvec3.X], vecDraw[vpvec3.Y])
			gc.Stroke()
		}
	case DrawTriangles:
		// todo
	}

	return img
}

// F64Demo draws a line for demo purposes, creates an image on the fly,
// with default point-of-view, perspective and other parameters, can typically
// be used for testing, not testing this package, but possibly testing out
// point transformations functions (Bezier curves is an example).
func F64Demo(line *vpline3.F64) *image.RGBA {
	rect := image.Rect(0, 0, DemoWidth, DemoHeight)
	ret := image.NewRGBA(rect)
	bg := color.RGBA{0xff, 0xff, 0xff, 0xff}

	for x := 0; x < DemoWidth; x++ {
		for y := 0; y < DemoHeight; y++ {
			ret.SetRGBA(x, y, bg)
		}
	}

	dir := vpvec3.F64New(vpnumber.F64Const1, -vpnumber.F64Const1, -vpnumber.F64Const1*float64(demoDirRatio))
	proj := F64Proj(line, ret, dir)

	return F64Draw(ret, line, proj, DrawLines, color.Black)
}
