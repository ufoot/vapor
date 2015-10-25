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

// HostInfo stores the static data of a Host.
type HostInfo struct {
	// 256-bit id, depends on the pubKey
	ID []byte
	// Human readable hostname, not necessarily the DNS hostname.
	HostTitle string
	// URL used to connect to this host
	URL string
	// Public Key used to communicate and sign/decryp messages
	PubKey []byte
	// Wether to use cryptographic checks
	CryptoEnable bool
}


// Host is a physical host, it is used to uniquely identify
// a host, it can be used to handle several apps or rings.
type Host struct {
	// Info about the host
	Host HostInfo

	localNodes []Node
}
