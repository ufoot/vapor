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
 * HostInfo contains static informations about a host.
 */
struct HostInfo {
  1: string HostTitle,
  2: string HostURL,
  3: binary HostPubKey,
  4: binary HostSig,
}

/**
 * NodeInfo contains static informations about a node.
 */
struct NodeInfo {
  1: binary NodeID,
  2: binary RingID,
  3: binary HostPubKey,
  4: binary NodeSig,
}

/**
 * RingConfig contains functional parameters of a ring.
 */
struct RingConfig {
  1: i32 BruijnM,
  2: i32 BruijnN,
  3: i32 NbCopy,
  4: i32 NbStep,
}

/**
 * RingInfo contains static informations about a ring.
 */
struct RingInfo {
  1: binary RingID,
  2: string RingTitle,
  3: string RingDescription,
  4: binary AppID,
  5: RingConfig Config,
  6: bool HasPassword,
  7: binary HostPubKey,
  8: binary RingSig,
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
  LookupData IAmPrev(
    1:binary key,
    2:binary keyShift,
    3:binary imaginaryNode
  ),
  LookupData Lookup(
    1:binary key,
    2:binary keyShift,
    3:binary imaginaryNode
  ),
}
