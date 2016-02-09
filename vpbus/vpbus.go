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

package vpbus

import (
	"github.com/ufoot/vapor/vpbuild"
	"github.com/ufoot/vapor/vpcommonapi"
	"github.com/ufoot/vapor/vpsys"
	"time"
)

// VpBus is an implementation of VpBusApi interface.
type VpBus struct {
	startTime time.Time
}

// New creates a new VpBus object, to act as a server callback.
func New() *VpBus {
	var ret VpBus

	ret.startTime = time.Now()

	return &ret
}

// Ping is just here to make the server pingable, Thriftly speaking.
func (bus *VpBus) Ping() (err error) {
	return nil
}

// GetPackage returns the version of the program.
func (bus *VpBus) GetPackage() (r *vpcommonapi.Package, err error) {
	var p vpcommonapi.Package

	p.Tarname = vpbuild.PackageTarname
	p.Name = vpbuild.PackageName
	p.Email = vpbuild.PackageEmail
	p.URL = vpbuild.PackageURL

	return &p, nil
}

// GetVersion returns the version of the program.
func (bus *VpBus) GetVersion() (r *vpcommonapi.Version, err error) {
	var v vpcommonapi.Version

	v.Major = vpbuild.VersionMajor
	v.Minor = vpbuild.VersionMinor
	v.Stamp = vpbuild.VersionStamp

	return &v, nil
}

// Uptime returns how many seconds the server has been up.
func (bus *VpBus) Uptime() (r int64, err error) {
	return time.Now().Unix() - bus.startTime.Unix(), nil
}

// Halt stops the server.
func (bus *VpBus) Halt() (err error) {
	vpsys.LogNotice("FIXME Halt")

	return nil
}
