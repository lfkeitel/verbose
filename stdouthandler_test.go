package verbose

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStdoutDefaults(t *testing.T) {
	sh := NewStdoutHandler()
	if sh.min != LogLevelDebug {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelDebug, sh.min)
	}
	if sh.max != LogLevelEmergency {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelEmergency, sh.max)
	}
	if sh.out != os.Stdout {
		t.Error("Incorrect default writer, not Stdout")
	}
}

func TestStdoutLevelSetting(t *testing.T) {
	sh := NewStdoutHandler()
	sh.SetLevel(LogLevelWarning)
	if sh.min != LogLevelWarning {
		t.Errorf("Min level not set correctly. Expected %d, got %d", LogLevelWarning, sh.min)
	}
	if sh.max != LogLevelWarning {
		t.Errorf("Max level not set correctly. Expected %d, got %d", LogLevelWarning, sh.max)
	}

	sh.SetMinLevel(LogLevelInfo)
	if sh.min != LogLevelInfo {
		t.Errorf("Min level not set correctly. Expected %d, got %d", LogLevelInfo, sh.min)
	}

	sh.SetMaxLevel(LogLevelAlert)
	if sh.max != LogLevelAlert {
		t.Errorf("Max level not set correctly. Expected %d, got %d", LogLevelAlert, sh.max)
	}

	if sh.Handles(LogLevelDebug) {
		t.Errorf("Incorrect Handles result. Expected false, got %t", sh.Handles(LogLevelDebug))
	}
	if sh.Handles(LogLevelEmergency) {
		t.Errorf("Incorrect Handles result. Expected false, got %t", sh.Handles(LogLevelEmergency))
	}

	if !sh.Handles(LogLevelCritical) {
		t.Errorf("Incorrect Handles result. Expected true, got %t", sh.Handles(LogLevelCritical))
	}
}

func TestStdoutWriteLog(t *testing.T) {
	buf := &bytes.Buffer{}
	msg := "My spoon is too big"
	expected := fmt.Sprintf(
		"%s%s: %s%s: %s%s\n",
		colors[LogLevelInfo],
		strings.ToUpper(LogLevelInfo.String()),
		ColorGreen,
		"logger",
		ColorReset,
		msg,
	)
	sh := NewStdoutHandler()
	sh.out = buf

	e := NewEntry(&Logger{name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	sh.WriteLog(e)

	result := buf.String()[30:]
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected %s, got %s", expected, result)
	}
}
