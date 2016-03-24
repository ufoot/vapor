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

import (
	"testing"
)

const (
	testNbHostsNoSig            = 50
	testNbRingsNoSig            = 8
	testNbNodesPerHostRingNoSig = 250
	testNbHostsSig              = 2
	testNbRingsSig              = 2
	testNbNodesPerHostRingSig   = 2
)

func TestSetupLocalNoSig(t *testing.T) {
	hosts, rings, nodes, err := SetupLocal(testNbHostsNoSig, testNbRingsNoSig, testNbNodesPerHostRingNoSig, false)

	if err != nil {
		t.Error("unable to set up env (nosig)", err)
	}
	nbHosts := len(hosts)
	nbRings := len(rings)
	nbNodes := len(nodes)
	t.Logf("set up done (nosig), nbHosts=%d nbRings=%d nbNodes=%d", nbHosts, nbRings, nbNodes)
}

func TestSetupLocalSig(t *testing.T) {
	hosts, rings, nodes, err := SetupLocal(testNbHostsSig, testNbRingsSig, testNbNodesPerHostRingSig, true)

	if err != nil {
		t.Error("unable to set up env (sig)", err)
	}
	nbHosts := len(hosts)
	nbRings := len(rings)
	nbNodes := len(nodes)
	t.Logf("set up done (sig), nbHosts=%d nbRings=%d nbNodes=%d", nbHosts, nbRings, nbNodes)
}
