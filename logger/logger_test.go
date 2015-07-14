package logger

import (
	"bytes"
	"strings"
	"testing"
)

func configLogger() *bytes.Buffer {
	var buf []byte
	buffer := bytes.NewBuffer(buf)
	Config(buffer, "")

	return buffer
}

func Test_LoggerInfo(t *testing.T) {

	buffer := configLogger()

	Info("Hello %s", "World")

	bufString := string(buffer.Bytes())
	expectedString := "INFO: Hello World"

	if !strings.Contains(bufString, expectedString) {
		t.Fatalf("Expected contains %s, but got %s", expectedString, bufString)
	}
}

func Test_Error(t *testing.T) {
	buffer := configLogger()
	Error("Hello %s", "World")

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

func Test_Warn(t *testing.T) {
	buffer := configLogger()
	Warn("Hello %s", "World")

	bufString := string(buffer.Bytes())
	expectedString := "WARN: Hello World"
	expectedStyle := "\x1b[0;33m"

	if !strings.Contains(bufString, expectedString) {
		t.Fatalf("Expected contains %s, but got %s", expectedString, bufString)
	}

	if !strings.Contains(bufString, expectedStyle) {
		t.Fatalf("Expected ASCII Color Yellow, but got %s", bufString)
	}
}

func Test_Debug(t *testing.T) {
	buffer := configLogger()

	Debug("Debug %s", "Message")

	if len(buffer.Bytes()) > 0 {
		t.Fatalf("Expected debug message hidden, but got %s", string(buffer.Bytes()))
	}
}

func Test_SetStyle(t *testing.T) {
	configLogger()

	SetStyle(1, 1) // Red, Bold

	bufString := Format("Style", "")
	expectedStyle := "\x1b[1;31m"

	if !strings.Contains(bufString, expectedStyle) {
		t.Fatalf("Expected style is %s (Red, Bold) but got %s", expectedStyle, bufString)
	}
}
