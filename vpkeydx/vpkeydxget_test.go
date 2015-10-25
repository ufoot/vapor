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
	"fmt"
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vpsys"
	"github.com/ufoot/vapor/vpvec2"
	"github.com/ufoot/vapor/vpvec3"
	"testing"
)

func init() {
	testSeedBuf = make([]byte, testSeedSize)
}

func TestGetX(t *testing.T) {
	keydx, err := GenX(testSeedBuf, testKeyName, testX)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate X key"))
	}
	t.Logf("generated X key %s", vpcrypto.BufToStr256(keydx))
	x, err := GetX(keydx)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to get X coord"))
	}
	if x != testX {
		t.Error(fmt.Errorf("x and testX differ %d %d", x, testX))
	}
}

func TestGetXY(t *testing.T) {
	keydx, err := GenXY(testSeedBuf, testKeyName, testX, testY)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate XY key"))
	}
	t.Logf("generated XY key %s", vpcrypto.BufToStr256(keydx))
	x, y, err := GetXY(keydx)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to get XY coord"))
	}
	if x != testX {
		t.Error(fmt.Errorf("x and testX differ %d %d", x, testX))
	}
	if y != testY {
		t.Error(fmt.Errorf("y and testY differ %d %d", y, testY))
	}
}

func TestGetXYZ(t *testing.T) {
	keydx, err := GenXYZ(testSeedBuf, testKeyName, testX, testY, testZ)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate XYZ key"))
	}
	t.Logf("generated XYZ key %s", vpcrypto.BufToStr256(keydx))
	x, y, z, err := GetXYZ(keydx)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to get XYZ coord"))
	}
	if x != testX {
		t.Error(fmt.Errorf("x and testX differ %d %d", x, testX))
	}
	if y != testY {
		t.Error(fmt.Errorf("y and testY differ %d %d", y, testY))
	}
	if z != testZ {
		t.Error(fmt.Errorf("z and testZ differ %d %d", z, testZ))
	}
}

func TestGetVec1(t *testing.T) {
	keydx, err := GenVec1(testSeedBuf, testKeyName, testX)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate Vec1 key"))
	}
	t.Logf("generated Vec1 key %s", vpcrypto.BufToStr256(keydx))
	x, err := GetVec1(keydx)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to get Vec1 coord"))
	}
	if x != testX {
		t.Error(fmt.Errorf("v and testX differ %d %d", x, testX))
	}
}

func TestGetVec2(t *testing.T) {
	keydx, err := GenVec2(testSeedBuf, testKeyName, vpvec2.I32New(testX, testY))
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate Vec1 key"))
	}
	t.Logf("generated Vec1 key %s", vpcrypto.BufToStr256(keydx))
	vec2, err := GetVec2(keydx)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to get Vec1 coord"))
	}
	if vec2[vpvec2.X] != testX {
		t.Error(fmt.Errorf("x and testX differ %d %d", vec2[vpvec2.X], testX))
	}
	if vec2[vpvec2.Y] != testY {
		t.Error(fmt.Errorf("y and testY differ %d %d", vec2[vpvec2.Y], testY))
	}
}

func TestGetVec3(t *testing.T) {
	keydx, err := GenVec3(testSeedBuf, testKeyName, vpvec3.I32New(testX, testY, testZ))
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to generate Vec1 key"))
	}
	t.Logf("generated Vec1 key %s", vpcrypto.BufToStr256(keydx))
	vec3, err := GetVec3(keydx)
	if err != nil {
		t.Error(vpsys.ErrorChain(err, "unable to get Vec1 coord"))
	}
	if vec3[vpvec3.X] != testX {
		t.Error(fmt.Errorf("x and testX differ %d %d", vec3[vpvec3.X], testX))
	}
	if vec3[vpvec3.Y] != testY {
		t.Error(fmt.Errorf("y and testY differ %d %d", vec3[vpvec3.Y], testY))
	}
	if vec3[vpvec3.Z] != testZ {
		t.Error(fmt.Errorf("z and testZ differ %d %d", vec3[vpvec3.Z], testZ))
	}
}
