package verbose

import (
	"os"
	"testing"
)

func TestTextDefaults(t *testing.T) {
	sh := NewTextTransport()
	if sh.min != LogLevelDebug {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelDebug, sh.min)
	}
	if sh.max != LogLevelFatal {
		t.Errorf("Incorrect default maximum. Expected %d, got %d", LogLevelFatal, sh.max)
	}
	if sh.out != os.Stderr {
		t.Error("Incorrect default writer, not Stderr")
	}
}

func TestTextLevelSetting(t *testing.T) {
	sh := NewTextTransport()
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
