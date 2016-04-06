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

package vpp2p

import (
	"fmt"
	"github.com/ufoot/vapor/go/vpapp"
	"github.com/ufoot/vapor/go/vpcommonapi"
	"github.com/ufoot/vapor/go/vpcrypto"
	"github.com/ufoot/vapor/go/vpp2papi"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"github.com/ufoot/vapor/go/vprand"
	"github.com/ufoot/vapor/go/vpsum"
	"time"
)

// Host is a physical host, it is used to uniquely identify
// a host, it can be used to handle several apps or rings.
type Host struct {
	// Info about the host
	Info vpp2papi.HostInfo

	key              *vpcrypto.Key
	localNodeCatalog *NodeCatalog
	startTime        time.Time
}

// NewHost returns a new host object
func NewHost(title, url string, useSig bool) (*Host, error) {
	var ret Host
	var pubKey []byte
	var sig []byte
	var err error
	var ok bool

	ret.Info = vpp2papi.HostInfo{HostTitle: title, HostURL: url, HostPubKey: nil, HostSig: nil}

	if useSig {
		ret.key, err = vpcrypto.NewKey()
		if err != nil {
			return nil, err
		}
		pubKey, err = ret.key.ExportPub()
		if err != nil {
			return nil, err
		}
		ok, err = vpp2pdat.CheckPubKey(pubKey)
		if err != nil || !ok {
			return nil, err
		}
		sig, err = ret.key.Sign(vpp2pdat.HostInfoSigBytes(&(ret.Info)))
		if err != nil {
			return nil, err
		}
		ok, err = vpp2pdat.CheckSig(pubKey)
		if err != nil || !ok {
			return nil, err
		}
	} else {
		pubKey = vpsum.IntToBuf128(vprand.Rand128(nil, nil))
		sig = []byte("")
	}

	ret.Info.HostPubKey = pubKey
	ret.Info.HostSig = sig

	_, err = vpp2pdat.CheckHostInfo(&(ret.Info))
	if err != nil {
		return nil, err
	}

	ret.localNodeCatalog = NewNodeCatalog()
	ret.startTime = time.Now()

	return &ret, nil
}

// CanSign returns true if the host has a key it can sign with.
func (host *Host) CanSign() bool {
	return host.key != nil
}

// IsSigned returns true if the has been self-signed.
// It does not check if the signature is valid.
func (host *Host) IsSigned() bool {
	return vpp2pdat.HostInfoIsSigned(&host.Info)
}

// CheckSig checks if the host signature is OK, if it's not, returns false and an error.
func (host *Host) CheckSig() (bool, error) {
	return vpp2pdat.HostInfoCheckSig(&host.Info)
}

// Ping is a simple ping function
func (host *Host) Ping() error {
	return nil
}

// Uptime is a simple uptime function, returns time since host was created.
func (host *Host) Uptime() (int64, error) {
	return time.Now().Unix() - host.startTime.Unix(), nil
}

// GetPackage returns the package version. Program general information.
func (host *Host) GetPackage() (*vpcommonapi.Package, error) {
	return vpapp.DefaultPackage(), nil
}

// GetVersion returns the version version. Program general information.
func (host *Host) GetVersion() (*vpcommonapi.Version, error) {
	return vpapp.DefaultVersion(), nil
}

// Status is called to get another host status.
func (host *Host) Status() (*vpp2papi.HostStatus, error) {
	ret := vpp2papi.NewHostStatus()

	// todo...

	return ret, nil
}

// Lookup searches for a key on a given ring.
func (host *Host) Lookup(request *vpp2papi.LookupRequest) (*vpp2papi.LookupResponse, error) {
	_, err := vpp2pdat.CheckContextInfo(request.Context)
	if err != nil {
		return nil, err
	}

	node := host.localNodeCatalog.GetNode(request.Context.TargetNodeID)
	if node == nil {
		return nil, fmt.Errorf("unable to find node locally")
	}

	ret := vpp2papi.NewLookupResponse()
	ret.Found, ret.NodesPath, err = node.Lookup(request.Key, request.KeyShift, request.ImaginaryNode)
	if err != nil {
		return nil, err
	}
	if ret.Found && ret.NodesPath != nil {
		ret.HostsRefs = GlobalHostCatalog().CreateHostsRefs(&(host.Info), nil, ret.NodesPath)
	}

	return ret, nil
}

// GetSuccessors is called to retrieve successors of a node.
func (host *Host) GetSuccessors(request *vpp2papi.GetSuccessorsRequest) (*vpp2papi.GetSuccessorsResponse, error) {
	_, err := vpp2pdat.CheckContextInfo(request.Context)
	if err != nil {
		return nil, err
	}

	node := host.localNodeCatalog.GetNode(request.Context.TargetNodeID)
	if node == nil {
		return nil, fmt.Errorf("unable to find node locally")
	}

	ret := vpp2papi.NewGetSuccessorsResponse()

	ret.SuccessorNodes = node.GetSuccessors()
	ret.HostsRefs = GlobalHostCatalog().CreateHostsRefs(&(host.Info), nil, ret.SuccessorNodes)

	return ret, nil
}

// GetPredecessor is called to retrieve the predecessor of a node.
func (host *Host) GetPredecessor(request *vpp2papi.GetPredecessorRequest) (*vpp2papi.GetPredecessorResponse, error) {
	var hostsRefs []*vpp2papi.NodeInfo

	_, err := vpp2pdat.CheckContextInfo(request.Context)
	if err != nil {
		return nil, err
	}

	node := host.localNodeCatalog.GetNode(request.Context.TargetNodeID)
	if node == nil {
		return nil, fmt.Errorf("unable to find node locally")
	}

	ret := vpp2papi.NewGetPredecessorResponse()

	ret.PredecessorNode = node.GetPredecessor()
	if ret.PredecessorNode != nil {
		hostsRefs := make([]*vpp2papi.NodeInfo, 1)
		hostsRefs[0] = ret.PredecessorNode
	}

	ret.HostsRefs = GlobalHostCatalog().CreateHostsRefs(&(host.Info), nil, hostsRefs)

	return ret, nil
}

// Sync is called to connect and synchronize on a host. Basically, it does
// a lookup and returns successors and predecessor.
func (host *Host) Sync(request *vpp2papi.SyncRequest) (*vpp2papi.SyncResponse, error) {
	var hostsRefs []*vpp2papi.NodeInfo

	_, err := vpp2pdat.CheckContextInfo(request.Context)
	if err != nil {
		return nil, err
	}

	node := host.localNodeCatalog.GetNode(request.Context.TargetNodeID)
	if node == nil {
		return nil, fmt.Errorf("unable to find node locally")
	}

	ret := vpp2papi.NewSyncResponse()
	ret.Found, ret.NodesPath, ret.SuccessorNodes, ret.PredecessorNode, err = node.Sync(request.Context.SourceNode, request.Context.SourceNode.NodeID, request.KeyShift, request.ImaginaryNode)
	if err != nil {
		return nil, err
	}
	if ret.NodesPath != nil && ret.SuccessorNodes != nil && ret.PredecessorNode != nil {
		hostsRefs = make([]*vpp2papi.NodeInfo, len(ret.NodesPath)+len(ret.SuccessorNodes)+1)
		for i, v := range ret.NodesPath {
			hostsRefs[i] = v
		}
		for i, v := range ret.SuccessorNodes {
			hostsRefs[i+len(ret.NodesPath)] = v
		}
		hostsRefs[len(ret.NodesPath)+len(ret.SuccessorNodes)] = ret.PredecessorNode
	}
	if ret.Found && ret.NodesPath != nil && ret.SuccessorNodes != nil && ret.PredecessorNode != nil {
		ret.HostsRefs = GlobalHostCatalog().CreateHostsRefs(&(host.Info), nil, hostsRefs)
	}

	return ret, nil
}
