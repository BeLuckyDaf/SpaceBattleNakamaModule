// Copyright 2020 Vladislav Smirnov

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Logger stores the file and the writer used for logging
type Logger struct {
	f *os.File
	w *bufio.Writer
}

// NewLogger creates and initializes a Logger
func NewLogger(filepath string) *Logger {
	l := new(Logger)
	var err error
	l.f, err = os.Create(filepath)
	if err != nil {
		return nil
	}
	l.w = bufio.NewWriter(l.f)
	if err != nil {
		return nil
	}
	return l
}

// Log is used to actually write log to file and terminal
func (l *Logger) Log(a ...interface{}) {
	t := time.Now().Format("01.02.06 15:04:05 --")
	l.w.WriteString(fmt.Sprintln(t, a))
	fmt.Println(a...)
	l.w.Flush()
}

// Close is used to finalize and close the file
func (l *Logger) Close() {
	l.f.Close()
}
