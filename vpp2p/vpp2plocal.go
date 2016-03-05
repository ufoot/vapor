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
	"github.com/ufoot/vapor/vpp2papi"
	"github.com/ufoot/vapor/vpp2pdat"
	"sync"
)

type LocalNodeCatalog struct {
	access sync.RWMutex
	nodes  map[[vpp2pdat.NodeIDNbBytes]byte]*Node
}

var globalNodeCatalog LocalNodeCatalog = *NewLocalNodeCatalog()

func NewLocalNodeCatalog() *LocalNodeCatalog {
	return &LocalNodeCatalog{nodes: make(map[[vpp2pdat.NodeIDNbBytes]byte]*Node)}
}

func GlobalNodeCatalog() *LocalNodeCatalog {
	return &globalNodeCatalog
}

func (c *LocalNodeCatalog) ConnectToNode(nodeID []byte) (*vpp2papi.VpP2pApi, error) {
	defer c.access.RUnlock()
	c.access.RLock()

	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	n := c.nodes[nodeIDBuf]
	if n == nil {
		return nil, fmt.Errorf("node does not exist")
	}

	return nil, nil
}

func (c *LocalNodeCatalog) RegisterNode(node *Node) error {
	defer c.access.Unlock()
	c.access.Lock()

	nodeID := node.Info.NodeID
	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	c.nodes[nodeIDBuf] = node

	return nil
}

func (c *LocalNodeCatalog) UnregisterNode(node *Node) error {
	defer c.access.Unlock()
	c.access.Lock()

	nodeID := node.Info.NodeID
	nodeIDBuf := vpp2pdat.NodeIDToBuf(nodeID)
	delete(c.nodes, nodeIDBuf)

	return nil
}
