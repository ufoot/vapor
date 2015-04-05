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

package vpmat4x4

const (
	// Col0Row0 is the index of colum 0 row 0 in a 4x4 column-major matrix.
	Col0Row0 = iota
	// Col0Row1 is the index of colum 0 row 1 in a 4x4 column-major matrix.
	Col0Row1
	// Col0Row2 is the index of colum 0 row 2 in a 4x4 column-major matrix.
	Col0Row2
	// Col0Row3 is the index of colum 0 row 3 in a 4x4 column-major matrix.
	Col0Row3
	// Col1Row0 is the index of colum 1 row 0 in a 4x4 column-major matrix.
	Col1Row0
	// Col1Row1 is the index of colum 1 row 1 in a 4x4 column-major matrix.
	Col1Row1
	// Col1Row2 is the index of colum 1 row 2 in a 4x4 column-major matrix.
	Col1Row2
	// Col1Row3 is the index of colum 1 row 3 in a 4x4 column-major matrix.
	Col1Row3
	// Col2Row0 is the index of colum 2 row 0 in a 4x4 column-major matrix.
	Col2Row0
	// Col2Row1 is the index of colum 2 row 1 in a 4x4 column-major matrix.
	Col2Row1
	// Col2Row2 is the index of colum 2 row 2 in a 4x4 column-major matrix.
	Col2Row2
	// Col2Row3 is the index of colum 2 row 3 in a 4x4 column-major matrix.
	Col2Row3
	// Col3Row0 is the index of colum 3 row 0 in a 4x4 column-major matrix.
	Col3Row0
	// Col3Row1 is the index of colum 3 row 1 in a 4x4 column-major matrix.
	Col3Row1
	// Col3Row2 is the index of colum 3 row 2 in a 4x4 column-major matrix.
	Col3Row2
	// Col3Row3 is the index of colum 3 row 3 in a 4x4 column-major matrix.
	Col3Row3
	// Size is the number of elements in a 4x4 matrix.
	Size
	// Width is the width a 4x4 matrix (max column index is Width-1).
	Width=4
	// Height is the width a 4x4 matrix (max row index is Height-1).
	Height=4
)
