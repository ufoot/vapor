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
	"github.com/ufoot/vapor/go/vplog"
	"golang.org/x/crypto/openpgp"
	"io"
	"io/ioutil"
	"time"
)

// SymEncrypt encrypts a message using a symmetric password/key.
func SymEncrypt(content, password []byte) ([]byte, error) {
	var byteWriter bytes.Buffer
	var err error
	var output io.WriteCloser
	var gzipOutput *gzip.Writer
	var ret []byte
	var hints openpgp.FileHints

	if len(content) <= 0 {
		return nil, errors.New("no data")
	}
	if len(password) < 16 {
		return nil, errors.New("password too weak, need at laast 128 bits (16 bytes)")
	}

	hints.IsBinary = true
	hints.ModTime = time.Now()

	output, err = openpgp.SymmetricallyEncrypt(&byteWriter, password, &hints, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to encrypt content")
	}
	gzipOutput = gzip.NewWriter(output)
	_, err = gzipOutput.Write(content)
	if err != nil {
		return nil, vperror.Chain(err, "unable to write gzip")
	}
	err = gzipOutput.Flush()
	if err != nil {
		return nil, vperror.Chain(err, "unable to flush gzip")
	}
	err = gzipOutput.Close()
	if err != nil {
		return nil, vperror.Chain(err, "unable to close gzip")
	}
	err = output.Close()
	if err != nil {
		return nil, vperror.Chain(err, "unable to close output")
	}

	ret = byteWriter.Bytes()

	return ret, nil
}

func symDecrypt(content, password []byte) ([]byte, error) {
	var err error
	var byteReader io.Reader
	var messageDetails *openpgp.MessageDetails
	var gzipReader *gzip.Reader
	var ret []byte
	var promptCalled = false

	// This is not very go-ish but when passphrase is wrong,
	// for instance, ReadMessage fails with nil pointers or
	// other low level errors, we just trap those.
	err = errors.New("Unable to decrypt")
	defer func() {
		if rec := recover(); rec != nil {
			vplog.LogDebugf("symDecrypt error decrypting %d bytes (catched)", len(content))
		}
	}()

	byteReader = bytes.NewReader(content)
	messageDetails, err = openpgp.ReadMessage(byteReader, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if promptCalled {
			return nil, errors.New("bad password")
		}
		promptCalled = true
		return password, nil
	}, nil)
	if err != nil {
		return nil, vperror.Chain(err, "unable to read PGP content")
	}

	if !messageDetails.IsEncrypted {
		return nil, errors.New("PGP content was not encrypted")
	}
	gzipReader, err = gzip.NewReader(messageDetails.UnverifiedBody)
	defer func() {
		if gzipReader != nil {
			_ = gzipReader.Close()
		}
	}()
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

	return ret, err
}

// SymDecrypt decrypts a message crypted using a symmetric password/key.
func SymDecrypt(content, password []byte) ([]byte, error) {
	ret, err := symDecrypt(content, password)
	if err != nil {
		return nil, vperror.Chain(err, "unable to decrypt")
	}
	if ret == nil {
		return nil, errors.New("unable to decrypt, no content")
	}
	if len(ret) == 0 {
		return nil, errors.New("unable to decrypt, zero length")
	}
	return ret, err
}
