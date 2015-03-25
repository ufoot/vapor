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

package vpcrypto

import (
	"encoding/hex"
	"testing"
)

var testKey *Key

func init() {
	testKey, _ = NewKey()
}

func TestGenerateID512(t *testing.T) {
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
}

func TestGenerateID256(t *testing.T) {
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
}

func TestGenerateID128(t *testing.T) {
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
}

func TestGenerateID64(t *testing.T) {
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
}

func TestGenerateID32(t *testing.T) {
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
}
