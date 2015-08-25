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
	"math/big"
)

// Proxy is used to communicate with another node.
type Proxy interface {
	// GetInfo gets the remote node informations.
	GetInfo() *NodeInfo
	// Lookup searchs for a key, might recursively call other
	Lookup(i *big.Int) ([]*NodeInfo, error)	
}

// LocalProxy is used to communicate with a local node, this is
// for simulation or to host several distinct virtual on a single
// physical node.
type LocalProxy struct {
	localNode *Node
	ring *Ring
}
