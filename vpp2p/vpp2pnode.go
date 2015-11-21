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
	// 256-bit id, generated from signing HostPubKey and RingID,
	// could be recalculated dynamically, but cached for efficiency.
	NodeID []byte
	// Refers to the physical host.
	HostPubKey []byte
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

	// The successor nodes within the ring, use 1st elem for direct successor..
	successor []*Node
	// A list of nodes preceeding m*Id (the 1st Bruijn node),
	// so that it contains about O(Log(n)) before stumbling on D.
	// The first element is actually D, the other ones go backwards on the ring.
	D []*Node
}

// NodeProxy is an interface used to perform node operations.
// All calls return the complete call stack
type NodeProxy interface {
	NodeID() []byte
	Lookup(key []byte) ([]*NodeInfo, error)
	Set(key, value []byte) ([]*NodeInfo, error)
	Get(key []byte) ([]byte, []*NodeInfo, error)
	Clear(key []byte) ([]*NodeInfo, error)
	// GetRange(key1, key2 []byte) TODO !
}

// LocalProxy is a proxy which contacts other nodes directly without
// using any network interface whatsoever.
type LocalProxy struct {
	localNode Node
}

// Lookup a key and return the path of nodes to this key.
func (LocalProxy) Lookup(key []byte) ([]*NodeInfo, error) {
	// pseudo code :
	// procedure m.LOOKUP(k, kshift, i)
	//   if k is in (m,successor] then return (successor)
	//   else if i is in (m,successor] then return (
	//     d.lookup(k,
	//              kshift<<1,
	//              i o topBit(kshift)))
	//   else return (successor.lookup(k,kshift,i))
	// Note : i can be chosen so that its low bits are top bits of k

	return nil, nil // todo
}

// Set a key and return the path of nodes to this key.
func (LocalProxy) Set(key, value []byte) ([]*NodeInfo, error) {
	return nil, nil // todo
}

// Get a key and returns the path of nodes to this key.
func (LocalProxy) Get(key []byte) ([]byte, []*NodeInfo, error) {
	return nil, nil, nil // todo
}

// Clear a key and returns the path of nodes to this key.
func (LocalProxy) Clear(key []byte) ([]*NodeInfo, error) {
	return nil, nil // todo
}
