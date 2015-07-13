// +build debug

package logger

// Debug
func (l Logger) Debug(message string, v ...interface{}) {
	l.SetStyle(Normal, Blue)
	l.Println(l.Format("Debug", message, v...))
}
