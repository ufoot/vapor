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

package vpmat4x3

import (
	"testing"
)

func TestConst(t *testing.T) {
	if Size == 12 {
		t.Logf("OK, 4x3 matrix size is %d", Size)
	} else {
		t.Errorf("wrong size for 4x3 matrix, got %d should be %d", Size, 12)
	}

	if Width*Height != Size {
		t.Errorf("wrong widht/height/size for 4x3 matrix, %d * %d != %d", Width, Height, Size)
	}
}
