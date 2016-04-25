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
	"Fatal",
}

var fileTmpl = `// This file was generated with level_generator. DO NOT EDIT

package verbose

import (
	"os"
	"fmt"
)
{{range .}}
// {{.}} - Log {{.}} message
func (e *Entry) {{.}}(v ...interface{}) {
    e.log(LogLevel{{.}}, fmt.Sprint(v...)){{if eq . "Fatal"}}
	os.Exit(1){{end}}
    return
}
{{end}}
// Panic - Log Panic message
func (e *Entry) Panic(v ...interface{}) {
    e.log(LogLevelEmergency, fmt.Sprint(v...))
    return
}

// Print - Log Print message
func (e *Entry) Print(v ...interface{}) {
    e.log(LogLevelInfo, fmt.Sprint(v...))
    return
}

// Printf friendly functions
{{range .}}
// {{.}}f - Log formatted {{.}} message
func (e *Entry) {{.}}f(m string, v ...interface{}) {
    e.log(LogLevel{{.}}, fmt.Sprintf(m, v...)){{if eq . "Fatal"}}
	os.Exit(1){{end}}
    return
}
{{end}}
// Panicf - Log formatted Panic message
func (e *Entry) Panicf(m string, v ...interface{}) {
    e.log(LogLevelEmergency, fmt.Sprintf(m, v...))
    return
}

// Printf - Log formatted Print message
func (e *Entry) Printf(m string, v ...interface{}) {
    e.log(LogLevelInfo, fmt.Sprintf(m, v...))
    return
}

// Println friendly functions
{{range .}}
// {{.}}ln - Log {{.}} message with newline
func (e *Entry) {{.}}ln(v ...interface{}) {
    e.log(LogLevel{{.}}, e.sprintlnn(v...)){{if eq . "Fatal"}}
	os.Exit(1){{end}}
    return
}
{{end}}
// Panicln - Log Panic message with newline
func (e *Entry) Panicln(v ...interface{}) {
    e.log(LogLevelEmergency, e.sprintlnn(v...))
    return
}

// Println - Log Print message with newline
func (e *Entry) Println(v ...interface{}) {
    e.log(LogLevelInfo, e.sprintlnn(v...))
    return
}
`

func init() {
	flag.StringVar(&outFile, "output", "log_levels.go", "Filename of output")
}

func main() {
	flag.Parse()

	tmpl, err := template.New("funcs").Parse(fileTmpl)
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
	tmpl.Execute(file, levels)
}
