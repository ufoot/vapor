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

//go:generate bash ./stamp.sh
	
// DrawMode defines how to draw stuff, including lines.
// Inspired from OpenGL GL_POINTS, GL_LINES or GL_TRIANGLES.
type DrawMode int

const (
	// DrawPoints causes all points to be drawn individually, as points.
	DrawPoints DrawMode = iota
	// DrawLines is used to to draw lines with odd index points considered
	// as the begin point of line number index/2 and even index points considered
	// as the end point of line (index-1)/2.
	DrawLines
	// DrawTriangles groups points by pack of three, and then draws a triangle
	// joining these three points.
	DrawTriangles
)

// DemoWidth is the width used for demo/test rendering.
const DemoWidth int = 1600

// DemoHeight is the height used for demo/test rendering.
const DemoHeight int = 900

const demoDirRatio int = 2
