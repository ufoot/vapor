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
	"github.com/ufoot/vapor/vpcrypto"
	"testing"
)

func TestLookup(t *testing.T) {
	key := vpcrypto.Checksum256([]byte("toto"))
	keyShift := vpcrypto.Checksum256([]byte("toto"))
	imaginaryNode := vpcrypto.Checksum256([]byte("titi"))
	ai := vpapp.CalcID(vpapp.DefaultPackage(), vpapp.DefaultVersion())
	h, err := NewHost("foo bar", "http://foo.com", []byte("sadhjkagfnjka"))
	if err != nil {
		t.Error("error creating host", err)
	}
	r, err := NewRing([]byte("tutu"), "foo bar", ai, []byte("abcd"), DefaultRingConfig())
	if err != nil {
		t.Error("error creating ring", err)
	}
	lp, err := NewLocalProxy(vpcrypto.Checksum256([]byte("tata")), h, r)
	if err != nil {
		t.Error("error creating local proxy", err)
	}

	path, err := lp.Lookup(key, keyShift, imaginaryNode)
	t.Log("todo...", path, err)
}
