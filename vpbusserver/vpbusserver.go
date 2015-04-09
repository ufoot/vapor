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
	"github.com/ufoot/vapor/vpbus"
	"github.com/ufoot/vapor/vpbusapi"
	"github.com/ufoot/vapor/vpsys"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) (*thrift.TSimpleServer, error) {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return nil, vpsys.ErrorChain(err, "unable to create server socket")
	}
	vpsys.LogNoticef("%T\n", transport)
	handler := vpbus.New()
	processor := vpbusapi.NewVpBusApiProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	vpsys.LogNoticef("Starting Thrift server on %s\n", addr)
	return server, server.Serve()
}

// RunDefault runs a server with default parameters (for testing purposes).
func RunDefault() (*thrift.TSimpleServer, error) {
	return runServer(thrift.NewTTransportFactory(), thrift.NewTBinaryProtocolFactoryDefault(), "127.0.0.1:9090")
}
