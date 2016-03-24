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
 * NodePeers contains informations about node peers,
 * those it should contact.
 */
struct NodePeers {
  1: list<NodeInfo> Successor,
  2: list<NodeInfo> D,
}

/**
 * NodeStatus contains details about a node.
 */
struct NodeStatus {
  1: NodeInfo Info,
  2: NodePeers Peers,
  3: NodeInfo Predecessor,
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
 * ContextInfo contains static informations about the program
 * calling a fonction, it gives context.
 */
struct ContextInfo {
  1: HostInfo sourceHost,
  2: RingInfo sourceRing,
  3: NodeInfo sourceNode,
  4: binary targetNodeID,
}

/**
 * Used to return data when calling status.
 */
struct HostStatus {
  1: HostInfo thisHostInfo,
  2: list<NodeStatus> localNodeStatus,
  3: list<RingInfo> ringsRef,
  4: list<HostInfo> hostsRef,
}

/**
 * Used to store results when doing Lookup-like requests.
 */
struct LookupData {
  1: bool found,
  2: list<NodeInfo> nodesPath,
  3: list<HostInfo> hostsRef,
}

/**
 * Used to store results when doing Sync requests.
 */
struct SyncData {
  1: bool accepted,
  2: list<NodeInfo> successorNodes,
  3: list<HostInfo> hostsRef,
}

/**
 * VpP2pApi is used to communicate between 2 Vapor nodes
 * in peer-to-peer mode.
 */
service VpP2pApi extends vpcommonapi.VpCommonApi
{
  HostStatus Status(
  ),
  SyncData Sync(
    1:ContextInfo context,
  ),
  LookupData Lookup(
    1:ContextInfo context,
    2:binary key,
    3:binary keyShift,
    4:binary imaginaryNode
  ),
}
