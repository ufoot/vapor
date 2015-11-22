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

package vptypes

import (
	"testing"
)

func TestSizeJSON(t *testing.T) {
	s := SizeMedium

	j, err := s.MarshalJSON()
	if err != nil {
		t.Errorf("unable to marshal s=%d", int(s))
	}
	t.Logf("marshalled %d is %s", int(s), string(j))
	s = SizeSmall
	err = s.UnmarshalJSON(j)
	if err != nil {
		t.Errorf("unable to unmarshal j=%s", string(j))
	}
	if s != SizeMedium {
		t.Errorf("bad unmarshalled size s=%d SizeMedium=%d", int(s), int(SizeMedium))
	}
}
