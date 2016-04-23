package verbose

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type FileHandler struct {
	min      LogLevel
	max      LogLevel
	path     string
	separate bool
	m        sync.Mutex
}

func NewFileHandler(path string) (*FileHandler, error) {
	f := &FileHandler{
		min:  LogLevelDebug,
		max:  LogLevelCustom,
		path: path,
		m:    sync.Mutex{},
	}

	// Determine of the path is a file or directory
	// We cannot assume the path exists yet
	stat, err := os.Stat(path)
	if err == nil { // Easiest, path exists
		f.separate = stat.IsDir()
	} else if os.IsNotExist(err) {
		// Typically an extension means it's a file
		ext := filepath.Ext(path)
		if ext == "" {
			// Attempt to create the directory
			if err := os.MkdirAll(path, 0755); err != nil {
				return nil, err
			}
			f.separate = true
		} else {
			// Attempt to create the file
			file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				return nil, err
			}
			file.Close()
			f.separate = false
		}
	}
	return f, nil
}

func (f *FileHandler) SetLevel(l LogLevel) {
	f.min = l
	f.max = l
}

func (f *FileHandler) SetMinLevel(l LogLevel) {
	if l > f.max {
		return
	}
	f.min = l
}

func (f *FileHandler) SetMaxLevel(l LogLevel) {
	if l < f.min {
		return
	}
	f.max = l
}

func (f *FileHandler) Handles(l LogLevel) bool {
	return (f.min <= l && l <= f.max)
}

func (f *FileHandler) WriteLog(l LogLevel, name, msg string) {
	now := time.Now().Format("2006-01-02 15:04:05 MST")
	var logfile string
	if !f.separate {
		logfile = f.path
	} else {
		logfile = fmt.Sprintf("%s-%s.log", strings.ToLower(l.String()), name)
		logfile = path.Join(f.path, logfile)
	}

	f.m.Lock()
	defer f.m.Unlock()

	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %s\n", err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(now + ": " + msg + "\n")
	if err != nil {
		fmt.Printf("Error writing to log file: %s\n", err.Error())
	}
	return
}
