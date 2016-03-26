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
	"testing"
)

func TestGlobalHostCatalog(t *testing.T) {
	var host1, host2 *Host
	var err error

	host1, err = NewHost(fmt.Sprintf("%s 1", testTitle), fmt.Sprintf("%s1", testURL), false)
	if err != nil {
		t.Error("unable to create host 1", err)
	}
	host2, err = NewHost(fmt.Sprintf("%s 2", testTitle), fmt.Sprintf("%s2", testURL), false)
	if err != nil {
		t.Error("unable to create host 2", err)
	}

	if GlobalHostCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog has host1, but it's not registered yet")
	}
	if GlobalHostCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog has host2, but it's not registered yet")
	}
	GlobalHostCatalog().RegisterHost(&(host1.Info))
	if !GlobalHostCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog does not have host1, but it has been registered")
	}
	if GlobalHostCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog has host2, but it's not registered yet")
	}
	GlobalHostCatalog().RegisterHost(&(host2.Info))
	if !GlobalHostCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog does not have host1, but it has been registered")
	}
	if !GlobalHostCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog does not have host2, but it has been registered")
	}
	if len(GlobalHostCatalog().List()) != 2 {
		t.Error("global catalog length should be 2")
	}
	GlobalHostCatalog().UnregisterHost(&(host1.Info))
	if GlobalHostCatalog().HasHost(host1.Info.HostPubKey) {
		t.Error("global catalog has host1, but it's not registered yet")
	}
	if !GlobalHostCatalog().HasHost(host2.Info.HostPubKey) {
		t.Error("global catalog does not have host2, but it has been registered")
	}
	if len(GlobalHostCatalog().List()) != 1 {
		t.Error("global catalog length should be 1")
	}
}
