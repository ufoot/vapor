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

const CONF_ID_SEP string = "."

func ConfIdSplit(id string) []string {
	split := strings.Split(strings.ToLower(id), CONF_ID_SEP)
	var ret []string
	var val string

	for _, val = range split {
		if len(val) >= 1 {
			ret = append(ret, val)
		}
	}

	return ret
}

func ConfIdJoin(id []string) string {
	var tmp []string
	var val string

	for _, val = range id {
		if len(val) >= 1 {
			tmp = append(tmp, val)
		}
	}

	return strings.ToLower(strings.Join(tmp, CONF_ID_SEP))
}

func ConfIdParentChildren(id string) (string, string) {
	var parent string
	var children string

	split := ConfIdSplit(id)

	parent = split[0]
	children = ConfIdJoin(split[1:])

	return parent, children
}
