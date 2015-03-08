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

package vploop

import (
	"time"
)

func MainLoop(handler LoopHandler) int {
	var i int64
	quit := make(chan bool)

	handler.Init(time.Now())
	defer handler.Quit(time.Now())
	for {
		now := time.Now()
		select {
		case q := <-quit:
			if q == true {
				return 0
			}
		default:
			handler.Do(now, i, quit)
			i = i + 1
		}
	}
}
