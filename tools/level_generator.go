// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var outFile string

var levels = [...]string{
	"Debug",
	"Info",
	"Notice",
	"Warning",
	"Error",
	"Critical",
	"Alert",
	"Emergency",
}

var header = `// This file was generated with level_generator. DO NOT EDIT

package verbose

import (
	"os"
	"fmt"
)
`

var funcTemplate = `
// {{.}} - Log {{.}} message
func (l *Logger) {{.}}(m string) {
    l.Log(LogLevel{{.}}, m)
    return
}

// {{.}}f - Log formatted {{.}} message
func (l *Logger) {{.}}f(m string, v ...interface{}) {
    l.Log(LogLevel{{.}}, fmt.Sprintf(m, v...))
    return
}
`

var fatalFuncs = `
// Fatal - Log Fatal message
func (l *Logger) Fatal(m string) {
    l.Log(LogLevelFatal, m)
    os.Exit(1)
    return
}

// Fatalf - Log formatted Fatal message
func (l *Logger) Fatalf(m string, v ...interface{}) {
    l.Log(LogLevelFatal, fmt.Sprintf(m, v...))
    os.Exit(1)
    return
}
`

func init() {
	flag.StringVar(&outFile, "output", "log_levels.go", "Filename of output")
}

func main() {
	flag.Parse()

	tmpl, err := template.New("funcs").Parse(funcTemplate)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(header)

	for _, l := range levels {
		tmpl.Execute(file, l)
	}

	file.WriteString(fatalFuncs)
}
