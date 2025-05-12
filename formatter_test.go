package verbose

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestJSONFormatter(t *testing.T) {
	msg := "My spoon is too big"
	data := Fields{
		"key1": "value1",
	}
	now := time.Now()

	type event struct {
		Timestamp time.Time
		Level     string
		Logger    string
		Message   string
		Data      struct {
			Key1 string
		}
	}

	expected := &event{
		now,
		"INFO",
		"logger",
		msg,
		struct {
			Key1 string
		}{"value1"},
	}

	formatter := NewJSONFormatter()

	e := NewEntry(&Logger{Name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	e.Timestamp = now
	e.Data = data

	var decoded event
	formatted := formatter.FormatByte(e)
	if formatted[len(formatted)-1] != '\n' {
		t.Fatal("JSON formatter doesn't end in a newline")
	}

	if err := json.Unmarshal(formatted, &decoded); err != nil {
		t.Errorf("JSON failed to decode into event struct: %s", err)
	}

	if expected.Timestamp.Format(time.RFC3339) != decoded.Timestamp.Format(time.RFC3339) {
		t.Errorf("JSON formatter wrong timestamp. Wanted %s, got %s",
			expected.Timestamp.Format(time.RFC3339),
			decoded.Timestamp.Format(time.RFC3339))
	}

	if expected.Data.Key1 != decoded.Data.Key1 {
		t.Errorf("JSON formatter wrong data. Wanted %s, got %s", expected.Data.Key1, decoded.Data.Key1)
	}
}

func TestLineFormatter(t *testing.T) {
	msg := "My spoon is too big"
	data := Fields{
		"key1": "value1",
	}
	now := time.Now()
	expected := fmt.Sprintf(
		`%s: %s: %s: %s | "key1": "value1"%s`,
		now.Format(time.RFC3339),
		strings.ToUpper(LogLevelInfo.String()),
		"logger",
		msg,
		"\n",
	)
	formatter := NewLineFormatter(false)

	e := NewEntry(&Logger{Name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	e.Timestamp = now
	e.Data = data

	result := formatter.Format(e)
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected `%s`, got `%s`", expected, result)
	}
}

func TestColoredLineFormatter(t *testing.T) {
	msg := "My spoon is too big"
	data := Fields{
		"key1": "value1",
	}
	now := time.Now()
	expected := fmt.Sprintf(
		`%s%s: %s%s: %s%s: %s%s | "key1": "value1"%s`,
		ColorGrey,
		now.Format(time.RFC3339),
		colors[LogLevelInfo],
		strings.ToUpper(LogLevelInfo.String()),
		ColorGreen,
		"logger",
		ColorReset,
		msg,
		"\n",
	)
	formatter := NewLineFormatter(true)

	e := NewEntry(&Logger{Name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	e.Timestamp = now
	e.Data = data

	result := formatter.Format(e)
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected `%s`, got `%s`", expected, result)
	}
}

func TestLogFmtFormatter(t *testing.T) {
	msg := "My spoon is too big"
	data := Fields{
		"key1": "value1",
	}
	now := time.Now()
	expected := `timestamp="%s" level=INFO logger="logger" msg="My spoon is too big" key1="value1"` + "\n"
	expected = fmt.Sprintf(expected, now.Format(time.RFC3339))
	formatter := NewLogFmtFormatter()

	e := NewEntry(&Logger{Name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	e.Timestamp = now
	e.Data = data

	result := formatter.Format(e)
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected `%s`, got `%s`", expected, result)
	}
}
