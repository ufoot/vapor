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
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

package vpconf

import (
	"strings"
)

// ConfIDSep is the separator used to separate ids.
const ConfIDSep string = "."

// ConfIDSplit splits a conf path into its members.
func ConfIDSplit(id string) []string {
	split := strings.Split(strings.ToLower(id), ConfIDSep)
	var ret []string
	var val string

	for _, val = range split {
		if len(val) >= 1 {
			ret = append(ret, val)
		}
	}

	return ret
}

// ConfIDJoin joins conf path members to form a path.
func ConfIDJoin(id []string) string {
	var tmp []string
	var val string

	for _, val = range id {
		if len(val) >= 1 {
			tmp = append(tmp, val)
		}
	}

	return strings.ToLower(strings.Join(tmp, ConfIDSep))
}

// ConfIDParentChildren returns the top-level parent
// of a conf path, and the child path associated.
// For instance foo.bar.num should return foo and bar.num.
func ConfIDParentChildren(id string) (string, string) {
	var parent string
	var children string

	split := ConfIDSplit(id)

	parent = split[0]
	children = ConfIDJoin(split[1:])

	return parent, children
}
