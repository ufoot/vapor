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
	"github.com/ufoot/vapor/go/vpp2papi"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"sync"
)

// HostInfoCatalog is structure used to contain locally-known hosts.
type HostInfoCatalog struct {
	access sync.RWMutex
	hosts  map[[vpp2pdat.HostPubKeyBufNbBytes]byte]*vpp2papi.HostInfo
}

type HostsRefsCreator interface {
	CreateHostsRefs(locallHost *vpp2papi.HostInfo, rings []*vpp2papi.RingInfo, nodes []*vpp2papi.NodeInfo) map[string]*vpp2papi.HostInfo
}

var globalHostInfoCatalog = NewHostInfoCatalog()

// NewHostInfoCatalog creates a new instance of a local host catalog
func NewHostInfoCatalog() *HostInfoCatalog {
	return &HostInfoCatalog{hosts: make(map[[vpp2pdat.HostPubKeyBufNbBytes]byte]*vpp2papi.HostInfo)}
}

// GlobalHostInfoCatalog returns a catalog containing all local hosts.
func GlobalHostInfoCatalog() *HostInfoCatalog {
	return globalHostInfoCatalog
}

// HasHost returns true if the host exists in the catalog.
// It's thread-safe.
func (c *HostInfoCatalog) HasHost(hostPubKey []byte) bool {
	return (c.GetHost(hostPubKey) != nil)
}

// GetHost returns a handler which makes possible API calls on it.
// It's thread-safe.
func (c *HostInfoCatalog) GetHost(hostPubKey []byte) *vpp2papi.HostInfo {
	hostPubKeyBuf := vpp2pdat.HostPubKeyToBuf(hostPubKey)

	defer c.access.RUnlock()
	c.access.RLock()

	n := c.hosts[hostPubKeyBuf]

	return n
}

// RegisterHost registers a host within the catalog.
// It's thread-safe.
func (c *HostInfoCatalog) RegisterHost(host *vpp2papi.HostInfo) {
	hostIDBuf := vpp2pdat.HostPubKeyToBuf(host.HostPubKey)

	defer c.access.Unlock()
	c.access.Lock()

	c.hosts[hostIDBuf] = host
}

// UnregisterHost unregisters a host within the catalog.
// It's thread-safe.
func (c *HostInfoCatalog) UnregisterHost(host *vpp2papi.HostInfo) {
	hostIDBuf := vpp2pdat.HostPubKeyToBuf(host.HostPubKey)

	defer c.access.Unlock()
	c.access.Lock()

	delete(c.hosts, hostIDBuf)
}

// CreateHostsRefs returns a list of known hosts for a given set
// of nodes and rings. The key index is the short string,
// there might be collisions but this should not be a
// major problem as internally, full keys are used.
func (c *HostInfoCatalog) CreateHostsRefs(localHost *vpp2papi.HostInfo, rings []*vpp2papi.RingInfo, nodes []*vpp2papi.NodeInfo) map[string]*vpp2papi.HostInfo {
	ret := make(map[string]*vpp2papi.HostInfo)

	defer c.access.RUnlock()
	c.access.RLock()

	if nodes != nil {
		for _, value := range nodes {
			if value != nil && value.HostPubKey != nil {
				hostPubKeyBuf := vpp2pdat.HostPubKeyToBuf(value.HostPubKey)
				host := c.hosts[hostPubKeyBuf]
				if host != nil {
					ret[vpp2pdat.HostPubKeyToShortString(value.HostPubKey)] = host
				}
			}
		}
	}
	if rings != nil {
		for _, value := range rings {
			if value != nil && value.HostPubKey != nil {
				hostPubKeyBuf := vpp2pdat.HostPubKeyToBuf(value.HostPubKey)
				host := c.hosts[hostPubKeyBuf]
				if host != nil {
					ret[vpp2pdat.HostPubKeyToShortString(value.HostPubKey)] = host
				}
			}
		}
	}
	// add the host after everything else, it's typically us, the local
	// host, and we want this to override everything else
	if localHost != nil {
		ret[vpp2pdat.HostPubKeyToShortString(localHost.HostPubKey)] = localHost
	}

	return ret
}

// HasAllHostsRefs returns true if all the refs given are already known
func (c *HostInfoCatalog) HasAllHostsRefs(refs map[string]*vpp2papi.HostInfo) bool {
	ret := true

	defer c.access.RUnlock()
	c.access.RLock()

	for _, value := range refs {
		if c.hosts[vpp2pdat.HostPubKeyToBuf(value.HostPubKey)] == nil {
			ret = false
		}
	}

	return ret
}

// UpdateHostsRefs updates the host catalog according to a set of refs.
func (c *HostInfoCatalog) UpdateHostsRefs(refs map[string]*vpp2papi.HostInfo) {
	if c.HasAllHostsRefs(refs) {
		return
	}

	defer c.access.Unlock()
	c.access.Lock()

	for _, value := range refs {
		c.hosts[vpp2pdat.HostPubKeyToBuf(value.HostPubKey)] = value
	}
}

// List returns a list of known hosts. It returns static
// data about the host, not the hosts themselves.
func (c *HostInfoCatalog) List() []*vpp2papi.HostInfo {
	defer c.access.RUnlock()
	c.access.RLock()

	i := 0
	ret := make([]*vpp2papi.HostInfo, len(c.hosts))
	for _, value := range c.hosts {
		ret[i] = value
		i++
	}

	return ret
}
