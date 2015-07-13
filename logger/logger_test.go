package logger

import (
	"bytes"
	"strings"
	"testing"
)

func createLogger() (*Logger, *bytes.Buffer) {
	var buf []byte
	buffer := bytes.NewBuffer(buf)
	log := New(buffer, "")

	return log, buffer
}

func Test_LoggerNotice(t *testing.T) {

	log, buffer := createLogger()

	log.Notice("Hello %s", "World")

	bufString := string(buffer.Bytes())
	expectedString := "NOTICE: Hello World"

	if !strings.Contains(bufString, expectedString) {
		t.Fatalf("Expected contains %s, but got %s", expectedString, bufString)
	}
}

func Test_LoggerError(t *testing.T) {
	log, buffer := createLogger()
	log.Error("Hello %s", "World")

	bufString := string(buffer.Bytes())
	expectedString := "ERROR: Hello World"
	expectedStyle := "\x1b[0;31m"

	if !strings.Contains(bufString, expectedString) {
		t.Fatalf("Expected contains %s, but got %s", expectedString, bufString)
	}

	if !strings.Contains(bufString, expectedStyle) {
		t.Fatalf("Expected ASCII Color Red, but got %s", bufString)
	}
}

func Test_LoggerDebug(t *testing.T) {
	log, buffer := createLogger()

	log.Debug("Debug %s", "Message")

	if len(buffer.Bytes()) > 0 {
		t.Fatalf("Expected debug message hidden, but got %s", string(buffer.Bytes()))
	}
}

func Test_LoggerSetStyle(t *testing.T) {
	log, _ := createLogger()

	log.SetStyle(1, 1) // Red, Bold

	bufString := log.Format("Style", "")
	expectedStyle := "\x1b[1;31m"

	if !strings.Contains(bufString, expectedStyle) {
		t.Fatalf("Expected style is %s (Red, Bold) but got %s", expectedStyle, bufString)
	}
}
