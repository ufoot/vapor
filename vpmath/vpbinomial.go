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

package vpmath

const binomialCacheSize = 5

var binomialCache [binomialCacheSize][binomialCacheSize]int

func init() {
	for n := range binomialCache {
		for i := range binomialCache[n] {
			binomialCache[n][i] = calcBinomial(n, i)
		}
	}
}

func calcBinomial(n, i int) int {
	switch {
	case n < i:
		return 0
	case n <= 1:
		return 1
	}

	return calcFact(n) / (calcFact(i) * calcFact(n-i))
}

// Calculates the Binomial for n,i, for Bertein polynomials
func Binomial(n, i int) int {
	if n < binomialCacheSize && i < binomialCacheSize {
		return binomialCache[n][i]
	}

	return calcBinomial(n, i)
}
