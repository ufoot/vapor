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

package vpsys

import (
	"fmt"
	"os"
	"path"
)

// Home returns the path of the home directory of the user.
// It is not necessarly the "HOME" environment variable, indeed,
// if this one is not available, it will try other places, call mkdir,
// do whatever it can to find some suitable place where the program
// could legitimately store some data.
func Home(program string) string {
	var home string
	var err error

	home = "."
	home = os.Getenv("HOME")
	err = os.MkdirAll(home, 0750)
	if err != nil && os.IsNotExist(err) {
		home = "."
	}

	if home == "." {
		home, err = os.Getwd()
		if err != nil {
			home = "."
		} else {
			err = os.MkdirAll(home, 0750)
			if err != nil && os.IsNotExist(err) {
				home = "."
			}
		}
	}

	if home == "." {
		home = os.TempDir()
		err = os.MkdirAll(home, 0750)
		if err != nil && os.IsNotExist(err) {
			home = "."
		}
	}

	home = path.Join(home, fmt.Sprintf(".%s", program))

	err = os.MkdirAll(home, 0750)
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	return home
}
