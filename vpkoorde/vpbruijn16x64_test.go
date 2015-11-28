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
	"testing"
)

func TestBruijn16x64Next(t *testing.T) {
	const m = 16
	const n = 64

	implGeneric := BruijnGenericNew(m, n)
	impl16x64 := Bruijn16x64New()

	i := vpcrypto.Checksum256([]byte("toto"))

	nfGeneric := implGeneric.NextFirst(i)
	nf16x64 := impl16x64.NextFirst(i)
	if implGeneric.Cmp(nfGeneric, nf16x64) == 0 {
		t.Logf("i=%s nf=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(nf16x64))
	} else {
		t.Errorf("bad nf i=%s big=%s hex=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(nfGeneric), vpcrypto.BufToStr256(nf16x64))
	}

	nlGeneric := implGeneric.NextFirst(i)
	nl16x64 := impl16x64.NextFirst(i)
	if implGeneric.Cmp(nlGeneric, nl16x64) == 0 {
		t.Logf("i=%s nl=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(nl16x64))
	} else {
		t.Errorf("bad nl i=%s big=%s hex=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(nlGeneric), vpcrypto.BufToStr256(nl16x64))
	}
}

func TestBruijn16x64Prev(t *testing.T) {
	const m = 16
	const n = 64

	implGeneric := BruijnGenericNew(m, n)
	impl16x64 := Bruijn16x64New()

	i := vpcrypto.Checksum256([]byte("toto"))

	pfGeneric := implGeneric.PrevFirst(i)
	pf16x64 := impl16x64.PrevFirst(i)
	if implGeneric.Cmp(pfGeneric, pf16x64) == 0 {
		t.Logf("i=%s pf=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(pf16x64))
	} else {
		t.Errorf("bad pf i=%s big=%s hex=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(pfGeneric), vpcrypto.BufToStr256(pf16x64))
	}

	plGeneric := implGeneric.PrevFirst(i)
	pl16x64 := impl16x64.PrevFirst(i)
	if implGeneric.Cmp(plGeneric, pl16x64) == 0 {
		t.Logf("i=%s pl=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(pl16x64))
	} else {
		t.Errorf("bad pl i=%s big=%s hex=%s", vpcrypto.BufToStr256(i), vpcrypto.BufToStr256(plGeneric), vpcrypto.BufToStr256(pl16x64))
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
