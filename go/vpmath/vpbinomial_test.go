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

package vpmath

import (
	"testing"
)

func TestBinomial(t *testing.T) {
	var binomialCheck = [4][4]int{{1, 0, 0, 0}, {1, 1, 0, 0}, {1, 2, 1, 0}, {1, 3, 3, 1}}
	const b21Check = 2
	const b31Check = 3
	var binomialNI int

	for n := range binomialCheck {
		for i := range binomialCheck[n] {
			binomialNI = Binomial(n, i)
			if binomialNI == binomialCheck[n][i] {
				t.Logf("OK, Binomial(%d,%d)=%d", n, i, binomialNI)
			} else {
				t.Errorf("problem, Binomial(%d,%d)=%d should have got %d", n, i, binomialNI, binomialCheck[n][i])
			}
		}
	}
}
