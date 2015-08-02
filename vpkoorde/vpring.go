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

import (
	"fmt"
	"math/big"
)

// Ring contains the informations about a ring.
type Ring struct {
	m int
	n int

	// This node information.
	Local *Node
	// The successor nodes within the ring, use 1st elem for direct successor..
	Successor []*Node
	// A list of nodes preceeding m*Id (the 1st Bruijn node),
	// so that it contains about O(Log(n)) before stumbling on D.
	// The first element is actually D, the other ones go backwards on the ring.
	D []*Node
}

// Lookup performs a lookup on a given key. Returns all the nodes which
// must be traversed to go it. If there's only one element, it means the
// node is local. If there's no element, it could not be joined. At all.
func (*Ring) Lookup(i big.Int) ([]*Node, error) {
	// pseudo code :
	// procedure m.LOOKUP(k, shift, i)
	//   if k is in (m,successor] then return (successor)
	//   else if is in (m,successor] then return (
	//     d.lookup(k,
	//              kshift<<1,
	//              i o topBit(kshift)))
	//   else return (successor.lookup(k,shift,i))

	return nil,fmt.Errorf("TODO")
}
