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

package vptypes

import (
	"github.com/ufoot/vapor/vperror"
)

// Size is a generic type to store small/medium/large values.
// Typically used for game settings.
type Size int

const (
	// SizeSmall represent an abritrary "small" value.
	SizeSmall Size = iota
	// SizeMedium represent an abritrary "medium" value.
	SizeMedium
	// SizeLarge represent an abritrary "large" value.
	SizeLarge
	// SizeRange is the number of available sizes.
	SizeRange
)

var sizeShortStr []string
var sizeLongStr []string

func init() {
	sizeShortStr = make([]string, SizeRange)
	sizeShortStr[SizeSmall] = "S"
	sizeShortStr[SizeMedium] = "M"
	sizeShortStr[SizeLarge] = "L"
	sizeLongStr = make([]string, SizeRange)
	sizeLongStr[SizeSmall] = "small"
	sizeLongStr[SizeMedium] = "medium"
	sizeLongStr[SizeLarge] = "large"
}

// MarshalJSON implements the json.Marshaler interface.
func (s *Size) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(int(*s), &sizeShortStr)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Size) UnmarshalJSON(data []byte) error {
	tmpInt, err := enumUnmarshalJSON(data, &sizeShortStr, &sizeLongStr)

	if err != nil {
		return vperror.Chain(err, "unable to unmarshal Size")
	}
	*s = Size(tmpInt)

	return nil
}

// String returns a readable form of the Size.
func (s *Size) String() string {
	return enumString(int(*s), &sizeLongStr)
}
