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
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/ufoot/vapor/vpsys"
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

	vpsys.LogNotice("SymEncrypt BEGIN")
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
		return nil, vpsys.ErrorChain(err, "unable to encrypt content")
	}
	gzipOutput = gzip.NewWriter(output)
	gzipOutput.Write(content)
	gzipOutput.Flush()
	gzipOutput.Close()
	output.Close()
	vpsys.LogNotice("SymEncrypt END")

	ret = byteWriter.Bytes()

	vpsys.LogNoticef("SymEncrypt REAL END %d", len(ret))

	return ret, nil
}

type proxyReader struct {
	reader io.Reader
}

func (proxy proxyReader) Read(p []byte) (int, error) {
	vpsys.LogNoticef("Read BEGIN %d", len(p))
	if len(p) >= 0 {
		panic("Want stack trace")
	}
	n, err := proxy.Read(p)
	vpsys.LogNoticef("Read END %d,%s", n, err)

	return n, err
}

func symDecrypt(content, password []byte) ([]byte, error) {
	var err error
	var byteReader io.Reader
	var proxy proxyReader
	var messageDetails *openpgp.MessageDetails
	var gzipReader *gzip.Reader
	var ret []byte

	vpsys.LogNoticef("symDecrypt BEGIN %d", len(content))

	// This is not very go-ish but when passphrase is wrong,
	// for instance, ReadMessage fails with nil pointers or
	// other low level errors, we just trap those.
	/*
		err = errors.New("Unable to decrypt")
		defer func() {
			if rec := recover(); rec != nil {
				vpsys.LogWarningf("symDecrypt error decrypting %d bytes",len(content))
			}
		}()
	*/

	vpsys.LogNotice("symDecrypt A")
	byteReader = bytes.NewReader(content)
	proxy.reader = byteReader
	messageDetails, err = openpgp.ReadMessage(proxy, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		return password, nil
	}, nil)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to read PGP content")
	}

	vpsys.LogNotice("symDecrypt B")
	if !messageDetails.IsEncrypted {
		return nil, errors.New("PGP content was not encrypted")
	}
	vpsys.LogNotice("symDecrypt C")
	gzipReader, err = gzip.NewReader(messageDetails.UnverifiedBody)
	defer gzipReader.Close()
	vpsys.LogNotice("symDecrypt D")
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to open GZIP within PGP encrypted content")
	}
	vpsys.LogNotice("symDecrypt E")
	ret, err = ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to read GZIP within PGP encrypted content")
	}
	vpsys.LogNotice("symDecrypt F")
	if len(ret) <= 0 {
		return nil, vpsys.ErrorChain(err, "no data in GZIP within PGP encrypted content")
	}
	vpsys.LogNotice("symDecrypt END")

	return ret, err
}

// SymDecrypt decrypts a message crypted using a symmetric password/key.
func SymDecrypt(content, password []byte) ([]byte, error) {
	ret, err := symDecrypt(content, password)
	if ret == nil || len(ret) == 0 {
		return nil, errors.New("unable to decrypt")
	}
	return ret, err
}
