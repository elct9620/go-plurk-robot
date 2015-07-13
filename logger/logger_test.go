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

func TestNotice(t *testing.T) {

	log, buffer := createLogger()

	log.Notice("Hello World")

	bufString := string(buffer.Bytes())
	expectedString := "NOTICE: Hello World"

	if !strings.Contains(bufString, expectedString) {
		t.Fatalf("Expected contains %s, but got %s", expectedString, bufString)
	}
}

func TestError(t *testing.T) {
	log, buffer := createLogger()
	log.Error("Hello World")

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

func TestDebug(t *testing.T) {
	log, buffer := createLogger()

	log.Debug("Debug Message")

	if len(buffer.Bytes()) > 0 {
		t.Fatalf("Expected debug message hidden, but got %s", string(buffer.Bytes()))
	}
}
