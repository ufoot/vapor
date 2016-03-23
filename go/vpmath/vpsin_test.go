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

package vpmath

import (
	"github.com/ufoot/vapor/go/vpnumber"
	"math"
	"testing"
)

func TestX32Sin(t *testing.T) {
	var f float64
	var x1 vpnumber.X32
	var x2 vpnumber.X32

	for f = -20.0; f < 20.0; f += 0.5 {
		x1 = X32Sin(vpnumber.F64ToX32(f))
		x2 = vpnumber.F64ToX32(math.Sin(f))
		if vpnumber.X32IsSimilar(x1, x2) {
			t.Logf("similar sin values for f=%f x1=%d x2=%d", f, x1, x2)
		} else {
			t.Errorf("inconsistent sin values for f=%f x1=%d x2=%d", f, x1, x2)
		}
	}
}

func TestX64Sin(t *testing.T) {
	var f float64
	var x1 vpnumber.X64
	var x2 vpnumber.X64

	for f = -20.0; f < 20.0; f += 0.5 {
		x1 = X64Sin(vpnumber.F64ToX64(f))
		x2 = vpnumber.F64ToX64(math.Sin(f))
		if vpnumber.X64IsSimilar(x1, x2) {
			t.Logf("similar sin values for f=%f x1=%d x2=%d", f, x1, x2)
		} else {
			t.Errorf("inconsistent sin values for f=%f x1=%d x2=%d", f, x1, x2)
		}
	}
}

func BenchmarkX32Sin(b *testing.B) {
	var x = vpnumber.I32ToX32(100)

	for i := 0; i < b.N; i++ {
		_ = X32Sin(x)
	}
}

func BenchmarkX64Sin(b *testing.B) {
	var x = vpnumber.I64ToX64(10000)

	for i := 0; i < b.N; i++ {
		_ = X64Sin(x)
	}
}

func BenchmarkFSin(b *testing.B) {
	var f = 1000000.0

	for i := 0; i < b.N; i++ {
		_ = math.Sin(f)
	}
}
