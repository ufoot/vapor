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
	"github.com/ufoot/vapor/go/vpp2pdat"
	"testing"
)

func TestGlobalCatalog(t *testing.T) {
	var host *Host
	var node1, node2 *Node
	var ring *Ring
	var err error

	host, err = NewHost(testTitle, testURL, false)
	if err != nil {
		t.Error("unable to create host", err)
	}
	ring, err = NewRing(host, testTitle, testDescription, testID, vpp2pdat.DefaultRingConfig(), nil, nil)
	if err != nil {
		t.Error("unable to create ring", err)
	}
	node1, err = NewNode(host, ring)
	if err != nil {
		t.Error("unable to create node1", err)
	}
	node2, err = NewNode(host, ring)
	if err != nil {
		t.Error("unable to create node2", err)
	}

	if GlobalNodeCatalog().HasNode(node1.Info.NodeID) {
		t.Error("global catalog has node1, but it's not started yet")
	}
	if host.localNodeCatalog.HasNode(node1.Info.NodeID) {
		t.Error("host local catalog has node1, but it's not started yet")
	}
	if GlobalNodeCatalog().HasNode(node2.Info.NodeID) {
		t.Error("global catalog has node2, but it's not started yet")
	}
	if host.localNodeCatalog.HasNode(node2.Info.NodeID) {
		t.Error("host local catalog has node2, but it's not started yet")
	}
	node1.Start()
	node2.Start()
	if !GlobalNodeCatalog().HasNode(node1.Info.NodeID) {
		t.Error("global catalog does not have node1, but it has been started")
	}
	if !host.localNodeCatalog.HasNode(node1.Info.NodeID) {
		t.Error("host local catalog does not have node1, but it has been started")
	}
	if !GlobalNodeCatalog().HasNode(node2.Info.NodeID) {
		t.Error("global catalog does not have node2, but it has been started")
	}
	if !host.localNodeCatalog.HasNode(node2.Info.NodeID) {
		t.Error("host local catalog does not have node2, but it has been started")
	}
	if len(GlobalNodeCatalog().List()) != 2 {
		t.Error("global catalog length should be 2")
	}
	if len(host.localNodeCatalog.List()) != 2 {
		t.Error("host local catalog length should be 2")
	}
	node1.Stop()
	if GlobalNodeCatalog().HasNode(node1.Info.NodeID) {
		t.Error("global catalog has node1, but it has been stopped")
	}
	if host.localNodeCatalog.HasNode(node1.Info.NodeID) {
		t.Error("host local catalog has node1, but has been stopped")
	}
	if !GlobalNodeCatalog().HasNode(node2.Info.NodeID) {
		t.Error("global catalog does not have node2, but it should still be started")
	}
	if !host.localNodeCatalog.HasNode(node2.Info.NodeID) {
		t.Error("host local catalog does not have node2, but it shoulto still be started")
	}
}
