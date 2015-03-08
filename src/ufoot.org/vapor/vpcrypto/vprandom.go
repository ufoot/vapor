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
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"math/big"
	"ufoot.org/vapor/vpsys"
)

const positiveMask64 = 0x7fffffffffffffff
const positiveMask32 = 0x7fffffff

var crc32_table *crc32.Table
var crc64_table *crc64.Table

func init() {
	crc64_table = crc64.MakeTable(crc64.ECMA)
	crc32_table = crc32.MakeTable(crc32.IEEE)
}

func ChecksumUToSN(checksum *big.Int, bits int) string {
	var ret string

	ret = hex.EncodeToString(checksum.Bytes())
	for len(ret) < bits>>2 {
		// performance killer but checksum rarely
		// start by 0 so this is a fairly rare case
		ret = "0" + ret
	}

	return ret[0 : bits>>2]
}

func ChecksumUToS512(checksum *big.Int) string {
	return ChecksumUToSN(checksum, 512)
}

func ChecksumUToS256(checksum *big.Int) string {
	return ChecksumUToSN(checksum, 256)
}

func ChecksumUToS128(checksum *big.Int) string {
	return ChecksumUToSN(checksum, 128)
}

func ChecksumUToS64(checksum uint64) string {
	return fmt.Sprintf("%016x", checksum)
}

func ChecksumUToS32(checksum uint32) string {
	return fmt.Sprintf("%08x", checksum)
}

func ChecksumSToUN(checksum string, bits int) (*big.Int, error) {
	var buf []byte
	var ret big.Int
	var err error

	if len(checksum) == bits>>2 {
		buf, err = hex.DecodeString(checksum)
		if err == nil {
			ret.SetBytes(buf[0 : bits>>3])
		} else {
			return nil, vpsys.ErrorChain(err, "unable to decode checksum hex string")
		}
	} else {
		return nil, errors.New("bad checksum size")
	}

	return &ret, nil
}

func ChecksumSToU512(checksum string) (*big.Int, error) {
	return ChecksumSToUN(checksum, 512)
}

func ChecksumSToU256(checksum string) (*big.Int, error) {
	return ChecksumSToUN(checksum, 256)
}

func ChecksumSToU128(checksum string) (*big.Int, error) {
	return ChecksumSToUN(checksum, 128)
}

func ChecksumSToU64(checksum string) (uint64, error) {
	var n int
	var err error
	var ret uint64

	n, err = fmt.Sscanf(checksum, "%16x", &ret)
	if err != nil {
		return 0, err
	}
	if n == 1 {
		err = errors.New("too few elements")
	}
	return ret, nil
}

func ChecksumSToU32(checksum string) (uint32, error) {
	var n int
	var err error
	var ret uint32

	n, err = fmt.Sscanf(checksum, "%8x", &ret)
	if err != nil {
		return 0, err
	}
	if n == 1 {
		err = errors.New("too few elements")
	}
	return ret, nil
}

func PredictableRandomS512(seed string) string {
	var ret string
	var tmp big.Int

	sum := sha512.Sum512([]byte(seed))
	tmp.SetBytes(sum[0:])
	ret = ChecksumUToS512(&tmp)

	return ret
}

func PredictableRandomS256(seed string) string {
	var ret string
	var tmp big.Int

	sum := sha256.Sum256([]byte(seed))
	tmp.SetBytes(sum[0:])
	ret = ChecksumUToS256(&tmp)

	return ret
}

func PredictableRandomS128(seed string) string {
	var ret string
	var tmp big.Int

	sum := sha1.Sum([]byte(seed))
	tmp.SetBytes(sum[0:])
	ret = ChecksumUToS128(&tmp)

	return ret
}

func PredictableRandomS64(seed string) string {
	var ret string

	sum := md5.Sum([]byte(seed))
	ret = hex.EncodeToString(sum[8:16])

	return ret
}

func PredictableRandomS32(seed string) string {
	var ret string

	sum := md5.Sum([]byte(seed))
	ret = hex.EncodeToString(sum[0:4])

	return ret
}

func PredictableRandomU64(seed int64) uint64 {
	buf := make([]byte, 8)

	binary.LittleEndian.PutUint64(buf, uint64(seed))

	return crc64.Checksum(buf, crc64_table)
}

func PredictableRandomU32(seed int32) uint32 {
	buf := make([]byte, 4)

	binary.LittleEndian.PutUint32(buf, uint32(seed))

	return crc32.Checksum(buf, crc32_table)
}

func PredictableRandom64(seed int64, limit int64) (int64, error) {
	if limit <= 2 {
		return 0, errors.New(fmt.Sprintf("limit is %d, should be greater or equal to 2", limit))
	}
	return int64(PredictableRandomU64(seed)&positiveMask64) % limit, nil
}

func PredictableRandom32(seed int32, limit int32) (int32, error) {
	if limit <= 2 {
		return 0, errors.New(fmt.Sprintf("limit is %d, should be greater or equal to 2", limit))
	}
	return int32(PredictableRandomU32(seed)&positiveMask32) % limit, nil
}
