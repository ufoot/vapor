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

const factCacheSize = 12

var factCache [factCacheSize]int

func init() {
	for i := range factCache {
		if i <= 1 {
			factCache[i] = 1
		} else {
			factCache[i] = i * factCache[i-1]
		}
	}
}

func calcFact(n int) int {
	if n > 1 {
		return n * calcFact(n-1)
	}

	return 1
}

// Fact returns the factorial of an integer.
// Operates on standard default int, caches
// small values for speed.
func Fact(n int) int {
	switch {
	case n >= factCacheSize:
		return n * Fact(n-1)
	case n >= 0:
		return factCache[n]
	}

	return 1
}
