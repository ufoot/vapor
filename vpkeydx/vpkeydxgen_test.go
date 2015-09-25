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

package vpkeydx

import (
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpsys"
	"testing"
)

const testSeedSize = 512

var testSeedBuf []byte

const testKeyName = "test"
const testX = 1000
const testY = 1000000
const testZ = 1000000000

func init() {
	testSeedBuf = make([]byte, testSeedSize)
}

func TestGen(t *testing.T) {
	keydx, err := Gen(testSeedBuf, testKeyName)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate flat key"))
	}
	t.Logf("generated key %s", vpcrypto.BufToStr256(keydx))
}

func TestGen1d(t *testing.T) {
	keydx, err := Gen1d(testSeedBuf, testKeyName, testX)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate 1d key"))
	}
	t.Logf("generated key %s", vpcrypto.BufToStr256(keydx))
}

func TestGen2d(t *testing.T) {
	keydx, err := Gen2d(testSeedBuf, testKeyName, testX, testY)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate 2d key"))
	}
	t.Logf("generated key %s", vpcrypto.BufToStr256(keydx))
}

func TestGen3d(t *testing.T) {
	keydx, err := Gen3d(testSeedBuf, testKeyName, testX, testY, testZ)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate 3d key"))
	}
	t.Logf("generated key %s", vpcrypto.BufToStr256(keydx))
}

func BenchmarkGen(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = Gen(testSeedBuf, testKeyName)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate flat key"))
		}
	}
}

func BenchmarkGen1d(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = Gen1d(testSeedBuf, testKeyName, testX)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate 1d key"))
		}
	}
}

func BenchmarkGen2d(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = Gen2d(testSeedBuf, testKeyName, testX, testY)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate 2d key"))
		}
	}
}

func BenchmarkGen3d(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = Gen3d(testSeedBuf, testKeyName, testX, testY, testZ)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate 3d key"))
		}
	}
}
