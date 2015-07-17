// Simple logger
package logger

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

const (
	Normal = iota
	Bold
	_
	_
	Underline
)

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	NoColor
)

type NullOutput struct{}

func (NullOutput) Write(p []byte) (int, error) {
	return len(p), nil
}

var prefix string = ""
var outputStyle string = ""
var output io.Writer = NullOutput{}

var l *log.Logger = log.New(output, "", 0)

// New Logger
func Config(out io.Writer, p string) {
	if out == nil {
		out = NullOutput{}
	}
	l = log.New(out, "", 0)
	prefix = p
}

func SetStyle(style int, color int) {
	outputStyle = fmt.Sprintf("\x1b[%0d;3%dm", style, color)
}

// Notice
func Info(message string, v ...interface{}) {
	SetStyle(Normal, NoColor)
	l.Println(Format("Info", message, v...))
}

func FInfo(output io.Writer, message string, v ...interface{}) {
	SetStyle(Normal, NoColor)
	fmt.Fprintln(output, Format("Info", message, v...))
}

// Warn
func Warn(message string, v ...interface{}) {
	SetStyle(Normal, Yellow)
	l.Println(Format("Warn", message, v...))
}

func FWarn(output io.Writer, message string, v ...interface{}) {
	SetStyle(Normal, Yellow)
	fmt.Fprintln(output, Format("Warn", message, v...))
}

// Error
func Error(message string, v ...interface{}) {
	SetStyle(Normal, Red)
	l.Println(Format("Error", message, v...))
}

func FError(output io.Writer, message string, v ...interface{}) {
	SetStyle(Normal, Red)
	fmt.Fprintln(output, Format("Error", message, v...))
}

// Format Output String
func Format(logType string, message string, v ...interface{}) string {
	now := time.Now().Local().Format("2006/01/02 03:04:05")
	return fmt.Sprintf("%s%s [ %s ] %6s: %s", outputStyle, now, prefix, strings.ToUpper(logType), fmt.Sprintf(message, v...))
}
