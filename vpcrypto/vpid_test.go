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

package vpcrypto

import (
	"encoding/hex"
	"math/big"
	"testing"
)

var testKey *Key

func init() {
	testKey, _ = NewKey()
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

func TestGenerateID512(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}

	id, sig, z, err := GenerateID512(testKey, nil, 0)
	if err == nil {
		t.Logf("GenerateID512 (short) OK id=%s sig=%s z=%d", IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID512")
	}
	id, sig, z, err = GenerateID512(testKey, nil, 3)
	if err == nil {
		t.Logf("GenerateID512 (long) OK id=%s sig=%s z=%d", IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID512")
	}
	id, sig, z, err = GenerateID512(testKey, co, 1)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID512 (checked) OK id=%s sig=%s z=%d", IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate checked  ID512 id=%s", id.String())
	}
	id, sig, z, err = GenerateID512(testKey, fo, 1)
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(id) {
		t.Logf("GenerateID512 (filtered) OK id=%s sig=%s z=%d", IntToStr512(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate filtered  ID512 id=%s", id.String())
	}
	id, sig, z, err = GenerateID512(nil, co, 1)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID512 (checked) OK id=%s sig=nil z=%d", IntToStr512(id), z)
	} else {
		t.Errorf("unable to gererate checked  ID512 id=%s", id.String())
	}
}

func TestGenerateID256(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}

	id, sig, z, err := GenerateID256(testKey, nil, 0)
	if err == nil {
		t.Logf("GenerateID256 (short) OK id=%s sig=%s z=%d", IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID256")
	}
	id, sig, z, err = GenerateID256(testKey, nil, 3)
	if err == nil {
		t.Logf("GenerateID256 (long) OK id=%s sig=%s z=%d", IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID256")
	}
	id, sig, z, err = GenerateID256(testKey, co, 1)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID256 (checked) OK id=%s sig=%s z=%d", IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate checked  ID256 id=%s", id.String())
	}
	id, sig, z, err = GenerateID256(testKey, fo, 1)
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(id) {
		t.Logf("GenerateID256 (filtered) OK id=%s sig=%s z=%d", IntToStr256(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate filtered ID256 id=%s", id.String())
	}
	id, sig, z, err = GenerateID256(nil, co, 1)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID256 (checked) OK id=%s sig=nil z=%d", IntToStr256(id), z)
	} else {
		t.Errorf("unable to gererate checked  ID256 id=%s", id.String())
	}
}

func TestGenerateID128(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}

	id, sig, z, err := GenerateID128(testKey, nil, 0)
	if err == nil {
		t.Logf("GenerateID128 (short) OK id=%s sig=%s z=%d", IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID128")
	}
	id, sig, z, err = GenerateID128(testKey, nil, 3)
	if err == nil {
		t.Logf("GenerateID128 (long) OK id=%s sig=%s z=%d", IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID128")
	}
	id, sig, z, err = GenerateID128(testKey, co, 1)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID128 (checked) OK id=%s sig=%s z=%d", IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate checked  ID128 id=%s", id.String())
	}
	id, sig, z, err = GenerateID128(testKey, fo, 1)
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(id) {
		t.Logf("GenerateID128 (filtered) OK id=%s sig=%s z=%d", IntToStr128(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate filtered ID128 id=%s", id.String())
	}
	id, sig, z, err = GenerateID128(nil, co, 1)
	if err == nil && co.Check(id) {
		t.Logf("GenerateID128 (checked) OK id=%s sig=nil z=%d", IntToStr128(id), z)
	} else {
		t.Errorf("unable to gererate checked  ID128 id=%s", id.String())
	}
}

func TestGenerateID64(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var bigInt big.Int

	id, sig, z, err := GenerateID64(testKey, nil, 0)
	if err == nil {
		t.Logf("GenerateID64 (short) OK id=%s sig=%s z=%d", IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID64")
	}
	id, sig, z, err = GenerateID64(testKey, nil, 3)
	if err == nil {
		t.Logf("GenerateID64 (long) OK id=%s sig=%s z=%d", IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID64")
	}
	id, sig, z, err = GenerateID64(testKey, co, 1)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID64 (checked) OK id=%s sig=%s z=%d", IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate checked  ID64 id=%d", id)
	}
	id, sig, z, err = GenerateID64(testKey, fo, 1)
	bigInt.SetUint64(uint64(id))
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID64 (filtered) OK id=%s sig=%s z=%d", IntToStr64(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate filtered ID64 id=%d", id)
	}
	id, sig, z, err = GenerateID64(nil, co, 1)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID64 (checked) OK id=%s sig=nil z=%d", IntToStr64(id), z)
	} else {
		t.Errorf("unable to gererate checked  ID64 id=%d", id)
	}
}

func TestGenerateID32(t *testing.T) {
	fo := &filterOnly{filterCheckerMod, t}
	co := &checkerOnly{filterCheckerMod, t}
	var bigInt big.Int

	id, sig, z, err := GenerateID32(testKey, nil, 0)
	if err == nil {
		t.Logf("GenerateID32 (short) OK id=%s sig=%s z=%d", IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID32")
	}
	id, sig, z, err = GenerateID32(testKey, nil, 3)
	if err == nil {
		t.Logf("GenerateID32 (long) OK id=%s sig=%s z=%d", IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to gererate ID32")
	}
	id, sig, z, err = GenerateID32(testKey, co, 1)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID32 (checked) OK id=%s sig=%s z=%d", IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate checked ID32 id=%d", id)
	}
	id, sig, z, err = GenerateID32(testKey, fo, 1)
	bigInt.SetUint64(uint64(id))
	// note, checking with co what has been filter with fo
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID32 (filtered) OK id=%s sig=%s z=%d", IntToStr32(id), hex.EncodeToString(sig), z)
	} else {
		t.Errorf("unable to gererate filtered ID32 id=%d", id)
	}
	id, sig, z, err = GenerateID32(nil, co, 1)
	bigInt.SetUint64(uint64(id))
	if err == nil && co.Check(&bigInt) {
		t.Logf("GenerateID32 (checked) OK id=%s sig=nil z=%d", IntToStr32(id), z)
	} else {
		t.Errorf("unable to gererate checked  ID32 id=%d", id)
	}
}
