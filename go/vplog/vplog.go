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
	"bufio"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"strings"
	"sync"
)

const basename = "log.txt"
const stderrPriority = PriorityDebug
const problemPriority = PriorityWarning
const syslogPriority = PriorityWarning
const flushPriority = PriorityNotice
const sepHeaderContent = "-"

// 1st parameter of the Logger.Output func
const outputCalldepth = 4

type stderrWriter struct {
}

func (sw stderrWriter) Write(p []byte) (n int, err error) {
	fmt.Fprintf(os.Stderr, "%s", string(p))
	return len(p), nil
}

// Log is a default implementation of the Logger interface.
// It basically logs informations in two places,
// which are the console (stderr) and syslog.
type Log struct {
	p            Priority
	stderrBuffer *bufio.Writer
	stderrLogger *log.Logger
	syslogLogger *log.Logger
	flushMutex   sync.Mutex
}

func newLog(program string, writer io.Writer) *Log {
	var logger Log
	var err error

	prefix := fmt.Sprintf("%s: ", program)
	logger.p = PriorityInfo

	logger.stderrBuffer = bufio.NewWriter(writer)

	logger.stderrLogger = log.New(logger.stderrBuffer, prefix, log.Ltime)
	logger.syslogLogger, err = syslog.NewLogger(syslog.Priority(int(syslogPriority))|syslog.LOG_SYSLOG, log.Lshortfile)
	if err != nil {
		panic(err)
	}

	logger.Flush()

	return &logger
}

// NewLog Constructs a new log object, note that this is called under the hood
// by the global shared object constructor.
func NewLog(program string) *Log {
	return (newLog(program, os.Stderr))
}

func (l *Log) output(p Priority, ps, ms string) {
	var line string

	for _, line = range strings.Split(ms, "\n") {
		if len(line) > 0 {
			line = sepHeaderContent + " " + line
			if p <= stderrPriority {
				if p <= problemPriority {
					l.stderrLogger.Output(outputCalldepth, ps+line)
				} else {
					l.stderrLogger.Output(outputCalldepth, line)
				}
			}
			if p <= syslogPriority {
				l.syslogLogger.Output(outputCalldepth, ps+line)
			}
			if p <= flushPriority {
				l.Flush()
			}
		}
	}
}

// Logp logs a message on all relevant channels.
// EOL is added at the end, you do not need to provide it.
func (l *Log) Logp(p Priority, v ...interface{}) {
	if p <= l.p {
		ps := PriorityString(p) + " "
		ms := fmt.Sprintln(v...)
		l.output(p, ps, ms)
	}
}

// Logpf logs a message on all relevant channels, using a printf-like syntax.
// EOL is added at the end, you do not need to provide it.
func (l *Log) Logpf(p Priority, format string, v ...interface{}) {
	if p <= l.p {
		ps := PriorityString(p) + " "
		ms := fmt.Sprintf(format, v...) + "\n"
		l.output(p, ps, ms)
	}
}

// SetPriority sets the priority above which message won't be displayed any more.
func (l *Log) SetPriority(p Priority) {
	l.p = p
}

// GetPriority returns the priority under which message won't be displayed any more.
func (l *Log) GetPriority() Priority {
	return l.p
}

// Flush flushes all channels for which it makes sense to be flushed.
// This is automatically called if priority is CRIT, ERR, WARNING or NOTICE.
func (l *Log) Flush() {
	l.flushMutex.Lock()
	l.stderrBuffer.Flush()
	l.flushMutex.Unlock()
}

var globalLog *Log

func getGlobalLog(program string) *Log {
	if globalLog == nil {
		globalLog = NewLog(program)
	}

	return globalLog
}

func logInitWithWriter(program string, writer io.Writer) {
	globalLog = newLog(program, writer)
}

// LogInit initializes the log system. This is not mandatory, you might use
// functions such as LogWarning right away, the log file will be
// opened on-the-fly if needed. However, you might prefer to have the
// file opened at the very beginning of the program, without waiting
// for an artificial event. This is why this function is here.
// It can also be used to force a given file to be used, by default
// the program name "vapor" is used but it can be overridden.
func LogInit(program string) {
	getGlobalLog(program)
}

// LogCrit logs a critical message, no formatting.
// Uses the default global logging backend.
func LogCrit(v ...interface{}) {
	LoggerCrit(getGlobalLog(PackageTarname), v...)
}

// LogCritf logs a critical message, formatting "à la" printf.
// Uses the default global logging backend.
func LogCritf(f string, v ...interface{}) {
	LoggerCritf(getGlobalLog(PackageTarname), f, v...)
}

// LogErr logs an error message, no formatting.
// Uses the default global logging backend.
func LogErr(v ...interface{}) {
	LoggerErr(getGlobalLog(PackageTarname), v...)
}

// LogErrf logs an error message, formatting "à la" printf.
// Uses the default global logging backend.
func LogErrf(f string, v ...interface{}) {
	LoggerErrf(getGlobalLog(PackageTarname), f, v...)
}

// LogWarning logs a warning message, no formatting.
// Uses the default global logging backend.
func LogWarning(v ...interface{}) {
	LoggerWarning(getGlobalLog(PackageTarname), v...)
}

// LogWarningf logs a warning message, formatting "à la" printf.
// Uses the default global logging backend.
func LogWarningf(f string, v ...interface{}) {
	LoggerWarningf(getGlobalLog(PackageTarname), f, v...)
}

// LogNotice logs a notice message, no formatting.
// Uses the default global logging backend.
func LogNotice(v ...interface{}) {
	LoggerNotice(getGlobalLog(PackageTarname), v...)
}

// LogNoticef logs a notice message, formatting "à la" printf.
// Uses the default global logging backend.
func LogNoticef(f string, v ...interface{}) {
	LoggerNoticef(getGlobalLog(PackageTarname), f, v...)
}

// LogInfo logs an information message, no formatting.
// Uses the default global logging backend.
func LogInfo(v ...interface{}) {
	LoggerInfo(getGlobalLog(PackageTarname), v...)
}

// LogInfof logs an information message, formatting "à la" printf.
// Uses the default global logging backend.
func LogInfof(f string, v ...interface{}) {
	LoggerInfof(getGlobalLog(PackageTarname), f, v...)
}

// LogDebug logs a debug message, no formatting.
// Uses the default global logging backend.
func LogDebug(v ...interface{}) {
	LoggerDebug(getGlobalLog(PackageTarname), v...)
}

// LogDebugf logs a debug message, formatting "à la" printf.
// Uses the default global logging backend.
func LogDebugf(f string, v ...interface{}) {
	LoggerDebugf(getGlobalLog(PackageTarname), f, v...)
}

// LogSetPriority sets the global, default logging level.
// Uses the default global logging backend.
func LogSetPriority(p Priority) {
	getGlobalLog(PackageTarname).SetPriority(p)
}

// LogGetPriority returns the global, default logging level.
// Uses the default global logging backend.
func LogGetPriority() Priority {
	return getGlobalLog(PackageTarname).GetPriority()
}

// LogFlush flushes the global logging system, more precisely, flushes
// stderr.
func LogFlush() {
	getGlobalLog(PackageTarname).Flush()
}
