package verbose

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var colors = map[LogLevel]Color{
	LogLevelDebug:     ColorBlue,
	LogLevelInfo:      ColorCyan,
	LogLevelNotice:    ColorCyan,
	LogLevelWarning:   ColorMagenta,
	LogLevelError:     ColorRed,
	LogLevelCritical:  ColorRed,
	LogLevelAlert:     ColorRed,
	LogLevelEmergency: ColorRed,
	LogLevelFatal:     ColorRed,
	LogLevelCustom:    ColorWhite,
}

type StdOutHandler struct {
	min LogLevel
	max LogLevel
	out io.Writer // Usually os.Stdout, mainly used for testing
}

func NewStdOutHandler() *StdOutHandler {
	return &StdOutHandler{
		min: LogLevelDebug,
		max: LogLevelCustom,
		out: os.Stdout,
	}
}

func (s *StdOutHandler) SetLevel(l LogLevel) {
	s.min = l
	s.max = l
}

func (s *StdOutHandler) SetMinLevel(l LogLevel) {
	if l > s.max {
		return
	}
	s.min = l
}

func (s *StdOutHandler) SetMaxLevel(l LogLevel) {
	if l < s.min {
		return
	}
	s.max = l
}

func (s *StdOutHandler) Handles(l LogLevel) bool {
	return (s.min <= l && l <= s.max)
}

func (s *StdOutHandler) WriteLog(l LogLevel, name, msg string) {
	now := time.Now().Format("2006-01-02 15:04:05 MST")
	fmt.Fprintf(
		s.out,
		"%s%s: %s%s: %s%s: %s%s\n",
		ColorGrey,
		now,
		colors[l],
		strings.ToUpper(l.String()),
		ColorGreen,
		name,
		ColorReset,
		msg,
	)
}
