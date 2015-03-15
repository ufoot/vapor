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

func TestConvert512(t *testing.T) {
	checksum_int := big.NewInt(1)
	checksum_bytes := IntToBuf512(checksum_int)
	checksum_str := BufToStr512(checksum_bytes)
	t.Logf("checksum 512 str = %s (len %d)", checksum_str, len(checksum_str))
	checksum_int2, err := StrToInt512(checksum_str)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksum_int.Cmp(checksum_int2) != 0 {
		t.Error("checksum ints 512 differ")
	}
	checksum_str2 := IntToStr512(checksum_int)
	t.Logf("checksum 512 str2 = %s (len %d)", checksum_str2, len(checksum_str2))
	if checksum_str != checksum_str2 {
		t.Error("checksums strings 512 differ")
	}
	checksum_bytes2, err2 := StrToBuf512(checksum_str)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksum_int3, err3 := BufToInt512(checksum_bytes2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksum_int.Cmp(checksum_int3) != 0 {
		t.Error("checksum ints 512 differ")
	}
}

func TestConvert256(t *testing.T) {
	checksum_int := big.NewInt(1)
	checksum_bytes := IntToBuf256(checksum_int)
	checksum_str := BufToStr256(checksum_bytes)
	t.Logf("checksum 256 str = %s (len %d)", checksum_str, len(checksum_str))
	checksum_int2, err := StrToInt256(checksum_str)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksum_int.Cmp(checksum_int2) != 0 {
		t.Error("checksum ints 256 differ")
	}
	checksum_str2 := IntToStr256(checksum_int)
	t.Logf("checksum 256 str2 = %s (len %d)", checksum_str2, len(checksum_str2))
	if checksum_str != checksum_str2 {
		t.Error("checksums strings 256 differ")
	}
	checksum_bytes2, err2 := StrToBuf256(checksum_str)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksum_int3, err3 := BufToInt256(checksum_bytes2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksum_int.Cmp(checksum_int3) != 0 {
		t.Error("checksum ints 256 differ")
	}
}

func TestConvert128(t *testing.T) {
	checksum_int := big.NewInt(1)
	checksum_bytes := IntToBuf128(checksum_int)
	checksum_str := BufToStr128(checksum_bytes)
	t.Logf("checksum 128 str = %s (len %d)", checksum_str, len(checksum_str))
	checksum_int2, err := StrToInt128(checksum_str)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksum_int.Cmp(checksum_int2) != 0 {
		t.Error("checksum ints 128 differ")
	}
	checksum_str2 := IntToStr128(checksum_int)
	t.Logf("checksum 128 str2 = %s (len %d)", checksum_str2, len(checksum_str2))
	if checksum_str != checksum_str2 {
		t.Error("checksums strings 128 differ")
	}
	checksum_bytes2, err2 := StrToBuf128(checksum_str)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksum_int3, err3 := BufToInt128(checksum_bytes2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksum_int.Cmp(checksum_int3) != 0 {
		t.Error("checksum ints 128 differ")
	}
}

func TestConvert64(t *testing.T) {
	checksum_int := uint64(1)
	checksum_bytes := IntToBuf64(checksum_int)
	checksum_str := BufToStr64(checksum_bytes)
	t.Logf("checksum 64 str = %s (len %d)", checksum_str, len(checksum_str))
	checksum_int2, err := StrToInt64(checksum_str)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksum_int != checksum_int2 {
		t.Error("checksum ints 64 differ")
	}
	checksum_str2 := IntToStr64(checksum_int)
	t.Logf("checksum 64 str2 = %s (len %d)", checksum_str2, len(checksum_str2))
	if checksum_str != checksum_str2 {
		t.Error("checksums strings 64 differ")
	}
	checksum_bytes2, err2 := StrToBuf64(checksum_str)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksum_int3, err3 := BufToInt64(checksum_bytes2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksum_int != checksum_int3 {
		t.Error("checksum ints 64 differ")
	}
}

func TestConvert32(t *testing.T) {
	checksum_int := uint32(1)
	checksum_bytes := IntToBuf32(checksum_int)
	checksum_str := BufToStr32(checksum_bytes)
	t.Logf("checksum 32 str = %s (len %d)", checksum_str, len(checksum_str))
	checksum_int2, err := StrToInt32(checksum_str)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksum_int != checksum_int2 {
		t.Error("checksum ints 32 differ")
	}
	checksum_str2 := IntToStr32(checksum_int)
	t.Logf("checksum 32 str2 = %s (len %d)", checksum_str2, len(checksum_str2))
	if checksum_str != checksum_str2 {
		t.Error("checksums strings 32 differ")
	}
	checksum_bytes2, err2 := StrToBuf32(checksum_str)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksum_int3, err3 := BufToInt32(checksum_bytes2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksum_int != checksum_int3 {
		t.Error("checksum ints 32 differ")
	}
}

func TestChecksum512(t *testing.T) {
	var rand512_n [10]string

	const expected_rand512_n_0 string = "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7"
	const expected_rand512_n_9 string = "ce9933d84226e879fe45d36e778a8c8974cb4846a29192bde6a1ec895838959f802da42c7d92d7148efabdf1402b521191d8772a25adf632369ee9c25d348609"

	for i, _ := range rand512_n {
		if i == 0 {
			rand512_n[i] = BufToStr512(Checksum512([]byte("foo")))
		} else {
			rand512_n[i] = BufToStr512(Checksum512([]byte(rand512_n[i-1])))
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

func TestChecksum256(t *testing.T) {
	var rand256_n [10]string

	const expected_rand256_n_0 string = "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
	const expected_rand256_n_9 string = "437ae920503271030b1f5d2d2c20ba10f85bb50ba088b61b2c570b902e9f100f"

	for i, _ := range rand256_n {
		if i == 0 {
			rand256_n[i] = BufToStr256(Checksum256([]byte("foo")))
		} else {
			rand256_n[i] = BufToStr256(Checksum256([]byte(rand256_n[i-1])))
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

func TestChecksum128(t *testing.T) {
	var rand128_n [10]string

	const expected_rand128_n_0 string = "c7b5ea3f0fdbc95d0dd47f3c5bc275da"
	const expected_rand128_n_9 string = "9b3fc706cc7f619d1c64c479c8d78c56"

	for i, _ := range rand128_n {
		if i == 0 {
			rand128_n[i] = BufToStr128(Checksum128([]byte("foo")))
		} else {
			rand128_n[i] = BufToStr128(Checksum128([]byte(rand128_n[i-1])))
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

func TestChecksum64(t *testing.T) {
	var rand64_n [10]string

	const expected_rand64_n_0 string = "edef654fccc4a4d8"
	const expected_rand64_n_9 string = "417827c81852366a"

	for i, _ := range rand64_n {
		if i == 0 {
			rand64_n[i] = BufToStr64(Checksum64([]byte("foo")))
		} else {
			rand64_n[i] = BufToStr64(Checksum64([]byte(rand64_n[i-1])))
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

func TestChecksum32(t *testing.T) {
	var rand32_n [10]string

	const expected_rand32_n_0 string = "acbd18db"
	const expected_rand32_n_9 string = "9d232ba4"

	for i, _ := range rand32_n {
		if i == 0 {
			rand32_n[i] = BufToStr32(Checksum32([]byte("foo")))
		} else {
			rand32_n[i] = BufToStr32(Checksum32([]byte(rand32_n[i-1])))
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

func TestPseudoRand512(t *testing.T) {
	limit := big.NewInt(1000)
	const expected_n = 489

	n := PseudoRand512([]byte("bar"), limit).Int64()
	if n != expected_n {
		t.Errorf("n is %d, should be %d", n, expected_n)
	}
	_ = PseudoRand512([]byte(""), big.NewInt(0))
}

func TestPseudoRand256(t *testing.T) {
	limit := big.NewInt(1000)
	const expected_n = 929

	n := PseudoRand256([]byte("bar"), limit).Int64()
	if n != expected_n {
		t.Errorf("n is %d, should be %d", n, expected_n)
	}
	_ = PseudoRand256([]byte(""), big.NewInt(0))
}

func TestPseudoRand128(t *testing.T) {
	limit := big.NewInt(1000)
	const expected_n = 472

	n := PseudoRand128([]byte("bar"), limit).Int64()
	if n != expected_n {
		t.Errorf("n is %d, should be %d", n, expected_n)
	}
	_ = PseudoRand128([]byte(""), big.NewInt(0))
}

func TestPseudoRand64(t *testing.T) {
	const limit = 1000
	const expected_n = 760

	n := PseudoRand64(123, limit)
	if n != expected_n {
		t.Errorf("n is %d, should be %d", n, expected_n)
	}
	_ = PseudoRand64(456, 0)
}

func TestPseudoRand32(t *testing.T) {
	const limit = 1000
	const expected_n = 104

	n := PseudoRand32(123, limit)
	if n != expected_n {
		t.Errorf("n is %d, should be %d", n, expected_n)
	}
	_ = PseudoRand32(456, 0)
}

func BenchmarkChecksum512(b *testing.B) {
	data := make([]byte, 2048)

	for i := 0; i < b.N; i++ {
		_ = Checksum512(data)
	}
}

func BenchmarkChecksum256(b *testing.B) {
	data := make([]byte, 2048)

	for i := 0; i < b.N; i++ {
		_ = Checksum256(data)
	}
}

func BenchmarkChecksum128(b *testing.B) {
	data := make([]byte, 2048)

	for i := 0; i < b.N; i++ {
		_ = Checksum128(data)
	}
}

func BenchmarkChecksum64(b *testing.B) {
	data := make([]byte, 2048)

	for i := 0; i < b.N; i++ {
		_ = Checksum64(data)
	}
}

func BenchmarkChecksum32(b *testing.B) {
	data := make([]byte, 2048)

	for i := 0; i < b.N; i++ {
		_ = Checksum32(data)
	}
}

func BenchmarkPseudoRand512(b *testing.B) {
	data := make([]byte, 1024)
	limit := big.NewInt(10000)

	for i := 0; i < b.N; i++ {
		_ = PseudoRand512(data, limit)
	}
}

func BenchmarkPseudoRand256(b *testing.B) {
	data := make([]byte, 1024)
	limit := big.NewInt(10000)

	for i := 0; i < b.N; i++ {
		_ = PseudoRand256(data, limit)
	}
}

func BenchmarkPseudoRand128(b *testing.B) {
	data := make([]byte, 1024)
	limit := big.NewInt(10000)

	for i := 0; i < b.N; i++ {
		_ = PseudoRand128(data, limit)
	}
}

func BenchmarkPseudoRand64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PseudoRand64(1, 1000)
	}
}

func BenchmarkPseudoRand32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PseudoRand32(1, 1000)
	}
}
