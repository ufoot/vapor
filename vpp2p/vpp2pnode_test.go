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
	"github.com/ufoot/vapor/vpapp"
	"github.com/ufoot/vapor/vpid"
	"github.com/ufoot/vapor/vpsum"
	"testing"
)

func TestNewNode(t *testing.T) {
	var host *Host
	var node *Node
	var err error
	ringID := vpsum.Checksum128([]byte("myring"))
	var zeroes, zeroes2 int

	host, err = NewHost(testTitle, testURL, true)
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	node, err = NewNode(host, ringID)
	if err != nil {
		t.Error("unable to create node with a valid pubKey", err)
	}
	if node.IsSigned() == false {
		t.Error("node is unsigned, when it should be")
	}
	zeroes = vpid.ZeroesInBuf(vpsum.Checksum256(node.Info.NodeSig))
	if zeroes < NodeKeyZeroes {
		t.Errorf("Node created, but not enough zeroes in sig (%d)", zeroes)
	}
	t.Logf("Node created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = NodeInfoCheckSig(&(node.Info))
	if err != nil {
		t.Error("wrong sig", err)
	}
	if zeroes != zeroes2 {
		t.Errorf("NodeInfoCheckSig returned bad number of zeroes %d!=%d", zeroes, zeroes2)
	}

	host, err = NewHost(testTitle, testURL, false)
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	node, err = NewNode(host, ringID)
	if err != nil {
		t.Error("unable to create node with a valid pubKey", err)
	}
	if node.IsSigned() == true {
		t.Error("node is signed, when it should not be")
	}
	zeroes = vpid.ZeroesInBuf(vpsum.Checksum256(node.Info.NodeSig))
	t.Logf("Node created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = NodeInfoCheckSig(&(node.Info))
	if err == nil {
		t.Error("sig reported as good when it's not")
	}
}

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
