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

package vplog

import (
	"bytes"
	"fmt"
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

func initLog() *bytes.Buffer {
	var testBuffer bytes.Buffer
	logInitWithWriter("test", &testBuffer)
	return &testBuffer
}

func checkContains(t *testing.T, haystack string, needle string) {
	if strings.Contains(haystack, needle) {
		t.Logf("log contains \"%s\"", needle)
	} else {
		report := haystack
		for _, chr := range "\n\r\t" {
			report = strings.Replace(report, string(chr), " ", -1)
		}
		t.Errorf("log \"%s\" does not contains \"%s\"", report, needle)
	}
}

func TestLogCrit(t *testing.T) {
	buffer := initLog()

	LogCrit(fmt.Sprintf("%s", checkCritString))
	LogCritf("%s", checkCritStringf)
	// no flush as Crit must be auto-flushed

	content := buffer.String()

	checkContains(t, content, checkCritString)
	for _, line := range strings.Split(checkCritStringf, "\n") {
		checkContains(t, content, line)
	}
}

func TestLogErr(t *testing.T) {
	buffer := initLog()

	LogErr(fmt.Sprintf("%s", checkErrString))
	LogErrf("%s", checkErrStringf)
	// no flush as Err must be auto-flushed

	content := buffer.String()

	checkContains(t, content, checkErrString)
	for _, line := range strings.Split(checkErrStringf, "\n") {
		checkContains(t, content, line)
	}
}

func TestLogWarning(t *testing.T) {
	buffer := initLog()

	LogWarning(fmt.Sprintf("%s", checkWarningString))
	LogWarningf("%s", checkWarningStringf)
	// no flush as Warning must be auto-flushed

	content := buffer.String()

	checkContains(t, content, checkWarningString)
	for _, line := range strings.Split(checkWarningStringf, "\n") {
		checkContains(t, content, line)
	}
}

func TestLogNotice(t *testing.T) {
	buffer := initLog()

	LogNotice(fmt.Sprintf("%s", checkNoticeString))
	LogNoticef("%s", checkNoticeStringf)
	// no flush as Notice must be auto-flushed

	content := buffer.String()

	checkContains(t, content, checkNoticeString)
	for _, line := range strings.Split(checkNoticeStringf, "\n") {
		checkContains(t, content, line)
	}
}

func TestLogInfo(t *testing.T) {
	buffer := initLog()

	LogInfo(fmt.Sprintf("%s", checkInfoString))
	LogInfof("%s", checkInfoStringf)
	LogFlush()

	content := buffer.String()

	checkContains(t, content, checkInfoString)
	for _, line := range strings.Split(checkInfoStringf, "\n") {
		checkContains(t, content, line)
	}
}

func TestLogDebug(t *testing.T) {
	buffer := initLog()

	p := LogGetPriority()
	LogSetPriority(PriorityDebug)

	LogDebug(fmt.Sprintf("%s", checkDebugString))
	LogDebugf("%s", checkDebugStringf)
	LogFlush()

	LogSetPriority(p)

	content := buffer.String()

	checkContains(t, content, checkDebugString)
	for _, line := range strings.Split(checkDebugStringf, "\n") {
		checkContains(t, content, line)
	}
}
