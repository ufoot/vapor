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
	"testing"
)

func TestGlobalHostInfoCatalog(t *testing.T) {
	var host1, host2, host3 *Host
	var ring *Ring
	var node1, node2, node3 *Node
	var err error

	host1, err = NewHost(fmt.Sprintf("%s 1", testTitle), fmt.Sprintf("%s1", testURL), false, GlobalHostInfoCatalog())
	if err != nil {
		t.Error("unable to create host 1", err)
	}
	host2, err = NewHost(fmt.Sprintf("%s 2", testTitle), fmt.Sprintf("%s2", testURL), false, GlobalHostInfoCatalog())
	if err != nil {
		t.Error("unable to create host 2", err)
	}
	host3, err = NewHost(fmt.Sprintf("%s 3", testTitle), fmt.Sprintf("%s3", testURL), false, GlobalHostInfoCatalog())
	if err != nil {
		t.Error("unable to create host 3", err)
	}

	if GlobalHostInfoCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog has host1, but it's not registered yet")
	}
	if GlobalHostInfoCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog has host2, but it's not registered yet")
	}
	GlobalHostInfoCatalog().RegisterHost(&(host1.Info))
	if !GlobalHostInfoCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog does not have host1, but it has been registered")
	}
	if GlobalHostInfoCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog has host2, but it's not registered yet")
	}
	GlobalHostInfoCatalog().RegisterHost(&(host2.Info))
	if !GlobalHostInfoCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog does not have host1, but it has been registered")
	}
	if !GlobalHostInfoCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog does not have host2, but it has been registered")
	}
	if len(GlobalHostInfoCatalog().List()) != 2 {
		t.Error("global catalog length should be 2")
	}
	GlobalHostInfoCatalog().UnregisterHost(&(host1.Info))
	if GlobalHostInfoCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog has host1, but it's not registered yet")
	}
	if !GlobalHostInfoCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog does not have host2, but it has been registered")
	}
	if len(GlobalHostInfoCatalog().List()) != 1 {
		t.Error("global catalog length should be 1")
	}
	ring, err = NewRing(host1, testTitle, testDescription, testID, vpp2pdat.DefaultRingConfig(), nil, nil)
	if err != nil {
		t.Error("unable to create ring", err)
	}
	node1, err = NewNode(host1, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node1", err)
	}
	node2, err = NewNode(host2, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node2", err)
	}
	node3, err = NewNode(host3, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node3", err)
	}
	nodesList := make([]*vpp2papi.NodeInfo, 3)
	nodesList[0] = node1.Status.Info
	nodesList[1] = node2.Status.Info
	nodesList[2] = node3.Status.Info
	ringsList := make([]*vpp2papi.RingInfo, 1)
	ringsList[0] = &(ring.Info)
	refs := GlobalHostInfoCatalog().CreateHostsRefs(&(host1.Info), ringsList, nodesList)
	GlobalHostInfoCatalog().UpdateHostsRefs(refs)
}
