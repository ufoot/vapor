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
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/ufoot/vapor/vpline3"
	"github.com/ufoot/vapor/vpvec3"
	"testing"
)

const x0 float64 = 10
const x1 float64 = 100
const y0 float64 = -100
const y1 float64 = -10
const z0 float64 = 90
const z1 float64 = 180

func TestF64Square(t *testing.T) {
	var square vpline3.F64
	const pngName string = "square.png"

	square = append(square, *vpvec3.F64New(x0, y0, z0))
	square = append(square, *vpvec3.F64New(x0, y1, z0))
	square = append(square, *vpvec3.F64New(x0, y1, z0))
	square = append(square, *vpvec3.F64New(x1, y1, z0))
	square = append(square, *vpvec3.F64New(x1, y1, z0))
	square = append(square, *vpvec3.F64New(x1, y0, z0))
	square = append(square, *vpvec3.F64New(x1, y0, z0))
	square = append(square, *vpvec3.F64New(x0, y0, z0))

	img := F64Demo(&square)
	err := draw2dimg.SaveToPngFile(pngName, img)
	if err == nil {
		t.Logf("saved \"%s\", %dx%d", pngName, img.Bounds().Max.X, img.Bounds().Max.Y)
	} else {
		t.Errorf("error saving example file \"%s\"", pngName)
	}
}
