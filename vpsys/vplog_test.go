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

package vpsys

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const checkCritString = "crit1"
const checkCritStringf = "crit2"
const checkErrString = "err1"
const checkErrStringf = "err2"
const checkWarningString = "warning1"
const checkWarningStringf = "warning2"
const checkNoticeString = "notice1"
const checkNoticeStringf = "notice2"
const checkInfoString = "info1"
const checkInfoStringf = "info2"
const checkDebugString = "debug1"
const checkDebugStringf = "debug2"

func checkContains(t *testing.T, filename string, text string) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Errorf("unable to open %s, %s\n", filename, err)
	}
	if strings.Contains(string(content), text) {
		t.Logf("%s contains %s\n", filename, text)
	} else {
		t.Errorf("%s does not contain %s\n", filename, text)
	}
}

func TestLogCrit(t *testing.T) {
	LogCrit(fmt.Sprintf("%s\n", checkCritString))
	LogCritf("%s", checkCritStringf)
	// no flush as Crit must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkCritString)
	checkContains(t, filename, checkCritStringf)
}

func TestLogErr(t *testing.T) {
	LogErr(fmt.Sprintf("%s\n", checkErrString))
	LogErrf("%s", checkErrStringf)
	// no flush as Err must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkErrString)
	checkContains(t, filename, checkErrStringf)
}

func TestLogWarning(t *testing.T) {
	LogWarning(fmt.Sprintf("%s\n", checkWarningString))
	LogWarningf("%s", checkWarningStringf)
	// no flush as Warning must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkWarningString)
	checkContains(t, filename, checkWarningStringf)
}

func TestLogNotice(t *testing.T) {
	LogNotice(fmt.Sprintf("%s\n", checkNoticeString))
	LogNoticef("%s", checkNoticeStringf)
	// no flush as Notice must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkNoticeString)
	checkContains(t, filename, checkNoticeStringf)
}

func TestLogInfo(t *testing.T) {
	LogInfo(fmt.Sprintf("%s\n", checkInfoString))
	LogInfof("%s", checkInfoStringf)
	LogFlush()

	filename := LogFilename()
	checkContains(t, filename, checkInfoString)
	checkContains(t, filename, checkInfoStringf)
}

func TestLogDebug(t *testing.T) {
	p := LogGetPriority()
	LogSetPriority(PriorityDebug)

	LogDebug(fmt.Sprintf("%s\n", checkDebugString))
	LogDebugf("%s", checkDebugStringf)
	LogFlush()

	LogSetPriority(p)

	filename := LogFilename()
	checkContains(t, filename, checkDebugString)
	checkContains(t, filename, checkDebugStringf)
}
