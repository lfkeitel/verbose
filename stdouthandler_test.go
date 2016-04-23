package verbose

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStdOutDefaults(t *testing.T) {
	sh := NewStdOutHandler()
	if sh.min != LogLevelDebug {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelCustom, sh.min)
	}
	if sh.max != LogLevelCustom {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelCustom, sh.max)
	}
	if sh.out != os.Stdout {
		t.Error("Incorrect default writer, not Stdout")
	}
}

func TestStdOutLevelSetting(t *testing.T) {
	sh := NewStdOutHandler()
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

func TestStdOutWriteLog(t *testing.T) {
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
	sh := NewStdOutHandler()
	sh.out = buf

	sh.WriteLog(LogLevelInfo, "logger", msg)

	result := buf.String()[30:]
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected %s, got %s", expected, result)
	}
}
