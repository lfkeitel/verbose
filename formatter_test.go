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
	expected := `{"timestamp":"%s","level":"INFO","logger":"logger","message":"My spoon is too big","data":[{"key1":"value1"}]}`
	expected = fmt.Sprintf(expected, now.Format("2006-01-02 15:04:05 MST"))

	formatter := &JSONFormatter{}

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
		`%s: %s: %s: %s "key1": "value1"%s`,
		now.Format("2006-01-02 15:04:05 MST"),
		strings.ToUpper(LogLevelInfo.String()),
		"logger",
		msg,
		"\n",
	)
	formatter := &LineFormatter{}

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
		`%s%s: %s%s: %s%s: %s%s "key1": "value1"%s`,
		ColorGrey,
		now.Format("2006-01-02 15:04:05 MST"),
		colors[LogLevelInfo],
		strings.ToUpper(LogLevelInfo.String()),
		ColorGreen,
		"logger",
		ColorReset,
		msg,
		"\n",
	)
	formatter := &ColoredLineFormatter{}

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
