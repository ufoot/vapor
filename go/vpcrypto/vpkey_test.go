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
	"github.com/ufoot/vapor/go/vperror"
	"testing"
)

var benchKey *Key
var benchContent []byte
var benchSig []byte
var benchCrypted []byte

func init() {
	benchKey, _ = NewKey()
	benchContent = make([]byte, 1500)
	benchSig, _ = benchKey.Sign(benchContent)
	benchCrypted, _ = benchKey.Encrypt(benchContent)
}

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
	const contentStr string = "foo bar"
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

	content = []byte(contentStr)
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
	var key1, key2, key3 *Key
	var buf []byte
	var err error
	var encrypted []byte
	var decrypted []byte
	var decryptedStr string
	var content []byte
	const contentStr string = "foo bar"

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

	content = []byte(contentStr)
	encrypted, err = key2.Encrypt(content)
	if err != nil {
		t.Error(err)
	}
	t.Logf("len of encrypted is %d", len(encrypted))
	decrypted, err = key1.Decrypt(encrypted)
	if err != nil {
		t.Error(err)
	}
	decryptedStr = string(decrypted)
	if contentStr == decryptedStr {
		t.Log("decrypted message is same as source")
	} else {
		t.Errorf("encryption/decryption problem, results differ \"%s\" vs \"%s\"", contentStr, decryptedStr)
	}

	key3, err = NewKey()
	if err != nil {
		t.Error(err)
	}
	_, err = key3.Decrypt(encrypted)
	if err != nil {
		t.Log("OK, decrypt is impossible with a bad key")
	} else {
		t.Error("decrypt is possible with a bad key, this *should* be impossible")
	}
}

func BenchmarkNewKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewKey()
	}
}

func BenchmarkSign(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = benchKey.Sign(benchContent)
		if err != nil {
			b.Error(vperror.Chain(err, "unable to sign"))
		}
	}
}

func BenchmarkCheckSig(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = benchKey.CheckSig(benchContent, benchSig)
		if err != nil {
			b.Error(vperror.Chain(err, "unable to check sig"))
		}
	}
}

func BenchmarkEncrypt(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = benchKey.Encrypt(benchContent)
		if err != nil {
			b.Error(vperror.Chain(err, "unable to encrypt"))
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = benchKey.Decrypt(benchCrypted)
		if err != nil {
			b.Error(vperror.Chain(err, "unable to decrypt"))
		}
	}
}
