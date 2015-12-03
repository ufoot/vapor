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

package vpp2p

import (
	"github.com/ufoot/vapor/vpapp"
	"github.com/ufoot/vapor/vpcrypto"
)

// AppInfo contains details about the program.
type AppInfo struct {
	// Unique application ID, generated from other members
	AppID []byte
	// Details about package
	Package vpapp.Package
	// Details about version
	Version vpapp.Version
}

// NewAppInfo creates a new AppInfo object from package and version
func NewAppInfo(p *vpapp.Package, v *vpapp.Version) *AppInfo {
	var ret AppInfo

	ret.AppID = CalcAppID(p, v)
	ret.Package = *p
	ret.Version = *v

	return &ret
}

// CalcAppID generates an Application ID from Package and Version
func CalcAppID(p *vpapp.Package, v *vpapp.Version) []byte {
	var buf []byte

	buf = append(buf, p.Tarname...)
	buf = append(buf, p.Name...)
	buf = append(buf, p.Email...)
	buf = append(buf, p.URL...)
	buf = append(buf, byte(v.Major))
	buf = append(buf, byte(v.Minor))
	buf = append(buf, v.Stamp...)

	return vpcrypto.Checksum128(buf)
}
