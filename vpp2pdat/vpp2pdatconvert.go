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
	// NodeIDNbBytes gives the number of bytes for a node ID.
	NodeIDNbBytes = 32
	// NodeIDNbBytes gives the number of bytes for a ring ID.
	RingIDNbBytes = 64
)

func NodeIDToBuf(nodeID []byte) [NodeIDNbBytes]byte {
	var ret [NodeIDNbBytes]byte

	if nodeID == nil {
		return ret
	}

	if len(nodeID) < NodeIDNbBytes {
		copy(ret[NodeIDNbBytes-len(nodeID):NodeIDNbBytes], nodeID)
		return ret
	}

	copy(ret[0:NodeIDNbBytes], nodeID)
	return ret
}

func RingIDToBuf(ringID []byte) [RingIDNbBytes]byte {
	var ret [RingIDNbBytes]byte

	if ringID == nil {
		return ret
	}

	if len(ringID) < RingIDNbBytes {
		copy(ret[RingIDNbBytes-len(ringID):RingIDNbBytes], ringID)
		return ret
	}

	copy(ret[0:RingIDNbBytes], ringID)
	return ret
}
