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

package main

import (
	"encoding/base64"
	"github.com/ufoot/vapor/vpp2p"
)

// Base64Ring0 creates a new instance of the default directory ring
// and returns its major parameters as base64 encoded strings. These
// can typically be copy/pasted into the code to immortalize the
// reference instance of that ring.
func Base64Ring0() (string, string, string, string, error) {
	ring0, err := vpp2p.NewRing0()

	if err != nil {
		return "", "", "", "", err
	}

	base64RingID := base64.URLEncoding.EncodeToString(ring0.Info.RingID)
	base64AppID := base64.URLEncoding.EncodeToString(ring0.Info.AppID)
	base64HostPubKey := base64.URLEncoding.EncodeToString(ring0.Info.HostPubKey)
	base64RingSig := base64.URLEncoding.EncodeToString(ring0.Info.RingSig)

	return base64RingID, base64AppID, base64HostPubKey, base64RingSig, nil
}
