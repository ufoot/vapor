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

package vpp2p

import (
	"fmt"
	"github.com/ufoot/vapor/vplog"
)

// NodeInfo stores the static data for a Node.
type NodeInfo struct {
	// 256-bit id, generated randomly, this is what allows us to
	// have several nodes per host/ring.
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

	hostPtr *Host
	ringPtr *Ring

	// Successors is list of successing nodes within the ring,
	// use 1st elem for direct successor.
	Successor []NodeProxy
	// PredecessorInfo describes the previous node within the ring. This is not used
	// to walk along the ring but just to know what is the range of keys
	// this node is responsible for.
	PredecessorInfo *NodeInfo
	// D is a  list of nodes preceeding m*Id (the 1st Bruijn node),
	// so that it contains about O(Log(n)) before stumbling on D.
	// The first element is actually D, the other ones go backwards on the ring.
	D []NodeProxy
}

// NodeProxy is an interface used to perform node operations.
// All calls return the complete call stack
type NodeProxy interface {
	// Info returns information about the Node, this should be a
	// very fast function, with data kept locally, and should not
	// require any network calls.
	Info() *NodeInfo
	// Lookup searches for the host responsible for a given key
	// and returns the path found to it.
	Lookup(key, keyShift, imaginaryNode []byte) ([]*NodeInfo, error)
	// Set sets a key to a value, returns the path to the node.
	Set(key, keyShift, imaginaryNode, value []byte) ([]*NodeInfo, error)
	// Get returns the value of a key, and returns the path to the node.
	Get(key, keyShift, imaginaryNode []byte) ([]byte, []*NodeInfo, error)
	// Get clears a key, and returns the path to the node.
	Clear(key, keyShift, imaginaryNode []byte) ([]*NodeInfo, error)
}

// LocalProxy is a proxy which contacts other nodes directly without
// using any network interface whatsoever.
type LocalProxy struct {
	localNode Node
}

// NewLocalProxy creates a new local proxy.
func NewLocalProxy(nodeID []byte, hostPtr *Host, ringPtr *Ring) (*LocalProxy, error) {
	var ret LocalProxy

	ok, err := CheckID(nodeID)
	if err != nil || !ok {
		return nil, err
	}

	ret.localNode.Info.NodeID = ringPtr.walker.Filter(nodeID)
	ret.localNode.Info.HostPubKey = make([]byte, len(hostPtr.Info.HostPubKey))
	copy(ret.localNode.Info.HostPubKey, hostPtr.Info.HostPubKey)
	ret.localNode.Info.RingID = make([]byte, len(ringPtr.Info.RingID))
	copy(ret.localNode.Info.RingID, ringPtr.Info.RingID)
	ret.localNode.hostPtr = hostPtr
	ret.localNode.ringPtr = ringPtr

	return &ret, nil
}

// Info returns this node's info.
func (lp *LocalProxy) Info() *NodeInfo {
	return &(lp.localNode.Info)
}

// Lookup a key and return the path of nodes to this key.
func (lp *LocalProxy) Lookup(key, keyShift, imaginaryNode []byte) ([]*NodeInfo, error) {
	// pseudo code :
	// procedure m.LOOKUP(k, kshift, i)
	//   if k is in (m,successor] then return (successor)
	//   else if i is in (m,successor] then return (
	//     d.lookup(k,
	//              kshift<<1,
	//              i o topBit(kshift)))
	//   else return (successor.lookup(k,kshift,i))
	// Note : i can be chosen so that its low bits are top bits of k

	node := lp.localNode
	walker := lp.localNode.ringPtr.walker

	ret := make([]*NodeInfo, 1)
	ret[0] = &(node.Info)

	if node.Successor == nil || len(node.Successor) == 0 || node.D == nil || len(node.D) == 0 || node.PredecessorInfo == nil {
		// no successor -> we're alone !
		return ret, nil
	}

	if walker.GtLe(key, node.PredecessorInfo.NodeID, node.Info.NodeID) {
		// key is local
		return ret, nil
	}

	for i, successor := range node.Successor {
		var curInfo *NodeInfo
		if i == 0 {
			curInfo = &(node.Info)
		} else {
			curInfo = node.Successor[i-1].Info()
		}

		successorInfo := successor.Info()
		if walker.GtLe(key, curInfo.NodeID, successorInfo.NodeID) {
			// key is handled by successor
			return append(ret, successorInfo), nil
		}

		if walker.GtLe(imaginaryNode, curInfo.NodeID, successorInfo.NodeID) {
			for _, d := range node.D {
				upstreamPath, err := d.Lookup(key, walker.NextFirst(keyShift), walker.ForwardElem(keyShift, imaginaryNode, 1))
				if err != nil {
					vplog.LogDebug("error contacting remote host", err)
				}
				return append(ret, upstreamPath...), nil
			}
		}
	}

	// at this stage, key is not local, not handled by any direct host
	// we know and De Bruijn walking did not return anything interesting.
	// So we fall back on the default : ask next node...
	upstreamPath, err := node.Successor[0].Lookup(key, keyShift, imaginaryNode)
	if err != nil {
		return nil, err
	}
	if upstreamPath != nil {
		return append(ret, upstreamPath...), nil
	}

	return nil, fmt.Errorf("unable to Lookup remotes")
}

// Set a key and return the path of nodes to this key.
func (lp *LocalProxy) Set(key, keyShift, imaginaryNode, value []byte) ([]*NodeInfo, error) {
	return nil, nil // todo
}

// Get a key and returns the path of nodes to this key.
func (lp *LocalProxy) Get(key, keyShift, imaginaryNode []byte) ([]byte, []*NodeInfo, error) {
	return nil, nil, nil // todo
}

// Clear a key and returns the path of nodes to this key.
func (lp *LocalProxy) Clear(key, keyShift, imaginaryNode []byte) ([]*NodeInfo, error) {
	return nil, nil // todo
}
