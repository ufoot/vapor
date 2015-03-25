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
	"log/syslog"
)

// Priority is a logging priority (AKA level), the lower the number, the
// higher the priority.
type Priority int

const (
	// PriorityCrit should be used for critical messages.
	// Inspired from /usr/include/sys/syslog.h.
	// We don't use EMERG and ALERT
	PriorityCrit Priority = iota + Priority(int(syslog.LOG_CRIT))
	// PriorityErr should be used for error messages.
	PriorityErr
	// PriorityWarning should be used for warning messages.
	PriorityWarning
	// PriorityNotice should be used for notice messages.
	PriorityNotice
	// PriorityInfo should be used for information messages.
	PriorityInfo
	// PriorityDebug should be used for debugging messages.
	PriorityDebug
)

const critString = "CRIT"
const errString = "ERR"
const warningString = "WARNING"
const noticeString = "NOTICE"
const infoString = "INFO"
const debugString = "DEBUG"
const unknownString = "UNKNOWN"

// Logger is a custom logging interface, while the package with only one
// implementation, the interface is technically generic, and
// could be implemented/used elsewhere. This is design for
// quick-to-write global usage, not for universality/versatility.
type Logger interface {
	Log(p Priority, v ...interface{})
	Logf(p Priority, f string, v ...interface{})
	Filename() string
	SetPriority(p Priority)
	GetPriority() Priority
	Flush()
}

// PriorityString returns a readable string (English word) that corresponds
// to a given log priority.
func PriorityString(p Priority) string {
	if p <= PriorityCrit {
		p = PriorityCrit
	}
	if p >= PriorityDebug {
		p = PriorityDebug
	}
	switch p {
	case PriorityCrit:
		return critString
	case PriorityErr:
		return errString
	case PriorityWarning:
		return warningString
	case PriorityNotice:
		return noticeString
	case PriorityInfo:
		return infoString
	case PriorityDebug:
		return debugString
	}

	return unknownString
}

// LoggerCrit logs a message with critical level. No formatting.
func LoggerCrit(l Logger, v ...interface{}) {
	l.Log(PriorityCrit, v...)
}

// LoggerCritf logs a message with critical level. Formatting "à la" printf.
func LoggerCritf(l Logger, f string, v ...interface{}) {
	l.Logf(PriorityCrit, f, v...)
}

// LoggerErr logs a message with error level. No formatting.
func LoggerErr(l Logger, v ...interface{}) {
	l.Log(PriorityErr, v...)
}

// LoggerErrf logs a message with error level. Formatting "à la" printf.
func LoggerErrf(l Logger, f string, v ...interface{}) {
	l.Logf(PriorityErr, f, v...)
}

// LoggerWarning logs a message with warning level. No formatting.
func LoggerWarning(l Logger, v ...interface{}) {
	l.Log(PriorityWarning, v...)
}

// LoggerWarningf logs a message with warning level. Formatting "à la" printf.
func LoggerWarningf(l Logger, f string, v ...interface{}) {
	l.Logf(PriorityWarning, f, v...)
}

// LoggerNotice logs a message with notice level. No formatting.
func LoggerNotice(l Logger, v ...interface{}) {
	l.Log(PriorityNotice, v...)
}

// LoggerNoticef logs a message with notice level. Formatting "à la" printf.
func LoggerNoticef(l Logger, f string, v ...interface{}) {
	l.Logf(PriorityNotice, f, v...)
}

// LoggerInfo logs a message with info level. No formatting.
func LoggerInfo(l Logger, v ...interface{}) {
	l.Log(PriorityInfo, v...)
}

// LoggerInfof logs a message with info level. Formatting "à la" printf.
func LoggerInfof(l Logger, f string, v ...interface{}) {
	l.Logf(PriorityInfo, f, v...)
}

// LoggerDebug logs a message with debug level. No formatting.
func LoggerDebug(l Logger, v ...interface{}) {
	l.Log(PriorityDebug, v...)
}

// LoggerDebugf logs a message with debug level. Formatting "à la" printf.
func LoggerDebugf(l Logger, f string, v ...interface{}) {
	l.Logf(PriorityDebug, f, v...)
}
