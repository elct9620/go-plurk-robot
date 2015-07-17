// +build debug

package logger

import (
	"fmt"
	"io"
)

// Debug
func Debug(message string, v ...interface{}) {
	SetStyle(Normal, Blue)
	l.Println(Format("Debug", message, v...))
}

func FDebug(output io.Writer, message string, v ...interface{}) {
	SetStyle(Normal, Blue)
	fmt.Fprintln(output, Format("Debug", message, v...))
}
