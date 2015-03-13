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
	"testing"
)

func TestExportImport(t *testing.T) {
	var key1 *Key
	var buf1 []byte
	var len1 int
	var key2 *Key
	var buf2 []byte
	var len2 int
	var err error

	key1, err = NewKey()
	if err != nil {
		t.Error(err)
	}
	buf1, err = key1.ExportPub()
	if err != nil {
		t.Error(err)
	}
	len1 = len(buf1)
	t.Logf("len of exported key 1 is %d", len1)
	key2, err = ImportPubKey(buf1)
	if err != nil {
		t.Error(err)
	}
	buf2, err = key2.ExportPub()
	if err != nil {
		t.Error(err)
	}
	len2 = len(buf2)
	t.Logf("len of exported key 2 is %d", len2)
	if len1 == 0 || len2 == 0 {
		t.Errorf("zero len for exported key len1=%d len2=%d", len1, len2)
	}
	if len1 != len2 {
		t.Errorf("len for exported keys differ len1=%d len2=%d", len1, len2)
	}
}

func TestSig(t *testing.T) {
	var key1 *Key
	var buf []byte
	var key2 *Key
	var err error
	var sig []byte
	var content []byte
	const content_str string = "foo bar"
	var ok bool

	key1, err = NewKey()
	if err != nil {
		t.Error(err)
	}
	buf, err = key1.ExportPub()
	if err != nil {
		t.Error(err)
	}
	key2, err = ImportPubKey(buf)
	if err != nil {
		t.Error(err)
	}

	content = []byte(content_str)
	sig, err = key1.Sign(content)
	if err != nil {
		t.Error(err)
	}
	t.Logf("len of sig is %d", len(sig))
	ok, err = key2.CheckSig(content, sig)
	if err != nil {
		t.Error(err)
	}
	if ok == true {
		t.Log("signature is correct")
	} else {
		t.Error("bad signature")
	}
}

func TestEnc(t *testing.T) {
	var key1 *Key
	var buf []byte
	var key2 *Key
	var err error
	var encrypted []byte
	var decrypted []byte
	var decrypted_str string
	var content []byte
	const content_str string = "foo bar"

	key1, err = NewKey()
	if err != nil {
		t.Error(err)
	}
	buf, err = key1.ExportPub()
	if err != nil {
		t.Error(err)
	}
	key2, err = ImportPubKey(buf)
	if err != nil {
		t.Error(err)
	}

	content = []byte(content_str)
	encrypted, err = key2.Encrypt(content)
	if err != nil {
		t.Error(err)
	}
	t.Logf("len of encrypted is %d", len(encrypted))
	decrypted, err = key1.Decrypt(encrypted)
	if err != nil {
		t.Error(err)
	}
	decrypted_str = string(decrypted)
	if content_str == decrypted_str {
		t.Log("decrypted message is same as source")
	} else {
		t.Errorf("encryption/decryption problem, results differ \"%s\" vs \"%s\"", content_str, decrypted_str)
	}
}
