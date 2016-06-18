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

//go:generate bash ./stamp.sh
	
// LookupSplit builds an array of integers containing indices
// for a given lookup table. The idea is to consider the id is a
// sequence number within a N-dimension space. It arranges the
// indices so that items with minor (last parameters) values
// that are closed together are closed together in memory.
func LookupSplit(sizes []int, id int) []int {
	n := len(sizes)
	ret := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		s := sizes[i]
		ret[i] = id % s
		id /= s
	}

	return ret
}

// LookupJoin builds an index from a list of sub indexes
// for a given lookup table. The idea is to consider the id is a
// sequence number within a N-dimension space.
func LookupJoin(sizes, indexes []int) int {
	n := len(sizes)
	var ret int

	f := int(1)
	for i := n - 1; i >= 0; i-- {
		ret += indexes[i] * f
		f *= sizes[i]
	}

	return ret
}

// LookupSize returns the size (maximum index + 1)
// for a given lookup table.
func LookupSize(indexes []int) int {
	ret := 1

	for _, v := range indexes {
		ret *= v
	}

	return ret
}
