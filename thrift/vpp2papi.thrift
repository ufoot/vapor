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


# VpP2pApi
# Christian Mauduit (ufoot@ufoot.org)
#
# Thrift protocol between Golang backend server and C++ GUI client.

namespace cpp vpp2papi
namespace go vpp2papi
namespace php vpp2papi

include "vpcommonapi.thrift"

/**
 * VpP2pApi is used to communicate between 2 Vapor nodes
 * in peer-to-peer mode.
 */
service VpP2pApi extends vpcommonapi.VpCommonApi
{
  /**
   * Halt stops the server.
   */
  oneway void halt ()
}
