package verbose

import (
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
	expected := `{"timestamp":"%s","level":"INFO","logger":"logger","message":"My spoon is too big","data":{"key1":"value1"}}` + "\n"
	expected = fmt.Sprintf(expected, now.Format(time.RFC3339))

	formatter := NewJSONFormatter()

	e := NewEntry(&Logger{name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	e.Timestamp = now
	e.Data = data

	result := formatter.Format(e)
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected `%s`, got `%s`", expected, result)
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
	formatter := NewLineFormatter()

	e := NewEntry(&Logger{name: "logger"})
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
	formatter := NewColoredLineFormatter()

	e := NewEntry(&Logger{name: "logger"})
	e.Level = LogLevelInfo
	e.Message = msg
	e.Timestamp = now
	e.Data = data

	result := formatter.Format(e)
	if result != expected {
		t.Errorf("Incorrectly formatted message. Expected `%s`, got `%s`", expected, result)
	}
}
