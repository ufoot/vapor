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
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpcrypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"hash/crc32"
	"hash/crc64"
	"math/big"
	"ufoot.org/vapor/vpsys"
)

const positiveMask64 = 0x7fffffffffffffff
const positiveMask32 = 0x7fffffff

var crc32_table *crc32.Table
var crc64_table *crc64.Table
var big1 *big.Int

func init() {
	crc64_table = crc64.MakeTable(crc64.ECMA)
	crc32_table = crc32.MakeTable(crc32.IEEE)
	big1 = big.NewInt(1)
}

func intToBufN(checksum *big.Int, bits int) []byte {
	bytes := bits >> 3
	ret := checksum.Bytes()

	l := len(ret)
	switch {
	case l == bytes:
		return ret
	case l > bytes:
		return ret[0:bytes]
	}

	// l < bytes
	ret2 := make([]byte, bytes)
	for i := range ret {
		ret2[i+bytes-l] = ret[i]
	}
	return ret2
}

func intToStrN(checksum *big.Int, bits int) string {
	return bytesToStrN(intToBufN(checksum, bits), bits)
}

func bytesToIntN(checksum []byte, bits int) (*big.Int, error) {
	var ret big.Int
	bytes := bits >> 3

	if len(checksum) != bytes {
		return nil, errors.New("bad checksum size")
	}
	ret.SetBytes(checksum[0:bytes])

	return &ret, nil
}

func stringToIntN(checksum string, bits int) (*big.Int, error) {
	var buf []byte
	var ret *big.Int
	var err error

	buf, err = stringToBufN(checksum, bits)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to decode checksum hex string")
	}

	ret, err = bytesToIntN(buf, bits)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to get checksum bytes")
	}

	return ret, nil
}

func stringToBufN(checksum string, bits int) ([]byte, error) {
	var ret []byte
	var err error

	if len(checksum) != bits>>2 {
		return nil, errors.New("bad checksum size")
	}
	ret, err = hex.DecodeString(checksum)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to decode checksum hex string")
	}

	return ret, nil
}

func bytesToStrN(checksum []byte, bits int) string {
	l := len(checksum)
	bytes := bits >> 3
	switch {
	case l == bytes:
		return hex.EncodeToString(checksum)
	case l > bytes:
		return hex.EncodeToString(checksum[0:bytes])
	}

	// l < bytes
	checksum2 := make([]byte, bytes)
	for i := range checksum {
		checksum2[i+bytes-l] = checksum[i]
	}
	return hex.EncodeToString(checksum2)
}

func IntToBuf512(checksum *big.Int) []byte {
	return intToBufN(checksum, 512)
}

func IntToBuf256(checksum *big.Int) []byte {
	return intToBufN(checksum, 256)
}

func IntToBuf128(checksum *big.Int) []byte {
	return intToBufN(checksum, 128)
}

func IntToBuf64(checksum uint64) []byte {
	ret := make([]byte, 8)
	binary.BigEndian.PutUint64(ret, checksum)

	return ret
}

func IntToBuf32(checksum uint32) []byte {
	ret := make([]byte, 4)
	binary.BigEndian.PutUint32(ret, checksum)

	return ret
}

func IntToStr512(checksum *big.Int) string {
	return intToStrN(checksum, 512)
}

func IntToStr256(checksum *big.Int) string {
	return intToStrN(checksum, 256)
}

func IntToStr128(checksum *big.Int) string {
	return intToStrN(checksum, 128)
}

func IntToStr64(checksum uint64) string {
	return bytesToStrN(IntToBuf64(checksum), 64)
}

func IntToStr32(checksum uint32) string {
	return bytesToStrN(IntToBuf32(checksum), 32)
}

func BufToInt512(checksum []byte) (*big.Int, error) {
	return bytesToIntN(checksum, 512)
}

func BufToInt256(checksum []byte) (*big.Int, error) {
	return bytesToIntN(checksum, 256)
}

func BufToInt128(checksum []byte) (*big.Int, error) {
	return bytesToIntN(checksum, 128)
}

func BufToInt64(checksum []byte) (uint64, error) {
	if len(checksum) != 8 {
		return 0, errors.New("bad checksum size")
	}
	ret := binary.BigEndian.Uint64(checksum)

	return ret, nil
}

func BufToInt32(checksum []byte) (uint32, error) {
	if len(checksum) != 4 {
		return 0, errors.New("bad checksum size")
	}
	ret := binary.BigEndian.Uint32(checksum)

	return uint32(ret), nil
}

func StrToInt512(checksum string) (*big.Int, error) {
	return stringToIntN(checksum, 512)
}

func StrToInt256(checksum string) (*big.Int, error) {
	return stringToIntN(checksum, 256)
}

func StrToInt128(checksum string) (*big.Int, error) {
	return stringToIntN(checksum, 128)
}

func StrToInt64(checksum string) (uint64, error) {
	bytes, err := stringToBufN(checksum, 64)
	if err != nil {
		return 0, vpsys.ErrorChain(err, "can't convert string to bytes")
	}
	return BufToInt64(bytes)
}

func StrToInt32(checksum string) (uint32, error) {
	bytes, err := stringToBufN(checksum, 32)
	if err != nil {
		return 0, vpsys.ErrorChain(err, "can't convert string to bytes")
	}
	return BufToInt32(bytes)
}

func StrToBuf512(checksum string) ([]byte, error) {
	return stringToBufN(checksum, 512)
}

func StrToBuf256(checksum string) ([]byte, error) {
	return stringToBufN(checksum, 256)
}

func StrToBuf128(checksum string) ([]byte, error) {
	return stringToBufN(checksum, 128)
}

func StrToBuf64(checksum string) ([]byte, error) {
	return stringToBufN(checksum, 64)
}

func StrToBuf32(checksum string) ([]byte, error) {
	return stringToBufN(checksum, 32)
}

func BufToStr512(checksum []byte) string {
	return bytesToStrN(checksum, 512)
}

func BufToStr256(checksum []byte) string {
	return bytesToStrN(checksum, 256)
}

func BufToStr128(checksum []byte) string {
	return bytesToStrN(checksum, 128)
}

func BufToStr64(checksum []byte) string {
	return bytesToStrN(checksum, 64)
}

func BufToStr32(checksum []byte) string {
	return bytesToStrN(checksum, 32)
}

func Checksum512(data []byte) []byte {
	sum := sha512.Sum512(data)

	return sum[0:64]
}

func Checksum256(data []byte) []byte {
	sum := sha256.Sum256(data)

	return sum[0:32]
}

func Checksum128(data []byte) []byte {
	sum := sha1.Sum(data)

	return sum[2:18]
}

func Checksum64(data []byte) []byte {
	sum := md5.Sum(data)

	return sum[8:16]
}

func Checksum32(data []byte) []byte {
	sum := md5.Sum(data)

	return sum[0:4]
}

func PseudoRand512(seed []byte, n *big.Int) *big.Int {
	checksum, err := BufToInt512(Checksum512(seed))

	if err != nil {
		return nil
	}

	if n != nil && n.Cmp(big1) > 0 {
		var ret big.Int
		ret.Mod(checksum, n)
		return &ret
	}
	return checksum
}

func PseudoRand256(seed []byte, n *big.Int) *big.Int {
	checksum, err := BufToInt256(Checksum256(seed))

	if err != nil {
		return nil
	}

	if n != nil && n.Cmp(big1) > 0 {
		var ret big.Int
		ret.Mod(checksum, n)
		return &ret
	}
	return checksum
}

func PseudoRand128(seed []byte, n *big.Int) *big.Int {
	checksum, err := BufToInt128(Checksum128(seed))

	if err != nil {
		return nil
	}

	if n != nil && n.Cmp(big1) > 0 {
		var ret big.Int
		ret.Mod(checksum, n)
		return &ret
	}
	return checksum
}

func PseudoRand64(seed, n uint64) uint64 {
	checksum := crc64.Checksum(IntToBuf64(seed), crc64_table)

	if n > 1 {
		return checksum % n
	}
	return checksum
}

func PseudoRand32(seed, n uint32) uint32 {
	checksum := crc32.Checksum(IntToBuf32(seed), crc32_table)

	if n > 1 {
		return checksum % n
	}
	return checksum
}
