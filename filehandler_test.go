package verbose

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	testLogFile string = "test.log"
	testLogDir  string = "logs"
)

func cleanup() {
	os.Remove(testLogFile)
	os.RemoveAll(testLogDir)
}

func TestDefaults(t *testing.T) {
	defer cleanup()
	fh, err := NewFileHandler(testLogFile)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
	if fh.min != LogLevelDebug {
		t.Errorf("Incorrect default minimum. Expected %d, got %d", LogLevelDebug, fh.min)
	}
	if fh.max != LogLevelFatal {
		t.Errorf("Incorrect default maximum. Expected %d, got %d", LogLevelFatal, fh.max)
	}
	if fh.separate {
		t.Error("Incorrect separate field. Expected false, got true")
	}

	fh, err = NewFileHandler(testLogDir)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
	if !fh.separate {
		t.Error("Incorrect separate field. Expected true, got false")
	}
}

func TestFileHandlerLevelSetting(t *testing.T) {
	fh := &FileHandler{}
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
	// Test write to single file
	fh, err := NewFileHandler(testLogFile)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
	e := NewEntry(&Logger{name: "logger"})
	e.Level = LogLevelAlert
	e.Message = "What? No coffee!?"
	fh.WriteLog(e)

	stat, _ := os.Stat(testLogFile)
	if stat.Size() != 55 {
		t.Errorf("Incorrect log file size. Expected 55, got %d", stat.Size())
	}

	// Test write to directory
	fh, err = NewFileHandler(testLogDir)
	if err != nil {
		t.Fatalf("Error making file handler: %s", err.Error())
	}
	e = NewEntry(&Logger{name: "logger"})
	e.Level = LogLevelAlert
	e.Message = "What? No coffee!?"
	fh.WriteLog(e)

	stat, err = os.Stat(filepath.Join(testLogDir, "alert-logger.log"))
	if err != nil {
		t.Fatalf("Error stating log file: %s", err.Error())
	}
	if stat.Size() != 55 {
		t.Errorf("Incorrect log file size. Expected 55, got %d", stat.Size())
	}
}
