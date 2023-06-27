package log

import (
	"fmt"
	"time"
)

var IsTraceEnabled bool
var IsDebugEnabled bool
var IsInfoEnabled bool
var IsWarnEnabled bool
var IsErrorEnabled bool

const LAYOUT = "2023-01-01 00:00:00"

func SetLevel(level string) {
	switch level {
	case "error":
		SetErrorLevel()
	case "warn":
		SetWarnLevel()
	case "info":
		SetInfoLevel()
	case "debug":
		SetDebugLevel()
	case "trace":
		SetTraceLevel()
	}
}

func SetErrorLevel() {
	SetWarnLevel()
	IsWarnEnabled = false
}

func SetWarnLevel() {
	SetInfoLevel()
	IsInfoEnabled = false
}

func SetInfoLevel() {
	SetDebugLevel()
	IsDebugEnabled = false
}

func SetDebugLevel() {
	SetTraceLevel()
	IsTraceEnabled = false
}

func SetTraceLevel() {
	IsTraceEnabled = true
	IsDebugEnabled = true
	IsInfoEnabled = true
	IsWarnEnabled = true
	IsErrorEnabled = true
}

func Debug(messages ...any) {
	if IsDebugEnabled {
		Println("debug:", messages)
	}
}

func Trace(messages ...any) {
	if IsTraceEnabled {
		Println("trace:", messages)
	}
}

func Info(messages ...any) {
	if IsInfoEnabled {
		Println("info:", messages)
	}
}

func Warn(messages ...any) {
	if IsWarnEnabled {
		Println("warn:", messages)
	}
}

func Error(messages ...any) {
	if IsErrorEnabled {
		Println("error:", messages)
	}
}

func init() {
	SetWarnLevel()
}

func Println(level string, messages ...any) {
	fmt.Print(time.Now().Format(LAYOUT))
	fmt.Println("  ", level, messages)
}
