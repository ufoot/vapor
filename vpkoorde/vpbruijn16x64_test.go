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

func TestBruijn16x64Next(t *testing.T) {
	const m = 16
	const n = 64
	var bi big.Int
	var bCheck big.Int

	implGeneric := BruijnGenericNew(m, n)
	impl16x64 := Bruijn16x64New()

	i := vpcrypto.Checksum256([]byte("toto"))
	bi.SetBytes(i)
	t.Logf("i=%s", vpcrypto.IntToStr256(&bi))

	bnf, err := implGeneric.NextFirst(&bi)
	if err != nil {
		t.Error("unable to call BruijnBigNextFirst:", err)
	}
	nf, err := impl16x64.NextFirst(i)
	if err != nil {
		t.Error("unable to call BruijnHexNextFirst:", err)
	}
	bCheck.SetBytes(nf)
	if bnf.Cmp(&bCheck) != 0 {
		t.Errorf("bad nf big=%s hex=%s", vpcrypto.IntToStr256(bnf), vpcrypto.IntToStr256(&bCheck))
	}

	bnl, err := implGeneric.NextLast(&bi)
	if err != nil {
		t.Error("unable to call BruijnBigNextLast:", err)
	}
	nl, err := impl16x64.NextLast(i)
	if err != nil {
		t.Error("unable to call BruijnHexNextLast:", err)
	}
	bCheck.SetBytes(nl)
	if bnl.Cmp(&bCheck) != 0 {
		t.Errorf("bad nl big=%s hex=%s", vpcrypto.IntToStr256(bnl), vpcrypto.IntToStr256(&bCheck))
	}

	bnList, err := implGeneric.NextList(&bi)
	if err != nil {
		t.Error("unable to call BruijnBigNextLast:", err)
	}
	nList, err := impl16x64.NextList(i)
	if err != nil {
		t.Error("unable to call BruijnHexNextLast:", err)
	}
	for j, w := range bnList {
		bCheck.SetBytes(nList[j])
		if w.Cmp(&bCheck) != 0 {
			t.Errorf("j=%d bad nl big=%s hex=%s", j, vpcrypto.IntToStr256(w), vpcrypto.IntToStr256(&bCheck))
		}
	}
}

func TestBruijn16x64Prev(t *testing.T) {
	const m = 16
	const n = 64
	var bi big.Int
	var bCheck big.Int

	implGeneric := BruijnGenericNew(m, n)
	impl16x64 := Bruijn16x64New()

	i := vpcrypto.Checksum256([]byte("toto"))
	bi.SetBytes(i)
	t.Logf("i=%s", vpcrypto.IntToStr256(&bi))

	bnf, err := implGeneric.PrevFirst(&bi)
	if err != nil {
		t.Error("unable to call BruijnBigPrevFirst:", err)
	}
	nf, err := impl16x64.PrevFirst(i)
	if err != nil {
		t.Error("unable to call BruijnHexPrevFirst:", err)
	}
	bCheck.SetBytes(nf)
	if bnf.Cmp(&bCheck) != 0 {
		t.Errorf("bad nf big=%s hex=%s", vpcrypto.IntToStr256(bnf), vpcrypto.IntToStr256(&bCheck))
	}

	bnl, err := implGeneric.PrevLast(&bi)
	if err != nil {
		t.Error("unable to call BruijnBigPrevLast:", err)
	}
	nl, err := impl16x64.PrevLast(i)
	if err != nil {
		t.Error("unable to call BruijnHexPrevLast:", err)
	}
	bCheck.SetBytes(nl)
	if bnl.Cmp(&bCheck) != 0 {
		t.Errorf("bad nl big=%s hex=%s", vpcrypto.IntToStr256(bnl), vpcrypto.IntToStr256(&bCheck))
	}

	bnList, err := implGeneric.PrevList(&bi)
	if err != nil {
		t.Error("unable to call BruijnBigPrevLast:", err)
	}
	nList, err := impl16x64.PrevList(i)
	if err != nil {
		t.Error("unable to call BruijnHexPrevLast:", err)
	}
	for j, w := range bnList {
		bCheck.SetBytes(nList[j])
		if w.Cmp(&bCheck) != 0 {
			t.Errorf("j=%d bad nl big=%s hex=%s", j, vpcrypto.IntToStr256(w), vpcrypto.IntToStr256(&bCheck))
		}
	}
}

func BenchmarkBruijn16x64Next(b *testing.B) {
	impl16x64 := Bruijn16x64New()
	v := vpcrypto.Checksum256([]byte("toto"))
	for i := 0; i < b.N; i++ {
		impl16x64.NextFirst(v)
	}
}

func BenchmarkBruijn16x64Prev(b *testing.B) {
	impl16x64 := Bruijn16x64New()
	v := vpcrypto.Checksum256([]byte("toto"))
	for i := 0; i < b.N; i++ {
		impl16x64.PrevFirst(v)
	}
}
