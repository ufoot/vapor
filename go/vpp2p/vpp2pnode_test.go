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
	"bytes"
	"encoding/hex"
	"github.com/ufoot/vapor/go/vpid"
	"github.com/ufoot/vapor/go/vpp2papi"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"github.com/ufoot/vapor/go/vpsum"
	"strconv"
	"testing"
)

const testDescription = "This is a description, it is longer than a title"

var testID []byte

func init() {
	testID = []byte("1234567890abcdef")
}

func setupHostRing(t *testing.T, useSig bool) (*Host, *Ring, error) {
	var host *Host
	var ring *Ring
	var err error

	host, err = NewHost(testTitle, testURL, useSig, GlobalHostInfoCatalog())
	if err != nil {
		t.Error("unable to create host", t, err)
		return nil, nil, err
	}
	ring, err = NewRing(host, testTitle, testDescription, testID, vpp2pdat.DefaultRingConfig(), nil, nil)
	if err != nil {
		t.Error("unable to create ring", t, err)
		return nil, nil, err
	}

	return host, ring, nil
}

func TestNewNode(t *testing.T) {
	var host *Host
	var ring *Ring
	var node *Node
	var err error
	var zeroes, zeroes2 int

	host, ring, err = setupHostRing(t, true)
	node, err = NewNode(host, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node with a valid pubKey", err)
	}
	if node.IsSigned() == false {
		t.Error("node is unsigned, when it should be")
	}
	zeroes = vpid.ZeroesInBuf(vpsum.Checksum256(node.Status.Info.NodeSig))
	if zeroes < NodeKeyZeroes {
		t.Errorf("Node created, but not enough zeroes in sig (%d)", zeroes)
	}
	t.Logf("Node created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = node.CheckSig()
	if err != nil {
		t.Error("wrong sig", err)
	}
	if zeroes != zeroes2 {
		t.Errorf("NodeInfoCheckSig returned bad number of zeroes %d!=%d", zeroes, zeroes2)
	}
	node.Status.Info.NodeSig = node.Status.Info.HostPubKey
	_, err = node.CheckSig()
	if err == nil {
		t.Error("failed to report a broken sig", err)
	}

	host, ring, err = setupHostRing(t, false)
	node, err = NewNode(host, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node", err)
	}
	if node.IsSigned() == true {
		t.Error("node is signed, when it should not be")
	}
	zeroes = vpid.ZeroesInBuf(vpsum.Checksum256(node.Status.Info.NodeSig))
	t.Logf("Node created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = node.CheckSig()
	if err != nil {
		t.Error("sig reported as wrong when it's legal to have an empty sig")
	}
}

func TestGetImaginaryNode(t *testing.T) {
	var host *Host
	var ring *Ring
	var node1, node2 *Node
	var err error

	host, ring, err = setupHostRing(t, false)
	node1, err = NewNode(host, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node", err)
	}
	node2, err = NewNode(host, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node", err)
	}
	imaginaryNode := node1.GetImaginaryNode(node2.Status.Info.NodeID)
	t.Logf("node = %s", hex.EncodeToString(node1.Status.Info.NodeID))
	t.Logf("key =  %s", hex.EncodeToString(node2.Status.Info.NodeID))
	t.Logf("img =  %s", hex.EncodeToString(imaginaryNode))

	leftNode1 := node1.Status.Info.NodeID[0 : node1.ringPtr.Info.Config.NbStep/2]
	var leftNode1Int64 int64
	leftNode1Int64, err = strconv.ParseInt(hex.EncodeToString(leftNode1), 16, 64)
	if err != nil {
		t.Error("unable to parse int", leftNode1Int64, err)
	}

	leftImaginaryNode := imaginaryNode[0 : node1.ringPtr.Info.Config.NbStep/2]
	var leftImaginaryNodeInt64 int64
	leftImaginaryNodeInt64, err = strconv.ParseInt(hex.EncodeToString(leftImaginaryNode), 16, 64)
	if err != nil {
		t.Error("unable to parse int", leftImaginaryNode, err)
	}

	if leftImaginaryNodeInt64-leftNode1Int64 != 1 {
		t.Errorf("imaginaryNode=%x... not close enough to node1=%x...", leftImaginaryNodeInt64, leftNode1Int64)
	}

	leftNode2 := node2.Status.Info.NodeID[0 : 32-node1.ringPtr.Info.Config.NbStep/2]
	rightImaginaryNode := imaginaryNode[node1.ringPtr.Info.Config.NbStep/2 : 32]
	if bytes.Compare(leftNode2, rightImaginaryNode) != 0 {
		t.Errorf("imaginaryNode=...%s not matching to node2=%s...", hex.EncodeToString(rightImaginaryNode), hex.EncodeToString(leftNode2))
	}
}

func TestGetSync(t *testing.T) {
	var host *Host
	var ring *Ring
	var node1, node2, node3 *Node
	var err error

	host, ring, err = setupHostRing(t, false)
	node1, err = NewNode(host, ring, nil, GlobalNodeCatalog())
	if err != nil {
		t.Error("unable to create node", err)
	}
	for ; !(err == nil && node3 != nil && ring.walker.Cmp(node1.Status.Info.NodeID, node3.Status.Info.NodeID) < 0); node3, err = NewNode(host, ring, nil, GlobalNodeCatalog()) {
	}
	for ; !(err == nil && node2 != nil && ring.walker.GtLe(node2.Status.Info.NodeID, node1.Status.Info.NodeID, node3.Status.Info.NodeID)); node2, err = NewNode(host, ring, nil, GlobalNodeCatalog()) {
	}

	t.Logf("node1 = %s", hex.EncodeToString(node1.Status.Info.NodeID))
	t.Logf("node2 = %s", hex.EncodeToString(node2.Status.Info.NodeID))
	t.Logf("node3 = %s", hex.EncodeToString(node3.Status.Info.NodeID))

	imaginaryNode := node1.GetImaginaryNode(node2.Status.Info.NodeID)
	t.Logf("img =  %s", hex.EncodeToString(imaginaryNode))

	defer node1.Stop()
	node1.Start()
	defer node2.Stop()
	node2.Start()
	defer node3.Stop()
	node3.Start()

	var found bool
	var path []*vpp2papi.NodeInfo
	var successors []*vpp2papi.NodeInfo
	var predecessor *vpp2papi.NodeInfo
	found, path, successors, predecessor, err = node1.Sync(node2.Status.Info, node2.Status.Info.NodeID, node2.Status.Info.NodeID, imaginaryNode)
	if found != true {
		t.Error("unable to sync with self")
	}
	if path == nil {
		t.Error("path is nil for self")
	}
	if successors == nil {
		t.Error("successors is nil for self")
	}
	if predecessor == nil {
		t.Error("predecessor is nil for self")
	}
	if len(path) != 1 {
		t.Fatal("bad path len", len(path))
	}
	if len(successors) != 0 {
		t.Fatal("bad path len", len(successors))
	}
	if bytes.Compare(node2.Status.Info.NodeID, predecessor.NodeID) != 0 {
		t.Fatalf("bad predecessor %s %s", hex.EncodeToString(node2.Status.Info.NodeID), hex.EncodeToString(predecessor.NodeID))
	}
}

/*
func TestLookup(t *testing.T) {
	key := vpsum.Checksum256([]byte("toto"))
	keyShift := vpsum.Checksum256([]byte("toto"))
	imaginaryNode := vpsum.Checksum256([]byte("titi"))
	ai := vpapp.CalcID(vpapp.DefaultPackage(), vpapp.DefaultVersion())
	h, err := NewHost("foo bar", "http://foo.com", true)
	if err != nil {
		t.Error("error creating host", err)
	}
	r, err := NewRing(h, testTitle, testDescription, ai, nil, nil, nil)
	if err != nil {
		t.Error("error creating ring", err)
	}
	lp, err := NewLocalProxy(vpsum.Checksum256([]byte("tata")), h, r)
	if err != nil {
		t.Error("error creating local proxy", err)
	}

	path, err := lp.Lookup(key, keyShift, imaginaryNode)
	t.Log("todo...", path, err)
}
*/
