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

package vpsys

import (
	"errors"
	"testing"
)

func TestErrorChain(t *testing.T) {
	err1 := errors.New("foo")
	err2 := ErrorChain(err1, "bar")
	if err2.Error() == "bar (\"foo\")" {
		t.Logf("OK, err1=%s err2=%s", err1.Error(), err2.Error())
	} else {
		t.Errorf("Inconsistent errors, err1=%s err2=%s", err1.Error(), err2.Error())
	}
	err3 := ErrorChain(err1, "")
	if err3.Error() == "foo" {
		t.Logf("OK, err1=%s err3=%s", err1.Error(), err3.Error())
	} else {
		t.Errorf("Inconsistent errors, err1=%s err3=%s", err1.Error(), err3.Error())
	}
	err4 := ErrorChain(nil, "")
	if err4 == nil {
		t.Logf("OK, err4 is nil")
	} else {
		t.Errorf("Inconsistent errors, err4=%s", err4.Error())
	}
	err5 := ErrorChainf(err1, "bar %d", 1)
	if err5.Error() == "bar 1 (\"foo\")" {
		t.Logf("OK, err1=%s err5=%s", err1.Error(), err5.Error())
	} else {
		t.Errorf("Inconsistent errors, err1=%s err5=%s", err1.Error(), err5.Error())
	}
}
