// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
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
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vptypes"
	"math/big"
)

const MinSquareSizeSmall = 1
const MaxSquareSizeSmall = 5
const MinSquareSizeMedium = 2
const MaxSquareSizeMedium = 10
const MinSquareSizeLarge = 5
const MaxSquareSizeLarge = 25

const MinPlanetSizeSmall = 32
const MaxPlanetSizeSmall = 128
const MinPlanetSizeMedium = 64
const MaxPlanetSizeMedium = 512
const MinPlanetSizeLarge = 256
const MaxPlanetSizeLarge = 4096

const MinSystemSizeSmall = 1
const MaxSystemSizeSmall = 1
const MinSystemSizeMedium = 1
const MaxSystemSizeMedium = 7
const MinSystemSizeLarge = 2
const MaxSystemSizeLarge = 42

const NetworkIdSeconds = 5

type Sizes struct {
	SquareSize vptypes.Size
	PlanetSize vptypes.Size
	SystemSize vptypes.Size
}

// NetworkId generates a new level id, using cryptography to garantee
// two players will never generate the same level id.
func NetworkId(sizes Sizes, key *vpcrypto.Key) (*big.Int, []byte, int, error) {
	// todo : use sizes
	return vpcrypto.GenerateID512(key, nil, NetworkIdSeconds)
}

// LocalId generates a new level id, not suitable for inte
// two players will never generate the same level id.
func LocalId(sizes Sizes) *big.Int {
	// todo : use sizes
	return vpcrypto.Rand512(vpcrypto.NewRand(), nil)
}
