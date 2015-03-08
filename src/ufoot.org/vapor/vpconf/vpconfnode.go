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

package vpconf

import (
	"fmt"
	"strconv"
	"strings"
	"ufoot.org/vapor/vpsys"
)

type ConfNode struct {
	value    string
	sub_node *[]ConfNode
}

const CONF_TRUE string = "true"
const CONF_FALSE string = "false"
const CONF_YES string = "yes"
const CONF_NO string = "no"

func ConfParseBool(val string) bool {
	var err error
	var int_val int64

	val = strings.ToLower(val)

	if val == CONF_TRUE || val == CONF_YES {
		return true
	}
	if val == CONF_FALSE || val == CONF_NO {
		return false
	}

	int_val, err = strconv.ParseInt(val, 64, 10)
	if err == nil && int_val > 0 {
		return true
	}

	return false
}

func (n *ConfNode) ReadEnv(prefix string) int {
	var nb_read int

	return nb_read
}

func (n *ConfNode) ReadFile(filename string) (int, error) {
	var nb_read int
	var err error

	return nb_read, err
}

func (n *ConfNode) ReadArgv() int {
	var nb_read int

	return nb_read
}

func (n *ConfNode) Clear() {
}

func (n *ConfNode) Merge(other *ConfNode) int {
	var nb_merge int

	return nb_merge
}

func (n *ConfNode) SetString(key string, val string) error {
	var err error

	// todo

	return err
}

func (n *ConfNode) GetString(key string) (string, error) {
	var err error
	var val string

	// todo

	return val, err
}

func (n *ConfNode) SetInt32(key string, val int32) error {
	return n.SetString(key, fmt.Sprintf("%d", val))
}

func (n *ConfNode) GetInt32(key string) (int32, error) {
	var ret int64
	var err error
	var str_val string

	str_val, err = n.GetString(key)
	if err != nil {
		return 0, err
	}

	ret, err = strconv.ParseInt(str_val, 32, 10)

	return int32(ret), vpsys.ErrorChain(err, "unable to parse int")
}

func (n *ConfNode) SetFloat32(key string, val float32) error {
	return n.SetString(key, fmt.Sprintf("%f", val))
}

func (n *ConfNode) GetFloat32(key string) (float32, error) {
	var ret float64
	var err error
	var str_val string

	str_val, err = n.GetString(key)
	if err != nil {
		return 0, err
	}

	ret, err = strconv.ParseFloat(str_val, 32)

	return float32(ret), err
}

func (n *ConfNode) SetBool(key string, val bool) error {
	if val {
		return n.SetString(key, CONF_TRUE)
	} else {
		return n.SetString(key, CONF_FALSE)
	}
}

func (n *ConfNode) GetBool(key string) (bool, error) {
	var err error
	var str_val string

	str_val, err = n.GetString(key)
	if err != nil {
		return false, err
	}

	return ConfParseBool(str_val), nil
}
