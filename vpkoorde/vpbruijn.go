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

package vpkoorde

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
	NextFirst(key []byte) ([]byte, error)
	// NextLast returns the last Bruijn node pointing to this node.
	// Other nodes might be deduced by just decrementing this one with
	// a value of m**(n-1).
	NextLast(key []byte) ([]byte, error)
	// NextList returns the list of all Bruijn nodes pointed by
	// this node, the nodes following this one (we walk down the graph).
	NextList(key []byte) ([][]byte, error)
	// PrevFirst returns the first Bruijn node pointing to this node.
	// Other nodes might be deduced by just incrementing this one with
	// a value of m**(n-1).
	PrevFirst(key []byte) ([]byte, error)
	// PrevLast returns the last Bruijn node pointing to this node.
	// Other nodes might be deduced by just decrementing this one with
	// a value of m**(n-1).
	PrevLast(key []byte) ([]byte, error)
	// PrevList returns the list of all Bruijn nodes pointing to
	// this node, the nodes preceding this one (we walk up the graph).
	PrevList(key []byte) ([][]byte, error)
	// ForwardPath returns the path between two nodes. The path
	// might be non-optimized, it always contains m+1 elements, including
	// from and to destination. This is the default forward path in which
	// node n+1 is the node after n in the bruijn sequence.
	ForwardPath(from, to []byte) ([][]byte, error)
	// BackwardPath returns the path between two nodes. The path
	// might be non-optimized, it always contains m+1 elements, including
	// from and to destination. This is the alternative backward path in which
	// node n+1 is the node before n in the bruijn sequence.
	BackwardPath(from, to []byte) ([][]byte, error)
	// ForwardElem returns the path element between two nodes.
	// Index 0 is the from element, and n (number of elements as in Bruijn nodes)
	// the to element. Uses the forward, default path.
	ForwardElem(from, to []byte, i int) ([]byte, error)
	// BackwardElem returns the path element between two nodes.
	// Index 0 is the from element, and n (number of elements as in Bruijn nodes)
	// the to element. Uses the backward, alternative path.
	BackwardElem(from, to []byte, i int) ([]byte, error)
}

// BruijnNew creates a new BruijnWalker compatible object,
// which can be use to walk along De Bruijn networks.
func BruijnNew(m, n int) (BruijnWalker, error) {
	if m == 16 && n == 64 {
		return bruijn16x64{}, nil
	}
	if m >= GenericMinM && m <= GenericMaxM && n >= GenericMinN && n <= GenericMaxN {
		return bruijnGenericNew(m, n), nil
	}
	return nil, fmt.Errorf("no implementation for De Bruijn networks with values m=%d n=%d", m, n)
}
