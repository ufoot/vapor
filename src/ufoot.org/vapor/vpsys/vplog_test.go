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
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const checkCritString = "crit1"
const checkCritStringf = "crit2"
const checkCritStringln = "crit3"
const checkErrString = "err1"
const checkErrStringf = "err2"
const checkErrStringln = "err3"
const checkWarningString = "warning1"
const checkWarningStringf = "warning2"
const checkWarningStringln = "warning3"
const checkNoticeString = "notice1"
const checkNoticeStringf = "notice2"
const checkNoticeStringln = "notice3"
const checkInfoString = "info1"
const checkInfoStringf = "info2"
const checkInfoStringln = "info3"
const checkDebugString = "debug1"
const checkDebugStringf = "debug2"
const checkDebugStringln = "debug3"

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
	LogCritln(checkCritStringln)
	// no flush as Crit must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkCritString)
	checkContains(t, filename, checkCritStringf)
	checkContains(t, filename, checkCritStringln)
}

func TestLogErr(t *testing.T) {
	LogErr(fmt.Sprintf("%s\n", checkErrString))
	LogErrf("%s", checkErrStringf)
	LogErrln(checkErrStringln)
	// no flush as Err must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkErrString)
	checkContains(t, filename, checkErrStringf)
	checkContains(t, filename, checkErrStringln)
}

func TestLogWarning(t *testing.T) {
	LogWarning(fmt.Sprintf("%s\n", checkWarningString))
	LogWarningf("%s", checkWarningStringf)
	LogWarningln(checkWarningStringln)
	// no flush as Warning must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkWarningString)
	checkContains(t, filename, checkWarningStringf)
	checkContains(t, filename, checkWarningStringln)
}

func TestLogNotice(t *testing.T) {
	LogNotice(fmt.Sprintf("%s\n", checkNoticeString))
	LogNoticef("%s", checkNoticeStringf)
	LogNoticeln(checkNoticeStringln)
	// no flush as Notice must be auto-flushed

	filename := LogFilename()
	checkContains(t, filename, checkNoticeString)
	checkContains(t, filename, checkNoticeStringf)
	checkContains(t, filename, checkNoticeStringln)
}

func TestLogInfo(t *testing.T) {
	LogInfo(fmt.Sprintf("%s\n", checkInfoString))
	LogInfof("%s", checkInfoStringf)
	LogInfoln(checkInfoStringln)
	LogFlush()

	filename := LogFilename()
	checkContains(t, filename, checkInfoString)
	checkContains(t, filename, checkInfoStringf)
	checkContains(t, filename, checkInfoStringln)
}

func TestLogDebug(t *testing.T) {
	p := LogGetPriority()
	LogSetPriority(LOG_DEBUG)

	LogDebug(fmt.Sprintf("%s\n", checkDebugString))
	LogDebugf("%s", checkDebugStringf)
	LogDebugln(checkDebugStringln)
	LogFlush()

	LogSetPriority(p)

	filename := LogFilename()
	checkContains(t, filename, checkDebugString)
	checkContains(t, filename, checkDebugStringf)
	checkContains(t, filename, checkDebugStringln)
}
