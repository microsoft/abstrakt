// logger.go was adapted from https://github.com/microsoft/fabrikate

package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// lock is a global mutex lock to gain control of logrus.<SetLevel|SetOutput>
var lock = sync.RWMutex{}
var output io.Writer

// SetOutput sets the logger output. Used primarily for testing
func SetOutput(out io.Writer) {
	output = out
}

func outputStdOut() io.Writer {
	if output != nil {
		return output
	}
	return os.Stdout
}

func outputStdErr() io.Writer {
	if output != nil {
		return output
	}
	return os.Stderr
}

// SetLevelDebug sets the standard logger level to Debug
func SetLevelDebug() {
	lock.Lock()
	logrus.SetLevel(logrus.DebugLevel)
	lock.Unlock()
}

// SetLevelInfo sets the standard logger level to Info
func SetLevelInfo() {
	lock.Lock()
	logrus.SetLevel(logrus.InfoLevel)
	lock.Unlock()
}

// Trace logs a message at level Trace to stdout.
func Trace(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Trace(args...)
	lock.Unlock()
}

// Debug logs a message at level Debug to stdout.
func Debug(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Debug(args...)
	lock.Unlock()
}

// Debugf formats according to a format specifier and logs message at level Debug to stdout.
func Debugf(format string, args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Debugf(format, args...)
	lock.Unlock()
}

// Info logs a message at level Info to stdout.
func Info(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Info(args...)
	lock.Unlock()
}

// Infof logs a formatted message at level Info to stdout.
func Infof(format string, args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Infof(format, args...)
	lock.Unlock()
}

// Output logs a message to stdout without formatting.
func Output(args ...interface{}) {
	lock.Lock()
	fmt.Printf("%v", args...)
	fmt.Println(args...)
	lock.Unlock()
}

// Outputf formats according to a format specifier and writes to standard output.
func Outputf(format string, args ...interface{}) {
	lock.Lock()
	fmt.Printf(format, args...)
	lock.Unlock()
}

// Warn logs a message at level Warn to stdout.
func Warn(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Warn(args...)
	lock.Unlock()
}

// Warnf logs a formatted message at level Warn to stdout.
func Warnf(format string, args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdOut())
	logrus.Warnf(format, args...)
	lock.Unlock()
}

// Error logs a message at level Error to stderr.
func Error(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdErr())
	logrus.Error(args...)
	lock.Unlock()
}

// Errorf logs a formatted message at level Error to stderr.
func Errorf(format string, args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdErr())
	logrus.Errorf(format, args...)
	lock.Unlock()
}

// Fatal logs a message at level Fatal to stderr then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdErr())
	logrus.Fatal(args...)
	lock.Unlock()
}

// Fatalf logs a formatted message at level Fatal to stderr then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdErr())
	logrus.Fatalf(format, args...)
	lock.Unlock()
}

// Panic logs a message at level Panic to stderr; calls panic() after logging.
func Panic(args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdErr())
	logrus.Panic(args...)
	lock.Unlock()
}

// Panicf logs a formatted message at level Panic to stderr; calls panic() after logging.
func Panicf(format string, args ...interface{}) {
	lock.Lock()
	logrus.SetOutput(outputStdErr())
	logrus.Panicf(format, args...)
	lock.Unlock()
}

// PrintBuffer prints from buffer to either Debug or Info
func PrintBuffer(buffer *bytes.Buffer, logDebug bool) {
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		message := scanner.Text()
		if logDebug {
			Debug(message)
		} else {
			Info(message)
		}
	}
}

func init() {
	// Setup logger defaults
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	formatter.FullTimestamp = true
	logrus.SetFormatter(formatter)
	logrus.SetOutput(outputStdOut()) // Set output to stdout; set to stderr by default
	logrus.SetLevel(logrus.InfoLevel)
}
