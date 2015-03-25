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

func TestX32Convert(t *testing.T) {
	var i32 int32
	var i64 int64
	var x64 X64
	var f32 float32
	var f64 float64

	i32 = X32ToI32(X32Const1)
	if i32 != I32Const1 {
		t.Error("can't convert positive X32 to int32", i32)
	}
	i32 = X32ToI32(-X32Const1)
	if i32 != -I32Const1 {
		t.Error("can't convert negative X32 to int32", i32)
	}
	i64 = X32ToI64(X32Const1)
	if i64 != I64Const1 {
		t.Error("can't convert positive X32 to int64", i64)
	}
	i64 = X32ToI64(-X32Const1)
	if i64 != -I64Const1 {
		t.Error("can't convert negative X32 to int64", i64)
	}
	x64 = X32ToX64(X32Const1)
	if x64 != X64Const1 {
		t.Error("can't convert positive X32 to X64", x64)
	}
	x64 = X32ToX64(-X32Const1)
	if x64 != -X64Const1 {
		t.Error("can't convert negtive X32 to X64", x64)
	}
	f32 = X32ToF32(X32Const1)
	if f32 != F32Const1 {
		t.Error("can't convert positive X32 to float32", f32)
	}
	f32 = X32ToF32(-X32Const1)
	if f32 != -F32Const1 {
		t.Error("can't convert negative X32 to float32", f32)
	}
	f64 = X32ToF64(X32Const1)
	if f64 != F64Const1 {
		t.Error("can't convert positive X32 to float64", f64)
	}
	f64 = X32ToF64(-X32Const1)
	if f64 != -F64Const1 {
		t.Error("can't convert negative X32 to float64", f64)
	}
}

func TestX32Math(t *testing.T) {
	var x X32

	x = X32Min(-x32Const5, -X32Const1)
	if x != -x32Const5 {
		t.Error("can't find min of xint32 numbers", x)
	}
	x = X32Max(x32Const2, -x32Const3)
	if x != x32Const2 {
		t.Error("can't find max of xint32 numbers", x)
	}
	x = X32Abs(-x32Const2)
	if x != x32Const2 {
		t.Error("can't get abs of negative xint32 number", x)
	}
	x = X32Abs(x32Const3)
	if x != x32Const3 {
		t.Error("can't get abs of positive xint32 number", x)
	}
	x = X32Mul(x32Const2, x32Const3)
	if x != x32Const6 {
		t.Error("can't mul xint32 numbers", x)
	}
	x = X32Muln(x32Const2, x32Const3, x32Const4)
	if x != X32Mul(x32Const6, x32Const4) {
		t.Error("can't muln xint32 numbers", x)
	}
	x = X32Div(x32Const8, x32Const2)
	if x != x32Const4 {
		t.Error("can't div xint32 numbers", x)
	}
	x = X32Mulp(x32Const2, x32Const3)
	if x != x32Const6 {
		t.Error("can't mulp xint32 numbers", x)
	}
	x = X32Divp(x32Const8, x32Const2)
	if x != x32Const4 {
		t.Error("can't divp xint32 numbers", x)
	}
}

func TestX32Float(t *testing.T) {
	var e int
	var x X32

	e = X32Exponent(X32Const1)
	if e != 0 {
		t.Error("Exponent does not work on 1", e)
	}
	e = X32Exponent(X32Const0)
	if e != 0 {
		t.Error("Exponent does not work on 0", e)
	}
	e = X32Exponent(x32Const5)
	if e != 2 {
		t.Error("Exponent does not work on 5", e)
	}
	e = X32Exponent(-(x32Const5 >> 4))
	if e != -2 {
		t.Error("Exponent does not work on -5", e)
	}
	x = X32Mantis(X32Const1)
	if x != X32Const1 {
		t.Error("Mantis does not work on 1", x)
	}
	x = X32Mantis(x32Const5)
	if x != x32Const5>>2 {
		t.Error("Mantis does not work on 5", x)
	}
	x = X32Mantis(-(x32Const5 >> 4))
	if x != -x32Const5>>2 {
		t.Error("Mantis does not work on -5", x)
	}
	x = X32Round((x32Const5) >> 2)
	if x != X32Const1 {
		t.Error("Round problem on positive numbers", x)
	}
	x = X32Round((-x32Const5) >> 2)
	if x != -X32Const1 {
		t.Error("Round problem on negative numbers", x)
	}
	x = X32Floor(X32Const0)
	if x != X32Const0 {
		t.Error("Floor problem on zero", x)
	}
	x = X32Floor(x32Const5 >> 2)
	if x != X32Const1 {
		t.Error("Floor problem on positive numbers", x)
	}
	x = X32Floor((-x32Const5) >> 2)
	if x != -x32Const2 {
		t.Error("Floor problem on negative number", x)
	}
	x = X32Ceil(X32Const0)
	if x != X32Const0 {
		t.Error("Ceil problem on zero", x)
	}
	x = X32Ceil(x32Const5 >> 2)
	if x != x32Const2 {
		t.Error("Ceil problem on positive numbers", x)
	}
	x = X32Ceil((-x32Const5) >> 2)
	if x != -X32Const1 {
		t.Error("Ceil problem on negative numbers", x)
	}
}

func TestX32Similar(t *testing.T) {
	var x X32
	var f float64

	if !X32IsSimilar(X32Const1, X32Const1) {
		t.Error("Can't figure out same xint32 is similar")
	}
	if X32IsSimilar(X32Const1, -X32Const1) {
		t.Error("Can't figure out different xint32 are not similar")
	}
	f = float64(X32Const1) * 1.001
	x = X32(int32(f))
	if X32Const1 == x {
		t.Error("Similar values should not be the same", int32(X32Const1), x)
	}
	if !X32IsSimilar(X32Const1, x) {
		t.Error("Can't figure out xint32 is similar to 1", x)
	}
}

func TestX32Lerp(t *testing.T) {
	var x X32
	var x1 = F32ToX32(-2.0)
	var x2 = F32ToX32(8.0)
	var beta = F32ToX32(0.7)
	var lerp = F32ToX32(5.0)

	x = X32Lerp(x1, x2, beta)
	if !X32IsSimilar(x, lerp) {
		t.Errorf("bad lerp, got %x should be %x", x, lerp)
	}
	x = X32Lerp(x1, x2, -X32Const1)
	if x != x1 {
		t.Errorf("bad lerp on negative beta, got %x should be %x", x, x1)
	}
	x = X32Lerp(x1, x2, X32Const1+X32Const1)
	if x != x2 {
		t.Errorf("bad lerp on beta>1, got %x should be %x", x, x2)
	}
}

func TestX32Vec3JSON(t *testing.T) {
	x1 := F32ToX32(0.5)
	x2 := X32Const1

	var err error
	var jsonBuf []byte

	jsonBuf, err = x1.MarshalJSON()
	t.Logf("x1=%s x2=%s", x1.String(), x2.String())
	if err == nil {
		t.Logf("encoded JSON for X32 is \"%s\"", string(jsonBuf))
	} else {
		t.Error("unable to encode JSON for X32")
	}
	err = x2.UnmarshalJSON([]byte("nawak"))
	if err == nil {
		t.Error("able to decode JSON for X32, but json is not correct")
	}
	err = x2.UnmarshalJSON(jsonBuf)
	if err != nil {
		t.Error("unable to decode JSON for X32")
	}
	if !X32IsSimilar(x1, x2) {
		t.Error("unmarshalled vector is different from original")
	}
}

func BenchmarkX32Add(b *testing.B) {
	var x = x32Const2
	var y = x32Const5

	for i := 0; i < b.N; i++ {
		x = x + y
	}
}

func BenchmarkX32Sub(b *testing.B) {
	var x = x32Const2
	var y = x32Const5

	for i := 0; i < b.N; i++ {
		x = x - y
	}
}

func BenchmarkX32Mul(b *testing.B) {
	var x = x32Const2
	var y = x32Const5

	for i := 0; i < b.N; i++ {
		x = X32Mul(x, y)
	}
}

func BenchmarkX32Div(b *testing.B) {
	var x = x32Const2
	var y = x32Const5

	for i := 0; i < b.N; i++ {
		x = X32Div(x, y)
	}
}

func BenchmarkX32Mulp(b *testing.B) {
	var x = x32Const2
	var y = x32Const5

	for i := 0; i < b.N; i++ {
		x = X32Mulp(x, y)
	}
}

func BenchmarkX32Divp(b *testing.B) {
	var x = x32Const2
	var y = x32Const5

	for i := 0; i < b.N; i++ {
		x = X32Divp(x, y)
	}
}

func BenchmarkX32Exponent(b *testing.B) {
	var x X32 = 17

	for i := 0; i < b.N; i++ {
		_ = X32Exponent(x)
		x *= x
	}
}
