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
	"github.com/ufoot/vapor/go/vpp2papi"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"sync"
)

// NodeCatalog is structure used to contain locally-known hosts.
type NodeCatalog struct {
	access sync.RWMutex
	nodes  map[[vpp2pdat.NodeIDNbBytes]byte]*Node
}

var globalNodeCatalog = NewNodeCatalog()

// NewNodeCatalog creates a new instance of a local node catalog
func NewNodeCatalog() *NodeCatalog {
	return &NodeCatalog{nodes: make(map[[vpp2pdat.NodeIDNbBytes]byte]*Node)}
}

// GlobalNodeCatalog returns a catalog containing all local nodes.
func GlobalNodeCatalog() *NodeCatalog {
	return globalNodeCatalog
}

// ConnectToNode returns a handler which makes possible API calls on it.
// It's thread-safe.
func (c *NodeCatalog) ConnectToNode(nodeID []byte) (vpp2papi.VpP2pApi, error) {
	defer c.access.RUnlock()
	c.access.RLock()

	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	n := c.nodes[nodeIDBuf]
	if n == nil {
		return nil, fmt.Errorf("node does not exist")
	}

	return n.hostPtr, nil
}

// HasNode returns a handler which makes possible API calls on it.
// It's thread-safe.
func (c *NodeCatalog) HasNode(nodeID []byte) bool {
	defer c.access.RUnlock()
	c.access.RLock()

	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	n := c.nodes[nodeIDBuf]

	return n != nil
}

// RegisterNode registers a node within the catalog.
// It's thread-safe.
func (c *NodeCatalog) RegisterNode(node *Node) {
	defer c.access.Unlock()
	c.access.Lock()

	nodeID := node.Info.NodeID
	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	c.nodes[nodeIDBuf] = node
}

// UnregisterNode unregisters a node within the catalog.
// It's thread-safe.
func (c *NodeCatalog) UnregisterNode(node *Node) {
	defer c.access.Unlock()
	c.access.Lock()

	nodeID := node.Info.NodeID
	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	delete(c.nodes, nodeIDBuf)
}

// List returns a list of local nodes. It returns static
// data about the node, node the nodes themselves.
func (c *NodeCatalog) List() []*vpp2papi.NodeInfo {
	defer c.access.RUnlock()
	c.access.RLock()

	i := 0
	ret := make([]*vpp2papi.NodeInfo, len(c.nodes))
	for _, value := range c.nodes {
		ret[i] = &value.Info
		i++
	}

	return ret
}
