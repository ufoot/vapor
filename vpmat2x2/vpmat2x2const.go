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

package vpmat2x2

const (
	// Col0Row0 is the index of colum 0 row 0 in a 2x2 column-major matrix.
	Col0Row0 = iota
	// Col0Row1 is the index of colum 0 row 1 in a 2x2 column-major matrix.
	Col0Row1
	// Col1Row0 is the index of colum 1 row 0 in a 2x2 column-major matrix.
	Col1Row0
	// Col1Row1 is the index of colum 1 row 1 in a 2x2 column-major matrix.
	Col1Row1
	// Size is the number of elements in a 2x2 matrix.
	Size
)
