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

// AppPackage is used to identify the program package, think
// of it as old good old autotools PACKAGE (name on steroids).
type AppPackage struct {
	// Tarname, for instance myapp
	Tarname string
	// Readable name, for instance My App
	Name string
	// Email of maintainer
	Email string
	// URL of program, to get info about it
	URL string
}

// AppVersion is used to identify the program version.
type AppVersion struct {
	// Major version number
	Major int
	// Minor version number
	Minor int
	// Stamp, to differenciate between builds
	Stamp string
}

// AppInfo contains details about the program.
type AppInfo struct {
	// Unique application ID, generated from other members
	AppID []byte
	// Details about package
	Package AppPackage
	// Details about version
	Version AppVersion
}
