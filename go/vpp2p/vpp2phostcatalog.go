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

// HostCatalog is structure used to contain locally-known hosts.
type HostCatalog struct {
	access sync.RWMutex
	hosts  map[[vpp2pdat.HostPubKeyBufNbBytes]byte]*vpp2papi.HostInfo
}

var globalHostCatalog = NewHostCatalog()

// NewHostCatalog creates a new instance of a local host catalog
func NewHostCatalog() *HostCatalog {
	return &HostCatalog{hosts: make(map[[vpp2pdat.HostPubKeyBufNbBytes]byte]*vpp2papi.HostInfo)}
}

// GlobalHostCatalog returns a catalog containing all local hosts.
func GlobalHostCatalog() *HostCatalog {
	return globalHostCatalog
}

// HasHost returns true if the host exists in the catalog.
// It's thread-safe.
func (c *HostCatalog) HasHost(hostPubKey []byte) bool {
	return (c.GetHost(hostPubKey) != nil)
}

// GetHost returns a handler which makes possible API calls on it.
// It's thread-safe.
func (c *HostCatalog) GetHost(hostPubKey []byte) *vpp2papi.HostInfo {
	hostPubKeyBuf := vpp2pdat.HostPubKeyToBuf(hostPubKey)

	defer c.access.RUnlock()
	c.access.RLock()

	n := c.hosts[hostPubKeyBuf]

	return n
}

// RegisterHost registers a host within the catalog.
// It's thread-safe.
func (c *HostCatalog) RegisterHost(host *vpp2papi.HostInfo) {
	hostPubKey := host.HostPubKey
	hostPubKeyBuf := vpp2pdat.HostPubKeyToBuf(hostPubKey)

	defer c.access.Unlock()
	c.access.Lock()

	c.hosts[hostPubKeyBuf] = host
}

// UnregisterHost unregisters a host within the catalog.
// It's thread-safe.
func (c *HostCatalog) UnregisterHost(host *vpp2papi.HostInfo) {
	hostPubKey := host.HostPubKey
	hostPubKeyBuf := vpp2pdat.HostPubKeyToBuf(hostPubKey)

	defer c.access.Unlock()
	c.access.Lock()

	delete(c.hosts, hostPubKeyBuf)
}

// List returns a list of local hosts. It returns static
// data about the host, host the hosts themselves.
func (c *HostCatalog) List() []*vpp2papi.HostInfo {
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
