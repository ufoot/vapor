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

// HostInfo stores the static data of a Host.
type HostInfo struct {
	// Human readable hostname, not necessarily the DNS hostname.
	HostTitle string
	// URL used to connect to this host
	URL string
	// Public Key used to communicate and sign/decryp messages.
	// It's also the host unique identifier (AKA ID).
	PubKey []byte
}

// Host is a physical host, it is used to uniquely identify
// a host, it can be used to handle several apps or rings.
type Host struct {
	// Info about the host
	Info HostInfo

	localNodes []Node
}

// NewHost returns a new host object
func NewHost(title, url string, pubKey []byte) (*Host, error) {
	var ret Host

	ok, err := CheckTitle(title)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = CheckURL(url)
	if err != nil || !ok {
		return nil, err
	}
	ok, err = CheckPubKey(pubKey)
	if err != nil || !ok {
		return nil, err
	}

	ret.Info = HostInfo{title, url, pubKey}
	ret.localNodes = make([]Node, 0)

	return &ret, nil
}
