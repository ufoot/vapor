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
	"fmt"
	"github.com/ufoot/vapor/vpbruijn"
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpp2papi"
	"math/big"
)

// RingKeySeconds specifies how many seconds should we spend on creating
// ring keys in signed mode.
const RingKeySeconds = 5

// RingKeyZeroes specifies how many zeroes there should be at the end of
// a ring key sig in signed mode.
const RingKeyZeroes = 12

const (
	// DefaultBruijnM default for the m parameter (base) used for Koorde/Bruijn ops.
	DefaultBruijnM = 16
	// DefaultBruijnN default for the n parameter (number of elements) used for Koorde/Bruijn ops.
	DefaultBruijnN = 64
	// DefaultNbCopy default for the number of copies of a key we store within the network.
	DefaultNbCopy = 3
	// DefaultNbStep default to optimizes Bruijn walk by considering only this number
	// of steps in the worst case.
	DefaultNbStep = 8
)

// DefaultRingConfig returns a default ring configuration, with 256-bit keys
func DefaultRingConfig() *vpp2papi.RingConfig {
	return &vpp2papi.RingConfig{BruijnM: DefaultBruijnM, BruijnN: DefaultBruijnN, NbCopy: DefaultNbCopy, NbStep: DefaultNbStep}
}

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
	return SigBytesRing(ringID, r.ringTitle, r.ringDescription, r.appID, r.config, r.hasPassword)
}

// SigBytesRing returns the byte buffer that needs to be signed.
func SigBytesRing(ringID []byte, ringTitle, ringDescription string, appID []byte, config *vpp2papi.RingConfig, hasPassword bool) []byte {
	bufTitle := []byte(ringTitle)
	bufDescription := []byte(ringDescription)
	bufScalar := fmt.Sprintf("(%d,%d,%d,%d,%t)", config.BruijnM, config.BruijnN, config.NbCopy, config.NbStep, hasPassword)
	ret := make([]byte, len(ringID)+len(ringTitle)+len(ringDescription)+len(appID)+len(bufScalar))
	begin := 0
	end := begin + len(ringID)
	copy(ret[begin:end], ringID)
	begin += len(ringID)
	end = begin + len(bufTitle)
	copy(ret[begin:end], bufTitle)
	begin += len(bufTitle)
	end = begin + len(bufDescription)
	copy(ret[begin:end], bufDescription)
	begin += len(bufDescription)
	end = begin + len(appID)
	copy(ret[begin:end], appID)
	begin += len(appID)
	end = begin + len(bufScalar)
	copy(ret[begin:end], bufScalar)

	return ret
}

// NewRing creates a new ring from static data.
func NewRing(host *Host, ringTitle, ringDescription string, appID []byte, config *vpp2papi.RingConfig, fc vpcrypto.FilterChecker, passwordHash []byte) (*Ring, error) {
	var ret Ring
	var ok bool
	var err error
	var intRingID *big.Int
	var sig []byte

	ok, err = CheckTitle(ringTitle)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = CheckDescription(ringDescription)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = CheckID(appID)
	if err != nil || !ok {
		return nil, err
	}
	if config == nil {
		config = DefaultRingConfig()
	}
	ret.walker, err = vpbruijn.BruijnNew(int(config.BruijnM), int(config.BruijnN))
	if err != nil {
		return nil, err
	}
	if config.NbCopy <= 0 || config.NbCopy > config.BruijnM*config.BruijnN {
		return nil, fmt.Errorf("bad NbCopy param %d, should be between 0 and %d (the latter is BruijnN*BruijnM, while this number does not technically has a direct meaning, it should really be bigger than copies, just check your settings, usually, keeping more than a dozen copies is overkill)", config.NbCopy, config.BruijnN*config.BruijnM)
	}
	if config.NbStep < 1 || config.NbStep > config.BruijnN {
		return nil, fmt.Errorf("bad NbStep param %d, should be between 1 and BruijnN which is %d", config.NbStep, config.BruijnN)
	}
	if err != nil || !ok {
		return nil, err
	}

	hasPassword := (passwordHash != nil) && (len(passwordHash) > 0)
	rsa := ringStuffAppender{ringTitle: ringTitle, ringDescription: ringDescription, appID: appID, config: config, hasPassword: hasPassword}
	if host.CanSign() {
		intRingID, sig, _, err = vpcrypto.GenerateID512(host.key, fc, &rsa, RingKeySeconds, RingKeyZeroes)
		if err != nil {
			return nil, err
		}
	} else {
		intRingID, _, _, err = vpcrypto.GenerateID512(nil, fc, &rsa, RingKeySeconds, RingKeyZeroes)
		if err != nil {
			return nil, err
		}
		sig = []byte("")
	}

	ret.Info.RingID = vpcrypto.IntToBuf512(intRingID)
	ret.Info.RingTitle = ringTitle
	ret.Info.RingDescription = ringDescription
	ret.Info.AppID = appID
	ret.Info.Config = DefaultRingConfig()
	*ret.Info.Config = *config
	ret.Info.HasPassword = hasPassword
	ret.secret.PasswordHash = passwordHash
	ret.Info.HostPubKey = make([]byte, len(host.Info.HostPubKey))
	copy(ret.Info.HostPubKey, host.Info.HostPubKey)
	ret.Info.RingSig = sig

	ret.localNodes = make([]Node, 0)

	return &ret, nil
}

// IsSigned returns true if the ring has been signed by corresponding host.
func (ring *Ring) IsSigned() bool {
	return ring.Info.RingSig != nil && len(ring.Info.RingSig) > 0
}
