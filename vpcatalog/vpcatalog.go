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

package vpcatalog

//import (
//	"fmt"
//	"math/big"
//)

// Phost represents a physical, unique host.
// Typically identified by a physical endpoint (IP connection).
type Phost struct {
}

// Vhost represents a virtual host.
// Typically identified by a public key. Bound to an IP, but this
// could change over time, and a given IP could hide several
// vhosts behind multiplexing.
type Vhost struct {
}

// Pring represents a physical ring.
// Typically identified by a world/context ID, unique on the net.
type Pring struct {
}

// Vring represents a virtual ring.
// Typically identified by an ID deduced from the physical ring,
// but different for each ring.
type Vring struct {
}

// Pnode represents a physical node.
// Typically identified by a node ID, totally random.
// There can be several nodes per hosts, to have better sharding.
type Pnode struct {
}

// Vnode represents a virtual node.
// Typically identified by an ID deduced from the physical node,
// but different for each vnode.
// Vnode can be seen as the intersection of Vrings and Pnodes.
type Vnode struct {
}

// Catalog is the interface for catalogs which keep a list
// of hosts, rings and nodes.
type Catalog interface {
	GetPhosts() []Phost
	GetPhostByID() Phost
	GetPhostsByAppID() []Phost
	GetPhostsByURL() []Phost

	GetVhosts() []Vhost
	GetVhostByID() Vhost
	GetVhostsByPhostId() []Vhost
	GetVhostsByAppID() []Vhost
	GetVhostsByURL() []Vhost

	GetPrings() []Pring
	GetPringByID() Pring
	GetPringsByAppID() []Pring
	GetPringsByRingName() []Pring

	GetVrings() []Vring
	GetVringByID() Vring
	GetVringsByPringId() []Vring
	GetVringsByAppID() []Vring
	GetVringsByRingName() []Vring

	GetPnodes() []Pnode
	GetPnodeByID() Pnode
	GetPnodesByPhostID() []Pnode
	GetPnodesByPringID() []Pnode
	GetPnodesByAppID() []Pnode
	GetPnodesByRingName() []Pnode
	GetPnodesByURL() []Pnode

	GetVnodes() []Vnode
	GetVnodeByID() Vnode
	GetVnodesByVhostID() []Vnode
	GetVnodesByVringID() []Vnode
	GetVnodesByPnodeId() []Vnode
	GetVnodesByPhostID() []Vnode
	GetVnodesByPringID() []Vnode
	GetVnodesByAppID() []Vnode
	GetVnodesByRingName() []Vnode
	GetVnodesByURL() []Vnode
}
