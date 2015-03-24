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
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpnumber

import (
	"testing"
)

func TestX64Convert(t *testing.T) {
	var i32 int32
	var i64 int64
	var x32 X32
	var f32 float32
	var f64 float64

	i32 = X64ToI32(X64Const1)
	if i32 != I32Const1 {
		t.Error("can't convert positive X64 to int32", i32)
	}
	i32 = X64ToI32(-X64Const1)
	if i32 != -I32Const1 {
		t.Error("can't convert negative X64 to int32", i32)
	}
	i64 = X64ToI64(X64Const1)
	if i64 != I64Const1 {
		t.Error("can't convert positive X64 to int64", i64)
	}
	i64 = X64ToI64(-X64Const1)
	if i64 != -I64Const1 {
		t.Error("can't convert negative X64 to int64", i64)
	}
	x32 = X64ToX32(X64Const1)
	if x32 != X32Const1 {
		t.Error("can't convert positive X64 to X32", x32)
	}
	x32 = X64ToX32(-X64Const1)
	if x32 != -X32Const1 {
		t.Error("can't convert negtive X64 to X32", x32)
	}
	f32 = X64ToF32(X64Const1)
	if f32 != F32Const1 {
		t.Error("can't convert positive X64 to float32", f32)
	}
	f32 = X64ToF32(-X64Const1)
	if f32 != -F32Const1 {
		t.Error("can't convert negative X64 to float32", f32)
	}
	f64 = X64ToF64(X64Const1)
	if f64 != F64Const1 {
		t.Error("can't convert positive X64 to float64", f64)
	}
	f64 = X64ToF64(-X64Const1)
	if f64 != -F64Const1 {
		t.Error("can't convert negative X64 to float64", f64)
	}
}

func TestX64Math(t *testing.T) {
	var x X64

	x = X64Min(-x64Const5, -X64Const1)
	if x != -x64Const5 {
		t.Error("can't find min of xint64 numbers", x)
	}
	x = X64Max(x64Const2, -x64Const3)
	if x != x64Const2 {
		t.Error("can't find max of xint64 numbers", x)
	}
	x = X64Abs(-x64Const2)
	if x != x64Const2 {
		t.Error("can't get abs of negative xint64 number", x)
	}
	x = X64Abs(x64Const3)
	if x != x64Const3 {
		t.Error("can't get abs of positive xint64 number", x)
	}
	x = X64Mul(x64Const2, x64Const3)
	if x != x64Const6 {
		t.Error("can't mul xint64 numbers", x)
	}
	x = X64Muln(x64Const2, x64Const3, x64Const4)
	if x != X64Mul(x64Const6, x64Const4) {
		t.Error("can't muln xint64 numbers", x)
	}
	x = X64Div(x64Const8, x64Const2)
	if x != x64Const4 {
		t.Error("can't div xint64 numbers", x)
	}
	x = X64Mulp(x64Const2, x64Const3)
	if x != x64Const6 {
		t.Error("can't mulp xint64 numbers", x)
	}
	x = X64Divp(x64Const8, x64Const2)
	if x != x64Const4 {
		t.Error("can't divp xint64 numbers", x)
	}
}

func TestX64Float(t *testing.T) {
	var e int
	var x X64

	e = X64Exponent(X64Const1)
	if e != 0 {
		t.Error("Exponent does not work on 1", e)
	}
	e = X64Exponent(X64Const0)
	if e != 0 {
		t.Error("Exponent does not work on 0", e)
	}
	e = X64Exponent(x64Const5)
	if e != 2 {
		t.Error("Exponent does not work on 5", e)
	}
	e = X64Exponent(-(x64Const5 >> 4))
	if e != -2 {
		t.Error("Exponent does not work on -5", e)
	}
	x = X64Mantis(X64Const1)
	if x != X64Const1 {
		t.Error("Mantis does not work on 1", x)
	}
	x = X64Mantis(x64Const5)
	if x != x64Const5>>2 {
		t.Error("Mantis does not work on 5", x)
	}
	x = X64Mantis(-(x64Const5 >> 4))
	if x != -x64Const5>>2 {
		t.Error("Mantis does not work on -5", x)
	}
	x = X64Round((x64Const5) >> 2)
	if x != X64Const1 {
		t.Error("Round problem on positive numbers", x)
	}
	x = X64Round((-x64Const5) >> 2)
	if x != -X64Const1 {
		t.Error("Round problem on negative numbers", x)
	}
	x = X64Floor(X64Const0)
	if x != X64Const0 {
		t.Error("Floor problem on zero", x)
	}
	x = X64Floor(x64Const5 >> 2)
	if x != X64Const1 {
		t.Error("Floor problem on positive numbers", x)
	}
	x = X64Floor((-x64Const5) >> 2)
	if x != -x64Const2 {
		t.Error("Floor problem on negative number", x)
	}
	x = X64Ceil(X64Const0)
	if x != X64Const0 {
		t.Error("Ceil problem on zero", x)
	}
	x = X64Ceil(x64Const5 >> 2)
	if x != x64Const2 {
		t.Error("Ceil problem on positive numbers", x)
	}
	x = X64Ceil((-x64Const5) >> 2)
	if x != -X64Const1 {
		t.Error("Ceil problem on negative numbers", x)
	}
}

func TestX64Similar(t *testing.T) {
	var x X64
	var f float64

	if !X64IsSimilar(X64Const1, X64Const1) {
		t.Error("Can't figure out same xint64 is similar")
	}
	if X64IsSimilar(X64Const1, -X64Const1) {
		t.Error("Can't figure out different xint64 are not similar")
	}
	f = float64(X64Const1) * 1.00001
	x = X64(int64(f))
	if X64Const1 == x {
		t.Error("Similar values should not be the same", int64(X64Const1), x)
	}
	if !X64IsSimilar(X64Const1, x) {
		t.Error("Can't figure out xint64 is similar to 1", x)
	}
}

func TestX64Lerp(t *testing.T) {
	var x X64
	var x1 = F64ToX64(-2.0)
	var x2 = F64ToX64(8.0)
	var beta = F64ToX64(0.7)
	var lerp = F64ToX64(5.0)

	x = X64Lerp(x1, x2, beta)
	if !X64IsSimilar(x, lerp) {
		t.Errorf("bad lerp, got %x should be %x", x, lerp)
	}
	x = X64Lerp(x1, x2, -X64Const1)
	if x != x1 {
		t.Errorf("bad lerp on negative beta, got %x should be %x", x, x1)
	}
	x = X64Lerp(x1, x2, X64Const1+X64Const1)
	if x != x2 {
		t.Errorf("bad lerp on beta>1, got %x should be %x", x, x2)
	}
}

func TestX64Vec3JSON(t *testing.T) {
	x1 := F64ToX64(0.5)
	x2 := X64Const1

	var err error
	var jsonBuf []byte

	jsonBuf, err = x1.MarshalJSON()
	t.Logf("x1=%s x2=%s", x1.String(), x2.String())
	if err == nil {
		t.Logf("encoded JSON for X64 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X64")
	}
	err = x2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X64, but json is not correct")
	}
	err = x2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X64")
	}
	if !X64IsSimilar(x1, x2) {
		t.Error("unmarshalled vector is different from original")
	}
}

func BenchmarkX64Add(b *testing.B) {
	var x = x64Const2
	var y = x64Const5

	for i := 0; i < b.N; i++ {
		x = x + y
	}
}

func BenchmarkX64Sub(b *testing.B) {
	var x = x64Const2
	var y = x64Const5

	for i := 0; i < b.N; i++ {
		x = x - y
	}
}
func BenchmarkX64Mul(b *testing.B) {
	var x = x64Const2
	var y = x64Const5

	for i := 0; i < b.N; i++ {
		x = X64Mul(x, y)
	}
}

func BenchmarkX64Div(b *testing.B) {
	var x = x64Const2
	var y = x64Const5

	for i := 0; i < b.N; i++ {
		x = X64Div(x, y)
	}
}

func BenchmarkX64Mulp(b *testing.B) {
	var x = x64Const2
	var y = x64Const5

	for i := 0; i < b.N; i++ {
		x = X64Mulp(x, y)
	}
}

func BenchmarkX64Divp(b *testing.B) {
	var x = x64Const2
	var y = x64Const5

	for i := 0; i < b.N; i++ {
		x = X64Divp(x, y)
	}
}

func BenchmarkX4Exponent(b *testing.B) {
	var x X64 = -19

	for i := 0; i < b.N; i++ {
		_ = X64Exponent(x)
		x *= x
	}
}
