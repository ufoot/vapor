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

package vptimeout

import (
	"fmt"
	"time"
)

// Run runs the function f, but will stop everyting if timeout
// is reached. If f takes longer than timeout, a timeout error
// will be returned, and the goroutine might keep going until
// it's done, but the result will be ignored.
func Run(f func() error, d time.Duration) error {
	t := time.NewTimer(d)
	var err error
	done := make(chan bool, 1)

	go func() {
		err = f()
		done <- true
	}()

	for {
		select {
		case <-t.C:
			return fmt.Errorf("timeout after %s", d.String())
		case <-done:
			return err
		}
	}
}
