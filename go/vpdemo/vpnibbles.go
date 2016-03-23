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

package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/ufoot/vapor/go/vpbussrv"
	"github.com/ufoot/vapor/go/vperror"
	"github.com/ufoot/vapor/go/vplog"
	"time"
)

// NibblesState stores the game state.
type NibblesState struct {
	start time.Time
	level [][]int
}

// Duration returns the duration of an iteration, time between two Do calls.
func (state NibblesState) Duration() time.Duration {
	return time.Second / 10
}

// Init initializes the game state.
func (state NibblesState) Init(timestamp time.Time) error {
	vplog.LogNoticef("game state init")
	state.start = time.Now()

	return nil
}

// Do process stuff on a game state, typically called in game loop
// when receiving events.
func (state NibblesState) Do(timestamp time.Time, iteration int64, quit chan<- bool) {
	if state.start.Unix()+3 > time.Now().Unix() {
		vplog.LogNoticef("game state loop iteration=%d", iteration)
	} else {
		vplog.LogNoticef("game state end iteration=%d 1/2", iteration)
		quit <- true
		vplog.LogNoticef("game state end iteration=%d 2/2", iteration)
	}
}

// Quit should be called at the end of a game.
func (state NibblesState) Quit(timestamp time.Time) {
	vplog.LogNoticef("game state quit")
}

// NibblesServer stores the game server.
type NibblesServer struct {
	start  time.Time
	server *thrift.TSimpleServer
}

// Duration returns the duration of an iteration, time between two Do calls.
func (server NibblesServer) Duration() time.Duration {
	return time.Second / 2
}

// Init initializes the game server.
func (server NibblesServer) Init(timestamp time.Time) error {
	vplog.LogNoticef("game server init")
	var err error

	server.server, err = vpbussrv.NewDefault()
	if err != nil {
		return vperror.Chain(err, "unable to create vpbussrv Thrift server")
	}

	err = vpbussrv.AsyncServe(server.server)
	if err != nil {
		return vperror.Chain(err, "unable to start vpbussrv Thrift server")
	}

	return nil
}

// Do process stuff on a game server, typically called in game loop
// when receiving events.
func (server NibblesServer) Do(timestamp time.Time, iteration int64, quit chan<- bool) {
	if server.start.Unix()+3 > time.Now().Unix() {
		vplog.LogNoticef("game server loop iteration=%d", iteration)
	} else {
		vplog.LogNoticef("game server end iteration=%d 1/2", iteration)
		quit <- true
		vplog.LogNoticef("game server end iteration=%d 2/2", iteration)
	}
}

// Quit should be called at the end of a game.
func (server NibblesServer) Quit(timestamp time.Time) {
	vplog.LogNoticef("game server quit")
	if server.server != nil {
		server.server.Stop()
	}
	server.server = nil
}
