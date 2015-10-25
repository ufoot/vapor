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

package vpp2p

// NodeInfo stores the static data for a Node.
type NodeInfo struct {
	// 256-bit id, totally random, generated when instanciating node.
	NodeID []byte
	// Refers to the physical host. 
	HostID []byte
	// Refers to the ring on which this nodes participates.
	RingID []byte
}

// Node is the link between a ring (AKA a session) and a host (AKA a physical
// end-point).
type Node struct {
	// Info about the node
	Info NodeInfo

	ringPtr *Ring
	hostPtr *Host
}
