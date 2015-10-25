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
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
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

func TestGenX(t *testing.T) {
	keydx, err := GenX(testSeedBuf, testKeyName, testX)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate X key"))
	}
	t.Logf("generated X key %s", vpcrypto.BufToStr256(keydx))
}

func TestGenXY(t *testing.T) {
	keydx, err := GenXY(testSeedBuf, testKeyName, testX, testY)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate XY key"))
	}
	t.Logf("generated XY key %s", vpcrypto.BufToStr256(keydx))
}

func TestGenXYZ(t *testing.T) {
	keydx, err := GenXYZ(testSeedBuf, testKeyName, testX, testY, testZ)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate XYZ key"))
	}
	t.Logf("generated XYZ key %s", vpcrypto.BufToStr256(keydx))
}

func TestGenVec1(t *testing.T) {
	keydx, err := GenVec1(testSeedBuf, testKeyName, testX)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate Vec1 key"))
	}
	t.Logf("generated Vec1 key %s", vpcrypto.BufToStr256(keydx))
}

func TestGenVec2(t *testing.T) {
	keydx, err := GenVec2(testSeedBuf, testKeyName, vpvec2.I32New(testX, testY))
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate Vec2 key"))
	}
	t.Logf("generated Vec2 key %s", vpcrypto.BufToStr256(keydx))
}

func TestGenVec3(t *testing.T) {
	keydx, err := GenVec3(testSeedBuf, testKeyName, vpvec3.I32New(testX, testY, testZ))
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate Vec3 key"))
	}
	t.Logf("generated Vec3 key %s", vpcrypto.BufToStr256(keydx))
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

func BenchmarkGenX(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = GenX(testSeedBuf, testKeyName, testX)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate X key"))
		}
	}
}

func BenchmarkGenXY(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = GenXY(testSeedBuf, testKeyName, testX, testY)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate XY key"))
		}
	}
}

func BenchmarkGenXYZ(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		_, err = GenXYZ(testSeedBuf, testKeyName, testX, testY, testZ)
		if err != nil {
			b.Error(vpsys.ErrorChain(err, "unable to generate XYZ key"))
		}
	}
}
