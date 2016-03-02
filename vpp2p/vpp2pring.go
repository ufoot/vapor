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
	"encoding/base64"
	"fmt"
	"github.com/ufoot/vapor/vpapp"
	"github.com/ufoot/vapor/vpbruijn"
	"github.com/ufoot/vapor/vpid"
	"github.com/ufoot/vapor/vpp2papi"
	"github.com/ufoot/vapor/vpp2pdat"
	"github.com/ufoot/vapor/vpsum"
	"math/big"
)

const (
	// RingKeySeconds specifies how many seconds should we spend on creating
	// ring keys in signed mode.
	RingKeySeconds = 3
	// RingKeyZeroes specifies how many zeroes there should be at the end of
	// a ring key sig in signed mode.
	RingKeyZeroes = 12
)

// RingSecret stores the secret data of a Ring.
type RingSecret struct {
	// Password hash
	PasswordHash []byte
}

// Ring is a community, a network of related nodes, which communicate
// through hosts. It is the logical artefact used to relate several hosts/nodes
// together.
type Ring struct {
	// Info about the ring
	Info   vpp2papi.RingInfo
	secret RingSecret

	walker     vpbruijn.BruijnWalker
	localNodes []Node
}

type ringStuffAppender struct {
	ringTitle       string
	ringDescription string
	appID           []byte
	config          *vpp2papi.RingConfig
	hasPassword     bool
}

func (r *ringStuffAppender) Transform(ringID []byte) []byte {
	return vpp2pdat.SigBytesRing(ringID, r.ringTitle, r.ringDescription, r.appID, r.config, r.hasPassword)
}

// NewRing creates a new ring from static data.
func NewRing(host *Host, ringTitle, ringDescription string, appID []byte, config *vpp2papi.RingConfig, fc vpid.FilterChecker, passwordHash []byte) (*Ring, error) {
	var ret Ring
	var ok bool
	var err error
	var intRingID *big.Int
	var sig []byte

	ok, err = vpp2pdat.CheckTitle(ringTitle)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = vpp2pdat.CheckDescription(ringDescription)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = vpp2pdat.CheckID(appID)
	if err != nil || !ok {
		return nil, err
	}
	if config == nil {
		config = vpp2pdat.DefaultRingConfig()
	}
	ret.walker, err = vpbruijn.BruijnNew(int(config.BruijnM), int(config.BruijnN))
	if err != nil {
		return nil, err
	}
	ok, err = vpp2pdat.CheckRingConfig(config)
	if err != nil {
		return nil, err
	}
	if err != nil || !ok {
		return nil, err
	}

	hasPassword := (passwordHash != nil) && (len(passwordHash) > 0)
	rsa := ringStuffAppender{ringTitle: ringTitle, ringDescription: ringDescription, appID: appID, config: config, hasPassword: hasPassword}
	if host.CanSign() {
		intRingID, sig, _, err = vpid.GenerateID512(host.key, fc, &rsa, RingKeySeconds, RingKeyZeroes)
		if err != nil {
			return nil, err
		}
	} else {
		intRingID, _, _, err = vpid.GenerateID512(nil, fc, &rsa, RingKeySeconds, RingKeyZeroes)
		if err != nil {
			return nil, err
		}
		sig = []byte("")
	}

	ret.Info.RingID = vpsum.IntToBuf512(intRingID)
	ret.Info.RingTitle = ringTitle
	ret.Info.RingDescription = ringDescription
	ret.Info.AppID = appID
	ret.Info.Config = vpp2pdat.DefaultRingConfig()
	*ret.Info.Config = *config
	ret.Info.HasPassword = hasPassword
	ret.secret.PasswordHash = passwordHash
	ret.Info.HostPubKey = make([]byte, len(host.Info.HostPubKey))
	copy(ret.Info.HostPubKey, host.Info.HostPubKey)
	ret.Info.RingSig = sig

	ret.localNodes = make([]Node, 0)

	return &ret, nil
}

// RingFromInfo creates a new Ring object from its info static data, typically
// retrieved from the network, on an application directory.
func RingFromInfo(ringInfo *vpp2papi.RingInfo, passwordHash []byte) (*Ring, error) {
	var ret Ring
	var err error

	ret.Info = *ringInfo
	ret.Info.Config = vpp2pdat.DefaultRingConfig()
	*(ret.Info.Config) = *(ringInfo.Config)

	if ret.Info.HasPassword {
		if passwordHash == nil || len(passwordHash) == 0 {
			return nil, fmt.Errorf("Ring marked as having a password, but none supplied")
		}
		ret.secret.PasswordHash = make([]byte, len(passwordHash))
		copy(ret.secret.PasswordHash, passwordHash)
	}
	ret.walker, err = vpbruijn.BruijnNew(int(ret.Info.Config.BruijnM), int(ret.Info.Config.BruijnN))
	if err != nil {
		return nil, err
	}
	_, err = vpp2pdat.CheckRingConfig(ret.Info.Config)
	if err != nil {
		return nil, err
	}
	ret.localNodes = make([]Node, 0)

	_, err = ret.CheckSig()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// NewRing0 creates a new instance of the default directory ring.
func NewRing0() (*Ring, error) {
	var host0 *Host
	var ring0 *Ring
	var err error

	host0, err = NewHost(vpp2pdat.Host0Title, vpp2pdat.Host0URL, true)
	if err != nil {
		return nil, err
	}
	appID := vpapp.CalcID(vpapp.DefaultPackage(), vpapp.DefaultVersion())
	config := vpp2pdat.DefaultRingConfig()
	ring0, err = NewRing(host0, vpp2pdat.Ring0Title, vpp2pdat.Ring0Description, appID, config, nil, nil)

	return ring0, nil
}

// BuiltinRing0 creates an instance of the default directory ring.
func BuiltinRing0() (*Ring, error) {
	var info vpp2papi.RingInfo
	var err error

	info.RingID, err = base64.URLEncoding.DecodeString(vpp2pdat.Ring0Base64RingID)
	if err != nil {
		return nil, err
	}
	info.RingTitle = vpp2pdat.Ring0Title
	info.RingDescription = vpp2pdat.Ring0Description
	info.AppID, err = base64.URLEncoding.DecodeString(vpp2pdat.Ring0Base64AppID)
	if err != nil {
		return nil, err
	}
	info.Config = vpp2pdat.DefaultRingConfig()
	info.HasPassword = false
	info.HostPubKey, err = base64.URLEncoding.DecodeString(vpp2pdat.Ring0Base64HostPubKey)
	if err != nil {
		return nil, err
	}
	info.RingSig, err = base64.URLEncoding.DecodeString(vpp2pdat.Ring0Base64RingSig)
	if err != nil {
		return nil, err
	}

	return RingFromInfo(&info, nil)
}

// IsSigned returns true if the ring has been signed by corresponding host.
// It does not check if the signature is valid.
func (ring *Ring) IsSigned() bool {
	return vpp2pdat.RingInfoIsSigned(&ring.Info)
}

// CheckSig checks if the ring signature is OK, if it's not, returns 0 and an error.
// If it's OK, returns the number of zeroes in the signature hash.
func (ring *Ring) CheckSig() (int, error) {
	return vpp2pdat.RingInfoCheckSig(&ring.Info)
}
