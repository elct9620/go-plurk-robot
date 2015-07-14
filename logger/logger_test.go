package logger

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func configLogger() *bytes.Buffer {
	var buf []byte
	buffer := bytes.NewBuffer(buf)
	Config(buffer, "")

	return buffer
}

func Test_Info(t *testing.T) {

	buffer := configLogger()

	Info("Hello World")

	bufString := string(buffer.Bytes())
	expectedString := "INFO: Hello World"

	assert.Contains(t, bufString, expectedString)
}

func Test_Error(t *testing.T) {
	buffer := configLogger()
	Error("Hello %s", "World")

	bufString := string(buffer.Bytes())
	expectedString := "ERROR: Hello World"
	expectedStyle := "\x1b[0;31m"

	assert.Contains(t, bufString, expectedString)
	assert.Contains(t, bufString, expectedStyle)
}

func Test_Warn(t *testing.T) {
	buffer := configLogger()
	Warn("Hello %s", "World")

	bufString := string(buffer.Bytes())
	expectedString := "WARN: Hello World"
	expectedStyle := "\x1b[0;33m"

	assert.Contains(t, bufString, expectedString)
	assert.Contains(t, bufString, expectedStyle)
}

func Test_Debug(t *testing.T) {
	buffer := configLogger()

	Debug("Debug %s", "Message")

	assert.Len(t, buffer.Bytes(), 0)
}

func Test_SetStyle(t *testing.T) {
	configLogger()

	SetStyle(1, 1) // Red, Bold

	bufString := Format("Style", "")
	expectedStyle := "\x1b[1;31m"

	assert.Contains(t, bufString, expectedStyle)
}
