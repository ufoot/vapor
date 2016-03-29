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

package vpp2pdat

import (
	"bytes"
	"github.com/ufoot/vapor/go/vprand"
	"github.com/ufoot/vapor/go/vpsum"
	"testing"
)

func TestHostPubKeyToBuf(t *testing.T) {
	id := vpsum.Checksum128([]byte("foo"))
	b := HostPubKeyToBuf(id)
	if bytes.Compare(b[:], id) != 0 {
		t.Error("id and b differ")
	}
}

func TestNodeIDToBuf(t *testing.T) {
	id := vpsum.Checksum256([]byte("foo"))
	b := NodeIDToBuf(id)
	if bytes.Compare(b[:], id) != 0 {
		t.Error("id and b differ")
	}
}

func TestRingIDToBuf(t *testing.T) {
	id := vpsum.Checksum512([]byte("bar"))
	b := RingIDToBuf(id)
	if bytes.Compare(b[:], id) != 0 {
		t.Error("id and b differ")
	}
}

func TestHostPubKeyToShortString(t *testing.T) {
	id := vpsum.IntToBuf128(vprand.Rand128(nil, nil))
	s := HostPubKeyToShortString(id)
	t.Logf("HostPubKey short string: %s", s)
	if len(s) != HostPubKeyShortStringLen {
		t.Errorf("bad len %d for HostPubKey short string, should be %d", len(s), HostPubKeyShortStringLen)
	}
}

func TestNodeIDToShortString(t *testing.T) {
	id := vpsum.IntToBuf256(vprand.Rand256(nil, nil))
	s := NodeIDToShortString(id)
	t.Logf("NodeID short string: %s", s)
	if len(s) != NodeIDShortStringLen {
		t.Errorf("bad len %d for NodeID short string, should be %d", len(s), NodeIDShortStringLen)
	}
}

func TestRingIDToShortString(t *testing.T) {
	id := vpsum.IntToBuf512(vprand.Rand512(nil, nil))
	s := RingIDToShortString(id)
	t.Logf("Ring ID short string: %s", s)
	if len(s) != RingIDShortStringLen {
		t.Errorf("bad len %d for RingID short string, should be %d", len(s), RingIDShortStringLen)
	}
}
