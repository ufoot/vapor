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
	"github.com/ufoot/vapor/vpcrypto"
	"math/big"
	"testing"
)

func TestBruijnHexNext(t *testing.T) {
	const m = 2
	const n = 3
	var bi big.Int
	var bCheck big.Int

	i := vpcrypto.Checksum256([]byte("toto"))
	bi.SetBytes(i)
	t.Logf("i=%s", vpcrypto.IntToStr256(&bi))

	bnf, err := BruijnBigNextFirst(bruijnM, bruijnN, &bi)
	if err != nil {
		t.Error("unable to call BruijnBigNextFirst:", err)
	}
	nf, err := BruijnHexNextFirst(i)
	if err != nil {
		t.Error("unable to call BruijnHexNextFirst:", err)
	}
	bCheck.SetBytes(nf)
	if bnf.Cmp(&bCheck) != 0 {
		t.Errorf("bad nf big=%s hex=%s", vpcrypto.IntToStr256(bnf), vpcrypto.IntToStr256(&bCheck))
	}

	bnl, err := BruijnBigNextLast(bruijnM, bruijnN, &bi)
	if err != nil {
		t.Error("unable to call BruijnBigNextLast:", err)
	}
	nl, err := BruijnHexNextLast(i)
	if err != nil {
		t.Error("unable to call BruijnHexNextLast:", err)
	}
	bCheck.SetBytes(nl)
	if bnl.Cmp(&bCheck) != 0 {
		t.Errorf("bad nl big=%s hex=%s", vpcrypto.IntToStr256(bnl), vpcrypto.IntToStr256(&bCheck))
	}
}

func BenchmarkBruijnHexNext(b *testing.B) {
	v := vpcrypto.Checksum256([]byte("toto"))
	for i := 0; i < b.N; i++ {
		BruijnHexNextList(v)
	}
}

func BenchmarkBruijnHexPrev(b *testing.B) {
	v := vpcrypto.Checksum256([]byte("toto"))
	for i := 0; i < b.N; i++ {
		BruijnHexPrevList(v)
	}
}
