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

package vpp2p

import (
	"github.com/ufoot/vapor/vpbruijn"
)

const (
	// DefaultBruijnM default for the m parameter (base) used for Koorde/Bruijn ops.
	DefaultBruijnM = 16
	// DefaultBruijnN default for the n parameter (number of elements) used for Koorde/Bruijn ops.
	DefaultBruijnN = 64
	// DefaultNbCopy default for the number of copies of a key we store within the network.
	DefaultNbCopy = 3
	// DefaultNbSub default for the number sub virtual rings, used for redundancy, mostly.
	DefaultNbSub = 3
	// DefaultUseSig default for wether to use cryptographic signatures/checks.
	DefaultUseSig = false
	// DefaultSteps default to optimizes Bruijn walk by considering only this number
	// of steps in the worst case.
	DefaultSteps = 8
)

// RingConfig stores various ring technical parameters. Normally these
// do not change anything concerning the functionnal behavior of the
// program, it's just about performance, redundancy, stability.
type RingConfig struct {
	// BruijnM is the De Bruijn M (AKA base) used for De Bruijn networks.
	BruijnM int
	// BruijnN is the De Bruijn N (number of elems) used for De Bruijn networks.
	BruijnN int
	// NbBackup is the number of copies of a key we store within the network.
	NbCopy int
	// Number of sub virtual rings, used for redundancy, mostly
	NbSub int
	// UseSig tells wether to use cryptographic signatures/checks.
	UseSig bool
	// Steps optimizes Bruijn walk by considering only this number
	// of steps in the worst case.
	Steps int
}

// DefaultRingConfig returns a default ring configuration, with 256-bit keys
func DefaultRingConfig() *RingConfig {
	return &RingConfig{DefaultBruijnM, DefaultBruijnN, DefaultNbCopy, DefaultNbSub, DefaultUseSig, DefaultSteps}
}

// RingInfo stores the static data of a Ring.
type RingInfo struct {
	// N-bit id, totally random, create a new for a new session.
	RingID []byte
	// Human readable ring (short) description
	RingTitle string
	// App details
	App AppInfo
	// Password hash
	PasswordHash []byte
	// RingConfig contains technical parameters.
	Config RingConfig
}

// Ring is a community, a network of related nodes, which communicate
// through hosts. It is the logical artefact used to relate several hosts/nodes
// together.
type Ring struct {
	// Info about the ring
	Info RingInfo

	walker     vpbruijn.BruijnWalker
	localNodes []Node
}

// NewRing creates a new ring from static data.
func NewRing(ringID []byte, ringTitle string, app *AppInfo, passwordHash []byte, config *RingConfig) (*Ring, error) {
	var ret Ring
	var ok bool
	var err error

	ok, err = CheckID(ringID)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = CheckTitle(ringTitle)
	if err != nil || !ok {
		return nil, err
	}

	ret.localNodes = make([]Node, 0)

	ret.Info.RingID = ringID
	ret.Info.RingTitle = ringTitle
	ret.Info.App = *app
	ret.Info.PasswordHash = passwordHash
	ret.Info.Config = *config

	return &ret, nil
}
