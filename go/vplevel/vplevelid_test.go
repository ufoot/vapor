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

package vplevel

import (
	"github.com/ufoot/vapor/go/vpcrypto"
	"github.com/ufoot/vapor/go/vptypes"
	"math/big"
	"testing"
)

func TestFilterChecker(t *testing.T) {
	for i := 0; i < int(vptypes.SizeRange); i++ {
		for j := 0; j < int(vptypes.SizeRange); j++ {
			for k := 0; k < int(vptypes.SizeRange); k++ {
				n := big.NewInt(1000)
				setSize(n, squareIndex, vptypes.Size(i))
				setSize(n, planetIndex, vptypes.Size(j))
				setSize(n, systemIndex, vptypes.Size(k))
				checkI := int(getSize(n, squareIndex))
				checkJ := int(getSize(n, planetIndex))
				checkK := int(getSize(n, systemIndex))
				t.Logf("i/checkI=%d/%d j/checkJ=%d/%d k/checkK=%d/%d n=%d", i, checkI, j, checkJ, k, checkK, n.Int64())
				if i != checkI {
					t.Errorf("i=%d but checkI=%d", i, checkI)
				}
				if j != checkJ {
					t.Errorf("j=%d but checkJ=%d", j, checkJ)
				}
				if k != checkK {
					t.Errorf("k=%d but checkK=%d", k, checkK)
				}
			}
		}
	}
}

func TestNetworkID(t *testing.T) {
	var sizes Sizes

	sizes.SquareSize = vptypes.SizeMedium
	sizes.PlanetSize = vptypes.SizeSmall
	sizes.SystemSize = vptypes.SizeLarge

	key, err := vpcrypto.NewKey()
	if err == nil {
		ni, sig, err := NetworkID(sizes, key)
		if err == nil {
			t.Logf("Network ID generated n=%d sig=%s", ni.Int64(), string(sig))
			if getSize(ni, squareIndex) != sizes.SquareSize {
				t.Error("square size problem")
			}
			if getSize(ni, planetIndex) != sizes.PlanetSize {
				t.Error("planet size problem")
			}
			if getSize(ni, systemIndex) != sizes.SystemSize {
				t.Error("system size problem")
			}
		} else {
			t.Error("unable to generate Network ID", err)
		}
	} else {
		t.Error("unable to generate key", err)
	}
}

func TestLocalID(t *testing.T) {
	var sizes Sizes

	sizes.SquareSize = vptypes.SizeMedium
	sizes.PlanetSize = vptypes.SizeSmall
	sizes.SystemSize = vptypes.SizeLarge

	li, err := LocalID(sizes)

	if err == nil {
		t.Logf("Local ID generated n=%d", li.Int64())
		if getSize(li, squareIndex) != sizes.SquareSize {
			t.Error("square size problem")
		}
		if getSize(li, planetIndex) != sizes.PlanetSize {
			t.Error("planet size problem")
		}
		if getSize(li, systemIndex) != sizes.SystemSize {
			t.Error("system size problem")
		}
	} else {
		t.Error("unable to generate Local ID", err)
	}
}
