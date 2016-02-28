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

package vpapp

import (
	"github.com/ufoot/vapor/vpsum"
)

// CalcID generates an Application ID from Package and Version
func CalcID(p *Package, v *Version) []byte {
	var buf []byte

	buf = append(buf, p.Tarname...)
	//buf = append(buf, p.Name...)
	//buf = append(buf, p.Email...)
	//buf = append(buf, p.URL...)
	buf = append(buf, byte(v.Major))
	buf = append(buf, byte(v.Minor))
	//buf = append(buf, v.Stamp...)

	return vpsum.Checksum128(buf)
}
