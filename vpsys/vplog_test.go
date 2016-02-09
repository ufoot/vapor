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

package vpsys

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const checkCritString = "crit1"
const checkCritStringf = "crit2\ncrit3\n"
const checkErrString = "err1"
const checkErrStringf = "err2\nerr3\n"
const checkWarningString = "warning1"
const checkWarningStringf = "warning2\nwarning3\n"
const checkNoticeString = "notice1"
const checkNoticeStringf = "notice2\nnotice3\n"
const checkInfoString = "info1"
const checkInfoStringf = "info2\ninfo3\n"
const checkDebugString = "debug1"
const checkDebugStringf = "debug2\ndebug3\n"

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
	LogCrit(fmt.Sprintf("%s", checkCritString))
	LogCritf("%s", checkCritStringf)
	// no flush as Crit must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkCritString)
	for _, line := range strings.Split(checkCritStringf, "\n") {
		checkContains(t, filename, line)
	}
}

func TestLogErr(t *testing.T) {
	LogErr(fmt.Sprintf("%s", checkErrString))
	LogErrf("%s", checkErrStringf)
	// no flush as Err must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkErrString)
	for _, line := range strings.Split(checkErrStringf, "\n") {
		checkContains(t, filename, line)
	}
}

func TestLogWarning(t *testing.T) {
	LogWarning(fmt.Sprintf("%s", checkWarningString))
	LogWarningf("%s", checkWarningStringf)
	// no flush as Warning must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkWarningString)
	for _, line := range strings.Split(checkWarningStringf, "\n") {
		checkContains(t, filename, line)
	}
}

func TestLogNotice(t *testing.T) {
	LogNotice(fmt.Sprintf("%s", checkNoticeString))
	LogNoticef("%s", checkNoticeStringf)
	// no flush as Notice must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkNoticeString)
	for _, line := range strings.Split(checkNoticeStringf, "\n") {
		checkContains(t, filename, line)
	}
}

func TestLogInfo(t *testing.T) {
	LogInfo(fmt.Sprintf("%s", checkInfoString))
	LogInfof("%s", checkInfoStringf)
	LogFlush()

	filename := LogFilename()
	checkContains(t, filename, checkInfoString)
	for _, line := range strings.Split(checkInfoStringf, "\n") {
		checkContains(t, filename, line)
	}
}

func TestLogDebug(t *testing.T) {
	p := LogGetPriority()
	LogSetPriority(PriorityDebug)

	LogDebug(fmt.Sprintf("%s", checkDebugString))
	LogDebugf("%s", checkDebugStringf)
	LogFlush()

	LogSetPriority(p)

	filename := LogFilename()
	checkContains(t, filename, checkDebugString)
	for _, line := range strings.Split(checkDebugStringf, "\n") {
		checkContains(t, filename, line)
	}
}
