// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
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

// GetVersionMajor returns the major version number.
func (bus *VpBus) GetVersionMajor() (r int32, err error) {
	return vpbuild.VersionMajor, nil
}

// GetVersionMinor returns the minor version number.
func (bus *VpBus) GetVersionMinor() (r int32, err error) {
	return vpbuild.VersionMinor, nil
}

// GetVersionStamp returns the version stamp.
func (bus *VpBus) GetVersionStamp() (r string, err error) {
	return vpbuild.VersionStamp, nil
}

// GetPackageTarname returns the short package name.
func (bus *VpBus) GetPackageTarname() (r string, err error) {
	return vpbuild.PackageTarname, nil
}

// GetPackageName returns a human readable package name.
func (bus *VpBus) GetPackageName() (r string, err error) {
	return vpbuild.PackageName, nil
}

// GetPackageEmail returns the package maintainer's email.
func (bus *VpBus) GetPackageEmail() (r string, err error) {
	return vpbuild.PackageEmail, nil
}

// GetPackageURL returns the package home page address.
func (bus *VpBus) GetPackageURL() (r string, err error) {
	return vpbuild.PackageURL, nil
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
