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

// Config is use to keep major settings together.
type Config struct {
	// M is the m parameter (base) used for Koorde/Bruijn ops.
	M int
	// N is the n parameter (number of elements) used for Koorde/Bruijn ops.
	N int
	// UseSig tells wether to use cryptographic signatures/checks.
	UseSig bool
	// Steps optimizes Bruijn walk by considering only this number
	// of steps in the worst case.
	Steps int
}

