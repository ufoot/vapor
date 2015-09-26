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

package vpkeydx

import (
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpsys"
	"math/big"
)

// GetX gets the X (1st) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetX(keyID []byte) (int32, error) {
	keyInt, err := vpcrypto.BufToInt256(keyID)
	if err != nil {
		return 0, vpsys.ErrorChain(err, "unable to generate 1d keydx")
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
	keyInt, err := vpcrypto.BufToInt256(keyID)
	if err != nil {
		return 0, 0, vpsys.ErrorChain(err, "unable to generate 1d keydx")
	}
	xInt := big.NewInt(0)
	yInt := big.NewInt(0)
	for i := 0; i < n31; i++ {
		xInt = xInt.SetBit(xInt, i, keyInt.Bit(nOffset2+2*i))
		yInt = yInt.SetBit(yInt, i, keyInt.Bit(nOffset2+2*i+1))
	}

	return int32(xInt.Int64()), int32(yInt.Int64()), nil
}

// GetXYZ gets the X,Y,Z (3rd) coord value for a given key.
// Note that it can be used for any key, even possibly those which have
// not be generated with Gen1d.
func GetXYZ(keyID []byte) (int32, int32, int32, error) {
	keyInt, err := vpcrypto.BufToInt256(keyID)
	if err != nil {
		return 0, 0, 0, vpsys.ErrorChain(err, "unable to generate 1d keydx")
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
