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
	"encoding/hex"
	"testing"
)

var testKey *Key

func init() {
	testKey, _ = NewKey()
}

func TestId512(t *testing.T) {
	id, sig, z, err := Id512(testKey, nil, 0)
	if err == nil {
		t.Logf("Id512 (short) OK id=%s sig=%s z=%d", IntToString512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id512")
	}
	id, sig, z, err = Id512(testKey, nil, 3)
	if err == nil {
		t.Logf("Id512 (long) OK id=%s sig=%s z=%d", IntToString512(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id512")
	}
}

func TestId256(t *testing.T) {
	id, sig, z, err := Id256(testKey, nil, 0)
	if err == nil {
		t.Logf("Id256 (short) OK id=%s sig=%s z=%d", IntToString256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id256")
	}
	id, sig, z, err = Id256(testKey, nil, 3)
	if err == nil {
		t.Logf("Id256 (long) OK id=%s sig=%s z=%d", IntToString256(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id256")
	}
}

func TestId128(t *testing.T) {
	id, sig, z, err := Id128(testKey, nil, 0)
	if err == nil {
		t.Logf("Id128 (short) OK id=%s sig=%s z=%d", IntToString128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id128")
	}
	id, sig, z, err = Id128(testKey, nil, 3)
	if err == nil {
		t.Logf("Id128 (long) OK id=%s sig=%s z=%d", IntToString128(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id128")
	}
}

func TestId64(t *testing.T) {
	id, sig, z, err := Id64(testKey, nil, 0)
	if err == nil {
		t.Logf("Id64 (short) OK id=%s sig=%s z=%d", IntToString64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id64")
	}
	id, sig, z, err = Id64(testKey, nil, 3)
	if err == nil {
		t.Logf("Id64 (long) OK id=%s sig=%s z=%d", IntToString64(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id64")
	}
}

func TestId32(t *testing.T) {
	id, sig, z, err := Id32(testKey, nil, 0)
	if err == nil {
		t.Logf("Id32 (short) OK id=%s sig=%s z=%d", IntToString32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id32")
	}
	id, sig, z, err = Id32(testKey, nil, 3)
	if err == nil {
		t.Logf("Id32 (long) OK id=%s sig=%s z=%d", IntToString32(id), hex.EncodeToString(sig), z)
	} else {
		t.Error("unable to generate Id32")
	}
}
