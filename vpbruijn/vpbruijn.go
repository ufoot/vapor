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

package vpbruijn

// Read https://en.wikipedia.org/wiki/De_Bruijn_graph for theory

// This is an interface over implementation of De Bruijn graphs.

import (
	"fmt"
)

// BruijnWalker allows walking along a Bruijn graph, going forward, backward,
// getting list of previous or next nodes.
type BruijnWalker interface {
	// M returns the m paramater (AKA base) for the Bruijn network.
	M() int
	// N returns the m paramater (number of elements) for the Bruijn network.
	N() int
	// NbBits returns the number of bits of keys for this Bruijn network.
	NbBits() int
	// NbBytes returns the number of bytes of keys for this Bruijn network.
	NbBytes() int
	// NextFirst returns the first Bruijn node pointed by this node.
	// Other nodes might be deduced by just incrementing this one.
	NextFirst(key []byte) []byte
	// NextLast returns the last Bruijn node pointing to this node.
	// Other nodes might be deduced by just decrementing this one with
	// a value of m**(n-1).
	NextLast(key []byte) []byte
	// NextList returns the list of all Bruijn nodes pointed by
	// this node, the nodes following this one (we walk down the graph).
	NextList(key []byte) [][]byte
	// PrevFirst returns the first Bruijn node pointing to this node.
	// Other nodes might be deduced by just incrementing this one with
	// a value of m**(n-1).
	PrevFirst(key []byte) []byte
	// PrevLast returns the last Bruijn node pointing to this node.
	// Other nodes might be deduced by just decrementing this one with
	// a value of m**(n-1).
	PrevLast(key []byte) []byte
	// PrevList returns the list of all Bruijn nodes pointing to
	// this node, the nodes preceding this one (we walk up the graph).
	PrevList(key []byte) [][]byte
	// ForwardPath returns the path between two nodes. The path
	// might be non-optimized, it always contains m+1 elements, including
	// from and to destination. This is the default forward path in which
	// node n+1 is the node after n in the bruijn sequence.
	ForwardPath(from, to []byte) [][]byte
	// BackwardPath returns the path between two nodes. The path
	// might be non-optimized, it always contains m+1 elements, including
	// from and to destination. This is the alternative backward path in which
	// node n+1 is the node before n in the bruijn sequence.
	BackwardPath(from, to []byte) [][]byte
	// ForwardElem returns the path element between two nodes.
	// Index 0 is the from element, and n (number of elements as in Bruijn nodes)
	// the to element. Uses the forward, default path.
	ForwardElem(from, to []byte, i int) []byte
	// BackwardElem returns the path element between two nodes.
	// Index 0 is the from element, and n (number of elements as in Bruijn nodes)
	// the to element. Uses the backward, alternative path.
	BackwardElem(from, to []byte, i int) []byte
	// Add 2 keys, if result is too big, will loop with Mod() and keep it
	// within the allowed key range.
	Add(x, y []byte) []byte
	// Sub substracts a key from another, if result is too small, will loop with
	// Mod() and keep it within the allowed key range.
	Sub(x, y []byte) []byte
	// Cmp compares two keys, returns 1 is y>x, -1 if x<y and 0 if they are equal.
	// It considers values reprensent keys on a ring, so if y is really greater
	// than x (that is, more than half-way considering the lenght of the ring) then
	// y is considered smaller than x. So one can have a>b and b>c and c>a.
	Cmp(x, y []byte) int
	// GeLt tells wether a key is between two other keys. It considers keys are
	// on a ring so if begin is before end, tests x>=begin and x<end, and if
	// end is before begin, tests x<end or x>=begin.
	GeLt(x, begin, end []byte) bool
	// GtLe tells wether a key is between two other keys. It considers keys are
	// on a ring so if begin is before end, tests x>begin and x<=end, and if
	// end is before begin, tests x<=end or x>begin.
	GtLe(x, begin, end []byte) bool
}

// BruijnNew creates a new BruijnWalker compatible object,
// which can be use to walk along De Bruijn networks.
func BruijnNew(m, n int) (BruijnWalker, error) {
	if m == 16 && n == 64 {
		return Bruijn16x64New(), nil
	}
	if m >= GenericMinM && m <= GenericMaxM && n >= GenericMinN && n <= GenericMaxN {
		return BruijnGenericNew(m, n), nil
	}
	return nil, fmt.Errorf("no implementation for De Bruijn networks with values m=%d n=%d", m, n)
}
