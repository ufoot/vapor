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

package vpsys

import (
	"testing"
)

func TestPriorityString(t *testing.T) {
	t.Logf("priority crit %d/%s", int(PriorityCrit), PriorityString(PriorityCrit))
	t.Logf("priority err %d/%s", int(PriorityErr), PriorityString(PriorityErr))
	t.Logf("priority warning %d/%s", int(PriorityWarning), PriorityString(PriorityWarning))
	t.Logf("priority notice %d/%s", int(PriorityNotice), PriorityString(PriorityNotice))
	t.Logf("priority info %d/%s", int(PriorityInfo), PriorityString(PriorityInfo))
	t.Logf("priority debug %d/%s", int(PriorityDebug), PriorityString(PriorityDebug))
}
