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

package vpkoorde

// Read https://en.wikipedia.org/wiki/De_Bruijn_graph for theory

// This is a limited implementation which supposes m=16 and n=64.
// Using m=16 (base 16) justifies the "hex" name, and using n members
// scales it up to 256 bits / 32 bytes.

import (
	"encoding/hex"
	"fmt"
)

const (
	bruijnM   = 16
	bruijnN   = 64
	nbBits    = 256
	nbBytes   = 32
	hexdigits = "0123456789abcdef"
)

func checkX(x []byte) error {
	if len(x) != nbBytes {
		return fmt.Errorf("bad key len=%d, should be %d", len(x), nbBytes)
	}
	return nil
}

func prepareNextHexC(x []byte) []byte {
	enc := make([]byte, bruijnN+1)

	hex.Encode(enc[0:bruijnN], x)

	return enc
}

func nextHexC(enc []byte, c byte) ([]byte, error) {
	ret := make([]byte, nbBytes)

	enc[bruijnN] = c

	_, err := hex.Decode(ret, enc[1:bruijnN+1])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func nextHexFirst(x []byte) ([]byte, error) {
	enc := prepareNextHexC(x)

	return nextHexC(enc, byte('0'))
}

func nextHexLast(x []byte) ([]byte, error) {
	enc := prepareNextHexC(x)

	return nextHexC(enc, byte('f'))
}

// BruijnHexNextFirst returns the first Bruijn node pointed by this node.
// Other nodes might be deduced by just incrementing this one.
func BruijnHexNextFirst(x []byte) ([]byte, error) {
	err := checkX(x)
	if err != nil {
		return nil, err
	}

	return nextHexFirst(x)
}

// BruijnHexNextLast returns the last Bruijn node pointing to this node.
// Other nodes might be deduced by just decrementing this one with
// a value of m**(n-1).
func BruijnHexNextLast(x []byte) ([]byte, error) {
	err := checkX(x)
	if err != nil {
		return nil, err
	}

	return nextHexLast(x)
}

// BruijnHexNextList returns the list of all Bruijn nodes pointed by
// this node, the nodes following this one (we walk down the graph).
func BruijnHexNextList(x []byte) ([][]byte, error) {
	err := checkX(x)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnM)
	enc := prepareNextHexC(x)
	for i, v := range []byte(hexdigits) {
		ret[i], err = nextHexC(enc, v)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func preparePrevHexC(x []byte) []byte {
	enc := make([]byte, bruijnN+1)

	hex.Encode(enc[1:bruijnN+1], x)

	return enc
}

func prevHexC(enc []byte, c byte) ([]byte, error) {
	ret := make([]byte, nbBytes)

	enc[0] = c

	_, err := hex.Decode(ret, enc[0:bruijnN])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func prevHexFirst(x []byte) ([]byte, error) {
	enc := preparePrevHexC(x)

	return prevHexC(enc, byte('0'))
}

func prevHexLast(x []byte) ([]byte, error) {
	enc := preparePrevHexC(x)

	return prevHexC(enc, byte('f'))
}

// BruijnHexPrevFirst returns the first Bruijn node pointing to this node.
// Other nodes might be deduced by just incrementing this one with
// a value of m**(n-1).
func BruijnHexPrevFirst(x []byte) ([]byte, error) {
	err := checkX(x)
	if err != nil {
		return nil, err
	}

	return prevHexFirst(x)
}

// BruijnHexPrevLast returns the last Bruijn node pointing to this node.
// Other nodes might be deduced by just decrementing this one with
// a value of m**(n-1).
func BruijnHexPrevLast(x []byte) ([]byte, error) {
	err := checkX(x)
	if err != nil {
		return nil, err
	}

	return prevHexLast(x)
}

// BruijnHexPrevList returns the list of all Bruijn nodes pointing to
// this node, the nodes preceding this one (we walk up the graph).
func BruijnHexPrevList(x []byte) ([][]byte, error) {
	err := checkX(x)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnM)
	enc := preparePrevHexC(x)
	for i, v := range []byte(hexdigits) {
		ret[i], err = prevHexC(enc, v)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func prepareComposeHex(x, y []byte) []byte {
	enc := make([]byte, bruijnN*2)

	hex.Encode(enc[0:bruijnN], x)
	hex.Encode(enc[bruijnN:bruijnN*2], y)

	return enc
}

func composeHex(enc []byte, i int) ([]byte, error) {
	ret := make([]byte, nbBytes)

	_, err := hex.Decode(ret, enc[i:bruijnN+i])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// BruijnHexForwardPath returns the path between two nodes. The path
// might be non-optimized, it always contains m+1 elements, including
// from and to destination. This is the default forward path in which
// node n+1 is the node after n in the bruijn sequence.
func BruijnHexForwardPath(from, to []byte) ([][]byte, error) {
	err := checkX(from)
	if err != nil {
		return nil, err
	}
	err = checkX(to)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnN)
	enc := prepareComposeHex(from, to)
	for i := range ret {
		ret[i], err = composeHex(enc, i)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// BruijnHexBackwardPath returns the path between two nodes. The path
// might be non-optimized, it always contains m+1 elements, including
// from and to destination. This is the alternative backward path in which
// node n+1 is the node before n in the bruijn sequence.
func BruijnHexBackwardPath(from, to []byte) ([][]byte, error) {
	err := checkX(from)
	if err != nil {
		return nil, err
	}
	err = checkX(to)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnN)
	enc := prepareComposeHex(from, to)
	for i := range ret {
		ret[i], err = composeHex(enc, bruijnN-i)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// BruijnHexForwardElem returns the path element between two nodes.
// Index 0 is the from element, and n (number of elements as in Bruijn nodes)
// the to element. Uses the forward, default path.
func BruijnHexForwardElem(from, to []byte, i int) ([]byte, error) {
	err := checkX(from)
	if err != nil {
		return nil, err
	}
	err = checkX(to)
	if err != nil {
		return nil, err
	}

	enc := prepareComposeHex(from, to)

	return composeHex(enc, i)
}

// BruijnHexBackwardElem returns the path element between two nodes.
// Index 0 is the from element, and n (number of elements as in Bruijn nodes)
// the to element. Uses the backward, alternative path.
func BruijnHexBackwardElem(from, to []byte, i int) ([]byte, error) {
	err := checkX(from)
	if err != nil {
		return nil, err
	}
	err = checkX(to)
	if err != nil {
		return nil, err
	}

	enc := prepareComposeHex(to, from)

	return composeHex(enc, bruijnN-i)
}
