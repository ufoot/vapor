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
	"math/big"
	"testing"
)

func TestChecksumConvert512(t *testing.T) {
	const checksum_val int64 = 1
	var checksum_int *big.Int
	var checksum_str string
	var err error

	checksum_int = big.NewInt(checksum_val)
	checksum_str = ChecksumUToS512(checksum_int)
	t.Logf("checksum 512 = %s", checksum_str)
	checksum_int, err = ChecksumSToU512(checksum_str)
	if err != nil {
		t.Errorf("checksum 512 error %s", err.Error())
	}
	if checksum_int.Int64() != checksum_val {
		t.Errorf("checksum 512 mismatch %d", checksum_int.Int64())
	}
}

func TestChecksumConvert256(t *testing.T) {
	const checksum_val int64 = 1
	var checksum_int *big.Int
	var checksum_str string
	var err error

	checksum_int = big.NewInt(checksum_val)
	checksum_str = ChecksumUToS256(checksum_int)
	t.Logf("checksum 256 = %s", checksum_str)
	checksum_int, err = ChecksumSToU256(checksum_str)
	if err != nil {
		t.Errorf("checksum 256 error %s", err.Error())
	}
	if checksum_int.Int64() != checksum_val {
		t.Errorf("checksum 256 mismatch %d", checksum_int.Int64())
	}
}

func TestChecksumConvert128(t *testing.T) {
	const checksum_val int64 = 1
	var checksum_int *big.Int
	var checksum_str string
	var err error

	checksum_int = big.NewInt(checksum_val)
	checksum_str = ChecksumUToS128(checksum_int)
	t.Logf("checksum 128 = %s", checksum_str)
	checksum_int, err = ChecksumSToU128(checksum_str)
	if err != nil {
		t.Errorf("checksum 128 error %s", err.Error())
	}
	if checksum_int.Int64() != checksum_val {
		t.Errorf("checksum 128 mismatch %d", checksum_int.Int64())
	}
}

func TestChecksumConvert64(t *testing.T) {
	const checksum_val int64 = 1
	var checksum_int uint64
	var checksum_str string
	var err error

	checksum_int = uint64(checksum_val)
	checksum_str = ChecksumUToS64(checksum_int)
	t.Logf("checksum 64 = %s", checksum_str)
	checksum_int, err = ChecksumSToU64(checksum_str)
	if err != nil {
		t.Errorf("checksum 64 error %s", err.Error())
	}
	if int64(checksum_int) != checksum_val {
		t.Errorf("checksum 64 mismatch %d", checksum_int)
	}
}

func TestChecksumConvert32(t *testing.T) {
	const checksum_val int64 = 1
	var checksum_int uint32
	var checksum_str string
	var err error

	checksum_int = uint32(checksum_val)
	checksum_str = ChecksumUToS32(checksum_int)
	t.Logf("checksum 32 = %s", checksum_str)
	checksum_int, err = ChecksumSToU32(checksum_str)
	if err != nil {
		t.Errorf("checksum 32 error %s", err.Error())
	}
	if int64(checksum_int) != checksum_val {
		t.Errorf("checksum 32 mismatch %d", checksum_int)
	}
}

func TestPredictableRandomS512(t *testing.T) {
	var rand512_n [10]string

	const expected_rand512_n_0 string = "31bca02094eb78126a517b206a88c73cfa9ec6f704c7030d18212cace820f025f00bf0ea68dbf3f3a5436ca63b53bf7bf80ad8d5de7d8359d0b7fed9dbc3ab99"
	const expected_rand512_n_9 string = "729831a37d87a90f7dc4816088fdcc01c9b6be0b02736a3d215486371df7224cc83c5cdcf8041c997f8c78011da4a075f632867e58782bbb615b15788898ca1b"

	for i, _ := range rand512_n {
		if i == 0 {
			rand512_n[i] = PredictableRandomS512("0")
		} else {
			rand512_n[i] = PredictableRandomS512(rand512_n[i-1])
		}
	}

	for i, v := range rand512_n {
		t.Logf("rand512_n[%d]=%s", i, v)
	}

	if rand512_n[0] != expected_rand512_n_0 {
		t.Errorf("rand512_n[0] is %s, should be %s", rand512_n[0], expected_rand512_n_0)
	}

	if rand512_n[9] != expected_rand512_n_9 {
		t.Errorf("rand512_n[9] is %s, should be %s", rand512_n[9], expected_rand512_n_9)
	}
}

func TestPredictableRandomS256(t *testing.T) {
	var rand256_n [10]string

	const expected_rand256_n_0 string = "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9"
	const expected_rand256_n_9 string = "ae82b5a97107add16a053e09337a79516669a2dff9f56060956d221db13ba0c2"

	for i, _ := range rand256_n {
		if i == 0 {
			rand256_n[i] = PredictableRandomS256("0")
		} else {
			rand256_n[i] = PredictableRandomS256(rand256_n[i-1])
		}
	}

	for i, v := range rand256_n {
		t.Logf("rand256_n[%d]=%s", i, v)
	}

	if rand256_n[0] != expected_rand256_n_0 {
		t.Errorf("rand256_n[0] is %s, should be %s", rand256_n[0], expected_rand256_n_0)
	}

	if rand256_n[9] != expected_rand256_n_9 {
		t.Errorf("rand256_n[9] is %s, should be %s", rand256_n[9], expected_rand256_n_9)
	}
}

func TestPredictableRandomS128(t *testing.T) {
	var rand128_n [10]string

	const expected_rand128_n_0 string = "b6589fc6ab0dc82cf12099d1c2d40ab9"
	const expected_rand128_n_9 string = "97d04f6dbef880d28b547c1d4489f804"

	for i, _ := range rand128_n {
		if i == 0 {
			rand128_n[i] = PredictableRandomS128("0")
		} else {
			rand128_n[i] = PredictableRandomS128(rand128_n[i-1])
		}
	}

	for i, v := range rand128_n {
		t.Logf("rand128_n[%d]=%s", i, v)
	}

	if rand128_n[0] != expected_rand128_n_0 {
		t.Errorf("rand128_n[0] is %s, should be %s", rand128_n[0], expected_rand128_n_0)
	}

	if rand128_n[9] != expected_rand128_n_9 {
		t.Errorf("rand128_n[9] is %s, should be %s", rand128_n[9], expected_rand128_n_9)
	}
}

func TestPredictableRandomS64(t *testing.T) {
	var rand64_n [10]string

	const expected_rand64_n_0 string = "66e7dff9f98764da"
	const expected_rand64_n_9 string = "8c51fc19135d77d5"

	for i, _ := range rand64_n {
		if i == 0 {
			rand64_n[i] = PredictableRandomS64("0")
		} else {
			rand64_n[i] = PredictableRandomS64(rand64_n[i-1])
		}
	}

	for i, v := range rand64_n {
		t.Logf("rand64_n[%d]=%s", i, v)
	}

	if rand64_n[0] != expected_rand64_n_0 {
		t.Errorf("rand64_n[0] is %s, should be %s", rand64_n[0], expected_rand64_n_0)
	}

	if rand64_n[9] != expected_rand64_n_9 {
		t.Errorf("rand64_n[9] is %s, should be %s", rand64_n[9], expected_rand64_n_9)
	}
}

func TestPredictableRandomS32(t *testing.T) {
	var rand32_n [10]string

	const expected_rand32_n_0 string = "cfcd2084"
	const expected_rand32_n_9 string = "df702ed0"

	for i, _ := range rand32_n {
		if i == 0 {
			rand32_n[i] = PredictableRandomS32("0")
		} else {
			rand32_n[i] = PredictableRandomS32(rand32_n[i-1])
		}
	}

	for i, v := range rand32_n {
		t.Logf("rand32_n[%d]=%s", i, v)
	}

	if rand32_n[0] != expected_rand32_n_0 {
		t.Errorf("rand32_n[0] is %s, should be %s", rand32_n[0], expected_rand32_n_0)
	}

	if rand32_n[9] != expected_rand32_n_9 {
		t.Errorf("rand32_n[9] is %s, should be %s", rand32_n[9], expected_rand32_n_9)
	}
}

func TestPredictableRandomU64(t *testing.T) {
	var rand64_n [10]uint64

	const expected_rand64_n_0 uint64 = 0xb66a73654282cac0
	const expected_rand64_n_9 uint64 = 0x6e48c9ffaab595fc

	for i, _ := range rand64_n {
		if i == 0 {
			rand64_n[i] = PredictableRandomU64(0)
		} else {
			rand64_n[i] = PredictableRandomU64(int64(rand64_n[i-1]))
		}
	}

	for i, v := range rand64_n {
		t.Logf("rand64_n[%d]=%016x", i, v)
	}

	if rand64_n[0] != expected_rand64_n_0 {
		t.Errorf("rand64_n[0] is %016x, should be %016x", rand64_n[0], expected_rand64_n_0)
	}

	if rand64_n[9] != expected_rand64_n_9 {
		t.Errorf("rand64_n[9] is %016x, should be %016x", rand64_n[9], expected_rand64_n_9)
	}
}

func TestPredictableRandomU32(t *testing.T) {
	var rand32_n [10]uint32

	const expected_rand32_n_0 uint32 = 0x2144df1c
	const expected_rand32_n_9 uint32 = 0xe9ec3db1

	for i, _ := range rand32_n {
		if i == 0 {
			rand32_n[i] = PredictableRandomU32(0)
		} else {
			rand32_n[i] = PredictableRandomU32(int32(rand32_n[i-1]))
		}
	}

	for i, v := range rand32_n {
		t.Logf("rand32_n[%d]=%08x", i, v)
	}

	if rand32_n[0] != expected_rand32_n_0 {
		t.Errorf("rand32_n[0] is %08x, should be %08x", rand32_n[0], expected_rand32_n_0)
	}

	if rand32_n[9] != expected_rand32_n_9 {
		t.Errorf("rand32_n[9] is %08x, should be %08x", rand32_n[9], expected_rand32_n_9)
	}
}

func TestPredictableRandom64(t *testing.T) {
	var err error
	var rand64 int64

	const check64Seed int64 = 1
	const check64Limit int64 = 1000000000000
	const check64Result int64 = 608648394112

	rand64, err = PredictableRandom64(check64Seed, check64Limit)

	if err != nil {
		t.Error("error generating number", err)
	}

	t.Logf("rand64=%d", rand64)
	if rand64 != check64Result {
		t.Errorf("rand64 is %d, should be %d", rand64, check64Result)
	}

	rand64, err = PredictableRandom64(check64Seed, 1)

	if err != nil {
		t.Log("an error has been raised because limit is <=2, this is normal... ", err)
	} else {
		t.Error("an error should have been generated for a call to random with limit<=2 ", err)
	}
}

func TestPredictableRandom32(t *testing.T) {
	var err error
	var rand32 int32

	const check32Seed int32 = 1
	const check32Limit int32 = 1000000
	const check32Result int32 = 730553

	rand32, err = PredictableRandom32(check32Seed, check32Limit)

	if err != nil {
		t.Error("error generating number", err)
	}

	t.Logf("rand32=%d", rand32)
	if rand32 != check32Result {
		t.Errorf("rand32 is %d, should be %d", rand32, check32Result)
	}

	rand32, err = PredictableRandom32(check32Seed, 1)

	if err != nil {
		t.Log("an error has been raised because limit is <=2, this is normal... ", err)
	} else {
		t.Error("an error should have been generated for a call to random with limit<=2 ", err)
	}
}

func BenchmarkChecksumConvert512(b *testing.B) {
	const checksum_val int64 = 1
	var checksum_int *big.Int

	checksum_int = big.NewInt(checksum_val)

	for i := 0; i < b.N; i++ {
		checksum_int, _ = ChecksumSToU512(ChecksumUToS512(checksum_int))
	}
}

func BenchmarkSRandom512(b *testing.B) {
	var rand512 string

	for i := 0; i < b.N; i++ {
		rand512 = PredictableRandomS512(rand512)
	}
}

func BenchmarkSRandom256(b *testing.B) {
	var rand256 string

	for i := 0; i < b.N; i++ {
		rand256 = PredictableRandomS256(rand256)
	}
}

func BenchmarkSRandom128(b *testing.B) {
	var rand128 string

	for i := 0; i < b.N; i++ {
		rand128 = PredictableRandomS128(rand128)
	}
}

func BenchmarkSRandom64(b *testing.B) {
	var rand64 string

	for i := 0; i < b.N; i++ {
		rand64 = PredictableRandomS64(rand64)
	}
}

func BenchmarkSRandom32(b *testing.B) {
	var rand32 string

	for i := 0; i < b.N; i++ {
		rand32 = PredictableRandomS32(rand32)
	}
}

func BenchmarkURandom64(b *testing.B) {
	var rand64 uint64

	for i := 0; i < b.N; i++ {
		rand64 = PredictableRandomU64(int64(rand64))
	}
}

func BenchmarkURandom32(b *testing.B) {
	var rand32 uint32

	for i := 0; i < b.N; i++ {
		rand32 = PredictableRandomU32(int32(rand32))
	}
}
