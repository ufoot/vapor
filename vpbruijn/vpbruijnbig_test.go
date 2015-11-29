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

package vpbruijn

import (
	"encoding/hex"
	"math/big"
	"testing"
)

const (
	testI       = 1000
	testNbBytes = 4
	testIString = "000003e8"
)

func TestConv(t *testing.T) {
	i := big.NewInt(testI)
	t.Logf("i=%d", i.Int64())
	b := bigToBytes(i, testNbBytes)
	s := hex.EncodeToString(b)
	t.Logf("b=%s", s)
	if s != testIString {
		t.Errorf("encoded string is wrong, got %s should be %s", s, testIString)
	}
	var bigMax big.Int
	bigMax.Exp(big.NewInt(0x100), big.NewInt(testNbBytes), nil)
	j := bytesToBig(b, &bigMax)
	t.Logf("j=%d", j.Int64())
	if i.Cmp(j) != 0 {
		t.Errorf("big->bytes->big failed, i=%d, j=%d", i.Int64(), j.Int64())
	}
}
