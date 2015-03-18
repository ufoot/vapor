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
	"bufio"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"path"
	"sync"
	"ufoot.org/vapor/vpbuild"
)

const basename = "log.txt"
const stdoutPriority = PriorityNotice
const filePriority = PriorityDebug
const syslogPriority = PriorityWarning
const flushPriority = PriorityNotice

type stdoutWriter struct {
}

func (sw stdoutWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", string(p))
	return len(p), nil
}

// The default logger, implements Logger. It basically logs informations
// in three places, which are the console (stdout), a log file (typically
// placed in the user's home directory) and syslog. The log file contains
// everything while the console and syslog will only display important
// messages. Under the hood, it uses the log.Logger object, so it's safe
// to call it in a multithread environment. It also uses bufferized streams
// so you should be able to brutalize it with massive data without a too
// big slow down. On important messages (notice and above) it's in autoflush
// mode so it could typically slow down in such cases but hey, those messages
// should be rare by design. A set of functions which require no args but
// what you need to log are here if you don't want to carry the logger object
// arround between all func calls. The implementation just initializes an
// internal shared global object. As it's safe to call it concurrently,
// it should fit most cases.
type Log struct {
	filename      string
	p             Priority
	f             *os.File
	w             io.Writer
	file_buffer   *bufio.Writer
	stdout_buffer *bufio.Writer
	file_logger   *log.Logger
	stdout_logger *log.Logger
	syslog_logger *log.Logger
	flush_mutex   sync.Mutex
}

// Constructs a new log object, note that this is called under the hood
// by the global shared object constructor.
func NewLog(program string) *Log {
	var logger Log
	var err error
	var s stdoutWriter

	prefix := fmt.Sprintf("%s: ", program)
	logger.filename = path.Join(Home(program), basename)
	logger.p = PriorityInfo

	logger.f, err = os.Create(logger.filename)
	if err != nil {
		panic(err)
	}

	logger.file_buffer = bufio.NewWriter(io.Writer(logger.f))
	logger.stdout_buffer = bufio.NewWriter(s)

	logger.file_logger = log.New(logger.file_buffer, prefix, log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.stdout_logger = log.New(logger.stdout_buffer, prefix, log.LstdFlags)
	logger.syslog_logger, err = syslog.NewLogger(syslog.Priority(int(syslogPriority))|syslog.LOG_SYSLOG, log.LstdFlags|log.Lshortfile)
	if err != nil {
		panic(err)
	}

	logger.Logf(PriorityNotice, "Log file: %s", logger.filename)
	logger.Flush()

	return &logger
}

// Logs a message on all relevant channels, no EOL
// added, if you want one, do provide one.
func (l *Log) Log(p Priority, v ...interface{}) {
	if p <= l.p {
		if p <= stdoutPriority {
			l.stdout_logger.Print(v...)
		}
		if p <= filePriority {
			l.file_logger.Print(v...)
		}
		if p <= syslogPriority {
			l.syslog_logger.Print(v...)
		}
		if p <= flushPriority {
			l.Flush()
		}
	}
}

// Logs a message on all relevant channels, using a printf-like
// syntax, no EOL  added, if you want one
// do provide one.
func (l *Log) Logf(p Priority, format string, v ...interface{}) {
	if p <= l.p {
		if p <= stdoutPriority {
			l.stdout_logger.Printf(format, v...)
		}
		if p <= filePriority {
			l.file_logger.Printf(format, v...)
		}
		if p <= syslogPriority {
			l.syslog_logger.Printf(format, v...)
		}
		if p <= flushPriority {
			l.Flush()
		}
	}
}

// Returns the path of the log file.
func (l *Log) Filename() string {
	return l.filename
}

// Sets the priority above which message won't be displayed any more.
func (l *Log) SetPriority(p Priority) {
	l.p = p
}

// Returns the priority under which message won't be displayed any more.
func (l *Log) GetPriority() Priority {
	return l.p
}

// Flushes all channels for which it makes sense to be flushed.
// This is automatically called if priority is CRIT, ERR, WARNING or NOTICE.
func (l *Log) Flush() {
	l.flush_mutex.Lock()
	l.file_buffer.Flush()
	l.stdout_buffer.Flush()
	// This is why we use a Mutex and a Lock, while buffers buried
	// under the log.Logger API might be thread-safe, the file
	// direct access is another story.
	l.f.Sync()
	l.flush_mutex.Unlock()
}

var globalLog *Log

func getGlobalLog(program string) *Log {
	if globalLog == nil {
		globalLog = NewLog(program)
	}

	return globalLog
}

// Initializes the log system. This is not mandatory, you might use
// functions such as LogWarning right away, the log file will be
// opened on-the-fly if needed. However, you might prefer to have the
// file opened at the very beginning of the program, without waiting
// for an artificial event. This is why this function is here.
// It can also be used to force a given file to be used, by default
// the program name "vapor" is used but it can be overridden.
func LogInit(program string) {
	getGlobalLog(program)
}

// Global critical log, no EOL added.
func LogCrit(v ...interface{}) {
	LoggerCrit(getGlobalLog(vpbuild.PACKAGE_TARNAME), v...)
}

// Global critical log, printf like interface.
func LogCritf(f string, v ...interface{}) {
	LoggerCritf(getGlobalLog(vpbuild.PACKAGE_TARNAME), f, v...)
}

// Global error log, no EOL added.
func LogErr(v ...interface{}) {
	LoggerErr(getGlobalLog(vpbuild.PACKAGE_TARNAME), v...)
}

// Global error log, printf like interface.
func LogErrf(f string, v ...interface{}) {
	LoggerErrf(getGlobalLog(vpbuild.PACKAGE_TARNAME), f, v...)
}

// Global warning log, no EOL added.
func LogWarning(v ...interface{}) {
	LoggerWarning(getGlobalLog(vpbuild.PACKAGE_TARNAME), v...)
}

// Global warning log, printf like interface.
func LogWarningf(f string, v ...interface{}) {
	LoggerWarningf(getGlobalLog(vpbuild.PACKAGE_TARNAME), f, v...)
}

// Global notice log, no EOL added.
func LogNotice(v ...interface{}) {
	LoggerNotice(getGlobalLog(vpbuild.PACKAGE_TARNAME), v...)
}

// Global notice log, printf like interface.
func LogNoticef(f string, v ...interface{}) {
	LoggerNoticef(getGlobalLog(vpbuild.PACKAGE_TARNAME), f, v...)
}

// Global info log, no EOL added.
func LogInfo(v ...interface{}) {
	LoggerInfo(getGlobalLog(vpbuild.PACKAGE_TARNAME), v...)
}

// Global info log, printf like interface.
func LogInfof(f string, v ...interface{}) {
	LoggerInfof(getGlobalLog(vpbuild.PACKAGE_TARNAME), f, v...)
}

// Global debug log, no EOL added.
func LogDebug(v ...interface{}) {
	LoggerDebug(getGlobalLog(vpbuild.PACKAGE_TARNAME), v...)
}

// Global debug log, printf like interface.
func LogDebugf(f string, v ...interface{}) {
	LoggerDebugf(getGlobalLog(vpbuild.PACKAGE_TARNAME), f, v...)
}

// Returns the path of the file used for global, default logging.
func LogFilename() string {
	return getGlobalLog(vpbuild.PACKAGE_TARNAME).Filename()
}

// Sets the global, default logging level.
func LogSetPriority(p Priority) {
	getGlobalLog(vpbuild.PACKAGE_TARNAME).SetPriority(p)
}

// Returns the global, default logging level.
func LogGetPriority() Priority {
	return getGlobalLog(vpbuild.PACKAGE_TARNAME).GetPriority()
}

// Flushes the global logging system, more precisely, flushes
// stdout and the log file.
func LogFlush() {
	getGlobalLog(vpbuild.PACKAGE_TARNAME).Flush()
}
