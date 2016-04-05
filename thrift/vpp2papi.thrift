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
namespace js vpp2papi

include "vpcommonapi.thrift"

/**
 * DefaultPort is the TCP port the service listens to, by default.
 */
const i32 DefaultPort = 8777;

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
  1: list<NodeInfo> Successors,
  2: list<NodeInfo> Ds,
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
  5: i32 CallTimeout,
  6: i32 SyncDelay,
  7: i32 DisconnectTimeout,
  8: i32 DataLifetime,
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
  1: HostInfo SourceHost,
  2: RingInfo SourceRing,
  3: NodeInfo SourceNode,
  4: binary TargetNodeID,
}

/**
 * Used to return data when calling status.
 */
struct HostStatus {
  1: HostInfo ThisHostInfo,
  2: list<NodeStatus> LocalNodeStatus,
  3: map<string,RingInfo> RingsRefs,
  4: map<string,HostInfo> HostsRefs,
}

/**
 * Used to store Lookup-like requests.
 */
struct LookupRequest {
    1:ContextInfo Context,
    2:binary Key,
    3:binary KeyShift,
    4:binary ImaginaryNode,
    5:binary Sig,
}

/**
 * Used to store results when doing Lookup-like requests.
 */
struct LookupResponse {
  1: bool Found,
  2: list<NodeInfo> NodesPath,
  3: map<string,HostInfo> HostsRefs,
}

/**
 * Used to store  GetSuccessors requests.
 */
struct GetSuccessorsRequest {
    1:ContextInfo Context,
    2:binary Sig,
}

/**
 * Used to store results when doing GetSuccessors requests.
 */
struct GetSuccessorsResponse {
  1: list<NodeInfo> SuccessorNodes,
  2: map<string,HostInfo> HostsRefs,
}

/**
 * Used to store  GetPredecessor requests.
 */
struct GetPredecessorRequest {
    1:ContextInfo Context,
    2:binary Sig,
}

/**
 * Used to store results when doing GetPredecessor requests.
 */
struct GetPredecessorResponse {
  1: NodeInfo PredecessorNode,
  2: map<string,HostInfo> HostsRefs,
}

/**
 * Used to store Sync requests.
 */
struct SyncRequest {
    1:ContextInfo Context,
    2:binary KeyShift,
    3:binary ImaginaryNode,
    4:binary Sig,
}

/**
 * Used to store results when doing Sync requests.
 */
struct SyncResponse {
  1: bool Found,
  2: list<NodeInfo> NodesPath,
  3: list<NodeInfo> SuccessorNodes,
  4: NodeInfo PredecessorNode,
  5: map<string,HostInfo> HostsRefs,
}

/**
 * VpP2pApi is used to communicate between 2 Vapor nodes
 * in peer-to-peer mode.
 */
service VpP2pApi extends vpcommonapi.VpCommonApi
{
  HostStatus Status(
  ),
  LookupResponse Lookup(
    1:LookupRequest request,
  ),
  GetSuccessorsResponse GetSuccessors(
    1:GetSuccessorsRequest request,
  ),
  GetPredecessorResponse GetPredecessor(
    1:GetPredecessorRequest request,
  ),
  SyncResponse Sync(
    1:SyncRequest request,
  ),
}
