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
	"encoding/hex"
	"github.com/ufoot/vapor/go/vpid"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"github.com/ufoot/vapor/go/vpsum"
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

	host, err = NewHost(testTitle, testURL, useSig)
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
	node, err = NewNode(host, ring, nil)
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
	node, err = NewNode(host, ring, nil)
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
	node1, err = NewNode(host, ring)
	if err != nil {
		t.Error("unable to create node", err)
	}
	node2, err = NewNode(host, ring)
	if err != nil {
		t.Error("unable to create node", err)
	}
	imaginaryNode := node1.GetImaginaryNode(node2.Status.Info.NodeID)
	t.Logf("node = %s", hex.EncodeToString(node1.Status.Info.NodeID))
	t.Logf("key =  %s", hex.EncodeToString(node2.Status.Info.NodeID))
	t.Logf("img =  %s", hex.EncodeToString(imaginaryNode))
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
