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


# VpBusApi
# Christian Mauduit (ufoot@ufoot.org)
#
# Thrift protocol between Golang backend server and C++ GUI client.

namespace cpp vpbusapi
namespace go vpbusapi
namespace js vpbusapi

include "vpcommonapi.thrift"

/**
 * DefaultPort is the TCP port the service listens to, by default.
 */
const i32 DefaultPort = 7888;

/**
 * VpBusApi is used to communicate between Vapor and Fumes.
 * Vapor is the Golang server and Fumes the C++ client.
 */
service VpBusApi extends vpcommonapi.VpCommonApi
{
  /**
   * Halt stops the server.
   */
  oneway void halt (
  )
}
