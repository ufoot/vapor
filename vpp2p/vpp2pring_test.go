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
	"github.com/ufoot/vapor/vpcrypto"
	"testing"
)

func TestNewRing(t *testing.T) {
	var host *Host
	var ring *Ring
	var err error
	var zeroes, zeroes2 int

	host, err = NewHost(testTitle, testURL, true)
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	ring, err = NewRing(host, testTitle, testDescription, testID, nil, nil, nil)
	if err != nil {
		t.Error("unable to create ring with a valid pubKey", err)
	}
	if ring.IsSigned() == false {
		t.Error("ring is unsigned, when it should be")
	}
	zeroes = vpcrypto.ZeroesInBuf(vpcrypto.Checksum512(ring.Info.RingSig))
	if zeroes < RingKeyZeroes {
		t.Errorf("Ring created, but not enough zeroes in sig (%d)", zeroes)
	}
	t.Logf("Ring created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = RingInfoCheckSig(&(ring.Info))
	if err != nil {
		t.Error("wrong sig", err)
	}
	if zeroes != zeroes2 {
		t.Errorf("RingInfoCheckSig returned bad number of zeroes %d!=%d", zeroes, zeroes2)
	}

	host, err = NewHost(testTitle, testURL, false)
	if err != nil {
		t.Error("unable to create host with a valid pubKey", err)
	}
	ring, err = NewRing(host, testTitle, testDescription, testID, nil, nil, nil)
	if err != nil {
		t.Error("unable to create ring with a valid pubKey", err)
	}
	if ring.IsSigned() == true {
		t.Error("ring is signed, when it should not be")
	}
	zeroes = vpcrypto.ZeroesInBuf(vpcrypto.Checksum512(ring.Info.RingSig))
	t.Logf("Ring created, number of zeroes in sig is %d", zeroes)
	zeroes2, err = RingInfoCheckSig(&(ring.Info))
	if err == nil {
		t.Error("sig reported as good when it's not")
	}
}
