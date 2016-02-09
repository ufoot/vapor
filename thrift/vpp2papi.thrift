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


# VpP2pApi
# Christian Mauduit (ufoot@ufoot.org)
#
# Thrift protocol between Golang backend server and C++ GUI client.

# http://diwakergupta.github.io/thrift-missing-guide/

namespace cpp vpp2papi
namespace go vpp2papi
namespace php vpp2papi

include "vpcommonapi.thrift"

/**
 * NodeInfo contains static informations about a node.
 */
struct NodeInfo {
  1: binary NodeID,
  2: binary HostPubKey,
  3: binary RingID,
}

/**
 * Used to store results when doing Lookup-like requests.
 */
struct LookupData {
  1: list<NodeInfo> path,
  2: i32 errcode,
}

/**
 * Used to store results when doing Get-like requests.
 */
struct GetData {
  1: binary value,
  2: list<NodeInfo> path,
  3: i32 errcode,
}

/**
 * VpP2pApi is used to communicate between 2 Vapor nodes
 * in peer-to-peer mode.
 */
service VpP2pApi extends vpcommonapi.VpCommonApi
{
  LookupData Lookup(
    1:binary key,
    2:binary keyShift,
    3:binary imaginaryNode
  ),
}
