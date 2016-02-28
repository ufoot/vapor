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

package vpkeydx

import (
	"github.com/ufoot/vapor/vperror"
	"github.com/ufoot/vapor/vpsum"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
	"math/big"
)

// GetX gets the X (1st) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetX(keyID []byte) (int32, error) {
	keyInt, err := vpsum.BufToInt256(keyID)
	if err != nil {
		return 0, vperror.Chain(err, "unable to generate 1d keydx")
	}
	xInt := big.NewInt(0)
	for i := 0; i < n31; i++ {
		xInt = xInt.SetBit(xInt, i, keyInt.Bit(nOffset1+i))
	}

	return int32(xInt.Int64()), nil
}

// GetXY gets the X,Y (1st and 2nd) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetXY(keyID []byte) (int32, int32, error) {
	keyInt, err := vpsum.BufToInt256(keyID)
	if err != nil {
		return 0, 0, vperror.Chain(err, "unable to generate 2d keydx")
	}
	xInt := big.NewInt(0)
	yInt := big.NewInt(0)
	for i := 0; i < n31; i++ {
		xInt = xInt.SetBit(xInt, i, keyInt.Bit(nOffset2+2*i))
		yInt = yInt.SetBit(yInt, i, keyInt.Bit(nOffset2+2*i+1))
	}

	return int32(xInt.Int64()), int32(yInt.Int64()), nil
}

// GetXYZ gets the X,Y,Z (1st, 2nd and 3rd) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetXYZ(keyID []byte) (int32, int32, int32, error) {
	keyInt, err := vpsum.BufToInt256(keyID)
	if err != nil {
		return 0, 0, 0, vperror.Chain(err, "unable to generate 3dd keydx")
	}
	xInt := big.NewInt(0)
	yInt := big.NewInt(0)
	zInt := big.NewInt(0)
	for i := 0; i < n31; i++ {
		xInt = xInt.SetBit(xInt, i, keyInt.Bit(nOffset3+3*i))
		yInt = yInt.SetBit(yInt, i, keyInt.Bit(nOffset3+3*i+1))
		zInt = zInt.SetBit(zInt, i, keyInt.Bit(nOffset3+3*i+2))
	}

	return int32(xInt.Int64()), int32(yInt.Int64()), int32(zInt.Int64()), nil
}

// GetVec1 gets the 1st (X) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetVec1(keyID []byte) (int32, error) {
	return GetX(keyID)
}

// GetVec2 gets the 1st and 2nd (X, Y) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetVec2(keyID []byte) (*vpvec2.I32, error) {
	x, y, err := GetXY(keyID)

	if err != nil {
		return nil, err
	}

	return vpvec2.I32New(x, y), nil
}

// GetVec3 gets the 1st, 2nd and 3rd (X,Y,Z) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetVec3(keyID []byte) (*vpvec3.I32, error) {
	x, y, z, err := GetXYZ(keyID)

	if err != nil {
		return nil, err
	}

	return vpvec3.I32New(x, y, z), nil
}
