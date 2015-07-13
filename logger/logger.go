// Simple logger
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
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

type Logger struct {
	prefix      string
	outputStyle string
	*log.Logger
}

// New Logger
func New(out io.Writer, prefix string) *Logger {
	if out == nil {
		out = os.Stdout
	}
	return &Logger{prefix: prefix, Logger: log.New(out, "", 0)}
}

func (l *Logger) SetStyle(style int, color int) {
	l.outputStyle = fmt.Sprintf("\x1b[%0d;3%dm", style, color)
}

// Notice
func (l *Logger) Notice(message string, v ...interface{}) {
	l.SetStyle(Normal, NoColor)
	l.Println(l.Format("Notice", message, v...))
}

// Error
func (l *Logger) Error(message string, v ...interface{}) {
	l.SetStyle(Normal, Red)
	l.Println(l.Format("Error", message, v...))
}

// Format Output String
func (l *Logger) Format(logType string, message string, v ...interface{}) string {
	now := time.Now().Local().Format("2006/01/02 03:04:05")
	return fmt.Sprintf("%s%s [ %s ] %6s: %s", l.outputStyle, now, l.prefix, strings.ToUpper(logType), fmt.Sprintf(message, v...))
}
