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

//go:generate bash ./stamp.sh
	
import (
	"fmt"
	"github.com/ufoot/vapor/go/vplog"
	"os"
)

func main() {
	base64RingID, base64AppID, base64HostPubKey, base64RingSig, err := Base64Ring0()

	if err != nil {
		vplog.LogErr(err)
		os.Exit(1)
	}

	fmt.Printf("// Ring0Base64RingID contains the ring ID of the default ring0, encoded in base64\n")
	fmt.Printf("Ring0Base64RingID = \"%s\"\n", base64RingID)
	fmt.Printf("// Ring0Base64AppID contains the application ID of the default ring0, encoded in base64\n")
	fmt.Printf("Ring0Base64AppID = \"%s\"\n", base64AppID)
	fmt.Printf("// Ring0Base64HostPubKey contains the public key of the default ring0, encoded in base64\n")
	fmt.Printf("Ring0Base64HostPubKey = \"%s\"\n", base64HostPubKey)
	fmt.Printf("// Ring0Base64RingSig contains the siignature of the default ring0, encoded in base64\n")
	fmt.Printf("Ring0Base64RingSig = \"%s\"\n", base64RingSig)

	return
}
