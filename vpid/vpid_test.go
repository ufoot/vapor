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

package vpid

import (
	"encoding/hex"
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpsum"
	"math/big"
	"testing"
)

var testKey *vpcrypto.Key

const testMaxSecondsShort = 1
const testMaxSecondsLong = 3
const testMinZeroes = 7

func init() {
	testKey, _ = vpcrypto.NewKey()
}

const filterCheckerMod = 31

type filterOnly struct {
	Modulo int64
	T      *testing.T
}

func (fo *filterOnly) Filter(id *big.Int) *big.Int {
	var ret big.Int

	ret.Mod(id, big.NewInt(fo.Modulo))
	ret.Sub(id, &ret)
	fo.T.Logf("filter %s %% %d = %s", id.String(), fo.Modulo, ret.String())

	return &ret
}

func (fo *filterOnly) Check(id *big.Int) bool {
	fo.T.Log("check passthrough")

	return true
}

type checkerOnly struct {
	Modulo int64
	T      *testing.T
}

func (co *checkerOnly) Filter(id *big.Int) *big.Int {
	co.T.Log("filter passthrough")

	return id
}

func (co *checkerOnly) Check(id *big.Int) bool {
	var tmp big.Int
	ret := false

	tmp.Mod(id, big.NewInt(co.Modulo))
	if tmp.Cmp(big.NewInt(0)) == 0 {
		ret = true
	}
	co.T.Logf("check %s %% %d = %s -> %t", id.String(), co.Modulo, tmp.String(), ret)

	return ret
}

type duplicateTransform struct {
}

func (dt *duplicateTransform) Transform(input []byte) []byte {
	ret := make([]byte, len(input)*2)
	copy(ret[0:len(input)], input)
	copy(ret[len(input):len(input)*2], input)
	return ret
}

func TestGenerateID512(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var dt duplicateTransform

	id, sig, z, err := GenerateID512(testKey, nil, nil, 0, 0)
	if err == nil {
		t.Logf("GenerateID512 (short) OK id=%s sig=%s z=%d", vpsum.IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID512")
	}
	id, sig, z, err = GenerateID512(testKey, nil, nil, testMaxSecondsLong, 0)
	if err == nil {
		t.Logf("GenerateID512 (long) OK id=%s sig=%s z=%d", vpsum.IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID512")
	}
	id, sig, z, err = GenerateID512(testKey, nil, nil, testMaxSecondsLong, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID512 (very long) OK id=%s sig=%s z=%d", vpsum.IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID512")
	}
	if z < testMinZeroes {
		t.Errorf("GenerateID512 generated z=%d<%d", z, testMinZeroes)
	}
	id, sig, z, err = GenerateID512(testKey, co, nil, testMaxSecondsShort, 0)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID512 (checked) OK id=%s sig=%s z=%d", vpsum.IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate checked  ID512 id=%s", id.String())
	}
	id, sig, z, err = GenerateID512(testKey, fo, nil, testMaxSecondsShort, 0)
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(id) {
		t.Logf("GenerateID512 (filtered) OK id=%s sig=%s z=%d", vpsum.IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate filtered  ID512 id=%s", id.String())
	}
	id, sig, z, err = GenerateID512(testKey, co, nil, testMaxSecondsShort, 0)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID512 (checked) OK id=%s sig=nil z=%d", vpsum.IntToStr512(id), z)
	} else {
		t.Errorf("unable to generate checked  ID512 id=%s", id.String())
	}
	id, sig, z, err = GenerateID512(testKey, nil, &dt, testMaxSecondsShort, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID512 (transformed) OK id=%s sig=nil z=%d", vpsum.IntToStr512(id), z)
	} else {
		t.Errorf("unable to generate checked  ID512 id=%s", id.String())
	}
	if ZeroesInBuf(vpsum.Checksum512(sig)) < testMinZeroes {
		t.Errorf("GenerateID512 generated z=%d<%d", z, testMinZeroes)
	}
}

func TestGenerateID256(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var dt duplicateTransform

	id, sig, z, err := GenerateID256(testKey, nil, nil, 0, 0)
	if err == nil {
		t.Logf("GenerateID256 (short) OK id=%s sig=%s z=%d", vpsum.IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID256")
	}
	id, sig, z, err = GenerateID256(testKey, nil, nil, testMaxSecondsLong, 0)
	if err == nil {
		t.Logf("GenerateID256 (long) OK id=%s sig=%s z=%d", vpsum.IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID256")
	}
	id, sig, z, err = GenerateID256(testKey, nil, nil, testMaxSecondsLong, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID256 (very long) OK id=%s sig=%s z=%d", vpsum.IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID256")
	}
	if z < testMinZeroes {
		t.Errorf("GenerateID256 generated z=%d<%d", z, testMinZeroes)
	}
	id, sig, z, err = GenerateID256(testKey, co, nil, testMaxSecondsShort, 0)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID256 (checked) OK id=%s sig=%s z=%d", vpsum.IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate checked  ID256 id=%s", id.String())
	}
	id, sig, z, err = GenerateID256(testKey, fo, nil, testMaxSecondsShort, 0)
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(id) {
		t.Logf("GenerateID256 (filtered) OK id=%s sig=%s z=%d", vpsum.IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate filtered ID256 id=%s", id.String())
	}
	id, sig, z, err = GenerateID256(testKey, co, nil, testMaxSecondsShort, 0)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID256 (checked) OK id=%s sig=nil z=%d", vpsum.IntToStr256(id), z)
	} else {
		t.Errorf("unable to generate checked  ID256 id=%s", id.String())
	}
	id, sig, z, err = GenerateID256(testKey, nil, &dt, testMaxSecondsShort, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID256 (transformed) OK id=%s sig=nil z=%d", vpsum.IntToStr256(id), z)
	} else {
		t.Errorf("unable to generate checked  ID256 id=%s", id.String())
	}
	if ZeroesInBuf(vpsum.Checksum256(sig)) < testMinZeroes {
		t.Errorf("GenerateID256 generated z=%d<%d", z, testMinZeroes)
	}
}

func TestGenerateID128(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var dt duplicateTransform

	id, sig, z, err := GenerateID128(testKey, nil, nil, 0, 0)
	if err == nil {
		t.Logf("GenerateID128 (short) OK id=%s sig=%s z=%d", vpsum.IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID128")
	}
	id, sig, z, err = GenerateID128(testKey, nil, nil, testMaxSecondsLong, 0)
	if err == nil {
		t.Logf("GenerateID128 (long) OK id=%s sig=%s z=%d", vpsum.IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID128")
	}
	id, sig, z, err = GenerateID128(testKey, nil, nil, testMaxSecondsLong, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID128 (very long) OK id=%s sig=%s z=%d", vpsum.IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID128")
	}
	if z < testMinZeroes {
		t.Errorf("GenerateID128 generated z=%d<%d", z, testMinZeroes)
	}
	id, sig, z, err = GenerateID128(testKey, co, nil, testMaxSecondsShort, 0)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID128 (checked) OK id=%s sig=%s z=%d", vpsum.IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate checked  ID128 id=%s", id.String())
	}
	id, sig, z, err = GenerateID128(testKey, fo, nil, testMaxSecondsShort, 0)
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(id) {
		t.Logf("GenerateID128 (filtered) OK id=%s sig=%s z=%d", vpsum.IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate filtered ID128 id=%s", id.String())
	}
	id, sig, z, err = GenerateID128(testKey, co, nil, testMaxSecondsShort, 0)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID128 (checked) OK id=%s sig=nil z=%d", vpsum.IntToStr128(id), z)
	} else {
		t.Errorf("unable to generate checked  ID128 id=%s", id.String())
	}
	id, sig, z, err = GenerateID128(testKey, nil, &dt, testMaxSecondsShort, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID128 (transformed) OK id=%s sig=nil z=%d", vpsum.IntToStr128(id), z)
	} else {
		t.Errorf("unable to generate checked  ID128 id=%s", id.String())
	}
	if ZeroesInBuf(vpsum.Checksum128(sig)) < testMinZeroes {
		t.Errorf("GenerateID128 generated z=%d<%d", z, testMinZeroes)
	}
}

func TestGenerateID64(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var dt duplicateTransform
	var bigInt big.Int

	id, sig, z, err := GenerateID64(testKey, nil, nil, 0, 0)
	if err == nil {
		t.Logf("GenerateID64 (short) OK id=%s sig=%s z=%d", vpsum.IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID64")
	}
	id, sig, z, err = GenerateID64(testKey, nil, nil, testMaxSecondsLong, 0)
	if err == nil {
		t.Logf("GenerateID64 (long) OK id=%s sig=%s z=%d", vpsum.IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID64")
	}
	id, sig, z, err = GenerateID64(testKey, nil, nil, testMaxSecondsLong, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID64 (very long) OK id=%s sig=%s z=%d", vpsum.IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID64")
	}
	if z < testMinZeroes {
		t.Errorf("GenerateID64 generated z=%d<%d", z, testMinZeroes)
	}
	id, sig, z, err = GenerateID64(testKey, co, nil, testMaxSecondsShort, 0)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID64 (checked) OK id=%s sig=%s z=%d", vpsum.IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate checked  ID64 id=%d", id)
	}
	id, sig, z, err = GenerateID64(testKey, fo, nil, testMaxSecondsShort, 0)
	bigInt.SetUint64(uint64(id))
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID64 (filtered) OK id=%s sig=%s z=%d", vpsum.IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate filtered ID64 id=%d", id)
	}
	id, sig, z, err = GenerateID64(testKey, co, nil, testMaxSecondsShort, 0)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID64 (checked) OK id=%s sig=nil z=%d", vpsum.IntToStr64(id), z)
	} else {
		t.Errorf("unable to generate checked  ID64 id=%d", id)
	}
	id, sig, z, err = GenerateID64(testKey, nil, &dt, testMaxSecondsShort, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID64 (transformed) OK id=%s sig=nil z=%d", vpsum.IntToStr64(id), z)
	} else {
		t.Errorf("unable to generate checked  ID64 id=%d", id)
	}
	if ZeroesInBuf(vpsum.Checksum64(sig)) < testMinZeroes {
		t.Errorf("GenerateID64 generated z=%d<%d", z, testMinZeroes)
	}
}

func TestGenerateID32(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var dt duplicateTransform
	var bigInt big.Int

	id, sig, z, err := GenerateID32(testKey, nil, nil, 0, 0)
	if err == nil {
		t.Logf("GenerateID32 (short) OK id=%s sig=%s z=%d", vpsum.IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID32")
	}
	id, sig, z, err = GenerateID32(testKey, nil, nil, testMaxSecondsLong, 0)
	if err == nil {
		t.Logf("GenerateID32 (long) OK id=%s sig=%s z=%d", vpsum.IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID32")
	}
	id, sig, z, err = GenerateID32(testKey, nil, nil, testMaxSecondsLong, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID32 (very long) OK id=%s sig=%s z=%d", vpsum.IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate ID32")
	}
	if z < testMinZeroes {
		t.Errorf("GenerateID32 generated z=%d<%d", z, testMinZeroes)
	}
	id, sig, z, err = GenerateID32(testKey, co, nil, testMaxSecondsShort, 0)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID32 (checked) OK id=%s sig=%s z=%d", vpsum.IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate checked ID32 id=%d", id)
	}
	id, sig, z, err = GenerateID32(testKey, fo, nil, testMaxSecondsShort, 0)
	bigInt.SetUint64(uint64(id))
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID32 (filtered) OK id=%s sig=%s z=%d", vpsum.IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to generate filtered ID32 id=%d", id)
	}
	id, sig, z, err = GenerateID32(testKey, co, nil, testMaxSecondsShort, 0)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID32 (checked) OK id=%s sig=nil z=%d", vpsum.IntToStr32(id), z)
	} else {
		t.Errorf("unable to generate checked  ID32 id=%d", id)
	}
	id, sig, z, err = GenerateID32(testKey, nil, &dt, testMaxSecondsShort, testMinZeroes)
	if err == nil {
		t.Logf("GenerateID32 (transformed) OK id=%s sig=nil z=%d", vpsum.IntToStr32(id), z)
	} else {
		t.Errorf("unable to generate checked  ID32 id=%d", id)
	}
	if ZeroesInBuf(vpsum.Checksum32(sig)) < testMinZeroes {
		t.Errorf("GenerateID32 generated z=%d<%d", z, testMinZeroes)
	}
}
