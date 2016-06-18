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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ufoot/vapor/go/vperror"
	"strings"
)

func enumMarshalJSON(val int, shortStr *[]string) ([]byte, error) {
	if val >= 0 && val < len(*shortStr) {
		ret, err := json.Marshal((*shortStr)[val])
		if err != nil {
			return nil, vperror.Chainf(err, "unable to marshal %s", (*shortStr)[val])
		}
		return ret, nil
	}

	return nil, fmt.Errorf("bad enum %d, should be in range [0,%d)", val, len(*shortStr))
}

func enumUnmarshalJSON(data []byte, shortStr, longStr *[]string) (int, error) {
	var tmpStr string
	var tmpInt int

	err := json.Unmarshal(data, &tmpInt)
	if err != nil {
		err := json.Unmarshal(data, &tmpStr)
		if err != nil {
			if err != nil {
				return -1, vperror.Chain(err, "unable to unmarshal enum")
			}
		}

		for i, v := range *shortStr {
			if strings.ToUpper(tmpStr) == strings.ToUpper(v) {
				return i, nil
			}
		}
		for i, v := range *longStr {
			if strings.ToLower(tmpStr) == strings.ToLower(v) {
				return i, nil
			}
		}
	} else {
		if tmpInt < 0 || tmpInt >= len(*shortStr) {
			return -1, fmt.Errorf("out of bound enum %d, should be in range [0,%d)", tmpInt, len(*shortStr))
		}
		return tmpInt, nil
	}

	return -1, errors.New("unable to unmarshal enum")
}

func enumString(val int, longStr *[]string) string {
	if val >= 0 && val < len(*longStr) {
		return (*longStr)[val]
	}

	return ""
}
