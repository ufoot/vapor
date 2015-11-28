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

import (
	"fmt"
	"math/big"
)

const (
	// GenericMinM is the minimum value for M (AKA base)
	// supported by the generic Bruijn implementation.
	GenericMinM = 2
	// GenericMinN is the minimum value for N (number of elems)
	// supported by the generic Bruijn implementation.
	GenericMinN = 2
	// GenericMaxM is the maximum value for M (AKA base)
	// supported by the generic Bruijn implementation.
	GenericMaxM = 100
	// GenericMaxN is the maximum value for N (number of elems)
	// supported by the generic Bruijn implementation.
	GenericMaxN = 1000
)

type bruijnGeneric struct {
	m       int
	n       int
	nbBits  int
	nbBytes int

	bigM *big.Int
	bigN *big.Int

	bigM1         *big.Int
	bigN1         *big.Int
	bigZero       *big.Int
	bigOne        *big.Int
	bigMax        *big.Int
	bigPrevStep   *big.Int
	bigPrevToLast *big.Int
}

// BruijnGenericNew creates a new Bruijn object capable of walking
// Bruijn graphcs with arbitrary m an n numbers. If m and n are outside
// allowed values, they will be increased/decreased to that they fit.
// This is a generic implementation, can be used for real usage or as a
// reference for optimized version. It can theorically be slower than optimized
// 2^n as it relies on a general purpose big integer machinery, in practice
// it shows decent results.
func BruijnGenericNew(m, n int) bruijnGeneric {
	var ret bruijnGeneric

	if m < GenericMinM {
		m = GenericMinM
	}
	if n < GenericMinN {
		n = GenericMinN
	}
	if m > GenericMaxM {
		m = GenericMaxM
	}
	if n > GenericMaxN {
		n = GenericMaxN
	}

	ret.m = m
	ret.n = n
	ret.nbBits = ret.n * ret.m
	ret.nbBytes = (ret.nbBits + 7) >> 3

	ret.bigM = big.NewInt(int64(m))
	ret.bigN = big.NewInt(int64(n))

	ret.bigM1 = big.NewInt(int64(m - 1))
	ret.bigN1 = big.NewInt(int64(n - 1))
	ret.bigZero = big.NewInt(0)
	ret.bigOne = big.NewInt(1)
	ret.bigMax = big.NewInt(0)
	ret.bigMax.Exp(ret.bigM, ret.bigN, nil)
	ret.bigPrevStep = big.NewInt(0)
	ret.bigPrevStep.Exp(ret.bigM, ret.bigN1, ret.bigMax)
	ret.bigPrevToLast = big.NewInt(0)
	ret.bigPrevToLast.Mul(ret.bigPrevStep, ret.bigM1)

	return ret
}

func (b bruijnGeneric) checkX(x *big.Int) error {
	if x.Cmp(b.bigZero) < 0 {
		return fmt.Errorf("Invalid x=%s (current node) for Bruijn graph with m=%d n=%d, should be >=0", x.String(), b.m, b.n)
	}
	if x.Cmp(b.bigMax) >= 0 {
		return fmt.Errorf("Invalid x=%s (current node) for Bruijn graph with m=%d n=%d, should be <=%s", x.String(), b.m, b.n, b.bigMax.String())
	}
	return nil
}

func (b bruijnGeneric) M() int {
	return b.m
}

func (b bruijnGeneric) N() int {
	return b.n
}

func (b bruijnGeneric) NbBits() int {
	return b.nbBits
}

func (b bruijnGeneric) NbBytes() int {
	return b.nbBytes
}

func (b bruijnGeneric) NextFirst(x []byte) ([]byte, error) {
	bigX := bytesToBig(x)
	err := b.checkX(bigX)
	if err != nil {
		return nil, err
	}

	bigRet := nextBigFirst(bigX, b.bigM, b.bigMax)

	return bigToBytes(bigRet, b.nbBytes), nil
}

func (b bruijnGeneric) NextLast(x []byte) ([]byte, error) {
	bigX := bytesToBig(x)
	err := b.checkX(bigX)
	if err != nil {
		return nil, err
	}

	bigRet := nextBigFirst(bigX, b.bigM, b.bigMax)
	bigRet.Add(bigRet, b.bigM1)

	return bigToBytes(bigRet, b.nbBytes), nil
}

func (b bruijnGeneric) NextList(x []byte) ([][]byte, error) {
	bigX := bytesToBig(x)
	err := b.checkX(bigX)
	if err != nil {
		return nil, err
	}

	cur := nextBigFirst(bigX, b.bigM, b.bigMax)
	ret := make([][]byte, b.m)
	for i := range ret {
		ret[i] = bigToBytes(cur, b.nbBytes)
		cur.Add(cur, b.bigOne)
	}

	return ret, nil
}

func (b bruijnGeneric) PrevFirst(x []byte) ([]byte, error) {
	bigX := bytesToBig(x)
	err := b.checkX(bigX)
	if err != nil {
		return nil, err
	}

	bigRet := prevBigFirst(bigX, b.bigM)

	return bigToBytes(bigRet, b.nbBytes), nil
}

func (b bruijnGeneric) PrevLast(x []byte) ([]byte, error) {
	bigX := bytesToBig(x)
	err := b.checkX(bigX)
	if err != nil {
		return nil, err
	}

	bigRet := prevBigFirst(bigX, b.bigM)
	bigRet.Add(bigRet, b.bigPrevToLast)

	return bigToBytes(bigRet, b.nbBytes), nil
}

func (b bruijnGeneric) PrevList(x []byte) ([][]byte, error) {
	bigX := bytesToBig(x)
	err := b.checkX(bigX)
	if err != nil {
		return nil, err
	}

	cur := prevBigFirst(bigX, b.bigM)
	ret := make([][]byte, b.m)
	for i := range ret {
		ret[i] = bigToBytes(cur, b.nbBytes)
		cur.Add(cur, b.bigPrevStep)
	}

	return ret, nil
}

func (b bruijnGeneric) ForwardPath(from, to []byte) ([][]byte, error) {
	bigFrom := bytesToBig(from)
	err := b.checkX(bigFrom)
	if err != nil {
		return nil, err
	}
	bigTo := bytesToBig(to)
	err = b.checkX(bigTo)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, b.n+1)

	for i := range ret {
		ret[i] = bigToBytes(composeBig(bigFrom, bigTo, b.bigM, b.bigMax, b.m, b.n, i), b.nbBytes)
	}

	return ret, nil
}

func (b bruijnGeneric) BackwardPath(from, to []byte) ([][]byte, error) {
	bigFrom := bytesToBig(from)
	err := b.checkX(bigFrom)
	if err != nil {
		return nil, err
	}
	bigTo := bytesToBig(to)
	err = b.checkX(bigTo)
	if err != nil {
		return nil, err
	}

	ret := make([][]byte, b.n+1)

	for i := range ret {
		ret[i] = bigToBytes(composeBig(bigTo, bigFrom, b.bigM, b.bigMax, b.m, b.n, b.n-i), b.nbBytes)
	}

	return ret, nil
}

func (b bruijnGeneric) ForwardElem(from, to []byte, i int) ([]byte, error) {
	bigFrom := bytesToBig(from)
	err := b.checkX(bigFrom)
	if err != nil {
		return nil, err
	}
	bigTo := bytesToBig(to)
	err = b.checkX(bigTo)
	if err != nil {
		return nil, err
	}

	return bigToBytes(composeBig(bigFrom, bigTo, b.bigM, b.bigMax, b.m, b.n, i), b.nbBytes), nil
}

func (b bruijnGeneric) BackwardElem(from, to []byte, i int) ([]byte, error) {
	bigFrom := bytesToBig(from)
	err := b.checkX(bigFrom)
	if err != nil {
		return nil, err
	}
	bigTo := bytesToBig(to)
	err = b.checkX(bigTo)
	if err != nil {
		return nil, err
	}

	return bigToBytes(composeBig(bigTo, bigFrom, b.bigM, b.bigMax, b.m, b.n, b.n-i), b.nbBytes), nil
}
