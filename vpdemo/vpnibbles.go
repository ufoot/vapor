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

package main

import (
	"github.com/ufoot/vapor/vpsys"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

// NibblesState stores the game state.
type NibblesState struct {
	level [][]int
}

// Duration returns the duration of an iteration, time between two Do calls.
func (state NibblesState) Duration() time.Duration {
	return time.Second / 10
}

// Init initializes the game state.
func (state NibblesState) Init(timestamp time.Time) {
	vpsys.LogNoticef("game init")
}

// Do process stuff on a game state, typically called in game loop
// when receiving events.
func (state NibblesState) Do(timestamp time.Time, iteration int64, quit chan<- bool) {
	if start.Unix()+3 > time.Now().Unix() {
		vpsys.LogNoticef("game loop iteration=%d", iteration)
	} else {
		vpsys.LogNoticef("game end iteration=%d 1/2", iteration)
		quit <- true
		vpsys.LogNoticef("game end iteration=%d 2/2", iteration)
	}
}

// Quit should be called at the end of a game.
func (state NibblesState) Quit(timestamp time.Time) {
	vpsys.LogNoticef("game quit")
}
