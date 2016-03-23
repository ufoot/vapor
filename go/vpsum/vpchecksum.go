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

package vpsum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/ufoot/vapor/go/vperror"
	"hash/crc32"
	"hash/crc64"
	"math/big"
)

const positiveMask64 = 0x7fffffffffffffff
const positiveMask32 = 0x7fffffff

var crc32Table *crc32.Table
var crc64Table *crc64.Table
var bigOne *big.Int

func init() {
	crc64Table = crc64.MakeTable(crc64.ECMA)
	crc32Table = crc32.MakeTable(crc32.IEEE)
	bigOne = big.NewInt(1)
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
	return bufToStrN(intToBufN(checksum, bits), bits)
}

func bufToIntN(checksum []byte, bits int) (*big.Int, error) {
	var ret big.Int
	bytes := bits >> 3

	if len(checksum) != bytes {
		return nil, errors.New("bad checksum size")
	}
	ret.SetBytes(checksum[0:bytes])

	return &ret, nil
}

func strToIntN(checksum string, bits int) (*big.Int, error) {
	var buf []byte
	var ret *big.Int
	var err error

	buf, err = strToBufN(checksum, bits)
	if err != nil {
		return nil, vperror.Chain(err, "unable to decode checksum hex string")
	}

	ret, err = bufToIntN(buf, bits)
	if err != nil {
		return nil, vperror.Chain(err, "unable to get checksum bytes")
	}

	return ret, nil
}

func strToBufN(checksum string, bits int) ([]byte, error) {
	var ret []byte
	var err error

	if len(checksum) != bits>>2 {
		return nil, errors.New("bad checksum size")
	}
	ret, err = hex.DecodeString(checksum)
	if err != nil {
		return nil, vperror.Chain(err, "unable to decode checksum hex string")
	}

	return ret, nil
}

func bufToStrN(checksum []byte, bits int) string {
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

// IntToBuf512 converts a 512 bits integer to a byte buffer.
// If the number is too low, the buffer is padded with heading zeroes.
// So the result is always a 64 bytes buffer.
func IntToBuf512(checksum *big.Int) []byte {
	return intToBufN(checksum, 512)
}

// IntToBuf256 converts a 256 bits integer to a byte buffer.
// If the number is too low, the buffer is padded with heading zeroes.
// So the result is always a 32 bytes buffer.
func IntToBuf256(checksum *big.Int) []byte {
	return intToBufN(checksum, 256)
}

// IntToBuf128 converts a 128 bits integer to a byte buffer.
// If the number is too low, the buffer is padded with heading zeroes.
// So the result is always a 16 bytes buffer.
func IntToBuf128(checksum *big.Int) []byte {
	return intToBufN(checksum, 128)
}

// IntToBuf64 converts a 64 bits integer to a byte buffer.
// If the number is too low, the buffer is padded with heading zeroes.
// So the result is always an 8 bytes buffer.
func IntToBuf64(checksum uint64) []byte {
	ret := make([]byte, 8)
	binary.BigEndian.PutUint64(ret, checksum)

	return ret
}

// IntToBuf32 converts a 32 bits integer to a byte buffer.
// If the number is too low, the buffer is padded with heading zeroes.
// So the result is always a 4 bytes buffer.
func IntToBuf32(checksum uint32) []byte {
	ret := make([]byte, 4)
	binary.BigEndian.PutUint32(ret, checksum)

	return ret
}

// IntToStr512 converts a 512 bits integer to a string.
// It uses an hexadecimal representation of the number.
// If the number is too low, the string is padded with heading zeroes.
// So the result is always a 128 bytes/runes string.
func IntToStr512(checksum *big.Int) string {
	return intToStrN(checksum, 512)
}

// IntToStr256 converts a 256 bits integer to a string.
// It uses an hexadecimal representation of the number.
// If the number is too low, the string is padded with heading zeroes.
// So the result is always a 64 bytes/runes string.
func IntToStr256(checksum *big.Int) string {
	return intToStrN(checksum, 256)
}

// IntToStr128 converts a 128 bits integer to a string.
// It uses an hexadecimal representation of the number.
// If the number is too low, the string is padded with heading zeroes.
// So the result is always a 32 bytes/runes string.
func IntToStr128(checksum *big.Int) string {
	return intToStrN(checksum, 128)
}

// IntToStr64 converts a 64 bits integer to a string.
// It uses an hexadecimal representation of the number.
// If the number is too low, the string is padded with heading zeroes.
// So the result is always a 16 bytes/runes string.
func IntToStr64(checksum uint64) string {
	return bufToStrN(IntToBuf64(checksum), 64)
}

// IntToStr32 converts a 32 bits integer to a string.
// It uses an hexadecimal representation of the number.
// If the number is too low, the string is padded with heading zeroes.
// So the result is always an 8 bytes/runes string.
func IntToStr32(checksum uint32) string {
	return bufToStrN(IntToBuf32(checksum), 32)
}

// BufToInt512 converts a 512 bits buffer to an integer.
// Buffer must be exactly 64 bytes, else it will fail.
func BufToInt512(checksum []byte) (*big.Int, error) {
	return bufToIntN(checksum, 512)
}

// BufToInt256 converts a 256 bits buffer to an integer.
// Buffer must be exactly 32 bytes, else it will fail.
func BufToInt256(checksum []byte) (*big.Int, error) {
	return bufToIntN(checksum, 256)
}

// BufToInt128 converts a 128 bits buffer to an integer.
// Buffer must be exactly 16 bytes, else it will fail.
func BufToInt128(checksum []byte) (*big.Int, error) {
	return bufToIntN(checksum, 128)
}

// BufToInt64 converts a 64 bits buffer to an integer.
// Buffer must be exactly 8 bytes, else it will fail.
func BufToInt64(checksum []byte) (uint64, error) {
	if len(checksum) != 8 {
		return 0, errors.New("bad checksum size")
	}
	ret := binary.BigEndian.Uint64(checksum)

	return ret, nil
}

// BufToInt32 converts a 32 bits buffer to an integer.
// Buffer must be exactly 4 bytes, else it will fail.
func BufToInt32(checksum []byte) (uint32, error) {
	if len(checksum) != 4 {
		return 0, errors.New("bad checksum size")
	}
	ret := binary.BigEndian.Uint32(checksum)

	return uint32(ret), nil
}

// StrToInt512 converts a 512 bits checksum string to an integer.
// The string must be exactly 128 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
func StrToInt512(checksum string) (*big.Int, error) {
	return strToIntN(checksum, 512)
}

// StrToInt256 converts a 256 bits checksum string to an integer.
// The string must be exactly 64 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
func StrToInt256(checksum string) (*big.Int, error) {
	return strToIntN(checksum, 256)
}

// StrToInt128 converts a 128 bits checksum string to an integer.
// The string must be exactly 32 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
func StrToInt128(checksum string) (*big.Int, error) {
	return strToIntN(checksum, 128)
}

// StrToInt64 converts a 64 bits checksum string to an integer.
// The string must be exactly 16 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
func StrToInt64(checksum string) (uint64, error) {
	bytes, err := strToBufN(checksum, 64)
	if err != nil {
		return 0, vperror.Chain(err, "can't convert string to bytes")
	}
	return BufToInt64(bytes)
}

// StrToInt32 converts a 32 bits checksum string to an integer.
// The string must be exactly 8 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
func StrToInt32(checksum string) (uint32, error) {
	bytes, err := strToBufN(checksum, 32)
	if err != nil {
		return 0, vperror.Chain(err, "can't convert string to bytes")
	}
	return BufToInt32(bytes)
}

// StrToBuf512 converts a 512 bits checksum string to a buffer.
// The string must be exactly 128 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
// The result buffer is twice shorter (64 bytes).
func StrToBuf512(checksum string) ([]byte, error) {
	return strToBufN(checksum, 512)
}

// StrToBuf256 converts a 256 bits checksum string to a buffer.
// The string must be exactly 64 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
// The result buffer is twice shorter (32 bytes).
func StrToBuf256(checksum string) ([]byte, error) {
	return strToBufN(checksum, 256)
}

// StrToBuf128 converts a 128 bits checksum string to a buffer.
// The string must be exactly 32 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
// The result buffer is twice shorter (16 bytes).
func StrToBuf128(checksum string) ([]byte, error) {
	return strToBufN(checksum, 128)
}

// StrToBuf64 converts a 64 bits checksum string to a buffer.
// The string must be exactly 16 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
// The result buffer is twice shorter (8 bytes).
func StrToBuf64(checksum string) ([]byte, error) {
	return strToBufN(checksum, 64)
}

// StrToBuf32 converts a 32 bits checksum string to a buffer.
// The string must be exactly 8 bytes/runes long and be readable
// as an hexadecimal number, else it will fail.
// The result buffer is twice shorter (4 bytes).
func StrToBuf32(checksum string) ([]byte, error) {
	return strToBufN(checksum, 32)
}

// BufToStr512 converts a 512 bits buffer to a checksum string.
// If the buffer does not have the right size (64 bytes),
// the string is truncated or padded with heading zeroes.
// So the result is always a 128 bytes/runes string.
func BufToStr512(checksum []byte) string {
	return bufToStrN(checksum, 512)
}

// BufToStr256 converts a 256 bits buffer to a checksum string.
// If the buffer does not have the right size (32 bytes),
// the string is truncated or padded with heading zeroes.
// So the result is always a 64 bytes/runes string.
func BufToStr256(checksum []byte) string {
	return bufToStrN(checksum, 256)
}

// BufToStr128 converts a 128 bits buffer to a checksum string.
// If the buffer does not have the right size (16 bytes),
// the string is truncated or padded with heading zeroes.
// So the result is always a 32 bytes/runes string.
func BufToStr128(checksum []byte) string {
	return bufToStrN(checksum, 128)
}

// BufToStr64 converts a 64 bits buffer to a checksum string.
// If the buffer does not have the right size (8 bytes),
// the string is truncated or padded with heading zeroes.
// So the result is always a 16 bytes/runes string.
func BufToStr64(checksum []byte) string {
	return bufToStrN(checksum, 64)
}

// BufToStr32 converts a 32 bits buffer to a checksum string.
// If the buffer does not have the right size (4 bytes),
// the string is truncated or padded with heading zeroes.
// So the result is always an 8 bytes/runes string.
func BufToStr32(checksum []byte) string {
	return bufToStrN(checksum, 32)
}

// Checksum512 calculates a 512 bits checksum of data.
// Internally, uses some cryptographic method such as SHA,
// but don't rely on this, just consider it's a checksum.
func Checksum512(data []byte) []byte {
	sum := sha512.Sum512(data)

	return sum[0:64]
}

// Checksum256 calculates a 256 bits checksum of data.
// Internally, uses some cryptographic method such as SHA,
// but don't rely on this, just consider it's a checksum.
func Checksum256(data []byte) []byte {
	sum := sha256.Sum256(data)

	return sum[0:32]
}

// Checksum128 calculates a 128 bits checksum of data.
// Internally, uses some cryptographic method such as SHA,
// but don't rely on this, just consider it's a checksum.
func Checksum128(data []byte) []byte {
	sum := sha1.Sum(data)

	return sum[2:18]
}

// Checksum64 calculates a 64 bits checksum of data.
// Internally, uses some cryptographic method such as MD5,
// but don't rely on this, just consider it's a checksum.
func Checksum64(data []byte) []byte {
	sum := md5.Sum(data)

	return sum[8:16]
}

// Checksum32 calculates a 32 bits checksum of data.
// Internally, uses some cryptographic method such as MD5,
// but don't rely on this, just consider it's a checksum.
func Checksum32(data []byte) []byte {
	sum := md5.Sum(data)

	return sum[0:4]
}

// PseudoRand512 returns a pseudo-random value on 512 bits.
// It is totally predictable, using seed as a input.
// Called with a given seed, always returns the same result.
// If n is nil or 0, returns a number between [0, 2^512).
// If n is greater than 0, returns a number between [0, n).
func PseudoRand512(seed []byte, n *big.Int) *big.Int {
	checksum, err := BufToInt512(Checksum512(seed))

	if err != nil {
		return nil
	}

	if n != nil && n.Cmp(bigOne) > 0 {
		var ret big.Int
		ret.Mod(checksum, n)
		return &ret
	}
	return checksum
}

// PseudoRand256 returns a pseudo-random value on 256 bits.
// It is totally predictable, using seed as a input.
// Called with a given seed, always returns the same result.
// If n is nil or 0, returns a number between [0, 2^256).
// If n is greater than 0, returns a number between [0, n).
func PseudoRand256(seed []byte, n *big.Int) *big.Int {
	checksum, err := BufToInt256(Checksum256(seed))

	if err != nil {
		return nil
	}

	if n != nil && n.Cmp(bigOne) > 0 {
		var ret big.Int
		ret.Mod(checksum, n)
		return &ret
	}
	return checksum
}

// PseudoRand128 returns a pseudo-random value on 128 bits.
// It is totally predictable, using seed as a input.
// Called with a given seed, always returns the same result.
// If n is nil or 0, returns a number between [0, 2^128).
// If n is greater than 0, returns a number between [0, n).
func PseudoRand128(seed []byte, n *big.Int) *big.Int {
	checksum, err := BufToInt128(Checksum128(seed))

	if err != nil {
		return nil
	}

	if n != nil && n.Cmp(bigOne) > 0 {
		var ret big.Int
		ret.Mod(checksum, n)
		return &ret
	}
	return checksum
}

// PseudoRand64 returns a pseudo-random value on 64 bits.
// It is totally predictable, using seed as a input.
// Called with a given seed, always returns the same result.
// If n is nil or 0, returns a number between [0, 2^64).
// If n is greater than 0, returns a number between [0, n).
func PseudoRand64(seed, n uint64) uint64 {
	checksum := crc64.Checksum(IntToBuf64(seed), crc64Table)

	if n > 1 {
		return checksum % n
	}
	return checksum
}

// PseudoRand32 returns a pseudo-random value on 32 bits.
// It is totally predictable, using seed as a input.
// Called with a given seed, always returns the same result.
// If n is nil or 0, returns a number between [0, 2^32).
// If n is greater than 0, returns a number between [0, n).
func PseudoRand32(seed, n uint32) uint32 {
	checksum := crc32.Checksum(IntToBuf32(seed), crc32Table)

	if n > 1 {
		return checksum % n
	}
	return checksum
}
