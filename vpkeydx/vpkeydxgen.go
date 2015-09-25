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
	"fmt"
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpsys"
)

// Gen generates a unique ID for a key, given a seed.
// A typical usage of seed is the vring (virtual ring) name,
// which can be used to differenciate key locations on various
// vrings.
func Gen(seed []byte, keyName string) ([]byte, error) {
	if len(seed) <= 0 {
		return nil, fmt.Errorf("seed is not long enough")
	}
	if len(keyName) <= 0 {
		return nil, fmt.Errorf("key is not long enough")
	}

	return vpcrypto.Checksum256(append(seed, []byte(keyName)...)), nil
}

// Gen1d generates a unique ID for a key, given a seed.
// A typical usage of seed is the vring (virtual ring) name,
// which can be used to differenciate key locations on various
// vrings.
// The x parameter has a special meaning, it is used on a 31-bit
// scale (from 0 to 0x7fffffff) as a prefix for the key.
// This is used to explicitely store some keys close to some node/host.
func Gen1d(seed []byte, keyName string, x int32) ([]byte, error) {
	kBuf, err := Gen(seed, keyName)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 1d keydx")
	}
	kInt, err := vpcrypto.BufToInt256(kBuf)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 1d keydx")
	}
	xInt, err := toBigInt31(x)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 1d keydx")
	}
	for i := 0; i < n31; i++ {
		kInt = kInt.SetBit(kInt, nOffset1+i, xInt.Bit(i))
	}
	return vpcrypto.IntToBuf256(kInt), nil
}

// Gen2d generates a unique ID for a key, given a seed.
// A typical usage of seed is the vring (virtual ring) name,
// which can be used to differenciate key locations on various
// vrings.
// The x and y parameter have a special meaning, they are used on a 31-bit
// scale (from 0 to 0x7fffffff) as a prefix for the key.
// This is used to explicitely store some keys close to some node/host.
// Technically they are interlaced so that globally, square-like shapes stick
// together, and avoid the all row 0 first, then all row 1, etc.
func Gen2d(seed []byte, keyName string, x int32, y int32) ([]byte, error) {
	kBuf, err := Gen(seed, keyName)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 2d keydx")
	}
	kInt, err := vpcrypto.BufToInt256(kBuf)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 2d keydx")
	}
	xInt, err := toBigInt31(x)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 2d keydx")
	}
	yInt, err := toBigInt31(y)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 2d keydx")
	}
	for i := 0; i < n31; i++ {
		kInt = kInt.SetBit(kInt, nOffset2+2*i, xInt.Bit(i))
		kInt = kInt.SetBit(kInt, nOffset2+2*i+1, yInt.Bit(i))
	}
	return vpcrypto.IntToBuf256(kInt), nil
}

// Gen3d generates a unique ID for a key, given a seed.
// A typical usage of seed is the vring (virtual ring) name,
// which can be used to differenciate key locations on various
// vrings.
// The x, y and z parameter have a special meaning, they are used on a 31-bit
// scale (from 0 to 0x7fffffff) as a prefix for the key.
// This is used to explicitely store some keys close to some node/host.
// Technically they are interlaced so that globally, cube-like shapes stick
// together, and avoid the all row 0 first, then all row 1, etc.
func Gen3d(seed []byte, keyName string, x int32, y int32, z int32) ([]byte, error) {
	kBuf, err := Gen(seed, keyName)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 3d keydx")
	}
	kInt, err := vpcrypto.BufToInt256(kBuf)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 3d keydx")
	}
	xInt, err := toBigInt31(x)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 2d keydx")
	}
	yInt, err := toBigInt31(y)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 3d keydx")
	}
	zInt, err := toBigInt31(z)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to generate 3d keydx")
	}
	for i := 0; i < n31; i++ {
		kInt = kInt.SetBit(kInt, nOffset3+3*i, xInt.Bit(i))
		kInt = kInt.SetBit(kInt, nOffset3+3*i+1, yInt.Bit(i))
		kInt = kInt.SetBit(kInt, nOffset3+3*i+2, zInt.Bit(i))
	}
	return vpcrypto.IntToBuf256(kInt), nil
}
