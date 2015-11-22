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
	"github.com/ufoot/vapor/vpsys"
	"github.com/ufoot/vapor/vptypes"
	"math/big"
)

// MinSquareSizeSmall is the minimum size, in meters, of a square
// when using the "small" settings.
const MinSquareSizeSmall = 1

// MaxSquareSizeSmall is the maximum size, in meters, of a square
// when using the "small" settings.
const MaxSquareSizeSmall = 5

// MinSquareSizeMedium is the minimum size, in meters, of a square
// when using the "medium" settings.
const MinSquareSizeMedium = 2

// MaxSquareSizeMedium is the maximum size, in meters, of a square
// when using the "medium" settings.
const MaxSquareSizeMedium = 10

// MinSquareSizeLarge is the minimum size, in meters, of a square
// when using the "large" settings.
const MinSquareSizeLarge = 5

// MaxSquareSizeLarge is the maximum size, in meters, of a square
// when using the "large" settings.
const MaxSquareSizeLarge = 25

// MinPlanetSizeSmall is the minimum size, in number of squares
// per side, of a planet, when using the "small" settings.
const MinPlanetSizeSmall = 50

// MaxPlanetSizeSmall is the maximum size, in number of squares
// per side, of a planet, when using the "small" settings.
const MaxPlanetSizeSmall = 80

// MinPlanetSizeMedium is the minimum size, in number of squares
// per side, of a planet, when using the "medium" settings.
const MinPlanetSizeMedium = 60

// MaxPlanetSizeMedium is the maximum size, in number of squares
// per side, of a planet, when using the "medium" settings.
const MaxPlanetSizeMedium = 300

// MinPlanetSizeLarge is the minimum size, in number of squares
// per side, of a planet, when using the "large" settings.
const MinPlanetSizeLarge = 100

// MaxPlanetSizeLarge is the maximum size, in number of squares
// per side, of a planet, when using the "large" settings.
const MaxPlanetSizeLarge = 1000

// MinSystemSizeSmall is the minimum size, in number of planets,
// of a system, when using the "small" settings.
const MinSystemSizeSmall = 1

// MaxSystemSizeSmall is the maximum size, in number of planets,
// of a system, when using the "small" settings.
const MaxSystemSizeSmall = 1

// MinSystemSizeMedium is the minimum size, in number of planets,
// of a system, when using the "medium" settings.
const MinSystemSizeMedium = 2

// MaxSystemSizeMedium is the maimum size, in number of planets,
// of a system, when using the "medium" settings.
const MaxSystemSizeMedium = 7

// MinSystemSizeLarge is the minimum size, in number of planets,
// of a system, when using the "large" settings.
const MinSystemSizeLarge = 3

// MaxSystemSizeLarge is the maximum size, in number of planets,
// of a system, when using the "large" settings.
const MaxSystemSizeLarge = 42

// NetworkIDSeconds is the number of seconds it should take to
// generate a system when playing network games.
const NetworkIDSeconds = 5

// Sizes is used to hold together the three major level size
// settings, concerning squares, planets and systems.
type Sizes struct {
	SquareSize vptypes.Size
	PlanetSize vptypes.Size
	SystemSize vptypes.Size
}

const expBaseN = 10

const (
	squareIndex int = iota
	planetIndex
	systemIndex
)

func setSize(id *big.Int, n int, s vptypes.Size) {
	bsr := big.NewInt(int64(vptypes.SizeRange))

	mulA := big.NewInt(0)
	mulA.Exp(bsr, big.NewInt(int64(n+expBaseN)), nil)
	mulB := big.NewInt(0)
	mulB.Mul(bsr, mulA)

	tmp := big.NewInt(0)
	tmp.Mod(id, mulB)
	tmp.Div(tmp, mulA)
	tmp.Set(big.NewInt(int64(s) - tmp.Int64()))
	tmp.Mul(tmp, mulA)
	id.Add(id, tmp)
}

func getSize(id *big.Int, n int) vptypes.Size {
	bsr := big.NewInt(int64(vptypes.SizeRange))

	mulA := big.NewInt(0)
	mulA.Exp(bsr, big.NewInt(int64(n+expBaseN)), nil)

	tmp := big.NewInt(0)
	tmp.Div(id, mulA)
	tmp.Mod(tmp, bsr)

	return vptypes.Size(tmp.Int64())
}

type levelFilterChecker struct {
	sizes Sizes
}

func (fc levelFilterChecker) Filter(id *big.Int) *big.Int {
	var ret big.Int

	vpsys.LogDebug("filter for levelid")

	ret.Set(id)

	setSize(&ret, squareIndex, fc.sizes.SquareSize)
	setSize(&ret, planetIndex, fc.sizes.PlanetSize)
	setSize(&ret, systemIndex, fc.sizes.SystemSize)

	return &ret
}

func (fc levelFilterChecker) Check(id *big.Int) bool {
	vpsys.LogDebug("check for levelid")

	if getSize(id, squareIndex) != fc.sizes.SquareSize {
		return false
	}
	if getSize(id, planetIndex) != fc.sizes.PlanetSize {
		return false
	}
	if getSize(id, systemIndex) != fc.sizes.SystemSize {
		return false
	}

	return true
}

// NetworkID generates a new level id, using cryptography to garantee
// two players will never generate the same level id.
func NetworkID(sizes Sizes, key *vpcrypto.Key) (*big.Int, []byte, error) {
	var fc levelFilterChecker

	fc.sizes = sizes

	ret, sig, _, err := vpcrypto.GenerateID512(key, fc, NetworkIDSeconds)

	return ret, sig, err
}

// LocalID generates a new level id, not suitable for inte
// two players will never generate the same level id.
func LocalID(sizes Sizes) (*big.Int, error) {
	var fc levelFilterChecker

	fc.sizes = sizes

	ret, _, _, err := vpcrypto.GenerateID512(nil, fc, NetworkIDSeconds)

	return ret, err
}
