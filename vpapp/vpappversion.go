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

package vpapp

import (
	"fmt"
	"github.com/ufoot/vapor/vpbuild"
)

// Version is used to identify the program version.
type Version struct {
	// Major version number
	Major int
	// Minor version number
	Minor int
	// Stamp, to differenciate between builds
	Stamp string
}

// NewVersion creates a new version object.
func NewVersion(major, minor int, stamp string) *Version {
	return &Version{major, minor, stamp}
}

// BuildVersion creates a new default object.
func BuildVersion() *Version {
	return NewVersion(vpbuild.VersionMajor, vpbuild.VersionMinor, vpbuild.VersionStamp)
}

// Compare compares two versions, returns -1 if v1<v2,
// 0 if they are equal, and +1 if v1>v2. Stamp is just
// ignored, only major and minor are considered.
func Compare(v1, v2 *Version) int {
	if v1.Major < v2.Major {
		return -1
	}
	if v1.Major > v2.Major {
		return +1
	}
	if v1.Minor < v2.Minor {
		return -1
	}
	if v1.Minor > v2.Minor {
		return +1
	}
	return 0
}

// Equal returns true if the versions are exactly the same.
// Stamp is still just ignored, only major and minor are considered.
func Equal(v1, v2 *Version) bool {
	return Compare(v1, v2) == 0
}

// String returns the version as a string
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%s", v.Major, v.Minor, v.Stamp)
}
