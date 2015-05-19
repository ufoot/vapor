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

package vpgen

import (
	"github.com/ufoot/vapor/vpnumber"
	"github.com/ufoot/vapor/vpvec2"
	"testing"
)

func TestF32Lerp(t *testing.T) {
	const fy1 float32 = 3.0
	const fy2 float32 = 10.0
	const fv1 float32 = -2.0
	const fv2 float32 = -0.5
	const fbegin float32 = -0.2
	const fend float32 = 1.2
	const fbigstep float32 = 0.2
	const flocalrange float32 = 0.002
	const flocalstep float32 = 0.001
	const check float32 = 0.3
	const dt float32 = 0.0001

	f1 := vpvec2.F32New(fy1, fv1)
	f2 := vpvec2.F32New(fy2, fv2)

	for beta := fbegin; beta <= fend; beta += fbigstep {
		for localbeta := beta - flocalrange; localbeta <= beta+flocalrange; localbeta += flocalstep {
			f := F32SuperLerp(f1, f2, localbeta)
			t.Logf("beta=%f f=%s", localbeta, f.String())
		}
	}

	c1 := F32SuperLerp(f1, f2, check-dt)
	c2 := F32SuperLerp(f1, f2, check+dt)
	t.Logf("check=%f f1=%s f2=%s", check, f1.String(), f2.String())
	v := (c2[0] - c1[0]) / (2 * dt)
	t.Logf("v=%f v1=%f v2=%f", v, c1[1], c2[1])
	if !vpnumber.F32IsSimilar(v, c1[1]) {
		t.Errorf("v=%f v1=%f", v, c1[1])
	}
	if !vpnumber.F32IsSimilar(v, c2[1]) {
		t.Errorf("v=%f v2=%f", v, c2[1])
	}
}
