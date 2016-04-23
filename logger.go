package verbose

import (
	"fmt"
	"os"
	"sync"
)

type LogLevel int

func (l LogLevel) String() string {
	if s, ok := levelString[l]; ok {
		return s
	}
	return ""
}

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

type Color string

const (
	ColorReset   Color = "\033[0m"
	ColorRed     Color = "\033[31m"
	ColorGreen   Color = "\033[32m"
	ColorYellow  Color = "\033[33m"
	ColorBlue    Color = "\033[34m"
	ColorMagenta Color = "\033[35m"
	ColorCyan    Color = "\033[36m"
	ColorWhite   Color = "\033[37m"
	ColorGrey    Color = "\x1B[90m"
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
}

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

func (l *Logger) AddHandler(h Handler) {
	if h == nil {
		return
	}
	l.m.Lock()
	l.handlers = append(l.handlers, h)
	l.m.Unlock()
}

func (l *Logger) Close() {
	removeLogger(l)
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) Log(level LogLevel, msg string) {
	for _, h := range l.handlers {
		if h.Handles(level) {
			h.WriteLog(level, l.name, msg)
		}
	}
}

func (l *Logger) Debug(m string) {
	l.Log(LogLevelDebug, m)
	return
}

func (l *Logger) Debugf(m string, v ...interface{}) {
	l.Log(LogLevelDebug, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Info(m string) {
	l.Log(LogLevelInfo, m)
	return
}

func (l *Logger) Infof(m string, v ...interface{}) {
	l.Log(LogLevelInfo, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Notice(m string) {
	l.Log(LogLevelNotice, m)
	return
}

func (l *Logger) Noticef(m string, v ...interface{}) {
	l.Log(LogLevelNotice, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Warning(m string) {
	l.Log(LogLevelWarning, m)
	return
}

func (l *Logger) Warningf(m string, v ...interface{}) {
	l.Log(LogLevelWarning, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Error(m string) {
	l.Log(LogLevelError, m)
	return
}

func (l *Logger) Errorf(m string, v ...interface{}) {
	l.Log(LogLevelError, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Critical(m string) {
	l.Log(LogLevelCritical, m)
	return
}

func (l *Logger) Criticalf(m string, v ...interface{}) {
	l.Log(LogLevelCritical, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Alert(m string) {
	l.Log(LogLevelAlert, m)
	return
}

func (l *Logger) Alertf(m string, v ...interface{}) {
	l.Log(LogLevelAlert, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Emergency(m string) {
	l.Log(LogLevelEmergency, m)
	return
}

func (l *Logger) Emergencyf(m string, v ...interface{}) {
	l.Log(LogLevelEmergency, fmt.Sprintf(m, v...))
	return
}

func (l *Logger) Fatal(m string) {
	l.Log(LogLevelFatal, m)
	os.Exit(1)
	return
}

func (l *Logger) Fatalf(m string, v ...interface{}) {
	l.Log(LogLevelFatal, fmt.Sprintf(m, v...))
	os.Exit(1)
	return
}
