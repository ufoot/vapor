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

package vpp2p

import (
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpp2papi"
	"github.com/ufoot/vapor/vpp2pdat"
	"github.com/ufoot/vapor/vprand"
	"github.com/ufoot/vapor/vpsum"
)

// Host is a physical host, it is used to uniquely identify
// a host, it can be used to handle several apps or rings.
type Host struct {
	// Info about the host
	Info vpp2papi.HostInfo

	key        *vpcrypto.Key
	localNodes []Node
}

// NewHost returns a new host object
func NewHost(title, url string, useSig bool) (*Host, error) {
	var ret Host
	var pubKey []byte
	var sig []byte

	ok, err := vpp2pdat.CheckTitle(title)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = vpp2pdat.CheckURL(url)
	if err != nil || !ok {
		return nil, err
	}

	if useSig {
		ret.key, err = vpcrypto.NewKey()
		if err != nil {
			return nil, err
		}
		pubKey, err = ret.key.ExportPub()
		if err != nil {
			return nil, err
		}
		ok, err = vpp2pdat.CheckPubKey(pubKey)
		if err != nil || !ok {
			return nil, err
		}
		sig, err = ret.key.Sign(vpp2pdat.SigBytesHost(title, url))
		if err != nil {
			return nil, err
		}
		ok, err = vpp2pdat.CheckSig(pubKey)
		if err != nil || !ok {
			return nil, err
		}
	} else {
		pubKey = vpsum.IntToBuf512(vprand.Rand512(vprand.NewRand(), nil))
		sig = []byte("")
	}

	ret.Info = vpp2papi.HostInfo{HostTitle: title, HostURL: url, HostPubKey: pubKey, HostSig: sig}
	ret.localNodes = make([]Node, 0)

	return &ret, nil
}

// CanSign returns true if the host has a key it can sign with.
func (host *Host) CanSign() bool {
	return host.key != nil
}

// IsSigned returns true if the has been self-signed.
// It does not check if the signature is valid.
func (host *Host) IsSigned() bool {
	return vpp2pdat.HostInfoIsSigned(&host.Info)
}

// CheckSig checks if the host signature is OK, if it's not, returns false and an error.
func (host *Host) CheckSig() (bool, error) {
	return vpp2pdat.HostInfoCheckSig(&host.Info)
}
