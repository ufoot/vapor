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

package vpp2pdat

const (
	// HostPubKeyBufNbBytes gives the number of bytes for a host pub key when used as a map key.
	HostPubKeyBufNbBytes = 16
	// NodeIDBufNbBytes gives the number of bytes for a node ID when used as a map key.
	NodeIDBufNbBytes = 32
	// RingIDBufNbBytes gives the number of bytes for a ring ID when used as a map key.
	RingIDBufNbBytes = 64
)

// HostPubKeyToBuf converts a slice to a fixed-length 16 bytes (128 bits) buffer.
func HostPubKeyToBuf(hostPubKey []byte) [HostPubKeyBufNbBytes]byte {
	var ret [HostPubKeyBufNbBytes]byte

	if hostPubKey == nil {
		return ret
	}

	if len(hostPubKey) < HostPubKeyBufNbBytes {
		copy(ret[HostPubKeyBufNbBytes-len(hostPubKey):HostPubKeyBufNbBytes], hostPubKey)
		return ret
	}

	copy(ret[0:HostPubKeyBufNbBytes], hostPubKey)
	return ret
}

// NodeIDToBuf converts a slice to a fixed-length 32 bytes (256 bits) buffer.
func NodeIDToBuf(nodeID []byte) [NodeIDBufNbBytes]byte {
	var ret [NodeIDBufNbBytes]byte

	if nodeID == nil {
		return ret
	}

	if len(nodeID) < NodeIDBufNbBytes {
		copy(ret[NodeIDBufNbBytes-len(nodeID):NodeIDBufNbBytes], nodeID)
		return ret
	}

	copy(ret[0:NodeIDBufNbBytes], nodeID)
	return ret
}

// RingIDToBuf converts a slice to a fixed-length 64 bytes (512 bits) buffer.
func RingIDToBuf(ringID []byte) [RingIDBufNbBytes]byte {
	var ret [RingIDBufNbBytes]byte

	if ringID == nil {
		return ret
	}

	if len(ringID) < RingIDBufNbBytes {
		copy(ret[RingIDBufNbBytes-len(ringID):RingIDBufNbBytes], ringID)
		return ret
	}

	copy(ret[0:RingIDBufNbBytes], ringID)
	return ret
}
