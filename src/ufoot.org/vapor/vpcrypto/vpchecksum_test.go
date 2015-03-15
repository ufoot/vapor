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
	checksumInt := big.NewInt(1)
	checksumBuf := IntToBuf512(checksumInt)
	checksumStr := BufToStr512(checksumBuf)
	t.Logf("checksum 512 str = %s (len %d)", checksumStr, len(checksumStr))
	checksumInt2, err := StrToInt512(checksumStr)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksumInt.Cmp(checksumInt2) != 0 {
		t.Error("checksum ints 512 differ")
	}
	checksumStr2 := IntToStr512(checksumInt)
	t.Logf("checksum 512 str2 = %s (len %d)", checksumStr2, len(checksumStr2))
	if checksumStr != checksumStr2 {
		t.Error("checksums strings 512 differ")
	}
	checksumBuf2, err2 := StrToBuf512(checksumStr)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksumInt3, err3 := BufToInt512(checksumBuf2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksumInt.Cmp(checksumInt3) != 0 {
		t.Error("checksum ints 512 differ")
	}
}

func TestConvert256(t *testing.T) {
	checksumInt := big.NewInt(1)
	checksumBuf := IntToBuf256(checksumInt)
	checksumStr := BufToStr256(checksumBuf)
	t.Logf("checksum 256 str = %s (len %d)", checksumStr, len(checksumStr))
	checksumInt2, err := StrToInt256(checksumStr)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksumInt.Cmp(checksumInt2) != 0 {
		t.Error("checksum ints 256 differ")
	}
	checksumStr2 := IntToStr256(checksumInt)
	t.Logf("checksum 256 str2 = %s (len %d)", checksumStr2, len(checksumStr2))
	if checksumStr != checksumStr2 {
		t.Error("checksums strings 256 differ")
	}
	checksumBuf2, err2 := StrToBuf256(checksumStr)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksumInt3, err3 := BufToInt256(checksumBuf2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksumInt.Cmp(checksumInt3) != 0 {
		t.Error("checksum ints 256 differ")
	}
}

func TestConvert128(t *testing.T) {
	checksumInt := big.NewInt(1)
	checksumBuf := IntToBuf128(checksumInt)
	checksumStr := BufToStr128(checksumBuf)
	t.Logf("checksum 128 str = %s (len %d)", checksumStr, len(checksumStr))
	checksumInt2, err := StrToInt128(checksumStr)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksumInt.Cmp(checksumInt2) != 0 {
		t.Error("checksum ints 128 differ")
	}
	checksumStr2 := IntToStr128(checksumInt)
	t.Logf("checksum 128 str2 = %s (len %d)", checksumStr2, len(checksumStr2))
	if checksumStr != checksumStr2 {
		t.Error("checksums strings 128 differ")
	}
	checksumBuf2, err2 := StrToBuf128(checksumStr)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksumInt3, err3 := BufToInt128(checksumBuf2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksumInt.Cmp(checksumInt3) != 0 {
		t.Error("checksum ints 128 differ")
	}
}

func TestConvert64(t *testing.T) {
	checksumInt := uint64(1)
	checksumBuf := IntToBuf64(checksumInt)
	checksumStr := BufToStr64(checksumBuf)
	t.Logf("checksum 64 str = %s (len %d)", checksumStr, len(checksumStr))
	checksumInt2, err := StrToInt64(checksumStr)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksumInt != checksumInt2 {
		t.Error("checksum ints 64 differ")
	}
	checksumStr2 := IntToStr64(checksumInt)
	t.Logf("checksum 64 str2 = %s (len %d)", checksumStr2, len(checksumStr2))
	if checksumStr != checksumStr2 {
		t.Error("checksums strings 64 differ")
	}
	checksumBuf2, err2 := StrToBuf64(checksumStr)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksumInt3, err3 := BufToInt64(checksumBuf2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksumInt != checksumInt3 {
		t.Error("checksum ints 64 differ")
	}
}

func TestConvert32(t *testing.T) {
	checksumInt := uint32(1)
	checksumBuf := IntToBuf32(checksumInt)
	checksumStr := BufToStr32(checksumBuf)
	t.Logf("checksum 32 str = %s (len %d)", checksumStr, len(checksumStr))
	checksumInt2, err := StrToInt32(checksumStr)
	if err != nil {
		t.Error("string to int error", err)
	}
	if checksumInt != checksumInt2 {
		t.Error("checksum ints 32 differ")
	}
	checksumStr2 := IntToStr32(checksumInt)
	t.Logf("checksum 32 str2 = %s (len %d)", checksumStr2, len(checksumStr2))
	if checksumStr != checksumStr2 {
		t.Error("checksums strings 32 differ")
	}
	checksumBuf2, err2 := StrToBuf32(checksumStr)
	if err2 != nil {
		t.Error("string to bytes error", err2)
	}
	checksumInt3, err3 := BufToInt32(checksumBuf2)
	if err3 != nil {
		t.Error("bytes to int error", err3)
	}
	if checksumInt != checksumInt3 {
		t.Error("checksum ints 32 differ")
	}
}

func TestChecksum512(t *testing.T) {
	var rand512N [10]string

	const expectedRand512N0 string = "f7fbba6e0636f890e56fbbf3283e524c6fa3204ae298382d624741d0dc6638326e282c41be5e4254d8820772c5518a2c5a8c0c7f7eda19594a7eb539453e1ed7"
	const expectedRand512N9 string = "ce9933d84226e879fe45d36e778a8c8974cb4846a29192bde6a1ec895838959f802da42c7d92d7148efabdf1402b521191d8772a25adf632369ee9c25d348609"

	for i := range rand512N {
		if i == 0 {
			rand512N[i] = BufToStr512(Checksum512([]byte("foo")))
		} else {
			rand512N[i] = BufToStr512(Checksum512([]byte(rand512N[i-1])))
		}
	}

	for i, v := range rand512N {
		t.Logf("rand512N[%d]=%s", i, v)
	}

	if rand512N[0] != expectedRand512N0 {
		t.Errorf("rand512N[0] is %s, should be %s", rand512N[0], expectedRand512N0)
	}

	if rand512N[9] != expectedRand512N9 {
		t.Errorf("rand512N[9] is %s, should be %s", rand512N[9], expectedRand512N9)
	}
}

func TestChecksum256(t *testing.T) {
	var rand256N [10]string

	const expectedRand256N0 string = "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
	const expectedRand256N9 string = "437ae920503271030b1f5d2d2c20ba10f85bb50ba088b61b2c570b902e9f100f"

	for i := range rand256N {
		if i == 0 {
			rand256N[i] = BufToStr256(Checksum256([]byte("foo")))
		} else {
			rand256N[i] = BufToStr256(Checksum256([]byte(rand256N[i-1])))
		}
	}

	for i, v := range rand256N {
		t.Logf("rand256N[%d]=%s", i, v)
	}

	if rand256N[0] != expectedRand256N0 {
		t.Errorf("rand256N[0] is %s, should be %s", rand256N[0], expectedRand256N0)
	}

	if rand256N[9] != expectedRand256N9 {
		t.Errorf("rand256N[9] is %s, should be %s", rand256N[9], expectedRand256N9)
	}
}

func TestChecksum128(t *testing.T) {
	var rand128N [10]string

	const expectedRand128N0 string = "c7b5ea3f0fdbc95d0dd47f3c5bc275da"
	const expectedRand128N9 string = "9b3fc706cc7f619d1c64c479c8d78c56"

	for i := range rand128N {
		if i == 0 {
			rand128N[i] = BufToStr128(Checksum128([]byte("foo")))
		} else {
			rand128N[i] = BufToStr128(Checksum128([]byte(rand128N[i-1])))
		}
	}

	for i, v := range rand128N {
		t.Logf("rand128N[%d]=%s", i, v)
	}

	if rand128N[0] != expectedRand128N0 {
		t.Errorf("rand128N[0] is %s, should be %s", rand128N[0], expectedRand128N0)
	}

	if rand128N[9] != expectedRand128N9 {
		t.Errorf("rand128N[9] is %s, should be %s", rand128N[9], expectedRand128N9)
	}
}

func TestChecksum64(t *testing.T) {
	var rand64N [10]string

	const expectedRand64N0 string = "edef654fccc4a4d8"
	const expectedRand64N9 string = "417827c81852366a"

	for i := range rand64N {
		if i == 0 {
			rand64N[i] = BufToStr64(Checksum64([]byte("foo")))
		} else {
			rand64N[i] = BufToStr64(Checksum64([]byte(rand64N[i-1])))
		}
	}

	for i, v := range rand64N {
		t.Logf("rand64N[%d]=%s", i, v)
	}

	if rand64N[0] != expectedRand64N0 {
		t.Errorf("rand64N[0] is %s, should be %s", rand64N[0], expectedRand64N0)
	}

	if rand64N[9] != expectedRand64N9 {
		t.Errorf("rand64N[9] is %s, should be %s", rand64N[9], expectedRand64N9)
	}
}

func TestChecksum32(t *testing.T) {
	var rand32N [10]string

	const expectedRand32N0 string = "acbd18db"
	const expectedRand32N9 string = "9d232ba4"

	for i := range rand32N {
		if i == 0 {
			rand32N[i] = BufToStr32(Checksum32([]byte("foo")))
		} else {
			rand32N[i] = BufToStr32(Checksum32([]byte(rand32N[i-1])))
		}
	}

	for i, v := range rand32N {
		t.Logf("rand32N[%d]=%s", i, v)
	}

	if rand32N[0] != expectedRand32N0 {
		t.Errorf("rand32N[0] is %s, should be %s", rand32N[0], expectedRand32N0)
	}

	if rand32N[9] != expectedRand32N9 {
		t.Errorf("rand32N[9] is %s, should be %s", rand32N[9], expectedRand32N9)
	}
}

func TestPseudoRand512(t *testing.T) {
	limit := big.NewInt(1000)
	const expectedN = 489

	n := PseudoRand512([]byte("bar"), limit).Int64()
	if n != expectedN {
		t.Errorf("n is %d, should be %d", n, expectedN)
	}
	_ = PseudoRand512([]byte(""), big.NewInt(0))
}

func TestPseudoRand256(t *testing.T) {
	limit := big.NewInt(1000)
	const expectedN = 929

	n := PseudoRand256([]byte("bar"), limit).Int64()
	if n != expectedN {
		t.Errorf("n is %d, should be %d", n, expectedN)
	}
	_ = PseudoRand256([]byte(""), big.NewInt(0))
}

func TestPseudoRand128(t *testing.T) {
	limit := big.NewInt(1000)
	const expectedN = 472

	n := PseudoRand128([]byte("bar"), limit).Int64()
	if n != expectedN {
		t.Errorf("n is %d, should be %d", n, expectedN)
	}
	_ = PseudoRand128([]byte(""), big.NewInt(0))
}

func TestPseudoRand64(t *testing.T) {
	const limit = 1000
	const expectedN = 760

	n := PseudoRand64(123, limit)
	if n != expectedN {
		t.Errorf("n is %d, should be %d", n, expectedN)
	}
	_ = PseudoRand64(456, 0)
}

func TestPseudoRand32(t *testing.T) {
	const limit = 1000
	const expectedN = 104

	n := PseudoRand32(123, limit)
	if n != expectedN {
		t.Errorf("n is %d, should be %d", n, expectedN)
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
