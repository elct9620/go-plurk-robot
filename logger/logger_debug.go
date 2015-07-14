// +build debug

package logger

// Debug
func Debug(message string, v ...interface{}) {
	SetStyle(Normal, Blue)
	l.Println(Format("Debug", message, v...))
}
