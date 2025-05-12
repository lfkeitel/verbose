package verbose

import (
	"fmt"
	"testing"
)

// testHandler is a special handler that will check for a correct LogLevel and message
type testHandler struct {
	tester     *testing.T
	msg        string
	loggerName string
	level      LogLevel
}

func newTestHandler(t *testing.T, l LogLevel, name string, msg ...interface{}) *testHandler {
	message := fmt.Sprintln(msg...)
	return &testHandler{
		tester:     t,
		msg:        message[:len(message)-1],
		loggerName: name,
		level:      l,
	}
}

func (t *testHandler) Handles(l LogLevel) bool { return (l == t.level) }
func (_ *testHandler) Close()                  {}

func (t *testHandler) WriteLog(e *Entry) {
	if e.Level != t.level {
		t.tester.Errorf("Handled incorrect level. Expected %d, got %d", t.level, e.Level)
	}
	if e.Logger.Name != t.loggerName {
		t.tester.Errorf("Incorrect logger name. Expected %s, got %s", t.loggerName, e.Logger.Name)
	}
	if e.Message != t.msg {
		t.tester.Errorf("Incorrect message. Expected %s, got %s", t.msg, e.Message)
	}
}

// TestLoggerHandlers ensure handlers are manipulated correctly
func TestLoggerHandlers(t *testing.T) {
	l1 := New()
	l1.AddTransport(&testHandler{})
	l1.AddTransport(&testHandler{})

	if len(l1.transports) != 2 {
		t.Errorf("Not enough handlers. Expected 2, got %d", len(l1.transports))
	}
}

// TestLoggingLevels creates a custom handler for every level to make sure everything
// is being processed correctly. Each message is a custom message made up for a
// common message and the LogLevel.String().
func TestLoggingLevels(t *testing.T) {
	testMsg := "The space ship is coming"
	logger := New()
	logger.Name = "logger1"
	logger.AddTransport(newTestHandler(t, LogLevelEmergency, "logger1", testMsg, LogLevelEmergency.String()))
	logger.AddTransport(newTestHandler(t, LogLevelAlert, "logger1", testMsg, LogLevelAlert.String()))
	logger.AddTransport(newTestHandler(t, LogLevelCritical, "logger1", testMsg, LogLevelCritical.String()))
	logger.AddTransport(newTestHandler(t, LogLevelError, "logger1", testMsg, LogLevelError.String()))
	logger.AddTransport(newTestHandler(t, LogLevelWarning, "logger1", testMsg, LogLevelWarning.String()))
	logger.AddTransport(newTestHandler(t, LogLevelNotice, "logger1", testMsg, LogLevelNotice.String()))
	logger.AddTransport(newTestHandler(t, LogLevelInfo, "logger1", testMsg, LogLevelInfo.String()))
	logger.AddTransport(newTestHandler(t, LogLevelDebug, "logger1", testMsg, LogLevelDebug.String()))

	logger.Emergency(fmt.Sprintf("%s %s", testMsg, LogLevelEmergency.String()))
	logger.Alert(fmt.Sprintf("%s %s", testMsg, LogLevelAlert.String()))
	logger.Critical(fmt.Sprintf("%s %s", testMsg, LogLevelCritical.String()))
	logger.Error(fmt.Sprintf("%s %s", testMsg, LogLevelError.String()))
	logger.Warning(fmt.Sprintf("%s %s", testMsg, LogLevelWarning.String()))
	logger.Notice(fmt.Sprintf("%s %s", testMsg, LogLevelNotice.String()))
	logger.Info(fmt.Sprintf("%s %s", testMsg, LogLevelInfo.String()))
	logger.Debug(fmt.Sprintf("%s %s", testMsg, LogLevelDebug.String()))
}
