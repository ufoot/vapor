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

package vpbruijn

//go:generate bash ./stamp.sh
	
// Read https://en.wikipedia.org/wiki/De_Bruijn_graph for theory

// This is a limited implementation which supposes m=16 and n=64.
// Using m=16 (base 16) justifies the "hex" name, and using n members
// scales it up to 256 bits / 32 bytes.

import (
	"encoding/hex"
	"math/rand"
)

const (
	bruijnM   = 16
	bruijnN   = 64
	nbBits    = 256
	nbBytes   = 32
	hexdigits = "0123456789abcdef"
)

type bruijn16x64 struct {
	implGeneric *bruijnGeneric
}

// Bruijn16x64New creates a new Bruijn object capable of walking
// Bruijn graphcs wih m (AKA base) =16 and n (number of elems) =64.
// This is a specific, hopefully optimized for this case, implementation.
func Bruijn16x64New() BruijnWalker {
	return bruijn16x64New()
}

func bruijn16x64New() *bruijn16x64 {
	return &bruijn16x64{bruijnGenericNew(16, 64)}
}

func (b *bruijn16x64) M() int {
	return bruijnM
}

func (b *bruijn16x64) N() int {
	return bruijnN
}

func (b *bruijn16x64) NbBits() int {
	return nbBits
}

func (b *bruijn16x64) NbBytes() int {
	return nbBytes
}

func (b *bruijn16x64) filterX(x []byte) []byte {
	l := len(x)
	if l >= nbBytes {
		return x[0:nbBytes]
	}

	return append(make([]byte, nbBytes-l), x...)
}

func (b *bruijn16x64) prepareNext16x64C(x []byte) []byte {
	enc := make([]byte, bruijnN+1)

	hex.Encode(enc[0:bruijnN], x)

	return enc
}

func (b *bruijn16x64) next16x64C(enc []byte, c byte) []byte {
	ret := make([]byte, nbBytes)

	enc[bruijnN] = c

	_, err := hex.Decode(ret, enc[1:bruijnN+1])
	if err != nil {
		// return 0, this really should not happen as next16x64 is
		// called internally only, and all keys are checked above,
		// here we reset stuff to 0 to make sure no random data
		// is passed upstream
		ret = make([]byte, nbBytes)
	}

	return ret
}

func (b *bruijn16x64) next16x64First(x []byte) []byte {
	enc := b.prepareNext16x64C(x)

	return b.next16x64C(enc, byte('0'))
}

func (b *bruijn16x64) next16x64Last(x []byte) []byte {
	enc := b.prepareNext16x64C(x)

	return b.next16x64C(enc, byte('f'))
}

func (b *bruijn16x64) NextFirst(x []byte) []byte {
	return b.next16x64First(b.filterX(x))
}

func (b *bruijn16x64) NextLast(x []byte) []byte {
	return b.next16x64Last(b.filterX(x))
}

func (b *bruijn16x64) NextList(x []byte) [][]byte {
	ret := make([][]byte, bruijnM)
	enc := b.prepareNext16x64C(b.filterX(x))
	for i, v := range []byte(hexdigits) {
		ret[i] = b.next16x64C(enc, v)
	}

	return ret
}

func (b *bruijn16x64) preparePrev16x64C(x []byte) []byte {
	enc := make([]byte, bruijnN+1)

	hex.Encode(enc[1:bruijnN+1], b.filterX(x))

	return enc
}

func (b *bruijn16x64) prev16x64C(enc []byte, c byte) []byte {
	ret := make([]byte, nbBytes)

	enc[0] = c

	_, err := hex.Decode(ret, enc[0:bruijnN])
	if err != nil {
		// return 0, this really should not happen as next16x64 is
		// called internally only, and all keys are checked above,
		// here we reset stuff to 0 to make sure no random data
		// is passed upstream
		ret = make([]byte, nbBytes)
	}

	return ret
}

func (b *bruijn16x64) prev16x64First(x []byte) []byte {
	enc := b.preparePrev16x64C(x)

	return b.prev16x64C(enc, byte('0'))
}

func (b *bruijn16x64) prev16x64Last(x []byte) []byte {
	enc := b.preparePrev16x64C(x)

	return b.prev16x64C(enc, byte('f'))
}

func (b *bruijn16x64) PrevFirst(x []byte) []byte {
	return b.prev16x64First(b.filterX(x))
}

func (b *bruijn16x64) PrevLast(x []byte) []byte {
	return b.prev16x64Last(b.filterX(x))
}

func (b *bruijn16x64) PrevList(x []byte) [][]byte {
	ret := make([][]byte, bruijnM)
	enc := b.preparePrev16x64C(b.filterX(x))
	for i, v := range []byte(hexdigits) {
		ret[i] = b.prev16x64C(enc, v)
	}

	return ret
}

func (b *bruijn16x64) prepareCompose16x64(x, y []byte) []byte {
	enc := make([]byte, bruijnN*2)

	hex.Encode(enc[0:bruijnN], x)
	hex.Encode(enc[bruijnN:bruijnN*2], y)

	return enc
}

func (b *bruijn16x64) compose16x64(enc []byte, i int) []byte {
	ret := make([]byte, nbBytes)

	_, err := hex.Decode(ret, enc[i:bruijnN+i])
	if err != nil {
		// return 0, this really should not happen as next16x64 is
		// called internally only, and all keys are checked above,
		// here we reset stuff to 0 to make sure no random data
		// is passed upstream
		ret = make([]byte, nbBytes)
	}

	return ret
}

func (b *bruijn16x64) ForwardPath(from, to []byte) [][]byte {
	ret := make([][]byte, bruijnN)
	enc := b.prepareCompose16x64(b.filterX(from), b.filterX(to))
	for i := range ret {
		ret[i] = b.compose16x64(enc, i)
	}

	return ret
}

func (b *bruijn16x64) BackwardPath(from, to []byte) [][]byte {
	ret := make([][]byte, bruijnN)
	enc := b.prepareCompose16x64(b.filterX(to), b.filterX(from))
	for i := range ret {
		ret[i] = b.compose16x64(enc, bruijnN-i)
	}

	return ret
}

func (b *bruijn16x64) ForwardElem(from, to []byte, i int) []byte {
	enc := b.prepareCompose16x64(b.filterX(from), b.filterX(to))

	return b.compose16x64(enc, i)
}

func (b *bruijn16x64) BackwardElem(from, to []byte, i int) []byte {
	enc := b.prepareCompose16x64(b.filterX(to), b.filterX(from))

	return b.compose16x64(enc, bruijnN-i)
}

func (b *bruijn16x64) Filter(x []byte) []byte {
	l := len(x)
	if l >= nbBytes {
		ret := make([]byte, nbBytes)
		copy(ret, x[l-nbBytes:l])
		return ret
	}

	return append(make([]byte, nbBytes-l), x...)
}

func (b *bruijn16x64) Rand(r *rand.Rand) []byte {
	return b.implGeneric.Rand(r)
}

func (b *bruijn16x64) Zero() []byte {
	return b.implGeneric.Zero()
}

func (b *bruijn16x64) Add(x, y []byte) []byte {
	return b.implGeneric.Add(x, y)
}

func (b *bruijn16x64) Sub(x, y []byte) []byte {
	return b.implGeneric.Sub(x, y)
}

func (b *bruijn16x64) Cmp(x, y []byte) int {
	return b.implGeneric.Cmp(x, y)
}

func (b *bruijn16x64) GeLt(x, begin, end []byte) bool {
	return b.implGeneric.GeLt(x, begin, end)
}

func (b *bruijn16x64) GtLe(x, begin, end []byte) bool {
	return b.implGeneric.GtLe(x, begin, end)
}

func (b *bruijn16x64) Incr(x []byte) []byte {
	if len(x) == nbBytes && x[nbBytes-1] < 255 {
		ret := make([]byte, nbBytes)
		copy(ret, x)
		ret[nbBytes-1] = x[nbBytes-1] + 1
		return ret
	}

	return b.implGeneric.Incr(x)
}

func (b *bruijn16x64) Decr(x []byte) []byte {
	if len(x) == nbBytes && x[nbBytes-1] > 0 {
		ret := make([]byte, nbBytes)
		copy(ret, x)
		ret[nbBytes-1] = x[nbBytes-1] - 1
		return ret
	}

	return b.implGeneric.Decr(x)
}

func (b *bruijn16x64) RingPos(x []byte) float64 {
	return b.implGeneric.RingPos(x)
}

func (b *bruijn16x64) RingRange(x, y []byte) float64 {
	return b.implGeneric.RingRange(x, y)
}

func (b *bruijn16x64) BytesToIntArray(x []byte) []int {
	return b.implGeneric.BytesToIntArray(x)
}

func (b *bruijn16x64) IntArrayToBytes(x []int) []byte {
	return b.implGeneric.IntArrayToBytes(x)
}
