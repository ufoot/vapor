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
	"bytes"
	"compress/gzip"
	"errors"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"golang.org/x/crypto/ripemd160"
	"io"
	"io/ioutil"
	"time"
	"ufoot.org/vapor/vpsys"
)

type Key struct {
	entity *openpgp.Entity
}

func init() {
	_ = ripemd160.New()
}

func NewKey() (*Key, error) {
	var entity *openpgp.Entity
	var err error
	var key_pair Key
	var byte_writer bytes.Buffer

	entity, err = openpgp.NewEntity("", "", "", nil)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to create a new OpenPGP entity")
	}

	err = entity.SerializePrivate(&byte_writer, nil)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to serialize private key")
	}

	key_pair.entity = entity

	return &key_pair, nil
}

func (key_pair Key) ExportPub() ([]byte, error) {
	var byte_writer bytes.Buffer
	var err error
	var pub_key []byte

	err = key_pair.entity.Serialize(&byte_writer)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to serialize public key")
	}

	pub_key = byte_writer.Bytes()

	return pub_key, nil
}

func ImportPubKey(key []byte) (*Key, error) {
	var byte_reader io.Reader
	var packet_reader *packet.Reader
	var err error
	var pub_key Key
	var entity *openpgp.Entity

	byte_reader = bytes.NewReader(key)
	packet_reader = packet.NewReader(byte_reader)
	entity, err = openpgp.ReadEntity(packet_reader)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to read entity from public key")
	}

	pub_key.entity = entity

	return &pub_key, nil
}

func (key Key) Sign(content []byte) ([]byte, error) {
	var byte_writer bytes.Buffer
	var byte_reader io.Reader
	var sig []byte
	var err error

	byte_reader = bytes.NewReader(content)
	err = openpgp.DetachSign(&byte_writer, key.entity, byte_reader, nil)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to sign content")
	}

	sig = byte_writer.Bytes()

	return sig, nil
}

func (key Key) CheckSig(content []byte, sig []byte) (bool, error) {
	var key_ring openpgp.EntityList
	var err error
	var byte_reader_content io.Reader
	var byte_reader_sig io.Reader

	key_ring = make([]*openpgp.Entity, 1)
	key_ring[0] = key.entity

	byte_reader_content = bytes.NewReader(content)
	byte_reader_sig = bytes.NewReader(sig)
	_, err = openpgp.CheckDetachedSignature(key_ring, byte_reader_content, byte_reader_sig)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (key Key) Encrypt(content []byte) ([]byte, error) {
	var key_ring openpgp.EntityList
	var byte_writer bytes.Buffer
	var err error
	var output io.WriteCloser
	var gzip_output *gzip.Writer
	var ret []byte
	var hints openpgp.FileHints

	if len(content) <= 0 {
		return nil, errors.New("no data")
	}

	hints.IsBinary = true
	hints.ModTime = time.Now()

	key_ring = make([]*openpgp.Entity, 1)
	key_ring[0] = key.entity

	output, err = openpgp.Encrypt(&byte_writer, key_ring, nil, &hints, nil)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to encrypt content")
	}
	gzip_output = gzip.NewWriter(output)
	gzip_output.Write(content)
	gzip_output.Close()
	output.Close()

	ret = byte_writer.Bytes()

	return ret, nil
}

func (key Key) Decrypt(content []byte) ([]byte, error) {
	var key_ring openpgp.EntityList
	var err error
	var byte_reader io.Reader
	var message_details *openpgp.MessageDetails
	var gzip_reader *gzip.Reader
	var ret []byte

	key_ring = make([]*openpgp.Entity, 1)
	key_ring[0] = key.entity

	byte_reader = bytes.NewReader(content)
	message_details, err = openpgp.ReadMessage(byte_reader, key_ring, nil, nil)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to read PGP content")
	}

	if !message_details.IsEncrypted {
		return nil, errors.New("PGP content was not encrypted")
	}
	gzip_reader, err = gzip.NewReader(message_details.UnverifiedBody)
	defer gzip_reader.Close()
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to open GZIP within PGP encrypted content")
	}
	ret, err = ioutil.ReadAll(gzip_reader)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to read GZIP within PGP encrypted content")
	}
	if len(ret) <= 0 {
		return nil, vpsys.ErrorChain(err, "no data in GZIP within PGP encrypted content")
	}
	return ret, nil
}
