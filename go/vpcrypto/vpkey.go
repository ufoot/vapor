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
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/ufoot/vapor/go/vperror"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"golang.org/x/crypto/ripemd160"
	"io"
	"io/ioutil"
	"time"
)

// Key contains a cryptographic key pair. It might only
// contain a public key if it has been imported from a public
// key export. In that case it can not be used for signing
// and decrypting messages.
type Key struct {
	entity *openpgp.Entity
}

func init() {
	_ = ripemd160.New()
}

// NewKey returns a new pair of cryptographic keys.
// Note that this function is rather slow, it can take up to one second
// or more, even on a powerfull computer, so don't generate keys often.
func NewKey() (*Key, error) {
	var entity *openpgp.Entity
	var err error
	var key Key
	var byteWriter bytes.Buffer

	// NewEntity params: name, comment, email
	entity, err = openpgp.NewEntity(PackageName, PackageURL, PackageEmail, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to create a new OpenPGP entity")
	}

	err = entity.SerializePrivate(&byteWriter, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to serialize private key")
	}

	key.entity = entity

	return &key, nil
}

// ExportPub exports the public key of a key pair.
func (key Key) ExportPub() ([]byte, error) {
	var byteWriter bytes.Buffer
	var err error
	var pubKey []byte

	err = key.entity.Serialize(&byteWriter)
	if err != nil {
		return nil, vperror.Chain(err, "unable to serialize public key")
	}

	pubKey = byteWriter.Bytes()

	return pubKey, nil
}

// ImportPubKey creates a key from an exported public key.
func ImportPubKey(key []byte) (*Key, error) {
	var byteReader io.Reader
	var packetReader *packet.Reader
	var err error
	var pubKey Key
	var entity *openpgp.Entity

	byteReader = bytes.NewReader(key)
	packetReader = packet.NewReader(byteReader)
	entity, err = openpgp.ReadEntity(packetReader)
	if err != nil {
		return nil, vperror.Chain(err, "unable to read entity from public key")
	}

	pubKey.entity = entity

	return &pubKey, nil
}

// Sign signs a content with a key.
// Note that the key must contain a private key, it is not possible
// to sign with a public key.
func (key Key) Sign(content []byte) ([]byte, error) {
	var byteWriter bytes.Buffer
	var byteReader io.Reader
	var sig []byte
	var err error

	byteReader = bytes.NewReader(content)
	err = openpgp.DetachSign(&byteWriter, key.entity, byteReader, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to sign content")
	}

	sig = byteWriter.Bytes()

	return sig, nil
}

// CheckSig checks a signature.
// This can be done with a public key, even if private key is not available.
func (key Key) CheckSig(content, sig []byte) (bool, error) {
	var keyRing openpgp.EntityList
	var err error
	var byteReaderContent io.Reader
	var byteReaderSig io.Reader

	keyRing = make([]*openpgp.Entity, 1)
	keyRing[0] = key.entity

	byteReaderContent = bytes.NewReader(content)
	byteReaderSig = bytes.NewReader(sig)
	_, err = openpgp.CheckDetachedSignature(keyRing, byteReaderContent, byteReaderSig)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Encrypt encrypts a message.
// This can be done with a public key, even if private key is not available.
func (key Key) Encrypt(content []byte) ([]byte, error) {
	var keyRing openpgp.EntityList
	var byteWriter bytes.Buffer
	var err error
	var output io.WriteCloser
	var gzipOutput *gzip.Writer
	var ret []byte
	var hints openpgp.FileHints

	if len(content) <= 0 {
		return nil, errors.New("no data")
	}

	hints.IsBinary = true
	hints.ModTime = time.Now()

	keyRing = make([]*openpgp.Entity, 1)
	keyRing[0] = key.entity

	output, err = openpgp.Encrypt(&byteWriter, keyRing, nil, &hints, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to encrypt content")
	}
	gzipOutput = gzip.NewWriter(output)
	_,err=gzipOutput.Write(content)
	if err!=nil {
		return nil,vperror.Chain(err, "unable to write gzip")
	}
	err=gzipOutput.Close()
	if err!=nil {
		return nil,vperror.Chain(err, "unablen to close gzip")
	}
	err=output.Close()
	if err!=nil {
		return nil,vperror.Chain(err, "unablen to close output")
	}

	ret = byteWriter.Bytes()

	return ret, nil
}

func (key Key) decrypt(content []byte) ([]byte, error) {
	var keyRing openpgp.EntityList
	var err error
	var byteReader io.Reader
	var messageDetails *openpgp.MessageDetails
	var gzipReader *gzip.Reader
	var ret []byte

	// This is not very go-ish but when passphrase is wrong,
	// for instance, ReadMessage fails with nil pointers or
	// other low level errors, we just trap those.
	defer func() {
		if rec := recover(); rec != nil {
			ret = nil
			err = errors.New("Unable to decrypt")
		}
	}()

	keyRing = make([]*openpgp.Entity, 1)
	keyRing[0] = key.entity

	byteReader = bytes.NewReader(content)
	messageDetails, err = openpgp.ReadMessage(byteReader, keyRing, nil, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to read PGP content")
	}

	if !messageDetails.IsEncrypted {
		return nil, errors.New("PGP content was not encrypted")
	}
	gzipReader, err = gzip.NewReader(messageDetails.UnverifiedBody)
	defer func() {
		if gzipReader!=nil {
		_=gzipReader.Close()
		}
	} ()
	if err != nil {
		return nil, vperror.Chain(err, "unable to open GZIP within PGP encrypted content")
	}
	ret, err = ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, vperror.Chain(err, "unable to read GZIP within PGP encrypted content")
	}
	if len(ret) <= 0 {
		return nil, vperror.Chain(err, "no data in GZIP within PGP encrypted content")
	}
	return ret, nil
}

// Decrypt decrypts a message.
// Note that the key must contain a private key, it is not possible
// to sign with a public key.
func (key Key) Decrypt(content []byte) ([]byte, error) {
	ret, err := key.decrypt(content)
	if ret == nil || len(ret) == 0 {
		return nil, errors.New("unable to decrypt")
	}
	return ret, err
}
