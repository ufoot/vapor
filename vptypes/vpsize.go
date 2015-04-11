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
	"encoding/json"
	"fmt"
	"github.com/ufoot/vapor/vpsys"
	"strings"
)

type Size int32

const (
	SizeSmall Size = iota
	SizeMedium
	SizeLarge
	SizeRange
)

const sizeSmallShortStr = "S"
const sizeMediumShortStr = "M"
const sizeLargeShortStr = "L"
const sizeSmallLongStr = "small"
const sizeMediumLongStr = "medium"
const sizeLargeLongStr = "large"

// MarshalJSON implements the json.Marshaler interface.
func (s *Size) MarshalJSON() ([]byte, error) {
	var js string

	switch *s {
	case SizeSmall:
		js = sizeSmallShortStr
	case SizeMedium:
		js = sizeMediumShortStr
	case SizeLarge:
		js = sizeLargeLongStr
	default:
		return nil, fmt.Errorf("bad Size %d", s)
	}
	ret, err := json.Marshal(js)
	if err != nil {
		return nil, vpsys.ErrorChainf(err, "unable to marshal Size %s", js)
	}

	return ret, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Size) UnmarshalJSON(data []byte) error {
	var tmpStr string
	var tmpInt int32

	err := json.Unmarshal(data, &tmpInt)
	if err != nil {
		err := json.Unmarshal(data, &tmpStr)
		if err != nil {
			if err != nil {
				return vpsys.ErrorChain(err, "unable to unmarshal Size")
			}
		}
		switch {
		case strings.ToUpper(tmpStr) == sizeSmallShortStr || strings.ToLower(tmpStr) == sizeSmallLongStr:
			*s = SizeSmall
			return nil
		case strings.ToUpper(tmpStr) == sizeMediumShortStr || strings.ToLower(tmpStr) == sizeMediumLongStr:
			*s = SizeMedium
			return nil
		case strings.ToUpper(tmpStr) == sizeLargeShortStr || strings.ToLower(tmpStr) == sizeLargeLongStr:
			*s = SizeLarge
			return nil
		default:
			return fmt.Errorf("unable to interpret '%s' as a Size", tmpStr)
		}
	}

	if tmpInt < int32(SizeSmall) || tmpInt > int32(SizeLarge) {
		return fmt.Errorf("out of bound Size %d", tmpInt)
	}

	*s = Size(tmpInt)

	return nil
}
