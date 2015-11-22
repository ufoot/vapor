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

package vploop

import (
	"testing"
	"time"
)

type dummyHandler struct {
	t     *testing.T
	start time.Time
}

func (state dummyHandler) Duration() time.Duration {
	return time.Second / 10
}

func (state dummyHandler) Init(timestamp time.Time) error {
	state.t.Logf("dummy init")

	return nil
}

func (state dummyHandler) Do(timestamp time.Time, iteration int64, quit chan<- bool) {
	if state.start.Unix()+3 > time.Now().Unix() {
		state.t.Logf("dummy loop iteration=%d", iteration)
	} else {
		state.t.Logf("dummy end iteration=%d 1/2", iteration)
		quit <- true
		state.t.Logf("dummy end iteration=%d 2/2", iteration)
	}
}

func (state dummyHandler) Quit(timestamp time.Time) {
	state.t.Logf("dummy quit")
}

func TestMainLoop(t *testing.T) {
	dh := dummyHandler{t, time.Now()}

	t.Log("main loop BEGIN")
	MainLoop(dh)
	t.Log("main loop END")
}
