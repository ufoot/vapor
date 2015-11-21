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

package vpbusserver

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	var server *thrift.TSimpleServer
	var err error

	server, err = NewDefault()
	if err == nil {
		t.Log("createdThrift server")
	} else {
		t.Error("unable to create Thrift server", err)
	}

	if server != nil {
		start := time.Now()
		if AsyncServe(server) == nil {
			for start.Unix()+3 > time.Now().Unix() {
				t.Log("letting server run for some time")
				time.Sleep(time.Second)
			}
		} else {
			t.Error("unable to start Thrift server", err)
		}
	}

	if server != nil {
		t.Log("stopping vpbusserver")
		server.Stop()
	} else {
		t.Error("server not started")
	}
}
