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

package vploop

//go:generate bash ./stamp.sh
	
import (
	"github.com/ufoot/vapor/go/vplog"
	"time"
)

// MainLoop implements the main game loop.
func MainLoop(handlers ...LoopHandler) int64 {
	quit := make(chan bool, len(handlers))
	channels := make([]<-chan time.Time, len(handlers))
	iterations := make(chan int64, len(handlers))
	var done int
	var localIterations int64
	var totalIterations int64

	vplog.LogNoticef("loop initial")

	for i, handler := range handlers {
		handler.Init(time.Now())
		channels[i] = time.Tick(handler.Duration())
		defer handler.Quit(time.Now())
	}

	vplog.LogNoticef("loop all begin")
	for i, channel := range channels {
		go func(i int, channel <-chan time.Time) {
			vplog.LogNoticef("loop %d begin", i)
			var iteration int64
			for {
				iteration = iteration + 1
				select {
				case q := <-quit:
					vplog.LogNoticef("quit received %d iteration=%d", i, iteration)
					if q == true {
						vplog.LogNoticef("loop %d end 1/3", i)
						quit <- true
						vplog.LogNoticef("loop %d end 2/3", i)
						iterations <- iteration
						vplog.LogNoticef("loop %d end 3/3", i)
						return
					}
				default:
					select {
					case timestamp := <-channel:
						handlers[i].Do(timestamp, iteration, quit)
					}
				}
			}
		}(i, channel)
	}
	vplog.LogNoticef("loop all end")

	for {
		select {
		case localIterations = <-iterations:
			vplog.LogNoticef("local iterations %d", localIterations)
			totalIterations += localIterations
			done++
		default:
			if done >= len(handlers) {
				vplog.LogNoticef("loop final")
				return totalIterations
			}
		}
	}
}
