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

type bruijn16x64 struct {
}

// Bruijn16x64New creates a new Bruijn object capable of walking
// Bruijn graphcs wih m (AKA base) =16 and n (number of elems) =64.
// This is a specific, hopefully optimized for this case, implementation.
func Bruijn16x64New() bruijn16x64 {
	return bruijn16x64{}
}

func (b bruijn16x64) M() int {
	return bruijnM
}

func (b bruijn16x64) N() int {
	return bruijnN
}

func (b bruijn16x64) NbBits() int {
	return nbBits
}

func (b bruijn16x64) NbBytes() int {
	return nbBytes
}

func (b bruijn16x64) checkX(x []byte) error {
	if len(x) != nbBytes {
		return fmt.Errorf("bad key len=%d, should be %d", len(x), nbBytes)
	}
	return nil
}

func (b bruijn16x64) prepareNext16x64C(x []byte) []byte {
	enc := make([]byte, bruijnN+1)

	hex.Encode(enc[0:bruijnN], x)

	return enc
}

func (b bruijn16x64) next16x64C(enc []byte, c byte) ([]byte, error) {
	ret := make([]byte, nbBytes)

	enc[bruijnN] = c

	_, err := hex.Decode(ret, enc[1:bruijnN+1])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (b bruijn16x64) next16x64First(x []byte) ([]byte, error) {
	enc := b.prepareNext16x64C(x)

	return b.next16x64C(enc, byte('0'))
}

func (b bruijn16x64) next16x64Last(x []byte) ([]byte, error) {
	enc := b.prepareNext16x64C(x)

	return b.next16x64C(enc, byte('f'))
}

func (b bruijn16x64) NextFirst(x []byte) ([]byte, error) {
	err := b.checkX(x)
	if err != nil {
		return nil, err
	}

	return b.next16x64First(x)
}

func (b bruijn16x64) NextLast(x []byte) ([]byte, error) {
	err := b.checkX(x)
	if err != nil {
		return nil, err
	}

	return b.next16x64Last(x)
}

func (b bruijn16x64) NextList(x []byte) ([][]byte, error) {
	err := b.checkX(x)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnM)
	enc := b.prepareNext16x64C(x)
	for i, v := range []byte(hexdigits) {
		ret[i], err = b.next16x64C(enc, v)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b bruijn16x64) preparePrev16x64C(x []byte) []byte {
	enc := make([]byte, bruijnN+1)

	hex.Encode(enc[1:bruijnN+1], x)

	return enc
}

func (b bruijn16x64) prev16x64C(enc []byte, c byte) ([]byte, error) {
	ret := make([]byte, nbBytes)

	enc[0] = c

	_, err := hex.Decode(ret, enc[0:bruijnN])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (b bruijn16x64) prev16x64First(x []byte) ([]byte, error) {
	enc := b.preparePrev16x64C(x)

	return b.prev16x64C(enc, byte('0'))
}

func (b bruijn16x64) prev16x64Last(x []byte) ([]byte, error) {
	enc := b.preparePrev16x64C(x)

	return b.prev16x64C(enc, byte('f'))
}

func (b bruijn16x64) PrevFirst(x []byte) ([]byte, error) {
	err := b.checkX(x)
	if err != nil {
		return nil, err
	}

	return b.prev16x64First(x)
}

func (b bruijn16x64) PrevLast(x []byte) ([]byte, error) {
	err := b.checkX(x)
	if err != nil {
		return nil, err
	}

	return b.prev16x64Last(x)
}

func (b bruijn16x64) PrevList(x []byte) ([][]byte, error) {
	err := b.checkX(x)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnM)
	enc := b.preparePrev16x64C(x)
	for i, v := range []byte(hexdigits) {
		ret[i], err = b.prev16x64C(enc, v)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b bruijn16x64) prepareCompose16x64(x, y []byte) []byte {
	enc := make([]byte, bruijnN*2)

	hex.Encode(enc[0:bruijnN], x)
	hex.Encode(enc[bruijnN:bruijnN*2], y)

	return enc
}

func (b bruijn16x64) compose16x64(enc []byte, i int) ([]byte, error) {
	ret := make([]byte, nbBytes)

	_, err := hex.Decode(ret, enc[i:bruijnN+i])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (b bruijn16x64) ForwardPath(from, to []byte) ([][]byte, error) {
	err := b.checkX(from)
	if err != nil {
		return nil, err
	}
	err = b.checkX(to)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnN)
	enc := b.prepareCompose16x64(from, to)
	for i := range ret {
		ret[i], err = b.compose16x64(enc, i)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b bruijn16x64) BackwardPath(from, to []byte) ([][]byte, error) {
	err := b.checkX(from)
	if err != nil {
		return nil, err
	}
	err = b.checkX(to)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, bruijnN)
	enc := b.prepareCompose16x64(from, to)
	for i := range ret {
		ret[i], err = b.compose16x64(enc, bruijnN-i)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b bruijn16x64) ForwardElem(from, to []byte, i int) ([]byte, error) {
	err := b.checkX(from)
	if err != nil {
		return nil, err
	}
	err = b.checkX(to)
	if err != nil {
		return nil, err
	}

	enc := b.prepareCompose16x64(from, to)

	return b.compose16x64(enc, i)
}

func (b bruijn16x64) BackwardElem(from, to []byte, i int) ([]byte, error) {
	err := b.checkX(from)
	if err != nil {
		return nil, err
	}
	err = b.checkX(to)
	if err != nil {
		return nil, err
	}

	enc := b.prepareCompose16x64(to, from)

	return b.compose16x64(enc, bruijnN-i)
}
