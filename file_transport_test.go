package verbose

import (
	"os"
	"testing"
)

const (
	testLogFile string = "test.log"
)

func cleanup() {
	os.Remove(testLogFile)
}

func TestDefaults(t *testing.T) {
	defer cleanup()
	fh, err := NewFileTransport(testLogFile)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
	if fh.min != LogLevelDebug {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelDebug, fh.min)
	}
	if fh.max != LogLevelFatal {
		t.Errorf("Incorrect default maximum. Expected %d, got %d", LogLevelFatal, fh.max)
	}

	fh, err = NewFileTransport(testLogFile)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
}

func TestFileHandlerLevelSetting(t *testing.T) {
	fh := &FileTransport{}
	fh.SetLevel(LogLevelWarning)
	if fh.min != LogLevelWarning {
		t.Errorf("Min level not set correctly. Expected %d, got %d", LogLevelWarning, fh.min)
	}
	if fh.max != LogLevelWarning {
		t.Errorf("Max level not set correctly. Expected %d, got %d", LogLevelWarning, fh.max)
	}

	fh.SetMinLevel(LogLevelInfo)
	if fh.min != LogLevelInfo {
		t.Errorf("Min level not set correctly. Expected %d, got %d", LogLevelInfo, fh.min)
	}

	fh.SetMaxLevel(LogLevelAlert)
	if fh.max != LogLevelAlert {
		t.Errorf("Max level not set correctly. Expected %d, got %d", LogLevelAlert, fh.max)
	}

	if fh.Handles(LogLevelDebug) {
		t.Errorf("Incorrect Handles result. Expected false, got %t", fh.Handles(LogLevelDebug))
	}
	if fh.Handles(LogLevelEmergency) {
		t.Errorf("Incorrect Handles result. Expected false, got %t", fh.Handles(LogLevelEmergency))
	}

	if !fh.Handles(LogLevelCritical) {
		t.Errorf("Incorrect Handles result. Expected true, got %t", fh.Handles(LogLevelCritical))
	}
}

func TestFileHandlerWriteLog(t *testing.T) {
	defer cleanup()

	fh, err := NewFileTransport(testLogFile)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
	e := NewEntry(&Logger{Name: "logger"})
	e.Level = LogLevelAlert
	e.Message = "What? No coffee!?"
	fh.WriteLog(e)

	stat, _ := os.Stat(testLogFile)
	if stat.Size() != 55 {
		t.Errorf("Incorrect log file size. Expected 55, got %d", stat.Size())
	}
}
