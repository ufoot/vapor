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

package vpp2psim

//go:generate bash ./stamp.sh
	
import (
	"fmt"
	"github.com/ufoot/vapor/go/vpp2p"
	"github.com/ufoot/vapor/go/vpp2pdat"
	"github.com/ufoot/vapor/go/vprand"
	"github.com/ufoot/vapor/go/vpsum"
)

const (
	// MaxNbHosts is a global limit of how many local hosts the
	// local simulation can use.
	MaxNbHosts = 10000
	// MaxNbRings is a global limit of how many local rings
	// local simulation can use, per host.
	MaxNbRings = 100
	// MaxNbNodesPerHostRing is a global limit of how many local nodes
	// local simulation can use, per (host,ring) pair.
	MaxNbNodesPerHostRing = 10000
	// MaxNbNodes is a global limit of how many local nodes
	// local simulation can use, per host.
	MaxNbNodes = 1000000
)

// SetupLocal sets up a number of local hosts with some nodes per hosts.
func SetupLocal(nbHosts, nbRings, nbNodesPerHostRing int, useSig bool) ([]*vpp2p.Host, []*vpp2p.Ring, []*vpp2p.Node, error) {
	var err error
	seed := vpsum.IntToStr32(vprand.Rand32(nil, 1000000000))

	if nbHosts < 1 || nbHosts > MaxNbHosts {
		return nil, nil, nil, fmt.Errorf("Bad nbHosts=%d, range is [1,%d]", nbHosts, MaxNbHosts)
	}
	if nbRings < 1 || nbRings > MaxNbRings {
		return nil, nil, nil, fmt.Errorf("Bad nbRings=%d, range is [1,%d]", nbRings, MaxNbRings)
	}
	if nbNodesPerHostRing < 1 || nbNodesPerHostRing > MaxNbNodesPerHostRing {
		return nil, nil, nil, fmt.Errorf("Bad nbNodesPerHostRing=%d, range is [1,%d]", nbNodesPerHostRing, MaxNbNodesPerHostRing)
	}
	nbNodes := nbHosts * nbRings * nbNodesPerHostRing
	if nbNodes < 1 || nbNodes > MaxNbNodes {
		return nil, nil, nil, fmt.Errorf("Bad nbNodes=%d, range is [1,%d]", nbNodes, MaxNbNodes)
	}

	hosts := make([]*vpp2p.Host, nbHosts)

	for i := range hosts {
		hosts[i], err = vpp2p.NewHost(fmt.Sprintf("Host %s/%d", seed, i),
			fmt.Sprintf("http://localhost:%04d/%s", 8080+i, seed),
			useSig, vpp2p.GlobalHostInfoCatalog())
		if err != nil {
			return nil, nil, nil, err
		}
	}

	rings := make([]*vpp2p.Ring, nbRings)

	for i := range rings {
		rings[i], err = vpp2p.NewRing(hosts[0], fmt.Sprintf("Ring %s/%d", seed, i), fmt.Sprintf("Simulation ring %s/%d", seed, i), vpsum.Checksum128([]byte(fmt.Sprintf("%s/%d", seed, i))), vpp2pdat.DefaultRingConfig(), nil, nil)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	nodes := make([]*vpp2p.Node, nbNodes)

	i := 0
	for _, v := range hosts {
		for _, w := range rings {
			for j := 0; j < nbNodesPerHostRing; j++ {
				nodes[i], err = vpp2p.NewNode(v, w, nil, vpp2p.GlobalNodeCatalog())
				if err != nil {
					return nil, nil, nil, err
				}
			}
			i++
		}
	}

	return hosts, rings, nodes, nil
}
