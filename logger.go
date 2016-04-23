package verbose

import (
	"fmt"
	"os"
	"sync"
)

// LogLevel is used to compare levels in a consistant manner
type LogLevel int

// String returns the stringified version of LogLevel.
// I.e., "Error" for LogLevelError, and "Debug" for LogLevelDebug
// It will return an empty string for any undefined level.
func (l LogLevel) String() string {
	if s, ok := levelString[l]; ok {
		return s
	}
	return ""
}

// These are the defined log levels
const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelNotice
	LogLevelWarning
	LogLevelError
	LogLevelCritical
	LogLevelAlert
	LogLevelEmergency
	LogLevelFatal
	LogLevelCustom
)

// LogLevel to stringified versions
var levelString = map[LogLevel]string{
	LogLevelDebug:     "Debug",
	LogLevelInfo:      "Info",
	LogLevelNotice:    "Notice",
	LogLevelWarning:   "Warning",
	LogLevelError:     "Error",
	LogLevelCritical:  "Critical",
	LogLevelAlert:     "Alert",
	LogLevelEmergency: "Emergency",
	LogLevelFatal:     "Fatal",
	LogLevelCustom:    "Custom",
}

// Color is an escaped color code for the terminal
type Color string

// Pre-defined colors
const (
	ColorReset   Color = "\033[0m"
	ColorRed     Color = "\033[31m"
	ColorGreen   Color = "\033[32m"
	ColorYellow  Color = "\033[33m"
	ColorBlue    Color = "\033[34m"
	ColorMagenta Color = "\033[35m"
	ColorCyan    Color = "\033[36m"
	ColorWhite   Color = "\033[37m"
	ColorGrey    Color = "\033[90m"
)

var (
	loggers      map[string]*Logger
	loggersMutex = sync.RWMutex{}
)

func init() {
	loggers = make(map[string]*Logger)
}

func addLogger(l *Logger) {
	loggersMutex.Lock()
	loggers[l.name] = l
	loggersMutex.Unlock()
}

func getLogger(n string) *Logger {
	loggersMutex.RLock()
	l := loggers[n]
	loggersMutex.RUnlock()
	return l
}

func removeLogger(l *Logger) {
	loggersMutex.Lock()
	delete(loggers, l.name)
	loggersMutex.Unlock()
}

// A Handler is an object that can be used by the Logger to log a message
type Handler interface {
	// Handles returns if it wants to handle a particular log level
	// This can be used to suppress the higher log levels in production
	Handles(level LogLevel) bool

	// WriteLog actually logs the message using any system the Handler wishes.
	// The Handler must accept a LogLevel l which can be used for furthur processing
	// of specific levels, the name of the logger, and the log message.
	WriteLog(l LogLevel, name, message string)

	// Close is used to give a handler a chance to close any open resources
	Close()
}

// A Logger takes a message and writes it to as many handlers as possible
type Logger struct {
	name     string
	handlers []Handler
	m        sync.RWMutex
}

// New will create a new Logger with name n. If with the same name
// already exists, it will be replaced with the new logger.
func New(n string) *Logger {
	l := &Logger{
		name: n,
		m:    sync.RWMutex{},
	}
	addLogger(l)
	return l
}

// Get returns and existing logger with name n or a new Logger if one
// doesn't exist. To ensure Loggers are never overwritten, it may be safer to
// always use this method.
func Get(n string) *Logger {
	l := getLogger(n)
	if l != nil {
		return l
	}
	return New(n)
}

// AddHandler will add Handler h to the logger. Handlers cannot be removed.
func (l *Logger) AddHandler(h Handler) {
	if h == nil {
		return
	}
	l.m.Lock()
	l.handlers = append(l.handlers, h)
	l.m.Unlock()
}

// Close calls Close() on all the handlers then removes itself from the logger registry
func (l *Logger) Close() {
	for _, h := range l.handlers {
		h.Close()
	}
	removeLogger(l)
}

// Name returns the name of the logger
func (l *Logger) Name() string {
	return l.name
}

// Log is the generic function to log a message with the handlers.
// All other logging functions are simply wrappers around this.
func (l *Logger) Log(level LogLevel, msg string) {
	for _, h := range l.handlers {
		if h.Handles(level) {
			h.WriteLog(level, l.name, msg)
		}
	}
}

// Debug - Log debug message
func (l *Logger) Debug(m string) {
	l.Log(LogLevelDebug, m)
	return
}

// Debugf - Log formatted debug message
func (l *Logger) Debugf(m string, v ...interface{}) {
	l.Log(LogLevelDebug, fmt.Sprintf(m, v...))
	return
}

// Info - Log informational message
func (l *Logger) Info(m string) {
	l.Log(LogLevelInfo, m)
	return
}

// Infof - Log formatted informational message
func (l *Logger) Infof(m string, v ...interface{}) {
	l.Log(LogLevelInfo, fmt.Sprintf(m, v...))
	return
}

// Notice - Log notice message
func (l *Logger) Notice(m string) {
	l.Log(LogLevelNotice, m)
	return
}

// Noticef - Log formatted notice message
func (l *Logger) Noticef(m string, v ...interface{}) {
	l.Log(LogLevelNotice, fmt.Sprintf(m, v...))
	return
}

// Warning - Log warning message
func (l *Logger) Warning(m string) {
	l.Log(LogLevelWarning, m)
	return
}

// Warningf - Log formatted warning message
func (l *Logger) Warningf(m string, v ...interface{}) {
	l.Log(LogLevelWarning, fmt.Sprintf(m, v...))
	return
}

// Error - Log error message
func (l *Logger) Error(m string) {
	l.Log(LogLevelError, m)
	return
}

// Errorf - Log formatted error message
func (l *Logger) Errorf(m string, v ...interface{}) {
	l.Log(LogLevelError, fmt.Sprintf(m, v...))
	return
}

// Critical - Log critical message
func (l *Logger) Critical(m string) {
	l.Log(LogLevelCritical, m)
	return
}

// Criticalf - Log formatted critical message
func (l *Logger) Criticalf(m string, v ...interface{}) {
	l.Log(LogLevelCritical, fmt.Sprintf(m, v...))
	return
}

// Alert - Log alert message
func (l *Logger) Alert(m string) {
	l.Log(LogLevelAlert, m)
	return
}

// Alertf - Log formatted alert message
func (l *Logger) Alertf(m string, v ...interface{}) {
	l.Log(LogLevelAlert, fmt.Sprintf(m, v...))
	return
}

// Emergency - Log emergency message
func (l *Logger) Emergency(m string) {
	l.Log(LogLevelEmergency, m)
	return
}

// Emergencyf - Log formatted emergency message
func (l *Logger) Emergencyf(m string, v ...interface{}) {
	l.Log(LogLevelEmergency, fmt.Sprintf(m, v...))
	return
}

// Fatal - Log fatal message, calls os.Exit(1)
func (l *Logger) Fatal(m string) {
	l.Log(LogLevelFatal, m)
	os.Exit(1)
	return
}

// Fatalf - Log formatted fatal message, calls os.Exit(1)
func (l *Logger) Fatalf(m string, v ...interface{}) {
	l.Log(LogLevelFatal, fmt.Sprintf(m, v...))
	os.Exit(1)
	return
}
