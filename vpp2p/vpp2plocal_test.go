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
	"github.com/ufoot/vapor/vpp2pdat"
	"testing"
)

func TestGlobalCatalog(t *testing.T) {
	var host *Host
	var node *Node
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
	node, err = NewNode(host, ring, nil)
	if err != nil {
		t.Error("unable to create node", err)
	}

	err = GlobalNodeCatalog().RegisterNode(node)
	if err != nil {
		t.Error("unable to register node")
	}

}
