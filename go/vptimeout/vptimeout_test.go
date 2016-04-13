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

package vptimeout

import (
	"fmt"
	"testing"
	"time"
)

const (
	baseDuration = 1500 * time.Millisecond
)

func TestRun(t *testing.T) {
	okF := func() error {
		t.Logf("inside okF")
		return nil
	}
	errF := func() error {
		t.Logf("inside errF")
		return fmt.Errorf("this is an error")
	}
	timeoutF := func() error {
		t.Logf("inside timeoutF (begin)")
		time.Sleep(2 * baseDuration)
		t.Logf("inside timeoutF (end)")
		return nil
	}
	var err error

	err = Run(okF, baseDuration)
	if err != nil {
		t.Errorf("error running okF within %s", baseDuration.String())
	}
	err = Run(errF, baseDuration)
	if err == nil {
		t.Errorf("no error (?) running errF within %s", baseDuration.String())
	}
	err = Run(timeoutF, baseDuration)
	if err == nil {
		t.Errorf("no error (?) running timeoutF within %s", baseDuration.String())
	}
}
