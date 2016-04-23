// This file was generated with level_generator. DO NOT EDIT

package verbose

import (
	"os"
	"fmt"
)

// Debug - Log Debug message
func (l *Logger) Debug(m string) {
    l.Log(LogLevelDebug, m)
    return
}

// Debugf - Log formatted Debug message
func (l *Logger) Debugf(m string, v ...interface{}) {
    l.Log(LogLevelDebug, fmt.Sprintf(m, v...))
    return
}

// Info - Log Info message
func (l *Logger) Info(m string) {
    l.Log(LogLevelInfo, m)
    return
}

// Infof - Log formatted Info message
func (l *Logger) Infof(m string, v ...interface{}) {
    l.Log(LogLevelInfo, fmt.Sprintf(m, v...))
    return
}

// Notice - Log Notice message
func (l *Logger) Notice(m string) {
    l.Log(LogLevelNotice, m)
    return
}

// Noticef - Log formatted Notice message
func (l *Logger) Noticef(m string, v ...interface{}) {
    l.Log(LogLevelNotice, fmt.Sprintf(m, v...))
    return
}

// Warning - Log Warning message
func (l *Logger) Warning(m string) {
    l.Log(LogLevelWarning, m)
    return
}

// Warningf - Log formatted Warning message
func (l *Logger) Warningf(m string, v ...interface{}) {
    l.Log(LogLevelWarning, fmt.Sprintf(m, v...))
    return
}

// Error - Log Error message
func (l *Logger) Error(m string) {
    l.Log(LogLevelError, m)
    return
}

// Errorf - Log formatted Error message
func (l *Logger) Errorf(m string, v ...interface{}) {
    l.Log(LogLevelError, fmt.Sprintf(m, v...))
    return
}

// Critical - Log Critical message
func (l *Logger) Critical(m string) {
    l.Log(LogLevelCritical, m)
    return
}

// Criticalf - Log formatted Critical message
func (l *Logger) Criticalf(m string, v ...interface{}) {
    l.Log(LogLevelCritical, fmt.Sprintf(m, v...))
    return
}

// Alert - Log Alert message
func (l *Logger) Alert(m string) {
    l.Log(LogLevelAlert, m)
    return
}

// Alertf - Log formatted Alert message
func (l *Logger) Alertf(m string, v ...interface{}) {
    l.Log(LogLevelAlert, fmt.Sprintf(m, v...))
    return
}

// Emergency - Log Emergency message
func (l *Logger) Emergency(m string) {
    l.Log(LogLevelEmergency, m)
    return
}

// Emergencyf - Log formatted Emergency message
func (l *Logger) Emergencyf(m string, v ...interface{}) {
    l.Log(LogLevelEmergency, fmt.Sprintf(m, v...))
    return
}

// Fatal - Log Fatal message
func (l *Logger) Fatal(m string) {
    l.Log(LogLevelFatal, m)
    os.Exit(1)
    return
}

// Fatalf - Log formatted Fatal message
func (l *Logger) Fatalf(m string, v ...interface{}) {
    l.Log(LogLevelFatal, fmt.Sprintf(m, v...))
    os.Exit(1)
    return
}
