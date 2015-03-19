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

// ConfNode is where the config is stored.
type ConfNode struct {
	value    string
	sub_node *[]ConfNode
}

// ConfTrue is the string used to represent true.
const ConfTrue string = "true"

// ConfFalse is the string used to represent false.
const ConfFalse string = "false"

// ConfYes is the string used to represent yes.
const ConfYes string = "yes"

// ConfNo is the string used to represent no.
const ConfNo string = "no"

// ConfParseBool parses a boolean in a string.
func ConfParseBool(val string) bool {
	var err error
	var int_val int64

	val = strings.ToLower(val)

	if val == ConfTrue || val == ConfYes {
		return true
	}
	if val == ConfFalse || val == ConfNo {
		return false
	}

	int_val, err = strconv.ParseInt(val, 64, 10)
	if err == nil && int_val > 0 {
		return true
	}

	return false
}

// ReadEnv reads environment vars and puts them
// in the conf node.
// Returns the number of keys read.
func (n *ConfNode) ReadEnv(prefix string) int {
	var nb_read int

	return nb_read
}

// ReadNode reads vars from a file and puts the values
// in the conf node.
// Returns the number of keys read.
func (n *ConfNode) ReadFile(filename string) (int, error) {
	var nb_read int
	var err error

	return nb_read, err
}

// ReadArgv reads vars from argv and puts the values
// in the conf node.
// Returns the number of keys read.
func (n *ConfNode) ReadArgv() int {
	var nb_read int

	return nb_read
}

// Clear clears the contents of a conf node.
func (n *ConfNode) Clear() {
}

// Merge merges two conf node objects.
// If two different values are available, the values
// of the arguments should be used.
// Returns the number of keys read.
func (n *ConfNode) Merge(other *ConfNode) int {
	var nb_merge int

	return nb_merge
}

// SetString sets a key to the given value.
func (n *ConfNode) SetString(key string, val string) error {
	var err error

	// todo

	return err
}

// GetString returns a given value.
func (n *ConfNode) GetString(key string) (string, error) {
	var err error
	var val string

	// todo

	return val, err
}

// SetInt32 set a key to the given int value.
func (n *ConfNode) SetInt32(key string, val int32) error {
	return n.SetString(key, fmt.Sprintf("%d", val))
}

// GetInt32 returns a given value, as an int32.
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

// SetFloat32 set a key to the given int value.
func (n *ConfNode) SetFloat32(key string, val float32) error {
	return n.SetString(key, fmt.Sprintf("%f", val))
}

// GetFloat32 returns a given value, as a float32.
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

// SetBool set a key to the given boolean value.
func (n *ConfNode) SetBool(key string, val bool) error {
	if val {
		return n.SetString(key, ConfTrue)
	} else {
		return n.SetString(key, ConfFalse)
	}
}

// GetBool returns a given value, as a boolean.
func (n *ConfNode) GetBool(key string) (bool, error) {
	var err error
	var str_val string

	str_val, err = n.GetString(key)
	if err != nil {
		return false, err
	}

	return ConfParseBool(str_val), nil
}
