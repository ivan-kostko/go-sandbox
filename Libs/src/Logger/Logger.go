// Logger project Logger.go
// Package Logger provides a common interface and wrapper implementation for logging libraries.
// Contains predefined loggers: StderrLogger and MockLogger
package Logger

import (
	"fmt"
)

//go:generate stringer -type=Level

// Level specifies a level of severity. The available levels are the eight severities described in RFC 5424 and None
type Level int8

const (
	None      Level = iota - 1
	Emergency       //      Emergency: system is unusable
	Alert           //      Alert: action must be taken immediately
	Critical        //      Critical: critical conditions
	Error           //      Error: error conditions
	Warning         //      Warning: warning conditions
	Notice          //      Notice: normal but significant condition
	Info            //      Informational: informational messages
	Debug           //      Debug: debug-level messages
)

// ILogger is a common interface for logging.
type ILogger interface {
	// Emergency logs with an emergency level.
	Emergency(args ...interface{})

	// Emergencyf logs with an emergency level.
	// Arguments are handled in the manner of fmt.Printf.
	Emergencyf(format string, args ...interface{})

	// Alert logs with an alert level.
	Alert(args ...interface{})

	// Alertf logs with an alert level.
	// Arguments are handled in the manner of fmt.Printf.
	Alertf(format string, args ...interface{})

	// Critical logs with a critical level.
	Critical(args ...interface{})

	// Criticalf logs with a critical level.
	// Arguments are handled in the manner of fmt.Printf.
	Criticalf(format string, args ...interface{})

	// Error logs with an error level.
	Error(args ...interface{})

	// Errorf logs with an error level.
	// Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})

	// Warning logs with a warning level.
	Warning(args ...interface{})

	// Warningf logs with a warning level.
	// Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...interface{})

	// Notice logs with a notice level.
	Notice(args ...interface{})

	// Noticef logs with a notice level.
	// Arguments are handled in the manner of fmt.Printf.
	Noticef(format string, args ...interface{})

	// Info logs with an info level.
	Info(args ...interface{})

	// Infof logs with an info level.
	// Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})

	// Debug logs with a debug level.
	Debug(args ...interface{})

	// Debugf logs with a debug level.
	// Arguments are handled in the manner of fmt.Printf.
	Debugf(format string, args ...interface{})

	// Log logs at the level passed in argument.
	Log(level Level, args ...interface{})

	// Logf logs at the level passed in argument.
	// Arguments are handled in the manner of fmt.Printf.
	Logf(level Level, format string, args ...interface{})
}

// LogContainer represents adapter structure (satisfying ILogger interface) and redirecting all calls of helper functions to Log() function
// Could be used for mocking and quick simple intro of any logger
// NB: For production loggers it is better to create its own adapter
type LogContainer struct {
	internalLogFunction func(level Level, args ...interface{})
}

// LogContainer fabric
func GetNewLogContainer(intLog func(level Level, args ...interface{})) ILogger {
	lc := new(LogContainer)
	lc.internalLogFunction = intLog
	return lc
}

// Emergency logs with an emergency level
func (lc *LogContainer) Emergency(args ...interface{}) {
	lc.Log(Emergency, args...)
}

// Emergencyf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Emergencyf(format string, args ...interface{}) {
	lc.Log(Emergency, fmt.Sprintf(format, args...))
}

// Alert logs with an emergency level
func (lc *LogContainer) Alert(args ...interface{}) {
	lc.Log(Alert, args...)
}

// Alertf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Alertf(format string, args ...interface{}) {
	lc.Log(Alert, fmt.Sprintf(format, args...))
}

// Critical logs with an emergency level
func (lc *LogContainer) Critical(args ...interface{}) {
	lc.Log(Critical, args...)
}

// Criticalf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Criticalf(format string, args ...interface{}) {
	lc.Log(Critical, fmt.Sprintf(format, args...))
}

// Error logs with an emergency level
func (lc *LogContainer) Error(args ...interface{}) {
	lc.Log(Error, args...)
}

// Errorf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Errorf(format string, args ...interface{}) {
	lc.Log(Error, fmt.Sprintf(format, args...))
}

// Warning logs with an emergency level
func (lc *LogContainer) Warning(args ...interface{}) {
	lc.Log(Warning, args...)
}

// Warningf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Warningf(format string, args ...interface{}) {
	lc.Log(Warning, fmt.Sprintf(format, args...))
}

// Notice logs with an emergency level
func (lc *LogContainer) Notice(args ...interface{}) {
	lc.Log(Notice, args...)
}

// Noticef logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Noticef(format string, args ...interface{}) {
	lc.Log(Notice, fmt.Sprintf(format, args...))
}

// Info logs with an emergency level
func (lc *LogContainer) Info(args ...interface{}) {
	lc.Log(Info, args...)
}

// Infof logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Infof(format string, args ...interface{}) {
	lc.Log(Info, fmt.Sprintf(format, args...))
}

// Debug logs with an emergency level
func (lc *LogContainer) Debug(args ...interface{}) {
	lc.Log(Debug, args...)
}

// Debugf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Debugf(format string, args ...interface{}) {
	lc.Log(Debug, fmt.Sprintf(format, args...))
}

// Log logs with an emergency level
func (lc *LogContainer) Log(level Level, args ...interface{}) {
	lc.internalLogFunction(level, args...)
}

// Logf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogContainer) Logf(level Level, format string, args ...interface{}) {
	lc.Log(level, fmt.Sprintf(format, args...))
}

// Represents configuration to create a new Logger
type LoggerConfig struct {
	Prefix string
}
