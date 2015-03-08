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
	"log/syslog"
)

// Logging priority (AKA level), the lower the number, the
// higher the priority.
type Priority int

const (
	// Inspired from /usr/include/sys/syslog.h.
	// We don't use EMERG and ALERT
	LOG_CRIT Priority = iota + Priority(int(syslog.LOG_CRIT))
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const critString = "CRIT"
const errString = "ERR"
const warningString = "WARNING"
const noticeString = "NOTICE"
const infoString = "INFO"
const debugString = "DEBUG"
const unknownString = "UNKNOWN"

// Custom logging interface, while the package with only one
// implementation, the interface is technically generic, and
// could be implemented/used elsewhere. This is design for
// quick-to-write global usage, not for universality/versatility.
type Logger interface {
	Log(p Priority, v ...interface{})
	Logf(p Priority, f string, v ...interface{})
	Logln(p Priority, v ...interface{})
	Filename() string
	SetPriority(p Priority)
	GetPriority() Priority
	Flush()
}

// Returns a readable string (English word) that corresponds
// to a given log priority.
func PriorityString(p Priority) string {
	if p <= LOG_CRIT {
		p = LOG_CRIT
	}
	if p >= LOG_DEBUG {
		p = LOG_DEBUG
	}
	switch p {
	case LOG_CRIT:
		return critString
	case LOG_ERR:
		return errString
	case LOG_WARNING:
		return warningString
	case LOG_NOTICE:
		return noticeString
	case LOG_INFO:
		return infoString
	case LOG_DEBUG:
		return debugString
	}

	return unknownString
}

func LoggerCrit(l Logger, v ...interface{}) {
	l.Log(LOG_CRIT, v...)
}

func LoggerCritf(l Logger, f string, v ...interface{}) {
	l.Logf(LOG_CRIT, f, v...)
}

func LoggerCritln(l Logger, v ...interface{}) {
	l.Logln(LOG_CRIT, v...)
}

func LoggerErr(l Logger, v ...interface{}) {
	l.Log(LOG_ERR, v...)
}

func LoggerErrf(l Logger, f string, v ...interface{}) {
	l.Logf(LOG_ERR, f, v...)
}

func LoggerErrln(l Logger, v ...interface{}) {
	l.Logln(LOG_ERR, v...)
}

func LoggerWarning(l Logger, v ...interface{}) {
	l.Log(LOG_WARNING, v...)
}

func LoggerWarningf(l Logger, f string, v ...interface{}) {
	l.Logf(LOG_WARNING, f, v...)
}

func LoggerWarningln(l Logger, v ...interface{}) {
	l.Logln(LOG_WARNING, v...)
}

func LoggerNotice(l Logger, v ...interface{}) {
	l.Log(LOG_NOTICE, v...)
}

func LoggerNoticef(l Logger, f string, v ...interface{}) {
	l.Logf(LOG_NOTICE, f, v...)
}

func LoggerNoticeln(l Logger, v ...interface{}) {
	l.Logln(LOG_NOTICE, v...)
}

func LoggerInfo(l Logger, v ...interface{}) {
	l.Log(LOG_INFO, v...)
}

func LoggerInfof(l Logger, f string, v ...interface{}) {
	l.Logf(LOG_INFO, f, v...)
}

func LoggerInfoln(l Logger, v ...interface{}) {
	l.Logln(LOG_INFO, v...)
}

func LoggerDebug(l Logger, v ...interface{}) {
	l.Log(LOG_DEBUG, v...)
}

func LoggerDebugf(l Logger, f string, v ...interface{}) {
	l.Logf(LOG_DEBUG, f, v...)
}

func LoggerDebugln(l Logger, v ...interface{}) {
	l.Logln(LOG_DEBUG, v...)
}
