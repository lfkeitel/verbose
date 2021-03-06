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

func (t *testHandler) Handles(l LogLevel) bool  { return (l == t.level) }
func (_ *testHandler) SetFormatter(_ Formatter) {}
func (_ *testHandler) Close()                   {}
func (_ *testHandler) SetLevel(_ LogLevel)      {}
func (_ *testHandler) SetMinLevel(_ LogLevel)   {}
func (_ *testHandler) SetMaxLevel(_ LogLevel)   {}

func (t *testHandler) WriteLog(e *Entry) {
	if e.Level != t.level {
		t.tester.Errorf("Handled incorrect level. Expected %d, got %d", t.level, e.Level)
	}
	if e.Logger.Name() != t.loggerName {
		t.tester.Errorf("Incorrect logger name. Expected %s, got %s", t.loggerName, e.Logger.Name())
	}
	if e.Message != t.msg {
		t.tester.Errorf("Incorrect message. Expected %s, got %s", t.msg, e.Message)
	}
}

// Delete all loggers
func clearLoggers() {
	if len(loggers) == 0 {
		return
	}
	loggersMutex.Lock()
	loggers = make(map[string]*Logger)
	loggersMutex.Unlock()
}

func TestLoggerNewGet(t *testing.T) {
	clearLoggers()
	l1 := New("logger 1")
	l2 := Get("logger 1")

	if l1.name != l2.name {
		t.Errorf("New and Get didn't return the same logger. Expected %s, got %s", l1.name, l2.name)
	}
}

// TestLoggerHandlers ensure handlers are manipulated correctly
func TestLoggerHandlers(t *testing.T) {
	clearLoggers()
	l1 := Get("logger 1")
	l1.AddHandler("h", &testHandler{})
	l1.AddHandler("h1", &testHandler{})

	if len(l1.handlers) != 2 {
		t.Errorf("Not enough handlers. Expected 2, got %d", len(l1.handlers))
	}

	h := l1.GetHandler("h")
	if h == nil {
		t.Error("No handler returned")
	}

	l1.RemoveHandler("h")
	if len(l1.handlers) != 1 {
		t.Errorf("Incorrect number of handlers. Expected 1, got %d", len(l1.handlers))
	}
}

// TestNewOverwrites makes sure that New() returns a new instance of Logger
// regardless of if it exists already
func TestNewOverwrites(t *testing.T) {
	clearLoggers()
	logger := Get("logger 1")
	logger.AddHandler("h", &testHandler{})
	logger.AddHandler("h1", &testHandler{})
	handlers := len(logger.handlers)

	newLogger := New("logger 1")
	if len(newLogger.handlers) == handlers {
		t.Error("New didn't return a new logger. Returned logger with same number of handlers")
	}
}

// TestLoggingLevels creates a custom handler for every level to make sure everything
// is being processed correctly. Each message is a custom message made up for a
// common message and the LogLevel.String().
func TestLoggingLevels(t *testing.T) {
	clearLoggers()
	testMsg := "The space ship is coming"
	logger := New("logger1")
	logger.AddHandler("h0", newTestHandler(t, LogLevelEmergency, "logger1", testMsg, LogLevelEmergency.String()))
	logger.AddHandler("h1", newTestHandler(t, LogLevelAlert, "logger1", testMsg, LogLevelAlert.String()))
	logger.AddHandler("h2", newTestHandler(t, LogLevelCritical, "logger1", testMsg, LogLevelCritical.String()))
	logger.AddHandler("h3", newTestHandler(t, LogLevelError, "logger1", testMsg, LogLevelError.String()))
	logger.AddHandler("h4", newTestHandler(t, LogLevelWarning, "logger1", testMsg, LogLevelWarning.String()))
	logger.AddHandler("h5", newTestHandler(t, LogLevelNotice, "logger1", testMsg, LogLevelNotice.String()))
	logger.AddHandler("h6", newTestHandler(t, LogLevelInfo, "logger1", testMsg, LogLevelInfo.String()))
	logger.AddHandler("h7", newTestHandler(t, LogLevelDebug, "logger1", testMsg, LogLevelDebug.String()))

	logger.Emergency(testMsg, " ", LogLevelEmergency.String())
	logger.Emergencyf("%s %s", testMsg, LogLevelEmergency.String())
	logger.Emergencyln(testMsg, LogLevelEmergency.String())
	logger.Alert(testMsg, " ", LogLevelAlert.String())
	logger.Alertf("%s %s", testMsg, LogLevelAlert.String())
	logger.Alertln(testMsg, LogLevelAlert.String())
	logger.Critical(testMsg, " ", LogLevelCritical.String())
	logger.Criticalf("%s %s", testMsg, LogLevelCritical.String())
	logger.Criticalln(testMsg, LogLevelCritical.String())
	logger.Error(testMsg, " ", LogLevelError.String())
	logger.Errorf("%s %s", testMsg, LogLevelError.String())
	logger.Errorln(testMsg, LogLevelError.String())
	logger.Warning(testMsg, " ", LogLevelWarning.String())
	logger.Warningf("%s %s", testMsg, LogLevelWarning.String())
	logger.Warningln(testMsg, LogLevelWarning.String())
	logger.Notice(testMsg, " ", LogLevelNotice.String())
	logger.Noticef("%s %s", testMsg, LogLevelNotice.String())
	logger.Noticeln(testMsg, LogLevelNotice.String())
	logger.Info(testMsg, " ", LogLevelInfo.String())
	logger.Infof("%s %s", testMsg, LogLevelInfo.String())
	logger.Infoln(testMsg, LogLevelInfo.String())
	logger.Debug(testMsg, " ", LogLevelDebug.String())
	logger.Debugf("%s %s", testMsg, LogLevelDebug.String())
	logger.Debugln(testMsg, LogLevelDebug.String())
}
