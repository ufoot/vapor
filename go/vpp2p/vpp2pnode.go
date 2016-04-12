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
	"github.com/ufoot/vapor/go/vpid"
	"github.com/ufoot/vapor/go/vplog"
	"github.com/ufoot/vapor/go/vpp2papi"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"github.com/ufoot/vapor/go/vpsum"
	"math/big"
	"sync"
)

const (
	// NodeKeySeconds specifies how many seconds should we spend on creating
	// node keys in signed mode.
	NodeKeySeconds = 1
	// NodeKeyZeroes specifies how many zeroes there should be at the end of
	// a node key sig in signed mode.
	NodeKeyZeroes = 8
)

// Node is the link between a Ring and a Host. Basically a node is a point
// on the ring, which joins on the ring using the host as a connexion tool.
type Node struct {
	// Status about the node
	Status vpp2papi.NodeStatus

	hostPtr *Host
	ringPtr *Ring

	registerers []NodeRegisterer
	up          bool

	successorsAccess  sync.RWMutex
	predecessorAccess sync.RWMutex
	dAccess           sync.RWMutex

	// Successors is list of successing nodes within the ring,
	// use 1st elem for direct successor.
	Successor []vpp2papi.VpP2pApi
	// D is a  list of nodes preceeding m*Id (the 1st Bruijn node),
	// so that it contains about O(Log(n)) before stumbling on D.
	// The first element is actually D, the other ones go backwards on the ring.
	D []vpp2papi.VpP2pApi
}

// NodeRegisterer is and interface that records node registring and unregistring.
type NodeRegisterer interface {
	// RegisterNode should be called when the node is started, ready for action.
	RegisterNode(node *Node)
	// UnregisterNode should be called when node is stopped, not responding any more.
	UnregisterNode(node *Node)
}

type ringIDAppender struct {
	ringID []byte
}

func (ria *ringIDAppender) Transform(nodeID []byte) []byte {
	ni := vpp2papi.NewNodeInfo()
	ni.NodeID = nodeID
	ni.RingID = ria.ringID
	return vpp2pdat.NodeInfoSigBytes(ni)
}

// NewNode builds a new node object. Host and Ring are required,
// nodeID is optional, by default a new nodeID is provided.
func NewNode(host *Host, ring *Ring, nodeID []byte, registerer NodeRegisterer) (*Node, error) {
	var ret Node
	var info vpp2papi.NodeInfo
	var err error
	var sig []byte

	ret.hostPtr = host
	ret.ringPtr = ring
	ret.Status.Info = &info

	ret.Status.Peers = vpp2papi.NewNodePeers()

	ret.resetSuccessors()
	ret.resetD()
	ret.resetPredecessor()

	// by doing this, nodes will always be (un)registerered within hosts
	// and the global node register. This is usefull when one wants to
	// quickly access them by reference.
	ret.registerers = make([]NodeRegisterer, 1)
	ret.registerers[0] = host.localNodeCatalog
	if registerer != nil {
		ret.registerers = append(ret.registerers, registerer)
	}

	ret.Status.Info.RingID = make([]byte, len(ring.Info.RingID))
	copy(ret.Status.Info.RingID, ring.Info.RingID)
	ret.Status.Info.HostPubKey = make([]byte, len(host.Info.HostPubKey))

	if nodeID != nil && len(nodeID) == vpp2pdat.NodeIDBufNbBytes {
		ret.Status.Info.NodeID = make([]byte, vpp2pdat.NodeIDBufNbBytes)
		for i, v := range nodeID {
			ret.Status.Info.NodeID[i] = v
		}
	} else {
		var intNodeID *big.Int

		ria := ringIDAppender{ringID: ring.Info.RingID}
		if host.CanSign() {
			intNodeID, sig, _, err = vpid.GenerateID256(host.key, nil, &ria, NodeKeySeconds, NodeKeyZeroes)
			if err != nil {
				return nil, err
			}
		} else {
			intNodeID, _, _, err = vpid.GenerateID256(nil, nil, &ria, NodeKeySeconds, NodeKeyZeroes)
			if err != nil {
				return nil, err
			}
			sig = []byte("")
		}
		ret.Status.Info.NodeID = vpsum.IntToBuf256(intNodeID)
	}

	copy(ret.Status.Info.HostPubKey, host.Info.HostPubKey)
	ret.Status.Info.NodeSig = sig

	_, err = vpp2pdat.CheckNodeInfo(ret.Status.Info)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (node *Node) resetSuccessors() {
	defer node.successorsAccess.Unlock()
	node.successorsAccess.Lock()

	node.Status.Peers.Successors = nil
}

func (node *Node) resetD() {
	defer node.dAccess.Unlock()
	node.dAccess.Lock()

	node.Status.Peers.D = nil
}

func (node *Node) resetPredecessor() {
	defer node.predecessorAccess.Unlock()
	node.predecessorAccess.Lock()

	node.Status.Predecessor = node.Status.Info
}

// Start starts the node, that is, makes it available and registers it into
// all the local node catalogs.
func (node *Node) Start() {
	node.up = true
	if node.registerers != nil {
		for _, r := range node.registerers {
			r.RegisterNode(node)
		}
	}
}

// Stop stops the node, that is, makes it unavailable and unregisters it from
// all the local node catalogs.
func (node *Node) Stop() {
	if node.registerers != nil {
		for _, r := range node.registerers {
			r.UnregisterNode(node)
		}
	}
	node.up = false
}

// Up tells wether the node is up or not.
func (node *Node) Up() bool {
	return node.up
}

// GetImaginaryNode returns the ID of the imaginary node
// for a lookup. This is to be used for the first lookup
// step. Technically you could use any number but choosing
// it properly makes search faster. Note that keyShift is
// just key on the first iteration.
func (node *Node) GetImaginaryNode(key []byte) []byte {
	walker := node.ringPtr.walker

	zeroID := walker.Zero()

	// shift stuff to the right, pad with zeroes
	nodePart := walker.BackwardElem(node.Status.Info.NodeID, zeroID, walker.N()-int(node.ringPtr.Info.Config.NbStep))
	// incr by one to make sure it's *after* us, regardless
	// of the fact what we're going to add might be smaller
	// than what we substract by putting zeroes
	nodePart = walker.Incr(nodePart)
	// shift stuff to the left, pad with zeroes
	nodePart = walker.ForwardElem(nodePart, zeroID, walker.N()-int(node.ringPtr.Info.Config.NbStep))
	// make sure the right part of the imaginary part
	// or the topmost bits of k, to stumble magically on k
	// faster than expected in a fully dense network
	keyPart := walker.BackwardElem(key, zeroID, int(node.ringPtr.Info.Config.NbStep))

	return walker.Add(nodePart, keyPart)
}

/*
// Lookup a key and return the path of nodes to this key.
func (node *Node) Lookup(key, keyShift, imaginaryNode []byte) ([]*vpp2papi.NodeInfo, error) {
	// pseudo code :
	// procedure m.LOOKUP(k, kshift, i)
	//   if k is in (m,successor] then return (successor)
	//   else if i is in (m,successor] then return (
	//     d.lookup(k,
	//              kshift<<1,
	//              i o topBit(kshift)))
	//   else return (successor.lookup(k,kshift,i))
	// Note : i can be chosen so that its low bits are top bits of k

		walker := node.ringPtr.walker

		ret := make([]*vpp2papi.NodeInfo, 1)
		ret[0] = &(node.Info)

		if node.Successor == nil || len(node.Successor) == 0 || node.D == nil || len(node.D) == 0 || node.PredecessorInfo == nil {
			// no successor -> we're alone!
			return ret, nil
		}

		if walker.GtLe(key, node.PredecessorInfo.NodeID, node.Info.NodeID) {
			// key is local
			return ret, nil
		}

		for i, successor := range node.Successor {
			var curInfo *vpp2papi.NodeInfo
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
*/

// IsSigned returns true if the node has been signed by corresponding host.
// It does not check if the signature is valid.
func (node *Node) IsSigned() bool {
	return vpp2pdat.NodeInfoIsSigned(node.Status.Info)
}

// CheckSig checks if the node signature is OK, if it's not, returns 0 and an error.
// If it's OK, returns the number of zeroes in the signature hash.
func (node *Node) CheckSig() (int, error) {
	return vpp2pdat.NodeInfoCheckSig(node.Status.Info)
}

// Lookup performs a lookup for a given key
func (node *Node) Lookup(key, keyShift, imaginaryNode []byte) (bool, []*vpp2papi.NodeInfo, error) {
	// todo

	return false, nil, nil
}

// GetSuccessors returns the successors of a given node
func (node *Node) GetSuccessors() []*vpp2papi.NodeInfo {
	defer node.successorsAccess.RUnlock()
	node.successorsAccess.RLock()

	ret := make([]*vpp2papi.NodeInfo, len(node.Status.Peers.Successors))
	for i, v := range node.Status.Peers.Successors {
		ret[i] = v
	}

	return ret
}

// GetD returns the d of a given node
func (node *Node) GetD() *vpp2papi.NodeInfo {
	defer node.dAccess.RUnlock()
	node.dAccess.RLock()

	return node.Status.Peers.D
}

// GetPredecessor returns the predecessor of a given node
func (node *Node) GetPredecessor() *vpp2papi.NodeInfo {
	defer node.predecessorAccess.RUnlock()
	node.predecessorAccess.RLock()

	return node.Status.Predecessor
}

// setPredecessor returns the predecessor of a given node
func (node *Node) setPredecessor(nodeInfo *vpp2papi.NodeInfo) {
	defer node.predecessorAccess.Unlock()
	node.predecessorAccess.Lock()

	node.Status.Predecessor = nodeInfo
}

func (node *Node) isKeyOnNode(key []byte) bool {
	defer node.predecessorAccess.RUnlock()
	node.predecessorAccess.RLock()

	walker := node.ringPtr.walker

	if walker.Cmp(node.Status.Predecessor.NodeID, node.Status.Info.NodeID) == 0 {
		// special case, if predecessor is self, we cover the whole ring,
		// it does not appear sensible to have GtLe return true if they
		// are the same because, factually, it's not between X and X. One
		// could consider between X and X, there's the complete ring, but
		// one could expect other quirks when we would really test a
		// narrow ring.
		return true
	}

	if walker.GtLe(key, node.Status.Predecessor.NodeID, node.Status.Info.NodeID) {
		return true
	}

	return false
}

func (node *Node) findKeyOnLocalNode(key []byte) *Node {
	// Check if it's on us
	if node.isKeyOnNode(key) {
		return node
	}
	// Check if it's on other local nodes, while this is not truely in
	// the Bruijn walking path, it would be stupid to just ping another
	// distant node when the information is local. This is typically
	// optimizing the case when there are many virtual nodes for few
	// physical hosts.
	for _, localNode := range node.hostPtr.localNodeCatalog.ListPtr() {
		if localNode.isKeyOnNode(key) {
			return localNode
		}
	}

	return nil
}

// Sync performs a lookup for a given key
func (node *Node) Sync(source *vpp2papi.NodeInfo, key, keyShift, imaginaryNode []byte) (bool, []*vpp2papi.NodeInfo, []*vpp2papi.NodeInfo, *vpp2papi.NodeInfo, error) {
	var found *Node
	var err error
	var ok bool

	found = node.findKeyOnLocalNode(key)
	if found != nil {
		var status *vpp2papi.HostStatus
		// todo, fix this, and get the *source* status
		status, err = node.hostPtr.Status()
		if err != nil {
			vplog.LogDebug("unable to join host requiring Sync")
		} else {
			if status != nil && status.ThisHostInfo != nil {
				ok, err = vpp2pdat.CheckHostInfo(status.ThisHostInfo)
				if ok && err == nil {
					found.setPredecessor(source)
				}
			}
		}

		nodesPath := make([]*vpp2papi.NodeInfo, 1)
		nodesPath[0] = found.Status.Info
		return true, nodesPath, found.GetSuccessors(), found.GetPredecessor(), nil
	}

	return false, nil, nil, nil, nil
}
